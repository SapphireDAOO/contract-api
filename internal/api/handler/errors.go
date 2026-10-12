package handler

import "errors"

var (
	ErrNotBase10                = errors.New("must be a base-10 integer")
	ErrZeroShare                = errors.New("share can not be zero")
	ErrEmptyInvoiceID           = errors.New("empty invoice id in path")
	ErrSettlementExpired        = errors.New("settlement window has expired")
	ErrFeeReceiverNotConfigured = errors.New("fee receiver service is not configured")
	ErrMetaNotIntermediated     = errors.New("meta invoices are intermediated only")
	ErrNoFeeReceiverState       = errors.New("is not in a state that takes fee receivers")

	ErrOnlyUSD            = errors.New("only USD is priced")
	ErrNoTokenRequested   = errors.New("no token requested")
	ErrUnknownToken       = errors.New("unknown token")
	ErrOracleUnavailable  = errors.New("oracle unavailable")
	ErrNonPositivePrice   = errors.New("oracle returned a non-positive price")
	ErrPriceNotInvertible = errors.New("price cannot be inverted")
)
