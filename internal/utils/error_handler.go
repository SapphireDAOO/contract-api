package utils

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
)

var RevertErrorDescriptions = map[string]string{
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
}

var RevertErrorStatusCodes = map[string]int{
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

func WriteHTTPErrorWithStatus(w http.ResponseWriter, statusCode int, err error, msg string) {
	reason := Reason(err)

	errorMessage := fmt.Sprintf(`{"error": "%s", "reason": "%s"}`, msg, strings.ReplaceAll(reason, `"`, `'`))
	http.Error(w, errorMessage, statusCode)
}

func WriteMappedRevertError(w http.ResponseWriter, err error, msg string) {
	reason := Reason(err)
	statusCode := RevertErrorStatusCodes[reason]
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	errorMessage := fmt.Sprintf(`{"error": "%s", "reason": "%s"}`, msg, strings.ReplaceAll(reason, `"`, `'`))
	http.Error(w, errorMessage, statusCode)
}

func Reason(err error) string {
	if data, ok := ethclient.RevertErrorData(err); ok {
		selector := data[:4]
		if reason, ok := RevertErrorDescriptions["0x"+hex.EncodeToString(selector)]; ok {
			return reason
		}
		return err.Error()
	}
	return "Call failed, but no revert data"
}
