package auth

import (
	"errors"
	"fmt"
)

var (
	ErrSecretTooShort  = fmt.Errorf("session secret must be at least %d characters", MinSecretLength)
	ErrNoDomains       = errors.New("at least one sign-in domain is required")
	ErrNonceGeneration = errors.New("generating nonce")
	ErrSessionSigning  = errors.New("signing session")

	ErrInvalidMessage   = errors.New("invalid sign-in message")
	ErrWrongDomain      = errors.New("sign-in message is for another domain")
	ErrWrongChain       = errors.New("sign-in message is for another chain")
	ErrMessageExpired   = errors.New("sign-in message is expired or not yet valid")
	ErrInvalidNonce     = errors.New("nonce was not issued by this API")
	ErrNonceExpired     = errors.New("nonce has expired")
	ErrNonceUsed        = errors.New("nonce has already been used")
	ErrInvalidSignature = errors.New("signature does not match the message's address")
)

// Sent to API clients as the response's error text.
var (
	ErrSignInNotConfigured = errors.New("Sign-in is not configured")
	ErrSignInRequired      = errors.New("Sign-in required")
	ErrInvalidSession      = errors.New("Session is invalid or expired")
)
