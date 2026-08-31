package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// MaxNoteLength mirrors MAX_NOTE_LENGTH on the website so both ends reject the
// same content.
const MaxNoteLength = 20

var (
	ErrNoteContentRequired = errors.New("Note content is required")
	ErrNoteTooLong         = fmt.Errorf("Notes are limited to %d characters", MaxNoteLength)
)

func ValidateNoteContent(content string) error {
	if content == "" {
		return ErrNoteContentRequired
	}

	if len([]rune(content)) > MaxNoteLength {
		return ErrNoteTooLong
	}

	return nil
}

func WriteNoteError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"error":   message,
	})
}

func WriteNoteSuccess(w http.ResponseWriter, body map[string]any) {
	body["success"] = true
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func ParseAddress(value string) (common.Address, bool) {
	trimmed := strings.TrimSpace(value)
	if !common.IsHexAddress(trimmed) {
		return common.Address{}, false
	}

	address := common.HexToAddress(trimmed)
	if address == (common.Address{}) {
		return common.Address{}, false
	}

	return address, true
}
