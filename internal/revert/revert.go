// Package revert maps contract revert selectors to human-readable reasons.
package revert

import (
	"encoding/hex"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Descriptions = map[string]string{
	"0xb12e2421": "The buyer and seller addresses cannot be the same.",
	"0xf4d678b8": "The account balance is insufficient to perform this action.",
	"0x34819f90": "The provided dispute resolution is invalid.",
	"0x487e4409": "The invoice is not in a valid state for this action.",
	"0x214510aa": "The native token payment is invalid for this invoice.",
	"0x56e7ec5f": "The specified payment token is not supported or invalid.",
	"0x453fb42d": "The seller's payout share is invalid.",
	"0x074bc935": "An invoice with this identifier already exists.",
	"0x715d9228": "The specified invoice does not exist.",
	"0xb09960c1": "A meta-invoice with this identifier already exists.",
	"0xea8e4eb5": "The caller is not authorized to perform this action.",
	"0x2c669f0a": "The price cannot be zero.",
	"0xdb8db569": "The price specified is too low.",

	// OracleManager
	UnsupportedToken: "The oracle has no price feed for this token.",
	"0x00bfc921":     "The oracle reported an invalid price.",
	"0x032b3d00":     "The sequencer is down; prices are unavailable.",
	"0x19abf40e":     "The oracle price is stale.",
	"0x1087e109":     "The oracle price feed is stale.",
}

// UnsupportedToken is OracleManager's UnsupportedToken() selector, which
// getUsdPerToken reverts with for a token it has no feed for.
const UnsupportedToken = "0x6a172882"

var StatusCodes = map[string]int{
	"The buyer and seller addresses cannot be the same.":          http.StatusBadRequest,
	"The account balance is insufficient to perform this action.": http.StatusBadRequest,
	"The provided dispute resolution is invalid.":                 http.StatusBadRequest,
	"The invoice is not in a valid state for this action.":        http.StatusConflict,
	"The native token payment is invalid for this invoice.":       http.StatusBadRequest,
	"The specified payment token is not supported or invalid.":    http.StatusBadRequest,
	"The seller's payout share is invalid.":                       http.StatusBadRequest,
	"An invoice with this identifier already exists.":             http.StatusConflict,
	"The specified invoice does not exist.":                       http.StatusNotFound,
	"A meta-invoice with this identifier already exists.":         http.StatusConflict,
	"The caller is not authorized to perform this action.":        http.StatusForbidden,
	"The price cannot be zero.":                                   http.StatusBadRequest,
	"The price specified is too low.":                             http.StatusBadRequest,
}

// Selector returns the four-byte custom error selector a call reverted with,
// or "" when the error is not a contract revert.
func Selector(err error) string {
	if err == nil {
		return ""
	}
	if data, ok := ethclient.RevertErrorData(err); ok && len(data) >= 4 {
		return "0x" + hex.EncodeToString(data[:4])
	}
	return ""
}

// Is reports whether err is the given custom error selector.
func Is(err error, selector string) bool {
	return Selector(err) == selector
}

// Reason maps a contract revert to its human-readable description, falling
// back to the error's own message.
func Reason(err error) string {
	if err == nil {
		return ""
	}
	if reason, ok := Descriptions[Selector(err)]; ok {
		return reason
	}
	return err.Error()
}
