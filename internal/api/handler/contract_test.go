package handler

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
)

const (
	usdcAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	zeroAddress = "0x0000000000000000000000000000000000000000"
)

func testTokens() config.Tokens {
	return config.Tokens{
		"ETH":  {Address: zeroAddress, Decimals: 18},
		"USDC": {Address: usdcAddress, Decimals: 6},
	}
}

func decodeError(t *testing.T, rec *httptest.ResponseRecorder) map[string]string {
	t.Helper()
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not a JSON error: %v (body %q)", err, rec.Body.String())
	}
	return body
}

func TestParseBigInt(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    string
		wantErr bool
	}{
		{name: "zero", value: "0", want: "0"},
		{name: "positive", value: "42", want: "42"},
		{name: "negative", value: "-42", want: "-42"},
		{name: "surrounding whitespace is trimmed", value: "  42  ", want: "42"},
		{
			name:  "a 216-bit invoice id",
			value: "105312291668557186697918027683670432318895095400549111254310977535",
			want:  "105312291668557186697918027683670432318895095400549111254310977535",
		},
		{name: "empty", value: "", wantErr: true},
		{name: "whitespace only", value: "   ", wantErr: true},
		{name: "not a number", value: "abc", wantErr: true},

		{name: "hex is rejected", value: "0x2a", wantErr: true},
		{name: "decimal point", value: "1.5", wantErr: true},
		{name: "trailing characters", value: "42abc", wantErr: true},
		{name: "internal whitespace", value: "4 2", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBigInt("invoiceId", tt.value)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("parseBigInt(%q) = %v, want an error", tt.value, got)
				}

				if !strings.Contains(err.Error(), "invoiceId") {
					t.Errorf("error = %q, want it to name the field", err)
				}
				if !strings.Contains(err.Error(), `"`+tt.value+`"`) {
					t.Errorf("error = %q, want it to quote the value", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseBigInt(%q) returned %v", tt.value, err)
			}
			if got.String() != tt.want {
				t.Errorf("parseBigInt(%q) = %s, want %s", tt.value, got, tt.want)
			}
		})
	}
}

func TestTxURL(t *testing.T) {
	const hash = "0xabc123"

	t.Run("with an explorer", func(t *testing.T) {
		h := &ContractHandler{ExplorerURL: "https://etherscan.io"}

		if got, want := h.txURL(hash), "https://etherscan.io/tx/0xabc123"; got != want {
			t.Errorf("txURL = %q, want %q", got, want)
		}
	})

	t.Run("without an explorer", func(t *testing.T) {
		h := &ContractHandler{}

		if got := h.txURL(hash); got != hash {
			t.Errorf("txURL = %q, want the bare hash", got)
		}
	})
}

func TestFeeReceiverReady(t *testing.T) {
	tests := []struct {
		name    string
		handler *ContractHandler
		want    bool
	}{
		{
			name:    "no fee receiver client",
			handler: &ContractHandler{ChainID: big.NewInt(1)},
		},
		{
			name:    "no chain id",
			handler: &ContractHandler{FeeReceiver: &feereceiver.Client{}},
		},
		{
			name:    "neither configured",
			handler: &ContractHandler{},
		},
		{
			name:    "both configured",
			handler: &ContractHandler{FeeReceiver: &feereceiver.Client{}, ChainID: big.NewInt(1)},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			got := tt.handler.feeReceiverReady(rec)

			if got != tt.want {
				t.Fatalf("feeReceiverReady = %v, want %v", got, tt.want)
			}
			if tt.want {
				if rec.Body.Len() != 0 {
					t.Errorf("feeReceiverReady wrote %q when ready", rec.Body.String())
				}
				return
			}

			if rec.Code != http.StatusServiceUnavailable {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
			}
			body := decodeError(t, rec)
			if !strings.Contains(body["error"], "fee receivers are unavailable") {
				t.Errorf("error = %q, want it to say the endpoint is unavailable", body["error"])
			}
		})
	}
}

func TestNewContractHandler(t *testing.T) {
	in := &ContractHandler{
		ExplorerURL: "https://etherscan.io",
		Tokens:      testTokens(),
		BaseUrl:     "https://pay.example.com/checkout?data=",
		ChainID:     big.NewInt(8453),
	}

	got := NewContractHandler(in)

	if got == in {
		t.Error("NewContractHandler returned the argument itself, want a copy")
	}
	if got.ExplorerURL != in.ExplorerURL {
		t.Errorf("ExplorerURL = %q, want %q", got.ExplorerURL, in.ExplorerURL)
	}
	if got.BaseUrl != in.BaseUrl {
		t.Errorf("BaseUrl = %q, want %q", got.BaseUrl, in.BaseUrl)
	}
	if got.ChainID == nil || got.ChainID.Cmp(in.ChainID) != 0 {
		t.Errorf("ChainID = %v, want %v", got.ChainID, in.ChainID)
	}
	if len(got.Tokens) != len(in.Tokens) {
		t.Errorf("Tokens has %d entries, want %d", len(got.Tokens), len(in.Tokens))
	}
}

func TestNewContractHandlerWithOptionalFieldsUnset(t *testing.T) {
	got := NewContractHandler(&ContractHandler{})

	if got == nil {
		t.Fatal("NewContractHandler returned nil")
	}
	if got.FeeReceiver != nil {
		t.Error("FeeReceiver was populated from an empty handler")
	}
	if got.Oracle != nil {
		t.Error("Oracle was populated from an empty handler")
	}
}
