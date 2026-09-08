package note

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateNoteContent(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr error
	}{
		{"empty", "", ErrNoteContentRequired},
		{"single character", "a", nil},
		{"at the limit", strings.Repeat("a", MaxNoteLength), nil},
		{"one over the limit", strings.Repeat("a", MaxNoteLength+1), ErrNoteTooLong},
		{"well over the limit", strings.Repeat("a", MaxNoteLength*10), ErrNoteTooLong},

		{"whitespace", " ", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNoteContent(tt.content)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ValidateNoteContent(%q) = %v, want %v", tt.content, err, tt.wantErr)
			}
		})
	}
}

func TestValidateNoteContentCountsRunes(t *testing.T) {
	atLimit := strings.Repeat("☃", MaxNoteLength)
	overLimit := strings.Repeat("☃", MaxNoteLength+1)

	if len(atLimit) <= MaxNoteLength {
		t.Fatalf("test string is %d bytes, which does not exercise the rune count", len(atLimit))
	}
	if err := ValidateNoteContent(atLimit); err != nil {
		t.Errorf("ValidateNoteContent(%d runes) = %v, want nil", MaxNoteLength, err)
	}
	if err := ValidateNoteContent(overLimit); !errors.Is(err, ErrNoteTooLong) {
		t.Errorf("ValidateNoteContent(%d runes) = %v, want ErrNoteTooLong", MaxNoteLength+1, err)
	}
}

func TestNoteErrorsAreDistinct(t *testing.T) {
	if errors.Is(ErrNoteContentRequired, ErrNoteTooLong) {
		t.Error("ErrNoteContentRequired and ErrNoteTooLong are indistinguishable")
	}
	if !strings.Contains(ErrNoteTooLong.Error(), "20") {
		t.Errorf("ErrNoteTooLong = %q, want it to state the limit", ErrNoteTooLong)
	}
}
