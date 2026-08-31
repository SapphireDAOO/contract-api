package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const (
	maxDecryptBatch = 50

	noteCallTimeout = 30 * time.Second
)

type noteRequest struct {
	Action    string   `json:"action"`
	InvoiceId string   `json:"invoiceId"`
	NoteId    string   `json:"noteId"`
	NoteIds   []string `json:"noteIds"`
	Author    string   `json:"author"`
	Viewer    string   `json:"viewer"`
	Content   string   `json:"content"`
	Share     bool     `json:"share"`
	Open      bool     `json:"open"`
}

type noteResult struct {
	NoteId  string  `json:"noteId"`
	Content *string `json:"content"`
}

// HandleNote is the single entry point for note actions, dispatching on the
// "action" field so the website can forward a request body unchanged.
func (h *ContractHandler) HandleNote(w http.ResponseWriter, r *http.Request) {
	var req noteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteNoteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	switch req.Action {
	case "create":
		h.createNote(w, r, &req)
	case "setOpened":
		h.setNoteOpened(w, r, &req)
	case "encrypt":
		h.encryptNote(w, &req)
	case "decrypt":
		h.decryptNotes(w, r, &req)
	default:
		utils.WriteNoteError(w, http.StatusBadRequest, "Unknown action")
	}
}

func (h *ContractHandler) createNote(w http.ResponseWriter, r *http.Request, req *noteRequest) {
	invoiceId, err := parseBigInt("invoiceId", req.InvoiceId)
	if err != nil {
		utils.WriteNoteError(w, http.StatusBadRequest, "invoiceId must be a base-10 integer")
		return
	}

	author, ok := utils.ParseAddress(req.Author)
	if !ok {
		utils.WriteNoteError(w, http.StatusBadRequest, "Invalid author address")
		return
	}

	content := strings.TrimSpace(req.Content)

	if err := utils.ValidateNoteContent(content); err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, utils.ErrNoteTooLong) {
			status = http.StatusRequestEntityTooLarge
		}
		utils.WriteNoteError(w, status, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), noteCallTimeout)
	defer cancel()

	encrypted, err := utils.ToEncryptedNoteBytes(content)
	if err != nil {
		utils.WriteNoteError(w, http.StatusInternalServerError, "Failed to encrypt note")
		return
	}

	txHash, err := h.Notes.CreateNote(ctx, invoiceId, author, encrypted, req.Share)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	utils.WriteNoteSuccess(w, map[string]any{"txHash": txHash.Hex()})
}

func (h *ContractHandler) setNoteOpened(w http.ResponseWriter, r *http.Request, req *noteRequest) {
	invoiceId, err := parseBigInt("invoiceId", req.InvoiceId)
	if err != nil {
		utils.WriteNoteError(w, http.StatusBadRequest, "invoiceId must be a base-10 integer")
		return
	}

	noteId, err := parseBigInt("noteId", req.NoteId)
	if err != nil {
		utils.WriteNoteError(w, http.StatusBadRequest, "noteId must be a base-10 integer")
		return
	}

	author, ok := utils.ParseAddress(req.Author)
	if !ok {
		utils.WriteNoteError(w, http.StatusBadRequest, "Invalid author address")
		return
	}

	// Only opening is recorded on chain; closing is a client-side state.
	if !req.Open {
		utils.WriteNoteSuccess(w, map[string]any{})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), noteCallTimeout)
	defer cancel()

	txHash, err := h.Notes.SetOpened(ctx, invoiceId, author, noteId)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	utils.WriteNoteSuccess(w, map[string]any{"txHash": txHash.Hex()})
}

func (h *ContractHandler) encryptNote(w http.ResponseWriter, req *noteRequest) {
	content := strings.TrimSpace(req.Content)

	if err := utils.ValidateNoteContent(content); err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, utils.ErrNoteTooLong) {
			status = http.StatusRequestEntityTooLarge
		}
		utils.WriteNoteError(w, status, err.Error())
		return
	}

	encrypted, err := utils.ToEncryptedNoteBytes(content)
	if err != nil {
		utils.WriteNoteError(w, http.StatusInternalServerError, "Failed to encrypt note")
		return
	}

	utils.WriteNoteSuccess(w, map[string]any{"payload": hexutil.Encode(encrypted)})
}

func (h *ContractHandler) decryptNotes(w http.ResponseWriter, r *http.Request, req *noteRequest) {
	invoiceId, err := parseBigInt("invoiceId", req.InvoiceId)
	if err != nil {
		utils.WriteNoteError(w, http.StatusBadRequest, "invoiceId must be a base-10 integer")
		return
	}

	if len(req.NoteIds) == 0 {
		utils.WriteNoteError(w, http.StatusBadRequest, "noteIds is required")
		return
	}

	if len(req.NoteIds) > maxDecryptBatch {
		utils.WriteNoteError(w, http.StatusRequestEntityTooLarge, "Too many notes requested")
		return
	}

	noteIds := make([]*big.Int, 0, len(req.NoteIds))
	for _, raw := range req.NoteIds {
		noteId, err := parseBigInt("noteId", raw)
		if err != nil {
			utils.WriteNoteError(w, http.StatusBadRequest, "noteId must be a base-10 integer")
			return
		}
		noteIds = append(noteIds, noteId)
	}

	viewer, hasViewer := utils.ParseAddress(req.Viewer)

	ctx, cancel := context.WithTimeout(r.Context(), noteCallTimeout)
	defer cancel()

	notes := make([]noteResult, len(noteIds))
	var wg sync.WaitGroup

	for i, noteId := range noteIds {
		wg.Add(1)
		go func(i int, noteId *big.Int) {
			defer wg.Done()

			result := noteResult{NoteId: noteId.String()}
			notes[i] = result

			note, err := h.Notes.GetNote(ctx, invoiceId, noteId)
			if err != nil {
				return
			}

			if !note.Share && !(hasViewer && note.Author == viewer) {
				return
			}

			content, err := utils.DecryptNoteBlob(note.Content)
			if err != nil {
				log.Printf("failed to decrypt note %s on invoice %s: %v",
					noteId.String(), invoiceId.String(), err)
				return
			}

			result.Content = &content
			notes[i] = result
		}(i, noteId)
	}

	wg.Wait()

	utils.WriteNoteSuccess(w, map[string]any{"notes": notes})
}
