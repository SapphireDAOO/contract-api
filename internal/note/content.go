package note

import (
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const MaxContentBytes = 4096

// ParseContent decodes the hex ciphertext a caller sends.
func ParseContent(value string) ([]byte, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return nil, ErrContentRequired
	}

	content, err := hexutil.Decode(trimmed)
	if err != nil {
		return nil, ErrContentHex
	}
	if len(content) == 0 {
		return nil, ErrContentRequired
	}
	if len(content) > MaxContentBytes {
		return nil, ErrContentTooLong
	}

	return content, nil
}
