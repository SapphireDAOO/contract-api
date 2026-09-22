package httpx

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// WriteFailure writes {"success": false, "error": message}.
func WriteFailure(w http.ResponseWriter, statusCode int, message string) {
	logErrorResponse(statusCode, message, message, nil)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"error":   message,
	}); err != nil {
		slog.Error("writing the failure envelope failed", "status", statusCode, "error", err)
	}
}

func WriteSuccess(w http.ResponseWriter, body map[string]any) {
	body["success"] = true
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("writing the success envelope failed", "error", err)
	}
}
