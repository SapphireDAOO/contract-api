package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
)

func AccessControlMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := os.Getenv("KEY")
		providedKey := r.Header.Get("X-API-KEY")

		if providedKey == "" {
			http.Error(w, "API key missing", http.StatusUnauthorized)
			return
		}

		hash := sha256.Sum256([]byte(providedKey))
		hashedKey := hex.EncodeToString(hash[:])

		if hashedKey != key {
			http.Error(w, "Forbbiden", http.StatusUnauthorized)
			return
		}

		next(w, r)
	})
}
