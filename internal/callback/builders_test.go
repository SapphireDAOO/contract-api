package callback

import (
	"encoding/json"
	"math/big"
	"strings"
	"testing"
	"time"
)

type stubToken struct {
	symbol   string
	decimals int
}

type stubTokens map[string]stubToken

func (s stubTokens) ByAddress(address string) (string, int, bool) {
	token, ok := s[strings.ToLower(address)]
	if !ok {
		return "", 0, false
	}
	return token.symbol, token.decimals, true
}

const (
	ethAddress  = "0x0000000000000000000000000000000000000000"
	usdcAddress = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	txURL       = "https://etherscan.io/tx/0xabc"
)

func testClient() *Client {
	return NewClient("https://market.example.com/cb", "key", stubTokens{
		ethAddress:  {symbol: "ETH", decimals: 18},
		usdcAddress: {symbol: "USDC", decimals: 6},
	})
}

func wei(t *testing.T, s string) *big.Int {
	t.Helper()
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		t.Fatalf("bad big.Int literal %q", s)
	}
	return n
}

func TestBuildPaymentReceivedCallbackPayload(t *testing.T) {
	client := testClient()
	before := time.Now().Add(10 * time.Minute).UnixMilli()

	raw, err := client.buildPaymentReceivedCallbackPayload(txURL, ethAddress, wei(t, "1500000000000000000"), 1700000000)
	if err != nil {
		t.Fatalf("buildPaymentReceivedCallbackPayload returned %v", err)
	}
	after := time.Now().Add(10 * time.Minute).UnixMilli()

	var payload paymentReceivedCallbackPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("payload is not valid JSON: %v", err)
	}

	if payload.Currency != "ETH" {
		t.Errorf("currency = %q, want ETH", payload.Currency)
	}
	if payload.Amount != "1.5" {
		t.Errorf("amount = %q, want 1.5", payload.Amount)
	}

	if payload.TransactionAmount != payload.Amount {
		t.Errorf("transactionAmount = %q, want it to match amount %q", payload.TransactionAmount, payload.Amount)
	}
	if payload.TransactionUrl != txURL {
		t.Errorf("transactionUrl = %q, want %q", payload.TransactionUrl, txURL)
	}
	if payload.TransactionTimestamp != 1700000000 {
		t.Errorf("transactionTimestamp = %d, want 1700000000", payload.TransactionTimestamp)
	}

	if payload.Releases < before || payload.Releases > after {
		t.Errorf("releases = %d, want it between %d and %d", payload.Releases, before, after)
	}
}

func TestBuildPaymentReceivedCallbackPayloadUsesTokenDecimals(t *testing.T) {
	client := testClient()

	raw, err := client.buildPaymentReceivedCallbackPayload(txURL, usdcAddress, wei(t, "1500000"), 1700000000)
	if err != nil {
		t.Fatalf("buildPaymentReceivedCallbackPayload returned %v", err)
	}

	var payload paymentReceivedCallbackPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("payload is not valid JSON: %v", err)
	}
	if payload.Currency != "USDC" {
		t.Errorf("currency = %q, want USDC", payload.Currency)
	}

	if payload.Amount != "1.5" {
		t.Errorf("amount = %q, want 1.5", payload.Amount)
	}
}

func TestBuildRefundCallbackPayload(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		amount      string
		refundShare string
		wantAmount  string
		wantCurrenc string
	}{
		{
			name:        "nil share refunds the whole amount",
			token:       ethAddress,
			amount:      "1000000000000000000",
			refundShare: "",
			wantAmount:  "1",
			wantCurrenc: "ETH",
		},
		{
			name:        "full share in basis points",
			token:       ethAddress,
			amount:      "1000000000000000000",
			refundShare: "10000",
			wantAmount:  "1",
			wantCurrenc: "ETH",
		},
		{
			name:        "half share in basis points",
			token:       ethAddress,
			amount:      "1000000000000000000",
			refundShare: "5000",
			wantAmount:  "0.5",
			wantCurrenc: "ETH",
		},
		{
			name:        "one percent share",
			token:       ethAddress,
			amount:      "1000000000000000000",
			refundShare: "100",
			wantAmount:  "0.01",
			wantCurrenc: "ETH",
		},
		{
			name:        "zero share refunds nothing",
			token:       ethAddress,
			amount:      "1000000000000000000",
			refundShare: "0",
			wantAmount:  "0",
			wantCurrenc: "ETH",
		},
		{
			name:        "six decimal token",
			token:       usdcAddress,
			amount:      "2000000",
			refundShare: "2500",
			wantAmount:  "0.5",
			wantCurrenc: "USDC",
		},
		{
			name:        "a share that does not divide evenly truncates",
			token:       usdcAddress,
			amount:      "1",
			refundShare: "3333",
			wantAmount:  "0",
			wantCurrenc: "USDC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := testClient()

			var share *big.Int
			if tt.refundShare != "" {
				share = wei(t, tt.refundShare)
			}

			raw, err := client.buildRefundCallbackPayload(tt.token, wei(t, tt.amount), share, txURL, 1700000000)
			if err != nil {
				t.Fatalf("buildRefundCallbackPayload returned %v", err)
			}

			var payload refundCallbackPayload
			if err := json.Unmarshal(raw, &payload); err != nil {
				t.Fatalf("payload is not valid JSON: %v", err)
			}
			if payload.Amount != tt.wantAmount {
				t.Errorf("amount = %q, want %q", payload.Amount, tt.wantAmount)
			}
			if payload.Currency != tt.wantCurrenc {
				t.Errorf("currency = %q, want %q", payload.Currency, tt.wantCurrenc)
			}
			if payload.TransactionUrl != txURL {
				t.Errorf("transactionUrl = %q, want %q", payload.TransactionUrl, txURL)
			}
			if payload.TransactionTimestamp != 1700000000 {
				t.Errorf("transactionTimestamp = %d, want 1700000000", payload.TransactionTimestamp)
			}
		})
	}
}

