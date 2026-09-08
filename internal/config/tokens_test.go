package config

import (
	"slices"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

const (
	usdcAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	daiAddress  = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	zeroAddress = "0x0000000000000000000000000000000000000000"
)

func testTokens() Tokens {
	return Tokens{
		"ETH":  {Address: zeroAddress, Decimals: 18},
		"USDC": {Address: usdcAddress, Decimals: 6},
	}
}

func TestTokensAddress(t *testing.T) {
	tokens := testTokens()

	tests := []struct {
		name   string
		symbol string
		want   string
		wantOK bool
	}{
		{"exact match", "USDC", usdcAddress, true},
		{"lowercase", "usdc", usdcAddress, true},
		{"mixed case", "UsDc", usdcAddress, true},
		{"native token resolves to the zero address", "ETH", zeroAddress, true},
		{"unknown symbol", "DOGE", "", false},
		{"empty symbol", "", "", false},
		{"whitespace is not trimmed here", " USDC", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tokens.Address(tt.symbol)

			if ok != tt.wantOK {
				t.Fatalf("Address(%q) ok = %v, want %v", tt.symbol, ok, tt.wantOK)
			}
			if !ok {
				if got != (common.Address{}) {
					t.Errorf("Address(%q) = %s on a miss, want the zero address", tt.symbol, got)
				}
				return
			}
			if got != common.HexToAddress(tt.want) {
				t.Errorf("Address(%q) = %s, want %s", tt.symbol, got, tt.want)
			}
		})
	}
}

func TestTokensByAddress(t *testing.T) {
	tokens := testTokens()

	tests := []struct {
		name         string
		address      string
		wantSymbol   string
		wantDecimals int
		wantOK       bool
	}{
		{"checksummed", usdcAddress, "USDC", 6, true},
		{"lowercase", strings.ToLower(usdcAddress), "USDC", 6, true},
		{"uppercase hex", "0x" + strings.ToUpper(usdcAddress[2:]), "USDC", 6, true},
		{"the zero address is the native token", zeroAddress, "ETH", 18, true},
		{"unconfigured address", daiAddress, "", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			symbol, decimals, ok := tokens.ByAddress(tt.address)

			if ok != tt.wantOK {
				t.Fatalf("ByAddress(%q) ok = %v, want %v", tt.address, ok, tt.wantOK)
			}
			if symbol != tt.wantSymbol {
				t.Errorf("symbol = %q, want %q", symbol, tt.wantSymbol)
			}
			if decimals != tt.wantDecimals {
				t.Errorf("decimals = %d, want %d", decimals, tt.wantDecimals)
			}
		})
	}
}

func TestTokensAddressRoundTrip(t *testing.T) {
	tokens := testTokens()

	for symbol := range tokens {
		address, ok := tokens.Address(symbol)
		if !ok {
			t.Fatalf("Address(%q) missed a configured symbol", symbol)
		}

		got, _, ok := tokens.ByAddress(address.Hex())
		if !ok {
			t.Fatalf("ByAddress(%s) missed the address Address(%q) returned", address, symbol)
		}
		if got != symbol {
			t.Errorf("round trip of %q produced %q", symbol, got)
		}
	}
}

func TestTokensSymbols(t *testing.T) {
	tokens := Tokens{
		"USDC": {Address: usdcAddress, Decimals: 6},
		"ETH":  {Address: zeroAddress, Decimals: 18},
		"DAI":  {Address: daiAddress, Decimals: 18},
	}

	got := tokens.Symbols()

	want := []string{"DAI", "ETH", "USDC"}
	if !slices.Equal(got, want) {
		t.Errorf("Symbols() = %v, want %v", got, want)
	}

	if !slices.IsSorted(got) {
		t.Errorf("Symbols() = %v, want it sorted", got)
	}
}

func TestTokensSymbolsEmpty(t *testing.T) {
	if got := (Tokens{}).Symbols(); len(got) != 0 {
		t.Errorf("Symbols() = %v, want no symbols", got)
	}
}

func TestTokensValidate(t *testing.T) {
	tests := []struct {
		name    string
		tokens  Tokens
		wantErr string
	}{
		{name: "valid", tokens: testTokens()},
		{
			name:    "empty table",
			tokens:  Tokens{},
			wantErr: "at least one payment token is required",
		},
		{
			name:    "nil table",
			tokens:  nil,
			wantErr: "at least one payment token is required",
		},
		{
			name:    "empty symbol",
			tokens:  Tokens{"": {Address: usdcAddress, Decimals: 6}},
			wantErr: "a symbol cannot be empty",
		},
		{
			name:    "whitespace symbol",
			tokens:  Tokens{"  ": {Address: usdcAddress, Decimals: 6}},
			wantErr: "a symbol cannot be empty",
		},
		{
			name:    "missing address",
			tokens:  Tokens{"USDC": {Decimals: 6}},
			wantErr: "is not a valid address",
		},
		{
			name:    "malformed address",
			tokens:  Tokens{"USDC": {Address: "0x123", Decimals: 6}},
			wantErr: "is not a valid address",
		},
		{
			name:    "negative decimals",
			tokens:  Tokens{"USDC": {Address: usdcAddress, Decimals: -1}},
			wantErr: "is out of range",
		},
		{
			name:    "decimals above the limit",
			tokens:  Tokens{"USDC": {Address: usdcAddress, Decimals: 78}},
			wantErr: "is out of range",
		},
		{
			name:   "decimals at the limit",
			tokens: Tokens{"USDC": {Address: usdcAddress, Decimals: 77}},
		},
		{
			name:   "zero decimals",
			tokens: Tokens{"USDC": {Address: usdcAddress, Decimals: 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tokens.validate()

			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("validate() returned %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("validate() returned nil, want an error containing %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestTokensValidateAllowsTheZeroAddress(t *testing.T) {
	tokens := Tokens{"ETH": {Address: zeroAddress, Decimals: 18}}

	if err := tokens.validate(); err != nil {
		t.Errorf("validate() rejected the native token: %v", err)
	}
}

func TestTokensValidateNamesTheSymbol(t *testing.T) {
	tokens := Tokens{
		"ETH":  {Address: zeroAddress, Decimals: 18},
		"DOGE": {Address: "not-an-address", Decimals: 8},
	}

	err := tokens.validate()

	if err == nil {
		t.Fatal("validate() returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "DOGE") {
		t.Errorf("error = %q, want it to name the DOGE entry", err)
	}
}
