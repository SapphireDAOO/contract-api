package blockchain

import "errors"

var (
	ErrRPCConnect          = errors.New("failed to connect to RPC")
	ErrWSSConnect          = errors.New("failed to connect to WSS")
	ErrChainID             = errors.New("failed to retrieve chain ID")
	ErrSignerNotConfigured = errors.New("signer key is not configured")
	ErrSignerNotHex        = errors.New("signer key is not valid hex")
	ErrSignerInvalid       = errors.New("signer key is not a valid private key")
)
