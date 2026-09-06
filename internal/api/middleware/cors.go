package middleware

import "net/http"

const (
	allowedMethods  = "POST, OPTIONS"
	allowedHeaders  = "Content-Type"
	preflightMaxAge = "86400"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Add("Vary", "Origin")

		next(w, r)
	}
}

func Preflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Max-Age", preflightMaxAge)
	w.Header().Add("Vary", "Origin")

	w.WriteHeader(http.StatusNoContent)
}
