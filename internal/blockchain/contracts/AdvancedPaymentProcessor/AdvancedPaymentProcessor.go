// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymentprocessor

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// IAdvancedPaymentProcessorInvoice is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorInvoice struct {
	InvoiceId             *big.Int
	Buyer                 common.Address
	Seller                common.Address
	Escrow                common.Address
	PaymentToken          common.Address
	State                 uint8
	PaidAt                *big.Int
	CreatedAt             *big.Int
	ReleaseWindow         uint32
	InvoiceExpiryDuration uint32
	TimeBeforeCancelation uint32
	Price                 *big.Int
	AmountPaid            *big.Int
	MetaInvoiceKey        [32]byte
}

// IAdvancedPaymentProcessorInvoiceCreationParam is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorInvoiceCreationParam struct {
	Seller                common.Address
	Buyer                 common.Address
	InvoiceExpiryDuration uint32
	TimeBeforeCancelation uint32
	ReleaseWindow         uint32
	Price                 *big.Int
}

// IAdvancedPaymentProcessorMetaInvoice is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorMetaInvoice struct {
	Buyer        common.Address
	Price        *big.Int
	Upper        *big.Int
	Lower        *big.Int
	PaymentToken common.Address
	InvoiceId    *big.Int
}

// PaymentprocessorMetaData contains all meta data concerning the Paymentprocessor contract.
var PaymentprocessorMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ownerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"marketplaceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nativeTokenAggregatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACCEPTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELATION_ACCEPTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELATION_REJECTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELATION_REQUESTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_DECIMAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_DISMISSED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_RESOLVED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_SETTLED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIATED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REFUNDED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REJECTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELEASED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptInvoice\",\"inputs\":[{\"name\":\"invoiceKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"invoiceKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimExpiredInvoiceRefunds\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeSalt\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"createDispute\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMetaInvoice\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"param\",\"type\":\"tuple[]\",\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam[]\",\"components\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSingleInvoice\",\"inputs\":[{\"name\":\"param\",\"type\":\"tuple\",\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam\",\"components\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"paidAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"createdAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metaInvoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketplace\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetaInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAdvancedPaymentProcessor.MetaInvoice\",\"components\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"upper\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetaInvoiceIdForSub\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextMetaInvoiceId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPredictedAddress\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenValueFromUsd\",\"inputs\":[{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleCancelationRequest\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accept\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payMetaInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"paySingleInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releasePayment\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestCancelation\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestCancelation\",\"inputs\":[{\"name\":\"invoiceKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"resolution\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"sellerShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMarketplace\",\"inputs\":[{\"name\":\"marketplaceAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalMetaInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalUniqueInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"CancelationRequestHandled\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CancelationRequested\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeCreated\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeDismissed\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeSettled\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"buyerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowCreated\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"escrow\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpiredInvoiceRefunded\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceAccepted\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCanceled\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCreated\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"invoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"paidAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"createdAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metaInvoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoicePaid\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"escrowAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceRejected\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetaInvoiceCreated\",\"inputs\":[{\"name\":\"metaInvoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"metaInvoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAdvancedPaymentProcessor.MetaInvoice\",\"components\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"upper\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReleased\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CancelationRequestDeadlinePassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisputeWindowExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBuyer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDisputeResolution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInvoiceState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMetaInvoicePayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNativePayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPaymentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSellersPayoutShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceResponseTimeExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceStillActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoShareAllocatedToBuyer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSubInvoiceCancelled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedBuyer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroEscrowBalance\",\"inputs\":[]}]",
	ID:  "Paymentprocessor",
}

// Paymentprocessor is an auto generated Go binding around an Ethereum contract.
type Paymentprocessor struct {
	abi abi.ABI
}

// NewPaymentprocessor creates a new instance of Paymentprocessor.
func NewPaymentprocessor() *Paymentprocessor {
	parsed, err := PaymentprocessorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Paymentprocessor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Paymentprocessor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address paymentProcessorStorageAddress, address ownerAddress, address marketplaceAddress, address nativeTokenAggregatorAddress) returns()
