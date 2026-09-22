package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"log/slog"
	"net/http"
	"os"
)

func AccessControlMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := os.Getenv("KEY")
		if key == "" {
			slog.Error("rejecting request", "reason", "KEY env var not set",
				"method", r.Method, "path", r.URL.Path)
			http.Error(w, "Server misconfigured", http.StatusInternalServerError)
			return
		}

		providedKey := r.Header.Get("X-API-KEY")
		if providedKey == "" {
			slog.Warn("request rejected", "reason", "X-API-KEY header missing",
				"method", r.Method, "path", r.URL.Path)
			http.Error(w, "API key missing", http.StatusUnauthorized)
			return
		}

		hash := sha256.Sum256([]byte(providedKey))
		hashedKey := hex.EncodeToString(hash[:])

		if subtle.ConstantTimeCompare([]byte(hashedKey), []byte(key)) != 1 {
			// The key itself is never logged, only that one was wrong.
			slog.Warn("request rejected", "reason", "X-API-KEY did not match",
				"method", r.Method, "path", r.URL.Path)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}
