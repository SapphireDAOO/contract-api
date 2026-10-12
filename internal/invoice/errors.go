package invoice

import "errors"

var (
	ErrNoInvoices          = errors.New("no invoice parameters provided")
	ErrNoPaymentTokens     = errors.New("at least one payment token is required")
	ErrUnknownPaymentToken = errors.New("unknown payment token")
	ErrOrderIDRequired     = errors.New("orderId is required")
	ErrInvalidAddress      = errors.New("is not a valid address")
	ErrPriceNotPositive    = errors.New("price must be greater than zero")

	ErrMissingOrderID       = errors.New("missing orderId")
	ErrMissingSeller        = errors.New("missing seller")
	ErrMissingPrice         = errors.New("missing price")
	ErrMissingPaymentTokens = errors.New("missing payment tokens")

	ErrEmptyInvoiceID   = errors.New("invoice id is empty")
	ErrInvalidInvoiceID = errors.New("invoice id is not valid base64url")
)
