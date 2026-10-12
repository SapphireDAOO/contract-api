package feereceiver

import "errors"

var (
	ErrConnect          = errors.New("connecting to the fee receiver")
	ErrInvalidProcessor = errors.New("processor must be simple or intermediated")
	ErrInvalidKind      = errors.New("kind must be single or meta")
	ErrUnknownToken     = errors.New("unknown token")
)
