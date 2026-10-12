package oraclemanager

import "errors"

var (
	ErrNotInitialized     = errors.New("oracle manager contract is not initialized")
	ErrPriceCountMismatch = errors.New("oracle returned a price count that does not match the tokens")
)
