package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"log"
	"net/http"
	"os"
)

func AccessControlMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := os.Getenv("KEY")
		if key == "" {
			log.Println("KEY env var not set; rejecting request")
			http.Error(w, "Server misconfigured", http.StatusInternalServerError)
			return
		}

		providedKey := r.Header.Get("X-API-KEY")
		if providedKey == "" {
			http.Error(w, "API key missing", http.StatusUnauthorized)
			return
		}

		hash := sha256.Sum256([]byte(providedKey))
		hashedKey := hex.EncodeToString(hash[:])

		if subtle.ConstantTimeCompare([]byte(hashedKey), []byte(key)) != 1 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}
