package invoice

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

// MetaPrefix marks a link as pointing at a meta invoice rather than a single
// one, so the two identifiers cannot be confused.
const MetaPrefix = "mt-"

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

func EncodeMetaID(metaInvoiceID *big.Int) string {
	encoded := EncodeID(metaInvoiceID)
	if encoded == "" {
		return ""
	}

	return MetaPrefix + encoded
}

func EncodeMetaIDString(metaInvoiceID string) string {
	encoded := EncodeIDString(metaInvoiceID)
	if encoded == "" {
		return ""
	}

	return MetaPrefix + encoded
}

func IsMetaID(encoded string) bool {
	return strings.HasPrefix(strings.TrimSpace(encoded), MetaPrefix)
}

// DecodeID reads either form, so a caller holding a link does not have to
// strip the prefix first. Use IsMetaID to tell them apart.
func DecodeID(encoded string) (*big.Int, error) {
	trimmed := strings.TrimPrefix(strings.TrimSpace(encoded), MetaPrefix)
	trimmed = strings.TrimRight(trimmed, "=")
	if trimmed == "" {
		return nil, fmt.Errorf("invoice id is empty")
	}

	b, err := base64.RawURLEncoding.DecodeString(trimmed)
	if err != nil {
		return nil, fmt.Errorf("invoice id is not valid base64url: %w", err)
	}

	return new(big.Int).SetBytes(b), nil
}
