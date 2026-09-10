package revert

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Selector returns the four-byte custom error selector a call reverted with,
// or "" when the error is not a contract revert.
func Selector(err error) string {
	if err == nil {
		return ""
	}
	if data, ok := ethclient.RevertErrorData(err); ok && len(data) >= 4 {
		return "0x" + hex.EncodeToString(data[:4])
	}
	return ""
}

// Is reports whether err is the given custom error selector.
func Is(err error, selector string) bool {
	return Selector(err) == selector
}

// Reason maps a contract revert to its human-readable description, falling
// back to the error's own message.
func Reason(err error) string {
	if err == nil {
		return ""
	}
	if reason, ok := Descriptions[Selector(err)]; ok {
		return reason
	}
	return err.Error()
}
