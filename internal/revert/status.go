package revert

import "net/http"

// StatusCodes maps a description to the HTTP status the API answers with. A
// reason absent here is a fault rather than something the caller can act on,
// and falls back to 500 in httpx.
var StatusCodes = map[string]int{
	// Bad input: the contract rejected what the caller sent.
	"The buyer and seller addresses cannot be the same.":                 http.StatusBadRequest,
	"The account balance is insufficient to perform this action.":        http.StatusBadRequest,
	"The provided dispute resolution is invalid.":                        http.StatusBadRequest,
	"The native token payment is invalid for this invoice.":              http.StatusBadRequest,
	"The seller's payout share is invalid.":                              http.StatusBadRequest,
	"The price cannot be zero.":                                          http.StatusBadRequest,
	"The price specified is too low.":                                    http.StatusBadRequest,
	"A meta-invoice must contain at least one invoice.":                  http.StatusBadRequest,
	"The escrow hold period cannot be zero.":                             http.StatusBadRequest,
	"The number of fee receivers does not match the number of invoices.": http.StatusBadRequest,
	"The payment amount does not match the meta-invoice total.":          http.StatusBadRequest,
	"The seller address is invalid.":                                     http.StatusBadRequest,
	"The invoice must accept at least one payment token.":                http.StatusBadRequest,
	"That payment token is not accepted for this invoice.":               http.StatusBadRequest,
	"The fee authorization signature is invalid.":                        http.StatusBadRequest,
	"The fee receiver address is invalid.":                               http.StatusBadRequest,
	"The payment amount is incorrect.":                                   http.StatusBadRequest,
	"The dispute decision window is invalid.":                            http.StatusBadRequest,
	"The seller cannot pay their own invoice.":                           http.StatusBadRequest,
	"This invoice does not accept a native token payment.":               http.StatusBadRequest,
	"The value sent is too low.":                                         http.StatusBadRequest,
	"The oracle has no price feed for this token.":                       http.StatusBadRequest,
	"The note content cannot be empty.":                                  http.StatusBadRequest,
	"The address provided is invalid.":                                   http.StatusBadRequest,
	"The fee rate is invalid.":                                           http.StatusBadRequest,
	"The fee signer address is invalid.":                                 http.StatusBadRequest,
	"The new owner cannot be the zero address.":                          http.StatusBadRequest,
	"The wrapped native token address is invalid.":                       http.StatusBadRequest,
	"The transaction target is invalid.":                                 http.StatusBadRequest,
	"The approval threshold is invalid.":                                 http.StatusBadRequest,
	"That address is not a signer.":                                      http.StatusBadRequest,
	"The approval threshold cannot be zero.":                             http.StatusBadRequest,

	// The caller is not permitted to take the action.
	"The caller is not authorized to perform this action.":  http.StatusForbidden,
	"The caller does not own this automation workflow.":     http.StatusForbidden,
	"This action can only be taken by the multisig itself.": http.StatusForbidden,
	"The caller is not a signer on this multisig.":          http.StatusForbidden,

	// The subject of the request does not exist.
	"The specified invoice does not exist.":     http.StatusNotFound,
	"The specified task does not exist.":        http.StatusNotFound,
	"The specified note does not exist.":        http.StatusNotFound,
	"The specified transaction does not exist.": http.StatusNotFound,

	// The request conflicts with the current on-chain state.
	"The invoice is not in a valid state for this action.":           http.StatusConflict,
	"An invoice with this identifier already exists.":                http.StatusConflict,
	"A meta-invoice with this identifier already exists.":            http.StatusConflict,
	"The invoice has expired.":                                       http.StatusConflict,
	"The acceptance window for this invoice has passed.":             http.StatusConflict,
	"A task with this identifier already exists.":                    http.StatusConflict,
	"The escrow hold period has not yet elapsed.":                    http.StatusConflict,
	"The invoice is no longer valid.":                                http.StatusConflict,
	"The invoice is not eligible for a refund.":                      http.StatusConflict,
	"The contract has already been initialized.":                     http.StatusConflict,
	"The contract is already paused.":                                http.StatusConflict,
	"There is no active emergency pause.":                            http.StatusConflict,
	"There is no pending ownership handover.":                        http.StatusConflict,
	"The contract is not paused.":                                    http.StatusConflict,
	"That address is already a signer.":                              http.StatusConflict,
	"The transaction has already been approved.":                     http.StatusConflict,
	"This signer has already approved the transaction.":              http.StatusConflict,
	"The transaction has already been executed.":                     http.StatusConflict,
	"There are not enough signers to meet the threshold.":            http.StatusConflict,
	"Removing that signer would drop the count below the threshold.": http.StatusConflict,
	"The transaction has not been approved.":                         http.StatusConflict,
	"The transaction has not been proposed.":                         http.StatusConflict,
	"The transaction has not been proposed or approved.":             http.StatusConflict,

	// Temporary: retrying later can succeed, which is what separates these
	// from a 4xx.
	"The contract is paused.":                        http.StatusServiceUnavailable,
	"The oracle reported an invalid price.":          http.StatusServiceUnavailable,
	"The sequencer is down; prices are unavailable.": http.StatusServiceUnavailable,
	"The oracle price is stale.":                     http.StatusServiceUnavailable,
	"The oracle price feed is stale.":                http.StatusServiceUnavailable,
}
