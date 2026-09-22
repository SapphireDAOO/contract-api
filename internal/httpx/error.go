// Package httpx writes the API's JSON responses.
package httpx

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/SapphireDAOO/contract-api/internal/revert"
)

// logErrorResponse records every error the API answers with, so an operator
// sees the failure even though the caller was told about it politely. A 5xx is
// our fault and logs as an error; a 4xx is the caller's and logs as a warning.
func logErrorResponse(statusCode int, msg, reason string, err error) {
	level := slog.LevelWarn
	if statusCode >= http.StatusInternalServerError {
		level = slog.LevelError
	}

	attrs := []any{"status", statusCode, "response", msg, "reason", reason}
	// The mapped reason is usually the revert description rather than the
	// error itself, so the original is kept when it adds anything.
	if err != nil && err.Error() != reason {
		attrs = append(attrs, "error", err)
	}

	slog.Log(context.Background(), level, "request failed", attrs...)
}

func writeJSONError(w http.ResponseWriter, statusCode int, msg, reason string, err error) {
	logErrorResponse(statusCode, msg, reason, err)

	body, err := json.Marshal(map[string]string{
		"error":  msg,
		"reason": reason,
	})
	if err != nil {
		slog.Error("marshalling the error response failed", "error", err)
		http.Error(w, msg, statusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(body); err != nil {
		slog.Error("writing the error response failed", "status", statusCode, "error", err)
	}
}

func WriteHTTPErrorWithStatus(w http.ResponseWriter, statusCode int, err error, msg string) {
	writeJSONError(w, statusCode, msg, revert.Reason(err), err)
}

func WriteMappedRevertError(w http.ResponseWriter, err error, msg string) {
	reason := revert.Reason(err)
	statusCode := revert.StatusCodes[reason]
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	writeJSONError(w, statusCode, msg, reason, err)
}
