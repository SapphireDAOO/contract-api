package middleware

import (
	"net/http"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/auth"
	"github.com/SapphireDAOO/contract-api/internal/httpx"
)

func RequireSession(sessions *auth.Service, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if sessions == nil {
			httpx.WriteFailure(w, http.StatusServiceUnavailable, auth.ErrSignInNotConfigured.Error())
			return
		}

		token, ok := bearerToken(r)
		if !ok {
			w.Header().Set("WWW-Authenticate", "Bearer")
			httpx.WriteFailure(w, http.StatusUnauthorized, auth.ErrSignInRequired.Error())
			return
		}

		address, err := sessions.SessionAddress(token)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Bearer error="invalid_token"`)
			httpx.WriteFailure(w, http.StatusUnauthorized, auth.ErrInvalidSession.Error())
			return
		}

		next(w, r.WithContext(auth.WithAddress(r.Context(), address)))
	}
}

func bearerToken(r *http.Request) (string, bool) {
	scheme, token, ok := strings.Cut(r.Header.Get("Authorization"), " ")
	token = strings.TrimSpace(token)
	if !ok || !strings.EqualFold(scheme, "Bearer") || token == "" {
		return "", false
	}
	return token, true
}