func (paymentprocessor *Paymentprocessor) PackConstructor(paymentProcessorStorageAddress common.Address, ownerAddress common.Address, marketplaceAddress common.Address, nativeTokenAggregatorAddress common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("", paymentProcessorStorageAddress, ownerAddress, marketplaceAddress, nativeTokenAggregatorAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackACCEPTED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4324efd7.
//
// Solidity: function ACCEPTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackACCEPTED() []byte {
	enc, err := paymentprocessor.abi.Pack("ACCEPTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackACCEPTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4324efd7.
//
// Solidity: function ACCEPTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackACCEPTED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("ACCEPTED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackBASISPOINTS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (paymentprocessor *Paymentprocessor) PackBASISPOINTS() []byte {
	enc, err := paymentprocessor.abi.Pack("BASIS_POINTS")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBASISPOINTS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (paymentprocessor *Paymentprocessor) UnpackBASISPOINTS(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("BASIS_POINTS", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCANCELATIONACCEPTED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d2c6615.
//
// Solidity: function CANCELATION_ACCEPTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackCANCELATIONACCEPTED() []byte {
	enc, err := paymentprocessor.abi.Pack("CANCELATION_ACCEPTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCANCELATIONACCEPTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d2c6615.
//
// Solidity: function CANCELATION_ACCEPTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackCANCELATIONACCEPTED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("CANCELATION_ACCEPTED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackCANCELATIONREJECTED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9934d4c.
//
// Solidity: function CANCELATION_REJECTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackCANCELATIONREJECTED() []byte {
	enc, err := paymentprocessor.abi.Pack("CANCELATION_REJECTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCANCELATIONREJECTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc9934d4c.
//
// Solidity: function CANCELATION_REJECTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackCANCELATIONREJECTED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("CANCELATION_REJECTED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackCANCELATIONREQUESTED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecc224ce.
//
// Solidity: function CANCELATION_REQUESTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackCANCELATIONREQUESTED() []byte {
	enc, err := paymentprocessor.abi.Pack("CANCELATION_REQUESTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCANCELATIONREQUESTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xecc224ce.
//
// Solidity: function CANCELATION_REQUESTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackCANCELATIONREQUESTED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("CANCELATION_REQUESTED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackCANCELED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackCANCELED() []byte {
	enc, err := paymentprocessor.abi.Pack("CANCELED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCANCELED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackCANCELED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("CANCELED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDEFAULTDECIMAL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackDEFAULTDECIMAL() []byte {
	enc, err := paymentprocessor.abi.Pack("DEFAULT_DECIMAL")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTDECIMAL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackDEFAULTDECIMAL(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("DEFAULT_DECIMAL", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDISPUTED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x620cac66.
//
// Solidity: function DISPUTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackDISPUTED() []byte {
	enc, err := paymentprocessor.abi.Pack("DISPUTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x620cac66.
//
// Solidity: function DISPUTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackDISPUTED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("DISPUTED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDISPUTEDISMISSED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1201920.
//
// Solidity: function DISPUTE_DISMISSED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackDISPUTEDISMISSED() []byte {
	enc, err := paymentprocessor.abi.Pack("DISPUTE_DISMISSED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTEDISMISSED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1201920.
//
// Solidity: function DISPUTE_DISMISSED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackDISPUTEDISMISSED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("DISPUTE_DISMISSED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDISPUTERESOLVED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x677a4ece.
//
// Solidity: function DISPUTE_RESOLVED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackDISPUTERESOLVED() []byte {
	enc, err := paymentprocessor.abi.Pack("DISPUTE_RESOLVED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTERESOLVED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x677a4ece.
//
// Solidity: function DISPUTE_RESOLVED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackDISPUTERESOLVED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("DISPUTE_RESOLVED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDISPUTESETTLED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55b9da62.
//
// Solidity: function DISPUTE_SETTLED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackDISPUTESETTLED() []byte {
	enc, err := paymentprocessor.abi.Pack("DISPUTE_SETTLED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTESETTLED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x55b9da62.
//
// Solidity: function DISPUTE_SETTLED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackDISPUTESETTLED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("DISPUTE_SETTLED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackINITIATED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88a03c53.
//
// Solidity: function INITIATED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackINITIATED() []byte {
	enc, err := paymentprocessor.abi.Pack("INITIATED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackINITIATED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88a03c53.
//
// Solidity: function INITIATED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackINITIATED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("INITIATED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackPAID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd31a03c2.
//
// Solidity: function PAID() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackPAID() []byte {
	enc, err := paymentprocessor.abi.Pack("PAID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPAID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd31a03c2.
//
// Solidity: function PAID() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackPAID(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("PAID", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackREFUNDED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94e15c8f.
//
// Solidity: function REFUNDED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackREFUNDED() []byte {
	enc, err := paymentprocessor.abi.Pack("REFUNDED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackREFUNDED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94e15c8f.
//
// Solidity: function REFUNDED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackREFUNDED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("REFUNDED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackREJECTED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb7390b7.
//
// Solidity: function REJECTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackREJECTED() []byte {
	enc, err := paymentprocessor.abi.Pack("REJECTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackREJECTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcb7390b7.
//
// Solidity: function REJECTED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackREJECTED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("REJECTED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackRELEASED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b52ea81.
//
// Solidity: function RELEASED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) PackRELEASED() []byte {
	enc, err := paymentprocessor.abi.Pack("RELEASED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRELEASED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b52ea81.
//
// Solidity: function RELEASED() view returns(uint8)
func (paymentprocessor *Paymentprocessor) UnpackRELEASED(data []byte) (uint8, error) {
	out, err := paymentprocessor.abi.Unpack("RELEASED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackAcceptInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40713eb5.
//
// Solidity: function acceptInvoice(bytes32[] invoiceKeys) returns()
func (paymentprocessor *Paymentprocessor) PackAcceptInvoice(invoiceKeys [][32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("acceptInvoice", invoiceKeys)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAcceptInvoice0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47a424b5.
//
// Solidity: function acceptInvoice(bytes32 invoiceKey) returns()
func (paymentprocessor *Paymentprocessor) PackAcceptInvoice0(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("acceptInvoice0", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfe06a2b.
//
// Solidity: function cancelInvoice(bytes32[] invoiceKeys) returns()
func (paymentprocessor *Paymentprocessor) PackCancelInvoice(invoiceKeys [][32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("cancelInvoice", invoiceKeys)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelInvoice0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0aa7e2e.
//
// Solidity: function cancelInvoice(bytes32 invoiceKey) returns()
func (paymentprocessor *Paymentprocessor) PackCancelInvoice0(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("cancelInvoice0", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (paymentprocessor *Paymentprocessor) PackCancelOwnershipHandover() []byte {
	enc, err := paymentprocessor.abi.Pack("cancelOwnershipHandover")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaimExpiredInvoiceRefunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa489897b.
//
// Solidity: function claimExpiredInvoiceRefunds(bytes32 invoiceKey) returns()
func (paymentprocessor *Paymentprocessor) PackClaimExpiredInvoiceRefunds(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("claimExpiredInvoiceRefunds", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCompleteOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (paymentprocessor *Paymentprocessor) PackCompleteOwnershipHandover(pendingOwner common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("completeOwnershipHandover", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackComputeSalt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c074054.
//
// Solidity: function computeSalt(address seller, address buyer, bytes32 invoiceKey) pure returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackComputeSalt(seller common.Address, buyer common.Address, invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("computeSalt", seller, buyer, invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackComputeSalt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c074054.
//
// Solidity: function computeSalt(address seller, address buyer, bytes32 invoiceKey) pure returns(bytes32)
func (paymentprocessor *Paymentprocessor) UnpackComputeSalt(data []byte) ([32]byte, error) {
	out, err := paymentprocessor.abi.Unpack("computeSalt", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCreateDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29092dc0.
//
// Solidity: function createDispute(bytes32 invoiceKey) returns()
func (paymentprocessor *Paymentprocessor) PackCreateDispute(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("createDispute", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b3fd37e.
//
// Solidity: function createMetaInvoice(address buyer, (address,address,uint32,uint32,uint32,uint256)[] param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackCreateMetaInvoice(buyer common.Address, param []IAdvancedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := paymentprocessor.abi.Pack("createMetaInvoice", buyer, param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b3fd37e.
//
// Solidity: function createMetaInvoice(address buyer, (address,address,uint32,uint32,uint32,uint256)[] param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) UnpackCreateMetaInvoice(data []byte) ([32]byte, error) {
	out, err := paymentprocessor.abi.Unpack("createMetaInvoice", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCreateSingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a38c874.
//
// Solidity: function createSingleInvoice((address,address,uint32,uint32,uint32,uint256) param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackCreateSingleInvoice(param IAdvancedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := paymentprocessor.abi.Pack("createSingleInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSingleInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a38c874.
//
// Solidity: function createSingleInvoice((address,address,uint32,uint32,uint32,uint256) param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) UnpackCreateSingleInvoice(data []byte) ([32]byte, error) {
	out, err := paymentprocessor.abi.Unpack("createSingleInvoice", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb802c8b.
//
// Solidity: function getInvoice(bytes32 invoiceKey) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32))
func (paymentprocessor *Paymentprocessor) PackGetInvoice(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("getInvoice", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcb802c8b.
//
// Solidity: function getInvoice(bytes32 invoiceKey) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32))
func (paymentprocessor *Paymentprocessor) UnpackGetInvoice(data []byte) (IAdvancedPaymentProcessorInvoice, error) {
	out, err := paymentprocessor.abi.Unpack("getInvoice", data)
	if err != nil {
		return *new(IAdvancedPaymentProcessorInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorInvoice)).(*IAdvancedPaymentProcessorInvoice)
	return out0, err
}

// PackGetMarketplace is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (paymentprocessor *Paymentprocessor) PackGetMarketplace() []byte {
	enc, err := paymentprocessor.abi.Pack("getMarketplace")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMarketplace is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (paymentprocessor *Paymentprocessor) UnpackGetMarketplace(data []byte) (common.Address, error) {
	out, err := paymentprocessor.abi.Unpack("getMarketplace", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ccef592.
//
// Solidity: function getMetaInvoice(bytes32 invoiceKey) view returns((address,uint256,uint256,uint256,address,uint256))
func (paymentprocessor *Paymentprocessor) PackGetMetaInvoice(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("getMetaInvoice", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ccef592.
//
// Solidity: function getMetaInvoice(bytes32 invoiceKey) view returns((address,uint256,uint256,uint256,address,uint256))
func (paymentprocessor *Paymentprocessor) UnpackGetMetaInvoice(data []byte) (IAdvancedPaymentProcessorMetaInvoice, error) {
	out, err := paymentprocessor.abi.Unpack("getMetaInvoice", data)
	if err != nil {
		return *new(IAdvancedPaymentProcessorMetaInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorMetaInvoice)).(*IAdvancedPaymentProcessorMetaInvoice)
	return out0, err
}

// PackGetMetaInvoiceIdForSub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5cbc9082.
//
// Solidity: function getMetaInvoiceIdForSub(bytes32 invoiceKey) view returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackGetMetaInvoiceIdForSub(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("getMetaInvoiceIdForSub", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMetaInvoiceIdForSub is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5cbc9082.
//
// Solidity: function getMetaInvoiceIdForSub(bytes32 invoiceKey) view returns(bytes32)
func (paymentprocessor *Paymentprocessor) UnpackGetMetaInvoiceIdForSub(data []byte) ([32]byte, error) {
	out, err := paymentprocessor.abi.Unpack("getMetaInvoiceIdForSub", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetNextInvoiceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint256)
func (paymentprocessor *Paymentprocessor) PackGetNextInvoiceId() []byte {
	enc, err := paymentprocessor.abi.Pack("getNextInvoiceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint256)
func (paymentprocessor *Paymentprocessor) UnpackGetNextInvoiceId(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("getNextInvoiceId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetNextMetaInvoiceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint256)
func (paymentprocessor *Paymentprocessor) PackGetNextMetaInvoiceId() []byte {
	enc, err := paymentprocessor.abi.Pack("getNextMetaInvoiceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextMetaInvoiceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint256)
func (paymentprocessor *Paymentprocessor) UnpackGetNextMetaInvoiceId(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("getNextMetaInvoiceId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetPredictedAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bec0c01.
//
// Solidity: function getPredictedAddress(bytes32 salt) view returns(address)
func (paymentprocessor *Paymentprocessor) PackGetPredictedAddress(salt [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("getPredictedAddress", salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPredictedAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4bec0c01.
//
// Solidity: function getPredictedAddress(bytes32 salt) view returns(address)
func (paymentprocessor *Paymentprocessor) UnpackGetPredictedAddress(data []byte) (common.Address, error) {
	out, err := paymentprocessor.abi.Unpack("getPredictedAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetTokenValueFromUsd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address paymentToken, uint256 price) view returns(uint256)
func (paymentprocessor *Paymentprocessor) PackGetTokenValueFromUsd(paymentToken common.Address, price *big.Int) []byte {
	enc, err := paymentprocessor.abi.Pack("getTokenValueFromUsd", paymentToken, price)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenValueFromUsd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address paymentToken, uint256 price) view returns(uint256)
func (paymentprocessor *Paymentprocessor) UnpackGetTokenValueFromUsd(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("getTokenValueFromUsd", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackHandleCancelationRequest is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9a6618a.
//
// Solidity: function handleCancelationRequest(bytes32 invoiceKey, bool accept) returns()
func (paymentprocessor *Paymentprocessor) PackHandleCancelationRequest(invoiceKey [32]byte, accept bool) []byte {
	enc, err := paymentprocessor.abi.Pack("handleCancelationRequest", invoiceKey, accept)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (paymentprocessor *Paymentprocessor) PackOwner() []byte {
	enc, err := paymentprocessor.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (paymentprocessor *Paymentprocessor) UnpackOwner(data []byte) (common.Address, error) {
	out, err := paymentprocessor.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackOwnershipHandoverExpiresAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (paymentprocessor *Paymentprocessor) PackOwnershipHandoverExpiresAt(pendingOwner common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("ownershipHandoverExpiresAt", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnershipHandoverExpiresAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (paymentprocessor *Paymentprocessor) UnpackOwnershipHandoverExpiresAt(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("ownershipHandoverExpiresAt", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPayMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97f2ef98.
//
// Solidity: function payMetaInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (paymentprocessor *Paymentprocessor) PackPayMetaInvoice(invoiceKey [32]byte, paymentToken common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("payMetaInvoice", invoiceKey, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPaySingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91a61baf.
//
// Solidity: function paySingleInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (paymentprocessor *Paymentprocessor) PackPaySingleInvoice(invoiceKey [32]byte, paymentToken common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("paySingleInvoice", invoiceKey, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (paymentprocessor *Paymentprocessor) PackPpStorage() []byte {
	enc, err := paymentprocessor.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (paymentprocessor *Paymentprocessor) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := paymentprocessor.abi.Unpack("ppStorage", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackReleasePayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7aa1ed58.
//
// Solidity: function releasePayment(bytes32 invoiceKey) returns()
func (paymentprocessor *Paymentprocessor) PackReleasePayment(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("releasePayment", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (paymentprocessor *Paymentprocessor) PackRenounceOwnership() []byte {
	enc, err := paymentprocessor.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestCancelation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0982cb6a.
//
// Solidity: function requestCancelation(bytes32 invoiceKey) returns()
func (paymentprocessor *Paymentprocessor) PackRequestCancelation(invoiceKey [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("requestCancelation", invoiceKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestCancelation0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7afd65c9.
//
// Solidity: function requestCancelation(bytes32[] invoiceKeys) returns()
func (paymentprocessor *Paymentprocessor) PackRequestCancelation0(invoiceKeys [][32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("requestCancelation0", invoiceKeys)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (paymentprocessor *Paymentprocessor) PackRequestOwnershipHandover() []byte {
	enc, err := paymentprocessor.abi.Pack("requestOwnershipHandover")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackResolveDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xff258f33.
//
// Solidity: function resolveDispute(bytes32 invoiceKey, uint8 resolution, uint256 sellerShare) returns()
func (paymentprocessor *Paymentprocessor) PackResolveDispute(invoiceKey [32]byte, resolution uint8, sellerShare *big.Int) []byte {
	enc, err := paymentprocessor.abi.Pack("resolveDispute", invoiceKey, resolution, sellerShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMarketplace is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddress) returns()
func (paymentprocessor *Paymentprocessor) PackSetMarketplace(marketplaceAddress common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("setMarketplace", marketplaceAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPriceFeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address aggregator) returns()
func (paymentprocessor *Paymentprocessor) PackSetPriceFeed(token common.Address, aggregator common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("setPriceFeed", token, aggregator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalMetaInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint256)
func (paymentprocessor *Paymentprocessor) PackTotalMetaInvoiceCreated() []byte {
	enc, err := paymentprocessor.abi.Pack("totalMetaInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalMetaInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint256)
func (paymentprocessor *Paymentprocessor) UnpackTotalMetaInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("totalMetaInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTotalUniqueInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint256)
func (paymentprocessor *Paymentprocessor) PackTotalUniqueInvoiceCreated() []byte {
	enc, err := paymentprocessor.abi.Pack("totalUniqueInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalUniqueInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint256)
func (paymentprocessor *Paymentprocessor) UnpackTotalUniqueInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := paymentprocessor.abi.Unpack("totalUniqueInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (paymentprocessor *Paymentprocessor) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PaymentprocessorCancelationRequestHandled represents a CancelationRequestHandled event raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequestHandled struct {
	InvoiceKey [32]byte
	Accepted   bool
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorCancelationRequestHandledEventName = "CancelationRequestHandled"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorCancelationRequestHandled) ContractEventName() string {
	return PaymentprocessorCancelationRequestHandledEventName
}

// UnpackCancelationRequestHandledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CancelationRequestHandled(bytes32 indexed invoiceKey, bool indexed accepted)
func (paymentprocessor *Paymentprocessor) UnpackCancelationRequestHandledEvent(log *types.Log) (*PaymentprocessorCancelationRequestHandled, error) {
	event := "CancelationRequestHandled"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorCancelationRequestHandled)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorCancelationRequested represents a CancelationRequested event raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequested struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorCancelationRequestedEventName = "CancelationRequested"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorCancelationRequested) ContractEventName() string {
	return PaymentprocessorCancelationRequestedEventName
}

// UnpackCancelationRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CancelationRequested(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackCancelationRequestedEvent(log *types.Log) (*PaymentprocessorCancelationRequested, error) {
	event := "CancelationRequested"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorCancelationRequested)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorDisputeCreated represents a DisputeCreated event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeCreated struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeCreatedEventName = "DisputeCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeCreated) ContractEventName() string {
	return PaymentprocessorDisputeCreatedEventName
}

// UnpackDisputeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeCreated(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackDisputeCreatedEvent(log *types.Log) (*PaymentprocessorDisputeCreated, error) {
	event := "DisputeCreated"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorDisputeCreated)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorDisputeDismissed represents a DisputeDismissed event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeDismissed struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeDismissedEventName = "DisputeDismissed"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeDismissed) ContractEventName() string {
	return PaymentprocessorDisputeDismissedEventName
}

// UnpackDisputeDismissedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeDismissed(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackDisputeDismissedEvent(log *types.Log) (*PaymentprocessorDisputeDismissed, error) {
	event := "DisputeDismissed"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorDisputeDismissed)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorDisputeResolved represents a DisputeResolved event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeResolved struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeResolvedEventName = "DisputeResolved"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeResolved) ContractEventName() string {
	return PaymentprocessorDisputeResolvedEventName
}

// UnpackDisputeResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeResolved(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackDisputeResolvedEvent(log *types.Log) (*PaymentprocessorDisputeResolved, error) {
	event := "DisputeResolved"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorDisputeResolved)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorDisputeSettled represents a DisputeSettled event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeSettled struct {
	InvoiceKey   [32]byte
	SellerAmount *big.Int
	BuyerAmount  *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeSettledEventName = "DisputeSettled"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeSettled) ContractEventName() string {
	return PaymentprocessorDisputeSettledEventName
}

// UnpackDisputeSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeSettled(bytes32 indexed invoiceKey, uint256 sellerAmount, uint256 buyerAmount)
func (paymentprocessor *Paymentprocessor) UnpackDisputeSettledEvent(log *types.Log) (*PaymentprocessorDisputeSettled, error) {
	event := "DisputeSettled"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorDisputeSettled)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorEscrowCreated represents a EscrowCreated event raised by the Paymentprocessor contract.
type PaymentprocessorEscrowCreated struct {
	InvoiceKey [32]byte
	Escrow     common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorEscrowCreatedEventName = "EscrowCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorEscrowCreated) ContractEventName() string {
	return PaymentprocessorEscrowCreatedEventName
}

// UnpackEscrowCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EscrowCreated(bytes32 indexed invoiceKey, address indexed escrow)
func (paymentprocessor *Paymentprocessor) UnpackEscrowCreatedEvent(log *types.Log) (*PaymentprocessorEscrowCreated, error) {
	event := "EscrowCreated"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorEscrowCreated)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorExpiredInvoiceRefunded represents a ExpiredInvoiceRefunded event raised by the Paymentprocessor contract.
type PaymentprocessorExpiredInvoiceRefunded struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorExpiredInvoiceRefundedEventName = "ExpiredInvoiceRefunded"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorExpiredInvoiceRefunded) ContractEventName() string {
	return PaymentprocessorExpiredInvoiceRefundedEventName
}

// UnpackExpiredInvoiceRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExpiredInvoiceRefunded(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackExpiredInvoiceRefundedEvent(log *types.Log) (*PaymentprocessorExpiredInvoiceRefunded, error) {
	event := "ExpiredInvoiceRefunded"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorExpiredInvoiceRefunded)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorInvoiceAccepted represents a InvoiceAccepted event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceAccepted struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoiceAcceptedEventName = "InvoiceAccepted"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoiceAccepted) ContractEventName() string {
	return PaymentprocessorInvoiceAcceptedEventName
}

// UnpackInvoiceAcceptedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceAccepted(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackInvoiceAcceptedEvent(log *types.Log) (*PaymentprocessorInvoiceAccepted, error) {
	event := "InvoiceAccepted"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorInvoiceAccepted)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCanceled struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoiceCanceled) ContractEventName() string {
	return PaymentprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackInvoiceCanceledEvent(log *types.Log) (*PaymentprocessorInvoiceCanceled, error) {
	event := "InvoiceCanceled"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorInvoiceCanceled)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorInvoiceCreated represents a InvoiceCreated event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCreated struct {
	InvoiceKey [32]byte
	Invoice    IAdvancedPaymentProcessorInvoice
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoiceCreated) ContractEventName() string {
	return PaymentprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(bytes32 indexed invoiceKey, (uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32) invoice)
func (paymentprocessor *Paymentprocessor) UnpackInvoiceCreatedEvent(log *types.Log) (*PaymentprocessorInvoiceCreated, error) {
	event := "InvoiceCreated"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorInvoiceCreated)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorInvoicePaid represents a InvoicePaid event raised by the Paymentprocessor contract.
type PaymentprocessorInvoicePaid struct {
	InvoiceKey    [32]byte
	PaymentToken  common.Address
	EscrowAddress common.Address
	Amount        *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoicePaidEventName = "InvoicePaid"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoicePaid) ContractEventName() string {
	return PaymentprocessorInvoicePaidEventName
}

// UnpackInvoicePaidEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoicePaid(bytes32 indexed invoiceKey, address paymentToken, address escrowAddress, uint256 amount)
func (paymentprocessor *Paymentprocessor) UnpackInvoicePaidEvent(log *types.Log) (*PaymentprocessorInvoicePaid, error) {
	event := "InvoicePaid"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorInvoicePaid)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorInvoiceRejected represents a InvoiceRejected event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceRejected struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoiceRejectedEventName = "InvoiceRejected"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoiceRejected) ContractEventName() string {
	return PaymentprocessorInvoiceRejectedEventName
}

// UnpackInvoiceRejectedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceRejected(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackInvoiceRejectedEvent(log *types.Log) (*PaymentprocessorInvoiceRejected, error) {
	event := "InvoiceRejected"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorInvoiceRejected)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorMetaInvoiceCreated represents a MetaInvoiceCreated event raised by the Paymentprocessor contract.
type PaymentprocessorMetaInvoiceCreated struct {
	MetaInvoiceKey [32]byte
	MetaInvoice    IAdvancedPaymentProcessorMetaInvoice
	Raw            *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorMetaInvoiceCreatedEventName = "MetaInvoiceCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorMetaInvoiceCreated) ContractEventName() string {
	return PaymentprocessorMetaInvoiceCreatedEventName
}

// UnpackMetaInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetaInvoiceCreated(bytes32 indexed metaInvoiceKey, (address,uint256,uint256,uint256,address,uint256) metaInvoice)
func (paymentprocessor *Paymentprocessor) UnpackMetaInvoiceCreatedEvent(log *types.Log) (*PaymentprocessorMetaInvoiceCreated, error) {
	event := "MetaInvoiceCreated"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorMetaInvoiceCreated)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorOwnershipHandoverCanceledEventName = "OwnershipHandoverCanceled"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorOwnershipHandoverCanceled) ContractEventName() string {
	return PaymentprocessorOwnershipHandoverCanceledEventName
}

// UnpackOwnershipHandoverCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (paymentprocessor *Paymentprocessor) UnpackOwnershipHandoverCanceledEvent(log *types.Log) (*PaymentprocessorOwnershipHandoverCanceled, error) {
	event := "OwnershipHandoverCanceled"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorOwnershipHandoverCanceled)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorOwnershipHandoverRequestedEventName = "OwnershipHandoverRequested"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorOwnershipHandoverRequested) ContractEventName() string {
	return PaymentprocessorOwnershipHandoverRequestedEventName
}

// UnpackOwnershipHandoverRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (paymentprocessor *Paymentprocessor) UnpackOwnershipHandoverRequestedEvent(log *types.Log) (*PaymentprocessorOwnershipHandoverRequested, error) {
	event := "OwnershipHandoverRequested"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorOwnershipHandoverRequested)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorOwnershipTransferred represents a OwnershipTransferred event raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorOwnershipTransferred) ContractEventName() string {
	return PaymentprocessorOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (paymentprocessor *Paymentprocessor) UnpackOwnershipTransferredEvent(log *types.Log) (*PaymentprocessorOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PaymentprocessorPaymentReleased represents a PaymentReleased event raised by the Paymentprocessor contract.
type PaymentprocessorPaymentReleased struct {
	InvoiceKey [32]byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorPaymentReleasedEventName = "PaymentReleased"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorPaymentReleased) ContractEventName() string {
	return PaymentprocessorPaymentReleasedEventName
}

// UnpackPaymentReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentReleased(bytes32 indexed invoiceKey)
func (paymentprocessor *Paymentprocessor) UnpackPaymentReleasedEvent(log *types.Log) (*PaymentprocessorPaymentReleased, error) {
	event := "PaymentReleased"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorPaymentReleased)
	if len(log.Data) > 0 {
		if err := paymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessor.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (paymentprocessor *Paymentprocessor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["AlreadyInitialized"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["AlreadyRefunded"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackAlreadyRefundedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["CancelationRequestDeadlinePassed"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackCancelationRequestDeadlinePassedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["DisputeWindowExpired"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackDisputeWindowExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["EscrowAddressMismatch"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackEscrowAddressMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidBuyer"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidBuyerError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidDisputeResolution"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidDisputeResolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidInvoiceStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidMetaInvoicePayment"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidMetaInvoicePaymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidNativePayment"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidNativePaymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidPaymentToken"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidPaymentTokenError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidSellersPayoutShare"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidSellersPayoutShareError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvoiceDoesNotExist"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvoiceDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvoiceExpired"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvoiceExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvoiceResponseTimeExpired"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvoiceResponseTimeExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvoiceStillActive"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvoiceStillActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNewOwnerIsZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NoHandoverRequest"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNoHandoverRequestError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NoShareAllocatedToBuyer"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNoShareAllocatedToBuyerError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NoSubInvoiceCancelled"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNoSubInvoiceCancelledError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["UnauthorizedBuyer"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackUnauthorizedBuyerError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["UnauthorizedSeller"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackUnauthorizedSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["ZeroEscrowBalance"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackZeroEscrowBalanceError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PaymentprocessorAlreadyInitialized represents a AlreadyInitialized error raised by the Paymentprocessor contract.
type PaymentprocessorAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInitialized()
func PaymentprocessorAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0x0dc149f07762891dbcea3fe72770f3d63a1863fc54b2f084e8c59ec476996927")
}

// UnpackAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInitialized()
func (paymentprocessor *Paymentprocessor) UnpackAlreadyInitializedError(raw []byte) (*PaymentprocessorAlreadyInitialized, error) {
	out := new(PaymentprocessorAlreadyInitialized)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "AlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorAlreadyRefunded represents a AlreadyRefunded error raised by the Paymentprocessor contract.
type PaymentprocessorAlreadyRefunded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyRefunded()
func PaymentprocessorAlreadyRefundedErrorID() common.Hash {
	return common.HexToHash("0xa85e6f1a941fe3f80fd3dc3c0277c926f4c2a65cbf0e07e2b7c7c061a53e018e")
}

// UnpackAlreadyRefundedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyRefunded()
func (paymentprocessor *Paymentprocessor) UnpackAlreadyRefundedError(raw []byte) (*PaymentprocessorAlreadyRefunded, error) {
	out := new(PaymentprocessorAlreadyRefunded)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "AlreadyRefunded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorCancelationRequestDeadlinePassed represents a CancelationRequestDeadlinePassed error raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequestDeadlinePassed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CancelationRequestDeadlinePassed()
func PaymentprocessorCancelationRequestDeadlinePassedErrorID() common.Hash {
	return common.HexToHash("0x7b8558d134e93a5746dc0f6f1c25f0c61fc7e5c54d855412b14dbf8e347a5c3c")
}

// UnpackCancelationRequestDeadlinePassedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CancelationRequestDeadlinePassed()
func (paymentprocessor *Paymentprocessor) UnpackCancelationRequestDeadlinePassedError(raw []byte) (*PaymentprocessorCancelationRequestDeadlinePassed, error) {
	out := new(PaymentprocessorCancelationRequestDeadlinePassed)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "CancelationRequestDeadlinePassed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorDisputeWindowExpired represents a DisputeWindowExpired error raised by the Paymentprocessor contract.
type PaymentprocessorDisputeWindowExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DisputeWindowExpired()
func PaymentprocessorDisputeWindowExpiredErrorID() common.Hash {
	return common.HexToHash("0x8d303ff4c212c555f083d7ab98ad2afb3f9e41e51bf19b66e226b2a7a08dccb5")
}

// UnpackDisputeWindowExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DisputeWindowExpired()
func (paymentprocessor *Paymentprocessor) UnpackDisputeWindowExpiredError(raw []byte) (*PaymentprocessorDisputeWindowExpired, error) {
	out := new(PaymentprocessorDisputeWindowExpired)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "DisputeWindowExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorEscrowAddressMismatch represents a EscrowAddressMismatch error raised by the Paymentprocessor contract.
type PaymentprocessorEscrowAddressMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowAddressMismatch()
func PaymentprocessorEscrowAddressMismatchErrorID() common.Hash {
	return common.HexToHash("0x2ea6c622f2e8642920a7e49165f929b4ad33e87c5b424fbf8170f87bea18b715")
}

// UnpackEscrowAddressMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowAddressMismatch()
func (paymentprocessor *Paymentprocessor) UnpackEscrowAddressMismatchError(raw []byte) (*PaymentprocessorEscrowAddressMismatch, error) {
	out := new(PaymentprocessorEscrowAddressMismatch)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "EscrowAddressMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidBuyer represents a InvalidBuyer error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidBuyer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBuyer()
func PaymentprocessorInvalidBuyerErrorID() common.Hash {
	return common.HexToHash("0xb1df0e0690eb45e806488a4d6f7f25217f93203bed1585525f530e489bcd8a73")
}

// UnpackInvalidBuyerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBuyer()
func (paymentprocessor *Paymentprocessor) UnpackInvalidBuyerError(raw []byte) (*PaymentprocessorInvalidBuyer, error) {
	out := new(PaymentprocessorInvalidBuyer)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidBuyer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidDisputeResolution represents a InvalidDisputeResolution error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidDisputeResolution struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDisputeResolution()
func PaymentprocessorInvalidDisputeResolutionErrorID() common.Hash {
	return common.HexToHash("0x34819f908388d0ed594d20c6802439086d46ba2397dca397160a28d8b2bd98b1")
}

// UnpackInvalidDisputeResolutionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDisputeResolution()
func (paymentprocessor *Paymentprocessor) UnpackInvalidDisputeResolutionError(raw []byte) (*PaymentprocessorInvalidDisputeResolution, error) {
	out := new(PaymentprocessorInvalidDisputeResolution)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidDisputeResolution", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidInvoiceState represents a InvalidInvoiceState error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidInvoiceState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInvoiceState()
func PaymentprocessorInvalidInvoiceStateErrorID() common.Hash {
	return common.HexToHash("0x487e4409b34dcf5275ed8908061cfcde1e134270e5620e0eaff4d68605de2cbc")
}

// UnpackInvalidInvoiceStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInvoiceState()
func (paymentprocessor *Paymentprocessor) UnpackInvalidInvoiceStateError(raw []byte) (*PaymentprocessorInvalidInvoiceState, error) {
	out := new(PaymentprocessorInvalidInvoiceState)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidInvoiceState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidMetaInvoicePayment represents a InvalidMetaInvoicePayment error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidMetaInvoicePayment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMetaInvoicePayment()
func PaymentprocessorInvalidMetaInvoicePaymentErrorID() common.Hash {
	return common.HexToHash("0x416aaaf8c429a968d675aa8e8264eb3b9fdb45afe5c3500dc019b0d22f250909")
}

// UnpackInvalidMetaInvoicePaymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMetaInvoicePayment()
func (paymentprocessor *Paymentprocessor) UnpackInvalidMetaInvoicePaymentError(raw []byte) (*PaymentprocessorInvalidMetaInvoicePayment, error) {
	out := new(PaymentprocessorInvalidMetaInvoicePayment)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidMetaInvoicePayment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidNativePayment represents a InvalidNativePayment error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidNativePayment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNativePayment()
func PaymentprocessorInvalidNativePaymentErrorID() common.Hash {
	return common.HexToHash("0x214510aac5bc5d45b2314d915edc9aa20e9ec869bcb7e6d50d8d068658a871c9")
}

// UnpackInvalidNativePaymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNativePayment()
func (paymentprocessor *Paymentprocessor) UnpackInvalidNativePaymentError(raw []byte) (*PaymentprocessorInvalidNativePayment, error) {
	out := new(PaymentprocessorInvalidNativePayment)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidNativePayment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidPaymentToken represents a InvalidPaymentToken error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidPaymentToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentToken()
func PaymentprocessorInvalidPaymentTokenErrorID() common.Hash {
	return common.HexToHash("0x56e7ec5fb32a6b8680de2f9f0c9db77ebacbce2dde5a21c3bd1e84f57d51c79c")
}

// UnpackInvalidPaymentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentToken()
func (paymentprocessor *Paymentprocessor) UnpackInvalidPaymentTokenError(raw []byte) (*PaymentprocessorInvalidPaymentToken, error) {
	out := new(PaymentprocessorInvalidPaymentToken)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidPaymentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvalidSellersPayoutShare represents a InvalidSellersPayoutShare error raised by the Paymentprocessor contract.
type PaymentprocessorInvalidSellersPayoutShare struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSellersPayoutShare()
func PaymentprocessorInvalidSellersPayoutShareErrorID() common.Hash {
	return common.HexToHash("0x453fb42ddfd1ecde870e9bd55d8b7f21b2333b613ee68779ba5f60498951666d")
}

// UnpackInvalidSellersPayoutShareError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSellersPayoutShare()
func (paymentprocessor *Paymentprocessor) UnpackInvalidSellersPayoutShareError(raw []byte) (*PaymentprocessorInvalidSellersPayoutShare, error) {
	out := new(PaymentprocessorInvalidSellersPayoutShare)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvalidSellersPayoutShare", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvoiceDoesNotExist represents a InvoiceDoesNotExist error raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceDoesNotExist()
func PaymentprocessorInvoiceDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x715d9228f420b2c4c07281fb8597619f2ca0c9d8cada84ce032e60ec6407b582")
}

// UnpackInvoiceDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceDoesNotExist()
func (paymentprocessor *Paymentprocessor) UnpackInvoiceDoesNotExistError(raw []byte) (*PaymentprocessorInvoiceDoesNotExist, error) {
	out := new(PaymentprocessorInvoiceDoesNotExist)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvoiceDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvoiceExpired represents a InvoiceExpired error raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceExpired()
func PaymentprocessorInvoiceExpiredErrorID() common.Hash {
	return common.HexToHash("0xf04e9cf09371be6ef375f7f016c1ac94b6c5c6a4d247bec67ae9568dd6b911b6")
}

// UnpackInvoiceExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceExpired()
func (paymentprocessor *Paymentprocessor) UnpackInvoiceExpiredError(raw []byte) (*PaymentprocessorInvoiceExpired, error) {
	out := new(PaymentprocessorInvoiceExpired)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvoiceExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvoiceResponseTimeExpired represents a InvoiceResponseTimeExpired error raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceResponseTimeExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceResponseTimeExpired()
func PaymentprocessorInvoiceResponseTimeExpiredErrorID() common.Hash {
	return common.HexToHash("0x9cbfeb15eb168d37400472a7d95f2587a354899148722e1e35d30b2b92288648")
}

// UnpackInvoiceResponseTimeExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceResponseTimeExpired()
func (paymentprocessor *Paymentprocessor) UnpackInvoiceResponseTimeExpiredError(raw []byte) (*PaymentprocessorInvoiceResponseTimeExpired, error) {
	out := new(PaymentprocessorInvoiceResponseTimeExpired)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvoiceResponseTimeExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInvoiceStillActive represents a InvoiceStillActive error raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceStillActive struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceStillActive()
func PaymentprocessorInvoiceStillActiveErrorID() common.Hash {
	return common.HexToHash("0xd44e574e39ad7117b3ba847f13b1ca60bb8bf0bea227a8d174246e499b696a89")
}

// UnpackInvoiceStillActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceStillActive()
func (paymentprocessor *Paymentprocessor) UnpackInvoiceStillActiveError(raw []byte) (*PaymentprocessorInvoiceStillActive, error) {
	out := new(PaymentprocessorInvoiceStillActive)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvoiceStillActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorNewOwnerIsZeroAddress represents a NewOwnerIsZeroAddress error raised by the Paymentprocessor contract.
type PaymentprocessorNewOwnerIsZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NewOwnerIsZeroAddress()
func PaymentprocessorNewOwnerIsZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x7448fbae245b5163a637f61fac94c5376c3e155928452ce47ee52d8c1b99587a")
}

// UnpackNewOwnerIsZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NewOwnerIsZeroAddress()
func (paymentprocessor *Paymentprocessor) UnpackNewOwnerIsZeroAddressError(raw []byte) (*PaymentprocessorNewOwnerIsZeroAddress, error) {
	out := new(PaymentprocessorNewOwnerIsZeroAddress)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "NewOwnerIsZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorNoHandoverRequest represents a NoHandoverRequest error raised by the Paymentprocessor contract.
type PaymentprocessorNoHandoverRequest struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoHandoverRequest()
func PaymentprocessorNoHandoverRequestErrorID() common.Hash {
	return common.HexToHash("0x6f5e8818469c73d5be4a0d17c371cde64695907022629c1d064c895f98d466a6")
}

// UnpackNoHandoverRequestError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoHandoverRequest()
func (paymentprocessor *Paymentprocessor) UnpackNoHandoverRequestError(raw []byte) (*PaymentprocessorNoHandoverRequest, error) {
	out := new(PaymentprocessorNoHandoverRequest)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "NoHandoverRequest", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorNoShareAllocatedToBuyer represents a NoShareAllocatedToBuyer error raised by the Paymentprocessor contract.
type PaymentprocessorNoShareAllocatedToBuyer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoShareAllocatedToBuyer()
func PaymentprocessorNoShareAllocatedToBuyerErrorID() common.Hash {
	return common.HexToHash("0xc83071b022f289f51071d8d05f77153f3e13e036b1ed2a069b7fea6ca272a13a")
}

// UnpackNoShareAllocatedToBuyerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoShareAllocatedToBuyer()
func (paymentprocessor *Paymentprocessor) UnpackNoShareAllocatedToBuyerError(raw []byte) (*PaymentprocessorNoShareAllocatedToBuyer, error) {
	out := new(PaymentprocessorNoShareAllocatedToBuyer)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "NoShareAllocatedToBuyer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorNoSubInvoiceCancelled represents a NoSubInvoiceCancelled error raised by the Paymentprocessor contract.
type PaymentprocessorNoSubInvoiceCancelled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoSubInvoiceCancelled()
func PaymentprocessorNoSubInvoiceCancelledErrorID() common.Hash {
	return common.HexToHash("0xfd1abbcf98a6588b06cc7041a3510ea199d0b2386ea129fc6fa9ea5d9b9e2624")
}

// UnpackNoSubInvoiceCancelledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoSubInvoiceCancelled()
func (paymentprocessor *Paymentprocessor) UnpackNoSubInvoiceCancelledError(raw []byte) (*PaymentprocessorNoSubInvoiceCancelled, error) {
	out := new(PaymentprocessorNoSubInvoiceCancelled)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "NoSubInvoiceCancelled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorNotAuthorized represents a NotAuthorized error raised by the Paymentprocessor contract.
type PaymentprocessorNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func PaymentprocessorNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (paymentprocessor *Paymentprocessor) UnpackNotAuthorizedError(raw []byte) (*PaymentprocessorNotAuthorized, error) {
	out := new(PaymentprocessorNotAuthorized)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorUnauthorized represents a Unauthorized error raised by the Paymentprocessor contract.
type PaymentprocessorUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func PaymentprocessorUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (paymentprocessor *Paymentprocessor) UnpackUnauthorizedError(raw []byte) (*PaymentprocessorUnauthorized, error) {
	out := new(PaymentprocessorUnauthorized)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorUnauthorizedBuyer represents a UnauthorizedBuyer error raised by the Paymentprocessor contract.
type PaymentprocessorUnauthorizedBuyer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnauthorizedBuyer()
func PaymentprocessorUnauthorizedBuyerErrorID() common.Hash {
	return common.HexToHash("0x4862d1fcd13f887f321a3de676b3c757033d1a6c68d65fc0d16c6d62066054de")
}

// UnpackUnauthorizedBuyerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnauthorizedBuyer()
func (paymentprocessor *Paymentprocessor) UnpackUnauthorizedBuyerError(raw []byte) (*PaymentprocessorUnauthorizedBuyer, error) {
	out := new(PaymentprocessorUnauthorizedBuyer)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "UnauthorizedBuyer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorUnauthorizedSeller represents a UnauthorizedSeller error raised by the Paymentprocessor contract.
type PaymentprocessorUnauthorizedSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnauthorizedSeller()
func PaymentprocessorUnauthorizedSellerErrorID() common.Hash {
	return common.HexToHash("0x32d2cc20a663e5b503d12c9521fdd9644ba773816f698377739ff938471e3aad")
}

// UnpackUnauthorizedSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnauthorizedSeller()
func (paymentprocessor *Paymentprocessor) UnpackUnauthorizedSellerError(raw []byte) (*PaymentprocessorUnauthorizedSeller, error) {
	out := new(PaymentprocessorUnauthorizedSeller)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "UnauthorizedSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorZeroEscrowBalance represents a ZeroEscrowBalance error raised by the Paymentprocessor contract.
type PaymentprocessorZeroEscrowBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroEscrowBalance()
func PaymentprocessorZeroEscrowBalanceErrorID() common.Hash {
	return common.HexToHash("0x2fda8c5d57ffec561eeac3957a854fbee13c20e88b1e9eb9217b0937dd3a1a36")
}

// UnpackZeroEscrowBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroEscrowBalance()
func (paymentprocessor *Paymentprocessor) UnpackZeroEscrowBalanceError(raw []byte) (*PaymentprocessorZeroEscrowBalance, error) {
	out := new(PaymentprocessorZeroEscrowBalance)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "ZeroEscrowBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}
