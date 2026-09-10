package httpx

import (
	"encoding/json"
	"net/http"
)

// WriteFailure writes {"success": false, "error": message}.
func WriteFailure(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"error":   message,
	})
}

func WriteSuccess(w http.ResponseWriter, body map[string]any) {
	body["success"] = true
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
