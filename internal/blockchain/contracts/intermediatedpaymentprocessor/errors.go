package intermediatedpaymentprocessor

import "errors"

var (
	ErrSingleInvoiceExpected    = errors.New("CreateInvoice expects exactly one parameter")
	ErrMultipleInvoicesExpected = errors.New("parameter has to be greater than one")
	ErrAddressNotInitialized    = errors.New("payment processor contract address is not initialized")
	ErrClientNotInitialized     = errors.New("blockchain client not initialized")
	ErrNoReceipt                = errors.New("no receipt returned for meta invoice creation")
	ErrMetaInvoiceEventMissing  = errors.New("MetaInvoiceCreated event not found in receipt logs")
	ErrUnsupportedAction        = errors.New("unsupported marketplace action")
	ErrTransactionReverted      = errors.New("transaction reverted")
)
