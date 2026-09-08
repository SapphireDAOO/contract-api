package feereceiver

import (
	"strings"
	"testing"

	"github.com/SapphireDAOO/contract-api/internal/config"
	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
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

func TestParseProcessorKind(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    pb.ProcessorKind
		wantErr bool
	}{
		{name: "simple", value: "simple", want: pb.ProcessorKind_PROCESSOR_KIND_SIMPLE},
		{name: "intermediated", value: "intermediated", want: pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED},
		{name: "uppercase", value: "SIMPLE", want: pb.ProcessorKind_PROCESSOR_KIND_SIMPLE},
		{name: "mixed case", value: "Intermediated", want: pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED},
		{name: "surrounding whitespace", value: "  simple  ", want: pb.ProcessorKind_PROCESSOR_KIND_SIMPLE},

		{name: "empty is rejected", value: "", wantErr: true},
		{name: "whitespace only is rejected", value: "   ", wantErr: true},
		{name: "unknown value", value: "escrow", wantErr: true},
		{name: "near miss", value: "simpl", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseProcessorKind(tt.value)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseProcessorKind(%q) = %v, want an error", tt.value, got)
				}
				if got != pb.ProcessorKind_PROCESSOR_KIND_UNSPECIFIED {
					t.Errorf("ParseProcessorKind(%q) = %v alongside an error, want UNSPECIFIED", tt.value, got)
				}
				if !strings.Contains(err.Error(), "simple or intermediated") {
					t.Errorf("error = %q, want it to list the accepted values", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseProcessorKind(%q) returned %v", tt.value, err)
			}
			if got != tt.want {
				t.Errorf("ParseProcessorKind(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestParseProcessorKindErrorQuotesTheValue(t *testing.T) {
	_, err := ParseProcessorKind("escrow")

	if err == nil {
		t.Fatal("ParseProcessorKind returned no error")
	}
	if !strings.Contains(err.Error(), `"escrow"`) {
		t.Errorf("error = %q, want it to quote the value sent", err)
	}
}

func TestParseInvoiceKind(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    pb.InvoiceKind
		wantErr bool
	}{
		{name: "single", value: "single", want: pb.InvoiceKind_INVOICE_KIND_SINGLE},
		{name: "meta", value: "meta", want: pb.InvoiceKind_INVOICE_KIND_META},

		{name: "empty defaults to single", value: "", want: pb.InvoiceKind_INVOICE_KIND_SINGLE},
		{name: "whitespace defaults to single", value: "   ", want: pb.InvoiceKind_INVOICE_KIND_SINGLE},
		{name: "uppercase", value: "META", want: pb.InvoiceKind_INVOICE_KIND_META},
		{name: "surrounding whitespace", value: "  meta  ", want: pb.InvoiceKind_INVOICE_KIND_META},
		{name: "unknown value", value: "batch", wantErr: true},
		{name: "near miss", value: "metas", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInvoiceKind(tt.value)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseInvoiceKind(%q) = %v, want an error", tt.value, got)
				}
				if got != pb.InvoiceKind_INVOICE_KIND_UNSPECIFIED {
					t.Errorf("ParseInvoiceKind(%q) = %v alongside an error, want UNSPECIFIED", tt.value, got)
				}
				if !strings.Contains(err.Error(), "single or meta") {
					t.Errorf("error = %q, want it to list the accepted values", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseInvoiceKind(%q) returned %v", tt.value, err)
			}
			if got != tt.want {
				t.Errorf("ParseInvoiceKind(%q) = %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}

func TestResolveFeeToken(t *testing.T) {
	tests := []struct {
		name    string
		symbol  string
		want    string
		wantErr bool
	}{
		{name: "empty is passed through", symbol: "", want: ""},
		{name: "whitespace is passed through", symbol: "   ", want: ""},
		{name: "known symbol", symbol: "USDC", want: usdcAddress},
		{name: "lowercase symbol", symbol: "usdc", want: usdcAddress},
		{name: "surrounding whitespace", symbol: "  USDC  ", want: usdcAddress},
		{name: "the native token", symbol: "ETH", want: zeroAddress},
		{name: "unknown symbol", symbol: "DOGE", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveFeeToken(testTokens(), tt.symbol)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("ResolveFeeToken(%q) = %q, want an error", tt.symbol, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ResolveFeeToken(%q) returned %v", tt.symbol, err)
			}
			if got != tt.want {
				t.Errorf("ResolveFeeToken(%q) = %q, want %q", tt.symbol, got, tt.want)
			}
		})
	}
}

func TestResolveFeeTokenReturnsAChecksummedAddress(t *testing.T) {
	got, err := ResolveFeeToken(testTokens(), "usdc")
	if err != nil {
		t.Fatalf("ResolveFeeToken returned %v", err)
	}

	if got != usdcAddress {
		t.Errorf("ResolveFeeToken = %q, want the checksummed %q", got, usdcAddress)
	}
}

func TestResolveFeeTokenUnknownSymbolListsKnown(t *testing.T) {
	_, err := ResolveFeeToken(testTokens(), "DOGE")

	if err == nil {
		t.Fatal("ResolveFeeToken returned no error")
	}
	if !strings.Contains(err.Error(), "DOGE") {
		t.Errorf("error = %q, want it to name the unknown symbol", err)
	}
	if !strings.Contains(err.Error(), "ETH, USDC") {
		t.Errorf("error = %q, want it to list the configured symbols", err)
	}
}

func TestResolveFeeTokenWithNoTokensConfigured(t *testing.T) {
	got, err := ResolveFeeToken(config.Tokens{}, "")
	if err != nil {
		t.Fatalf("ResolveFeeToken returned %v", err)
	}
	if got != "" {
		t.Errorf("ResolveFeeToken = %q, want an empty string", got)
	}

	if _, err := ResolveFeeToken(config.Tokens{}, "USDC"); err == nil {
		t.Error("ResolveFeeToken with no tokens configured returned no error for a named symbol")
	}
}
