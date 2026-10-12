package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/auth"
	"github.com/SapphireDAOO/contract-api/internal/httpx"
)

func (h *ContractHandler) AuthNonce(w http.ResponseWriter, r *http.Request) {
	if h.Auth == nil {
		httpx.WriteFailure(w, http.StatusServiceUnavailable, auth.ErrSignInNotConfigured.Error())
		return
	}

	nonce, expires, err := h.Auth.Nonce()
	if err != nil {
		slog.Error("issuing a sign-in nonce failed", "error", err)
		httpx.WriteFailure(w, http.StatusInternalServerError, "Could not issue a nonce")
		return
	}

	w.Header().Set("Cache-Control", "no-store")
	httpx.WriteSuccess(w, map[string]any{
		"nonce":     nonce,
		"expiresAt": expires.UTC().Format(time.RFC3339),
	})
}

func (h *ContractHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if h.Auth == nil {
		httpx.WriteFailure(w, http.StatusServiceUnavailable, auth.ErrSignInNotConfigured.Error())
		return
	}

	var req struct {
		Message   string `json:"message"`
		Signature string `json:"signature"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteFailure(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	session, err := h.Auth.SignIn(req.Message, req.Signature)
	if err != nil {
		status := http.StatusUnauthorized
		if errors.Is(err, auth.ErrInvalidMessage) {
			status = http.StatusBadRequest
		}
		httpx.WriteFailure(w, status, err.Error())
		return
	}

	w.Header().Set("Cache-Control", "no-store")
	httpx.WriteSuccess(w, map[string]any{
		"token":     session.Token,
		"address":   session.Address.Hex(),
		"expiresAt": session.ExpiresAt.UTC().Format(time.RFC3339),
	})
}
