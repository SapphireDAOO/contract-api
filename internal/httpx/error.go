// Package httpx writes the API's JSON error responses.
package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/SapphireDAOO/contract-api/internal/revert"
)

func writeJSONError(w http.ResponseWriter, statusCode int, msg, reason string) {
	body, err := json.Marshal(map[string]string{
		"error":  msg,
		"reason": reason,
	})
	if err != nil {
		http.Error(w, msg, statusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func WriteHTTPErrorWithStatus(w http.ResponseWriter, statusCode int, err error, msg string) {
	writeJSONError(w, statusCode, msg, revert.Reason(err))
}

func WriteMappedRevertError(w http.ResponseWriter, err error, msg string) {
	reason := revert.Reason(err)
	statusCode := revert.StatusCodes[reason]
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	writeJSONError(w, statusCode, msg, reason)
}
