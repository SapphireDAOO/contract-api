// Package note validates and encrypts invoice notes.
package note

import (
	"errors"
	"fmt"
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
