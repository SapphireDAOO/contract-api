package invoice

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func EncodeID(invoiceID *big.Int) string {
	if invoiceID == nil || invoiceID.Sign() < 0 {
		return ""
	}

	b := invoiceID.Bytes()
	if len(b) == 0 {
		b = []byte{0}
	}

	return base64.RawURLEncoding.EncodeToString(b)
}

func EncodeIDString(invoiceID string) string {
	trimmed := strings.TrimSpace(invoiceID)
	if trimmed == "" {
		return ""
	}

	id, ok := new(big.Int).SetString(trimmed, 10)
	if !ok {
		return ""
	}

	return EncodeID(id)
}

func DecodeID(encoded string) (*big.Int, error) {
	trimmed := strings.TrimRight(strings.TrimSpace(encoded), "=")
	if trimmed == "" {
		return nil, fmt.Errorf("invoice id is empty")
	}

	b, err := base64.RawURLEncoding.DecodeString(trimmed)
	if err != nil {
		return nil, fmt.Errorf("invoice id is not valid base64url: %w", err)
	}

	return new(big.Int).SetBytes(b), nil
}
