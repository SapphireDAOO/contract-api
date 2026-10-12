package callback

import "errors"

var (
	ErrInvalidAmount           = errors.New("invalid amount")
	ErrUnsupportedPaymentToken = errors.New("unsupported payment token")
	ErrURLNotConfigured        = errors.New("callback URL is not configured")
	ErrAPIKeyNotConfigured     = errors.New("callback API key is not configured")
	ErrOrderIDRequired         = errors.New("orderId is required for callback URL")
	ErrActionRequired          = errors.New("action is required for callback URL")
)
