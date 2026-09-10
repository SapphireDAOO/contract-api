package revert

// UnsupportedToken is OracleManager's UnsupportedToken() selector, which
// getUsdPerToken reverts with for a token it has no feed for.
const UnsupportedToken = "0x6a172882"

// Descriptions maps every custom error selector in the generated ABIs to the
// reason the API reports. Regenerating the ABIs can add, change or remove an
// error, so TestEveryABIErrorHasADescription fails when this drifts.
var Descriptions = map[string]string{
	// IntermediatedPaymentProcessor
	"0xb12e2421": "The buyer and seller addresses cannot be the same.",
	"0xf4d678b8": "The account balance is insufficient to perform this action.",
	"0x34819f90": "The provided dispute resolution is invalid.",
	"0x487e4409": "The invoice is not in a valid state for this action.",
	"0x214510aa": "The native token payment is invalid for this invoice.",
	"0x453fb42d": "The seller's payout share is invalid.",
	"0x074bc935": "An invoice with this identifier already exists.",
	"0x715d9228": "The specified invoice does not exist.",
	"0xb09960c1": "A meta-invoice with this identifier already exists.",
	"0x2c669f0a": "The price cannot be zero.",
	"0xdb8db569": "The price specified is too low.",
	"0x815ba404": "A meta-invoice must contain at least one invoice.",
	"0x705a7153": "The escrow hold period cannot be zero.",
	"0xa0c3f120": "The number of fee receivers does not match the number of invoices.",
	"0xc7632c7d": "The payment amount does not match the meta-invoice total.",
	"0x9589a27d": "The price oracle is not configured correctly.",
	"0xbab7ca35": "The seller address is invalid.",
	"0xf04e9cf0": "The invoice has expired.",
	"0xb883eab0": "The invoice must accept at least one payment token.",
	"0x8e6656fb": "That payment token is not accepted for this invoice.",
	"0x4ca249dc": "Contract deployment failed: the bytecode was empty.",
	"0xb06ebf3d": "Contract deployment failed.",

	// SimplePaymentProcessor. 0x1d5b1556 is InvalidInvoiceState(uint256), a
	// different selector from the intermediated processor's no-argument one.
	"0x2b8af0bb": "The acceptance window for this invoice has passed.",
	"0x6b22feb9": "A task with this identifier already exists.",
	"0xad2652ac": "The escrow hold period has not yet elapsed.",
	"0x47af6acc": "The payment amount is incorrect.",
	"0x39141cc3": "The dispute decision window is invalid.",
	"0x76f4a283": "The scheduler heap is in an invalid state.",
	"0x1d5b1556": "The invoice is not in a valid state for this action.",
	"0xff42dbfc": "The invoice is no longer valid.",
	"0xbb126ff1": "The invoice is not eligible for a refund.",
	"0x020175b1": "The seller cannot pay their own invoice.",
	"0xc325ae33": "The specified task does not exist.",
	"0xecb8b30d": "This invoice does not accept a native token payment.",
	"0x5033f274": "The value sent is too low.",

	// Declared by both payment processors
	"0xab35696f": "The contract is paused.",
	"0x667ecf9d": "The escrow withdrawal failed.",
	"0x1735eabe": "The fee authorization signature is invalid.",
	"0xd200485c": "The fee receiver address is invalid.",
	"0xab143c06": "A reentrant call was detected.",

	// OracleManager
	UnsupportedToken: "The oracle has no price feed for this token.",
	"0x00bfc921":     "The oracle reported an invalid price.",
	"0x032b3d00":     "The sequencer is down; prices are unavailable.",
	"0x19abf40e":     "The oracle price is stale.",
	"0x1087e109":     "The oracle price feed is stale.",

	// Notes
	"0x68b37036": "The note content cannot be empty.",
	"0xdc0b7363": "The specified note does not exist.",

	// PaymentAutomation
	"0xe6c4247b": "The address provided is invalid.",
	"0xbf241623": "The caller does not own this automation workflow.",

	// PaymentProcessorStorage
	"0x0dc149f0": "The contract has already been initialized.",
	"0x1785c681": "The contract is already paused.",
	"0x56d69198": "The fee rate is invalid.",
	"0x20d80102": "The fee signer address is invalid.",
	"0x7448fbae": "The new owner cannot be the zero address.",
	"0xdb469296": "There is no active emergency pause.",
	"0x6f5e8818": "There is no pending ownership handover.",
	"0x6cd60201": "The contract is not paused.",

	// NotAuthorized() and Unauthorized(), declared by several contracts. A
	// caller cannot act on them differently, so both read the same.
	"0xea8e4eb5": "The caller is not authorized to perform this action.",
	"0x82b42900": "The caller is not authorized to perform this action.",

	// Multisig
	"0xcc2fa558": "That address is already a signer.",
	"0x101f817a": "The transaction has already been approved.",
	"0xefa11217": "This signer has already approved the transaction.",
	"0x0dc10197": "The transaction has already been executed.",
	"0xacfdb444": "The multisig transaction failed to execute.",
	"0xc2ee9b9e": "There are not enough signers to meet the threshold.",
	"0x82d5d76a": "The transaction target is invalid.",
	"0xaabd5a09": "The approval threshold is invalid.",
	"0xda0357f7": "That address is not a signer.",
	"0x29c3b7ee": "This action can only be taken by the multisig itself.",
	"0xa1b035c8": "The caller is not a signer on this multisig.",
	"0x082d4dd3": "Removing that signer would drop the count below the threshold.",
	"0xf4124166": "The approval threshold cannot be zero.",
	"0x5451f114": "The specified transaction does not exist.",
	"0x34c68f2b": "The transaction has not been approved.",
	"0xd5a4fdee": "The transaction has not been proposed.",
	"0x76bf0404": "The transaction has not been proposed or approved.",
}
