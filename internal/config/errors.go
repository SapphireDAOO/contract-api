package config

import "errors"

// Field errors read after the field they describe, e.g. "contracts.notes is
// required", so each is a predicate.
var (
	ErrRequired          = errors.New("is required")
	ErrEmpty             = errors.New("is empty")
	ErrInvalidAddress    = errors.New("is not a valid address")
	ErrZeroAddress       = errors.New("cannot be the zero address")
	ErrInvalidURL        = errors.New("is not a valid URL")
	ErrUnsupportedScheme = errors.New("has an unsupported scheme")
	ErrTrailingSlash     = errors.New("must not end in a slash")
	ErrNegative          = errors.New("cannot be negative")
	ErrNotHostPort       = errors.New("must be host:port")
	ErrNotHost           = errors.New("must be a host such as app.example.com, not a URL")
	ErrKeyLength         = errors.New("must be 32 bytes of hex")
	ErrInvalidHex        = errors.New("is not valid hex")
	ErrOutOfRange        = errors.New("is out of range")
)

var (
	ErrReadConfig     = errors.New("reading config")
	ErrParseConfig    = errors.New("parsing config")
	ErrNoNetwork      = errors.New("no network selected")
	ErrUnknownNetwork = errors.New("is not defined")
	ErrNoTokens       = errors.New("at least one payment token is required")
	ErrEmptySymbol    = errors.New("a symbol cannot be empty")
)
