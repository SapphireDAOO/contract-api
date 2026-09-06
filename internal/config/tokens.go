package config

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Token is a payment token, named by its symbol in requests and resolved to
// the address deployed on the selected network.
type Token struct {
	Address  string `yaml:"address"`
	Decimals int    `yaml:"decimals"`
}

// Tokens maps a symbol to its token. Lookups are case-insensitive, so a caller
// may send "usdc" for a token configured as "USDC".
type Tokens map[string]Token

// Address resolves a symbol to its deployed address.
func (t Tokens) Address(symbol string) (common.Address, bool) {
	token, ok := t.lookup(symbol)
	if !ok {
		return common.Address{}, false
	}
	return common.HexToAddress(token.Address), true
}

// ByAddress resolves a deployed address back to its symbol and decimals, for
// rendering an amount that arrived in a chain event.
func (t Tokens) ByAddress(address string) (string, int, bool) {
	want := common.HexToAddress(address)
	for symbol, token := range t {
		if common.HexToAddress(token.Address) == want {
			return symbol, token.Decimals, true
		}
	}
	return "", 0, false
}

// Symbols lists the configured symbols, sorted, for error messages.
func (t Tokens) Symbols() []string {
	return slices.Sorted(maps.Keys(t))
}

func (t Tokens) lookup(symbol string) (Token, bool) {
	if token, ok := t[symbol]; ok {
		return token, true
	}
	for configured, token := range t {
		if strings.EqualFold(configured, symbol) {
			return token, true
		}
	}
	return Token{}, false
}

// validate checks the token table. The zero address is allowed: it is how the
// contracts denote the native token.
func (t Tokens) validate() error {
	if len(t) == 0 {
		return fmt.Errorf("tokens: at least one payment token is required")
	}
	for symbol, token := range t {
		if strings.TrimSpace(symbol) == "" {
			return fmt.Errorf("tokens: a symbol cannot be empty")
		}
		if !common.IsHexAddress(token.Address) {
			return fmt.Errorf("tokens.%s.address %q is not a valid address", symbol, token.Address)
		}
		if token.Decimals < 0 || token.Decimals > 77 {
			return fmt.Errorf("tokens.%s.decimals %d is out of range", symbol, token.Decimals)
		}
	}
	return nil
}