func TestBuildRefundCallbackPayloadDoesNotMutateTheAmount(t *testing.T) {
	client := testClient()
	amount := wei(t, "1000000000000000000")
	before := new(big.Int).Set(amount)

	if _, err := client.buildRefundCallbackPayload(ethAddress, amount, big.NewInt(5000), txURL, 0); err != nil {
		t.Fatalf("buildRefundCallbackPayload returned %v", err)
	}

	if amount.Cmp(before) != 0 {
		t.Errorf("the amount was mutated: got %s, want %s", amount, before)
	}
}

func TestBuildReleaseCallbackPayload(t *testing.T) {
	client := testClient()
	const receiver = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	raw, err := client.buildReleaseCallbackPayload(usdcAddress, receiver, wei(t, "2500000"), txURL, 1700000000)
	if err != nil {
		t.Fatalf("buildReleaseCallbackPayload returned %v", err)
	}

	var payload releaseCallbackPayload
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("payload is not valid JSON: %v", err)
	}

	if payload.Currency != "USDC" {
		t.Errorf("currency = %q, want USDC", payload.Currency)
	}
	if payload.Amount != "2.5" {
		t.Errorf("amount = %q, want 2.5", payload.Amount)
	}
	if payload.Address != receiver {
		t.Errorf("address = %q, want %q", payload.Address, receiver)
	}
	if payload.TransactionUrl != txURL {
		t.Errorf("transactionUrl = %q, want %q", payload.TransactionUrl, txURL)
	}
	if payload.TransactionTimestamp != 1700000000 {
		t.Errorf("transactionTimestamp = %d, want 1700000000", payload.TransactionTimestamp)
	}
}

func TestBuildersRejectANilAmount(t *testing.T) {
	client := testClient()

	builders := map[string]func() ([]byte, error){
		"paymentReceived": func() ([]byte, error) {
			return client.buildPaymentReceivedCallbackPayload(txURL, ethAddress, nil, 0)
		},
		"refund": func() ([]byte, error) {
			return client.buildRefundCallbackPayload(ethAddress, nil, big.NewInt(10000), txURL, 0)
		},
		"release": func() ([]byte, error) {
			return client.buildReleaseCallbackPayload(ethAddress, "0xabc", nil, txURL, 0)
		},
	}

	for name, build := range builders {
		t.Run(name, func(t *testing.T) {
			got, err := build()

			if err == nil {
				t.Fatalf("builder returned %s, want an error", got)
			}
			if !strings.Contains(err.Error(), "invalid amount") {
				t.Errorf("error = %q, want it to mention an invalid amount", err)
			}
			if got != nil {
				t.Errorf("builder returned %s alongside an error, want nil", got)
			}
		})
	}
}

func TestBuildersRejectAnUnknownToken(t *testing.T) {
	client := testClient()
	const unknown = "0x1111111111111111111111111111111111111111"

	builders := map[string]func() ([]byte, error){
		"paymentReceived": func() ([]byte, error) {
			return client.buildPaymentReceivedCallbackPayload(txURL, unknown, big.NewInt(1), 0)
		},
		"refund": func() ([]byte, error) {
			return client.buildRefundCallbackPayload(unknown, big.NewInt(1), nil, txURL, 0)
		},
		"release": func() ([]byte, error) {
			return client.buildReleaseCallbackPayload(unknown, "0xabc", big.NewInt(1), txURL, 0)
		},
	}

	for name, build := range builders {
		t.Run(name, func(t *testing.T) {
			got, err := build()

			if err == nil {
				t.Fatalf("builder returned %s, want an error", got)
			}
			if !strings.Contains(err.Error(), "unsupported payment token") {
				t.Errorf("error = %q, want it to mention an unsupported token", err)
			}
			if !strings.Contains(err.Error(), unknown) {
				t.Errorf("error = %q, want it to name the token address", err)
			}
		})
	}
}

func TestPayloadAmountsArePlainDecimalStrings(t *testing.T) {
	client := testClient()

	raw, err := client.buildReleaseCallbackPayload(ethAddress, "0xabc", wei(t, "1"), txURL, 0)
	if err != nil {
		t.Fatalf("buildReleaseCallbackPayload returned %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("payload is not valid JSON: %v", err)
	}

	amount, ok := payload["amount"].(string)
	if !ok {
		t.Fatalf("amount is %T, want a string", payload["amount"])
	}
	if strings.ContainsAny(amount, "eE") {
		t.Errorf("amount = %q, want plain decimal notation", amount)
	}
	if amount != "0.000000000000000001" {
		t.Errorf("amount = %q, want the smallest unit spelled out", amount)
	}
}
