package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/SapphireDAOO/contract-api/internal/auth"
	"github.com/SapphireDAOO/contract-api/internal/httpx"
	"github.com/SapphireDAOO/contract-api/internal/invoice"
	"github.com/SapphireDAOO/contract-api/internal/note"
)

const noteCallTimeout = 30 * time.Second

func (h *ContractHandler) WriteNote(w http.ResponseWriter, r *http.Request) {
	signedIn, ok := auth.AddressFrom(r.Context())
	if !ok {
		httpx.WriteFailure(w, http.StatusUnauthorized, auth.ErrSignInRequired.Error())
		return
	}

	var req struct {
		InvoiceId string `json:"invoiceId"`
		Author    string `json:"author"`
		Content   string `json:"content"`
		Share     bool   `json:"share"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteFailure(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	invoiceId, err := parseBigInt("invoiceId", req.InvoiceId)
	if err != nil {
		httpx.WriteFailure(w, http.StatusBadRequest, "invoiceId must be a base-10 integer")
		return
	}

	author, ok := signedInAuthor(w, req.Author, signedIn)
	if !ok {
		return
	}

	content, err := note.ParseContent(req.Content)
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, note.ErrContentTooLong) {
			status = http.StatusRequestEntityTooLarge
		}
		httpx.WriteFailure(w, status, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), noteCallTimeout)
	defer cancel()

	txHash, err := h.Notes.CreateNote(ctx, invoiceId, author, content, req.Share)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	httpx.WriteSuccess(w, map[string]any{"txHash": txHash.Hex()})
}

// signedInAuthor accepts a body "author" only when it is the signed-in
// account; leaving it out means the signed-in account.
func signedInAuthor(w http.ResponseWriter, raw string, signedIn common.Address) (common.Address, bool) {
	if strings.TrimSpace(raw) == "" {
		return signedIn, true
	}

	author, ok := invoice.ParseAddress(raw)
	if !ok {
		httpx.WriteFailure(w, http.StatusBadRequest, "Invalid author address")
		return common.Address{}, false
	}
	if author != signedIn {
		httpx.WriteFailure(w, http.StatusForbidden, "author is not the signed-in account")
		return common.Address{}, false
	}
	return author, true
}

// OpenNote records that the signed-in account has opened a note. Closing is a
// client-side state and is not written on chain.
func (h *ContractHandler) OpenNote(w http.ResponseWriter, r *http.Request) {
	signedIn, ok := auth.AddressFrom(r.Context())
	if !ok {
		httpx.WriteFailure(w, http.StatusUnauthorized, auth.ErrSignInRequired.Error())
		return
	}

	var req struct {
		InvoiceId string `json:"invoiceId"`
		Author    string `json:"author"`
		NoteId    string `json:"noteId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteFailure(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	invoiceId, err := parseBigInt("invoiceId", req.InvoiceId)
	if err != nil {
		httpx.WriteFailure(w, http.StatusBadRequest, "invoiceId must be a base-10 integer")
		return
	}

	noteId, err := parseBigInt("noteId", req.NoteId)
	if err != nil {
		httpx.WriteFailure(w, http.StatusBadRequest, "noteId must be a base-10 integer")
		return
	}

	author, ok := signedInAuthor(w, req.Author, signedIn)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), noteCallTimeout)
	defer cancel()

	txHash, err := h.Notes.SetOpened(ctx, invoiceId, author, noteId)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	httpx.WriteSuccess(w, map[string]any{"txHash": txHash.Hex()})
}
