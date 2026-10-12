package config

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func validateSignerKey(key string) error {
	trimmed := strings.TrimPrefix(strings.TrimSpace(key), "0x")
	if trimmed == "" {
		return fmt.Errorf("signerKey %w", ErrRequired)
	}
	if len(trimmed) != 64 {
		return fmt.Errorf("signerKey %w, got %d characters", ErrKeyLength, len(trimmed))
	}
	if _, err := hex.DecodeString(trimmed); err != nil {
		return fmt.Errorf("signerKey %w", ErrInvalidHex)
	}
	return nil
}
