package note

import "errors"

var (
	ErrContentRequired = errors.New("content is required")
	ErrContentHex      = errors.New("content must be 0x-prefixed hex")
	ErrContentTooLong  = errors.New("content is too large")
)
