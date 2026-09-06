package config

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func validateSignerKey(key string) error {
	trimmed := strings.TrimPrefix(strings.TrimSpace(key), "0x")
	if trimmed == "" {
		return fmt.Errorf("signerKey is required")
	}
	if len(trimmed) != 64 {
		return fmt.Errorf("signerKey must be 32 bytes of hex, got %d characters", len(trimmed))
	}
	if _, err := hex.DecodeString(trimmed); err != nil {
		return fmt.Errorf("signerKey is not valid hex")
	}
	return nil
}
