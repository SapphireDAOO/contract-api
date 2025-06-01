// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymentprocessor

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
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
var PaymentprocessorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ownerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"marketplaceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nativeTokenAggregatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACCEPTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELATION_ACCEPTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELATION_REJECTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELATION_REQUESTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCELED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_DECIMAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_DISMISSED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_RESOLVED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_SETTLED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INITIATED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REFUNDED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REJECTED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELEASED\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptInvoice\",\"inputs\":[{\"name\":\"invoiceKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"invoiceKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimExpiredInvoiceRefunds\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"computeSalt\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"createDispute\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMetaInvoice\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"param\",\"type\":\"tuple[]\",\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam[]\",\"components\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSingleInvoice\",\"inputs\":[{\"name\":\"param\",\"type\":\"tuple\",\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam\",\"components\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"paidAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"createdAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metaInvoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketplace\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetaInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIAdvancedPaymentProcessor.MetaInvoice\",\"components\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"upper\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetaInvoiceIdForSub\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextMetaInvoiceId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPredictedAddress\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenValueFromUsd\",\"inputs\":[{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleCancelationRequest\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accept\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payMetaInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"paySingleInvoice\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"releasePayment\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestCancelation\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestCancelation\",\"inputs\":[{\"name\":\"invoiceKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"resolution\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"sellerShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMarketplace\",\"inputs\":[{\"name\":\"marketplaceAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalMetaInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalUniqueInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"CancelationRequestHandled\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"accepted\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CancelationRequested\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeCreated\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeDismissed\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeSettled\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"buyerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowCreated\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"escrow\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExpiredInvoiceRefunded\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceAccepted\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCanceled\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCreated\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"invoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"paidAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"createdAt\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"releaseWindow\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"invoiceExpiryDuration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timeBeforeCancelation\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metaInvoiceKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoicePaid\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"escrowAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceRejected\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetaInvoiceCreated\",\"inputs\":[{\"name\":\"metaInvoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"metaInvoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIAdvancedPaymentProcessor.MetaInvoice\",\"components\":[{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"upper\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lower\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invoiceId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReleased\",\"inputs\":[{\"name\":\"invoiceKey\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRefunded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CancelationRequestDeadlinePassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisputeWindowExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBuyer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDisputeResolution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInvoiceState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMetaInvoicePayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNativePayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPaymentToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSellersPayoutShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceResponseTimeExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceStillActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoShareAllocatedToBuyer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoSubInvoiceCancelled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedBuyer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroEscrowBalance\",\"inputs\":[]}]",
}

// PaymentprocessorABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentprocessorMetaData.ABI instead.
var PaymentprocessorABI = PaymentprocessorMetaData.ABI

// Paymentprocessor is an auto generated Go binding around an Ethereum contract.
type Paymentprocessor struct {
	PaymentprocessorCaller     // Read-only binding to the contract
	PaymentprocessorTransactor // Write-only binding to the contract
	PaymentprocessorFilterer   // Log filterer for contract events
}

// PaymentprocessorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentprocessorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentprocessorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentprocessorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentprocessorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentprocessorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentprocessorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentprocessorSession struct {
	Contract     *Paymentprocessor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentprocessorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentprocessorCallerSession struct {
	Contract *PaymentprocessorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PaymentprocessorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentprocessorTransactorSession struct {
	Contract     *PaymentprocessorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PaymentprocessorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentprocessorRaw struct {
	Contract *Paymentprocessor // Generic contract binding to access the raw methods on
}

// PaymentprocessorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentprocessorCallerRaw struct {
	Contract *PaymentprocessorCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentprocessorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentprocessorTransactorRaw struct {
	Contract *PaymentprocessorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentprocessor creates a new instance of Paymentprocessor, bound to a specific deployed contract.
func NewPaymentprocessor(address common.Address, backend bind.ContractBackend) (*Paymentprocessor, error) {
	contract, err := bindPaymentprocessor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paymentprocessor{PaymentprocessorCaller: PaymentprocessorCaller{contract: contract}, PaymentprocessorTransactor: PaymentprocessorTransactor{contract: contract}, PaymentprocessorFilterer: PaymentprocessorFilterer{contract: contract}}, nil
}

// NewPaymentprocessorCaller creates a new read-only instance of Paymentprocessor, bound to a specific deployed contract.
func NewPaymentprocessorCaller(address common.Address, caller bind.ContractCaller) (*PaymentprocessorCaller, error) {
	contract, err := bindPaymentprocessor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorCaller{contract: contract}, nil
}

// NewPaymentprocessorTransactor creates a new write-only instance of Paymentprocessor, bound to a specific deployed contract.
func NewPaymentprocessorTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentprocessorTransactor, error) {
	contract, err := bindPaymentprocessor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorTransactor{contract: contract}, nil
}

// NewPaymentprocessorFilterer creates a new log filterer instance of Paymentprocessor, bound to a specific deployed contract.
func NewPaymentprocessorFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentprocessorFilterer, error) {
	contract, err := bindPaymentprocessor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorFilterer{contract: contract}, nil
}

// bindPaymentprocessor binds a generic wrapper to an already deployed contract.
func bindPaymentprocessor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentprocessorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymentprocessor *PaymentprocessorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymentprocessor.Contract.PaymentprocessorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymentprocessor *PaymentprocessorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.PaymentprocessorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymentprocessor *PaymentprocessorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.PaymentprocessorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paymentprocessor *PaymentprocessorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paymentprocessor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paymentprocessor *PaymentprocessorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paymentprocessor *PaymentprocessorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.contract.Transact(opts, method, params...)
}

// ACCEPTED is a free data retrieval call binding the contract method 0x4324efd7.
//
// Solidity: function ACCEPTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) ACCEPTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "ACCEPTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ACCEPTED is a free data retrieval call binding the contract method 0x4324efd7.
//
// Solidity: function ACCEPTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) ACCEPTED() (uint8, error) {
	return _Paymentprocessor.Contract.ACCEPTED(&_Paymentprocessor.CallOpts)
}

// ACCEPTED is a free data retrieval call binding the contract method 0x4324efd7.
//
// Solidity: function ACCEPTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) ACCEPTED() (uint8, error) {
	return _Paymentprocessor.Contract.ACCEPTED(&_Paymentprocessor.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorSession) BASISPOINTS() (*big.Int, error) {
	return _Paymentprocessor.Contract.BASISPOINTS(&_Paymentprocessor.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCallerSession) BASISPOINTS() (*big.Int, error) {
	return _Paymentprocessor.Contract.BASISPOINTS(&_Paymentprocessor.CallOpts)
}

// CANCELATIONACCEPTED is a free data retrieval call binding the contract method 0x7d2c6615.
//
// Solidity: function CANCELATION_ACCEPTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) CANCELATIONACCEPTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "CANCELATION_ACCEPTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANCELATIONACCEPTED is a free data retrieval call binding the contract method 0x7d2c6615.
//
// Solidity: function CANCELATION_ACCEPTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) CANCELATIONACCEPTED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELATIONACCEPTED(&_Paymentprocessor.CallOpts)
}

// CANCELATIONACCEPTED is a free data retrieval call binding the contract method 0x7d2c6615.
//
// Solidity: function CANCELATION_ACCEPTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) CANCELATIONACCEPTED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELATIONACCEPTED(&_Paymentprocessor.CallOpts)
}

// CANCELATIONREJECTED is a free data retrieval call binding the contract method 0xc9934d4c.
//
// Solidity: function CANCELATION_REJECTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) CANCELATIONREJECTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "CANCELATION_REJECTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANCELATIONREJECTED is a free data retrieval call binding the contract method 0xc9934d4c.
//
// Solidity: function CANCELATION_REJECTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) CANCELATIONREJECTED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELATIONREJECTED(&_Paymentprocessor.CallOpts)
}

// CANCELATIONREJECTED is a free data retrieval call binding the contract method 0xc9934d4c.
//
// Solidity: function CANCELATION_REJECTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) CANCELATIONREJECTED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELATIONREJECTED(&_Paymentprocessor.CallOpts)
}

// CANCELATIONREQUESTED is a free data retrieval call binding the contract method 0xecc224ce.
//
// Solidity: function CANCELATION_REQUESTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) CANCELATIONREQUESTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "CANCELATION_REQUESTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANCELATIONREQUESTED is a free data retrieval call binding the contract method 0xecc224ce.
//
// Solidity: function CANCELATION_REQUESTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) CANCELATIONREQUESTED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELATIONREQUESTED(&_Paymentprocessor.CallOpts)
}

// CANCELATIONREQUESTED is a free data retrieval call binding the contract method 0xecc224ce.
//
// Solidity: function CANCELATION_REQUESTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) CANCELATIONREQUESTED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELATIONREQUESTED(&_Paymentprocessor.CallOpts)
}

// CANCELED is a free data retrieval call binding the contract method 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) CANCELED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "CANCELED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CANCELED is a free data retrieval call binding the contract method 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) CANCELED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELED(&_Paymentprocessor.CallOpts)
}

// CANCELED is a free data retrieval call binding the contract method 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) CANCELED() (uint8, error) {
	return _Paymentprocessor.Contract.CANCELED(&_Paymentprocessor.CallOpts)
}

// DEFAULTDECIMAL is a free data retrieval call binding the contract method 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) DEFAULTDECIMAL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "DEFAULT_DECIMAL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DEFAULTDECIMAL is a free data retrieval call binding the contract method 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) DEFAULTDECIMAL() (uint8, error) {
	return _Paymentprocessor.Contract.DEFAULTDECIMAL(&_Paymentprocessor.CallOpts)
}

// DEFAULTDECIMAL is a free data retrieval call binding the contract method 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) DEFAULTDECIMAL() (uint8, error) {
	return _Paymentprocessor.Contract.DEFAULTDECIMAL(&_Paymentprocessor.CallOpts)
}

// DISPUTED is a free data retrieval call binding the contract method 0x620cac66.
//
// Solidity: function DISPUTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) DISPUTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "DISPUTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DISPUTED is a free data retrieval call binding the contract method 0x620cac66.
//
// Solidity: function DISPUTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) DISPUTED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTED(&_Paymentprocessor.CallOpts)
}

// DISPUTED is a free data retrieval call binding the contract method 0x620cac66.
//
// Solidity: function DISPUTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) DISPUTED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTED(&_Paymentprocessor.CallOpts)
}

// DISPUTEDISMISSED is a free data retrieval call binding the contract method 0xf1201920.
//
// Solidity: function DISPUTE_DISMISSED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) DISPUTEDISMISSED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "DISPUTE_DISMISSED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DISPUTEDISMISSED is a free data retrieval call binding the contract method 0xf1201920.
//
// Solidity: function DISPUTE_DISMISSED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) DISPUTEDISMISSED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTEDISMISSED(&_Paymentprocessor.CallOpts)
}

// DISPUTEDISMISSED is a free data retrieval call binding the contract method 0xf1201920.
//
// Solidity: function DISPUTE_DISMISSED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) DISPUTEDISMISSED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTEDISMISSED(&_Paymentprocessor.CallOpts)
}

// DISPUTERESOLVED is a free data retrieval call binding the contract method 0x677a4ece.
//
// Solidity: function DISPUTE_RESOLVED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) DISPUTERESOLVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "DISPUTE_RESOLVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DISPUTERESOLVED is a free data retrieval call binding the contract method 0x677a4ece.
//
// Solidity: function DISPUTE_RESOLVED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) DISPUTERESOLVED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTERESOLVED(&_Paymentprocessor.CallOpts)
}

// DISPUTERESOLVED is a free data retrieval call binding the contract method 0x677a4ece.
//
// Solidity: function DISPUTE_RESOLVED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) DISPUTERESOLVED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTERESOLVED(&_Paymentprocessor.CallOpts)
}

// DISPUTESETTLED is a free data retrieval call binding the contract method 0x55b9da62.
//
// Solidity: function DISPUTE_SETTLED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) DISPUTESETTLED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "DISPUTE_SETTLED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DISPUTESETTLED is a free data retrieval call binding the contract method 0x55b9da62.
//
// Solidity: function DISPUTE_SETTLED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) DISPUTESETTLED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTESETTLED(&_Paymentprocessor.CallOpts)
}

// DISPUTESETTLED is a free data retrieval call binding the contract method 0x55b9da62.
//
// Solidity: function DISPUTE_SETTLED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) DISPUTESETTLED() (uint8, error) {
	return _Paymentprocessor.Contract.DISPUTESETTLED(&_Paymentprocessor.CallOpts)
}

// INITIATED is a free data retrieval call binding the contract method 0x88a03c53.
//
// Solidity: function INITIATED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) INITIATED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "INITIATED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIATED is a free data retrieval call binding the contract method 0x88a03c53.
//
// Solidity: function INITIATED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) INITIATED() (uint8, error) {
	return _Paymentprocessor.Contract.INITIATED(&_Paymentprocessor.CallOpts)
}

// INITIATED is a free data retrieval call binding the contract method 0x88a03c53.
//
// Solidity: function INITIATED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) INITIATED() (uint8, error) {
	return _Paymentprocessor.Contract.INITIATED(&_Paymentprocessor.CallOpts)
}

// PAID is a free data retrieval call binding the contract method 0xd31a03c2.
//
// Solidity: function PAID() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) PAID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "PAID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAID is a free data retrieval call binding the contract method 0xd31a03c2.
//
// Solidity: function PAID() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) PAID() (uint8, error) {
	return _Paymentprocessor.Contract.PAID(&_Paymentprocessor.CallOpts)
}

// PAID is a free data retrieval call binding the contract method 0xd31a03c2.
//
// Solidity: function PAID() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) PAID() (uint8, error) {
	return _Paymentprocessor.Contract.PAID(&_Paymentprocessor.CallOpts)
}

// REFUNDED is a free data retrieval call binding the contract method 0x94e15c8f.
//
// Solidity: function REFUNDED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) REFUNDED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "REFUNDED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// REFUNDED is a free data retrieval call binding the contract method 0x94e15c8f.
//
// Solidity: function REFUNDED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) REFUNDED() (uint8, error) {
	return _Paymentprocessor.Contract.REFUNDED(&_Paymentprocessor.CallOpts)
}

// REFUNDED is a free data retrieval call binding the contract method 0x94e15c8f.
//
// Solidity: function REFUNDED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) REFUNDED() (uint8, error) {
	return _Paymentprocessor.Contract.REFUNDED(&_Paymentprocessor.CallOpts)
}

// REJECTED is a free data retrieval call binding the contract method 0xcb7390b7.
//
// Solidity: function REJECTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) REJECTED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "REJECTED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// REJECTED is a free data retrieval call binding the contract method 0xcb7390b7.
//
// Solidity: function REJECTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) REJECTED() (uint8, error) {
	return _Paymentprocessor.Contract.REJECTED(&_Paymentprocessor.CallOpts)
}

// REJECTED is a free data retrieval call binding the contract method 0xcb7390b7.
//
// Solidity: function REJECTED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) REJECTED() (uint8, error) {
	return _Paymentprocessor.Contract.REJECTED(&_Paymentprocessor.CallOpts)
}

// RELEASED is a free data retrieval call binding the contract method 0x9b52ea81.
//
// Solidity: function RELEASED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCaller) RELEASED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "RELEASED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RELEASED is a free data retrieval call binding the contract method 0x9b52ea81.
//
// Solidity: function RELEASED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorSession) RELEASED() (uint8, error) {
	return _Paymentprocessor.Contract.RELEASED(&_Paymentprocessor.CallOpts)
}

// RELEASED is a free data retrieval call binding the contract method 0x9b52ea81.
//
// Solidity: function RELEASED() view returns(uint8)
func (_Paymentprocessor *PaymentprocessorCallerSession) RELEASED() (uint8, error) {
	return _Paymentprocessor.Contract.RELEASED(&_Paymentprocessor.CallOpts)
}

// ComputeSalt is a free data retrieval call binding the contract method 0x9c074054.
//
// Solidity: function computeSalt(address seller, address buyer, bytes32 invoiceKey) pure returns(bytes32)
func (_Paymentprocessor *PaymentprocessorCaller) ComputeSalt(opts *bind.CallOpts, seller common.Address, buyer common.Address, invoiceKey [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "computeSalt", seller, buyer, invoiceKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeSalt is a free data retrieval call binding the contract method 0x9c074054.
//
// Solidity: function computeSalt(address seller, address buyer, bytes32 invoiceKey) pure returns(bytes32)
func (_Paymentprocessor *PaymentprocessorSession) ComputeSalt(seller common.Address, buyer common.Address, invoiceKey [32]byte) ([32]byte, error) {
	return _Paymentprocessor.Contract.ComputeSalt(&_Paymentprocessor.CallOpts, seller, buyer, invoiceKey)
}

// ComputeSalt is a free data retrieval call binding the contract method 0x9c074054.
//
// Solidity: function computeSalt(address seller, address buyer, bytes32 invoiceKey) pure returns(bytes32)
func (_Paymentprocessor *PaymentprocessorCallerSession) ComputeSalt(seller common.Address, buyer common.Address, invoiceKey [32]byte) ([32]byte, error) {
	return _Paymentprocessor.Contract.ComputeSalt(&_Paymentprocessor.CallOpts, seller, buyer, invoiceKey)
}

// GetInvoice is a free data retrieval call binding the contract method 0xcb802c8b.
//
// Solidity: function getInvoice(bytes32 invoiceKey) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32))
func (_Paymentprocessor *PaymentprocessorCaller) GetInvoice(opts *bind.CallOpts, invoiceKey [32]byte) (IAdvancedPaymentProcessorInvoice, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getInvoice", invoiceKey)

	if err != nil {
		return *new(IAdvancedPaymentProcessorInvoice), err
	}

	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorInvoice)).(*IAdvancedPaymentProcessorInvoice)

	return out0, err

}

// GetInvoice is a free data retrieval call binding the contract method 0xcb802c8b.
//
// Solidity: function getInvoice(bytes32 invoiceKey) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32))
func (_Paymentprocessor *PaymentprocessorSession) GetInvoice(invoiceKey [32]byte) (IAdvancedPaymentProcessorInvoice, error) {
	return _Paymentprocessor.Contract.GetInvoice(&_Paymentprocessor.CallOpts, invoiceKey)
}

// GetInvoice is a free data retrieval call binding the contract method 0xcb802c8b.
//
// Solidity: function getInvoice(bytes32 invoiceKey) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32))
func (_Paymentprocessor *PaymentprocessorCallerSession) GetInvoice(invoiceKey [32]byte) (IAdvancedPaymentProcessorInvoice, error) {
	return _Paymentprocessor.Contract.GetInvoice(&_Paymentprocessor.CallOpts, invoiceKey)
}

// GetMarketplace is a free data retrieval call binding the contract method 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (_Paymentprocessor *PaymentprocessorCaller) GetMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMarketplace is a free data retrieval call binding the contract method 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (_Paymentprocessor *PaymentprocessorSession) GetMarketplace() (common.Address, error) {
	return _Paymentprocessor.Contract.GetMarketplace(&_Paymentprocessor.CallOpts)
}

// GetMarketplace is a free data retrieval call binding the contract method 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (_Paymentprocessor *PaymentprocessorCallerSession) GetMarketplace() (common.Address, error) {
	return _Paymentprocessor.Contract.GetMarketplace(&_Paymentprocessor.CallOpts)
}

// GetMetaInvoice is a free data retrieval call binding the contract method 0x8ccef592.
//
// Solidity: function getMetaInvoice(bytes32 invoiceKey) view returns((address,uint256,uint256,uint256,address,uint256))
func (_Paymentprocessor *PaymentprocessorCaller) GetMetaInvoice(opts *bind.CallOpts, invoiceKey [32]byte) (IAdvancedPaymentProcessorMetaInvoice, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getMetaInvoice", invoiceKey)

	if err != nil {
		return *new(IAdvancedPaymentProcessorMetaInvoice), err
	}

	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorMetaInvoice)).(*IAdvancedPaymentProcessorMetaInvoice)

	return out0, err

}

// GetMetaInvoice is a free data retrieval call binding the contract method 0x8ccef592.
//
// Solidity: function getMetaInvoice(bytes32 invoiceKey) view returns((address,uint256,uint256,uint256,address,uint256))
func (_Paymentprocessor *PaymentprocessorSession) GetMetaInvoice(invoiceKey [32]byte) (IAdvancedPaymentProcessorMetaInvoice, error) {
	return _Paymentprocessor.Contract.GetMetaInvoice(&_Paymentprocessor.CallOpts, invoiceKey)
}

// GetMetaInvoice is a free data retrieval call binding the contract method 0x8ccef592.
//
// Solidity: function getMetaInvoice(bytes32 invoiceKey) view returns((address,uint256,uint256,uint256,address,uint256))
func (_Paymentprocessor *PaymentprocessorCallerSession) GetMetaInvoice(invoiceKey [32]byte) (IAdvancedPaymentProcessorMetaInvoice, error) {
	return _Paymentprocessor.Contract.GetMetaInvoice(&_Paymentprocessor.CallOpts, invoiceKey)
}

// GetMetaInvoiceIdForSub is a free data retrieval call binding the contract method 0x5cbc9082.
//
// Solidity: function getMetaInvoiceIdForSub(bytes32 invoiceKey) view returns(bytes32)
func (_Paymentprocessor *PaymentprocessorCaller) GetMetaInvoiceIdForSub(opts *bind.CallOpts, invoiceKey [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getMetaInvoiceIdForSub", invoiceKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMetaInvoiceIdForSub is a free data retrieval call binding the contract method 0x5cbc9082.
//
// Solidity: function getMetaInvoiceIdForSub(bytes32 invoiceKey) view returns(bytes32)
func (_Paymentprocessor *PaymentprocessorSession) GetMetaInvoiceIdForSub(invoiceKey [32]byte) ([32]byte, error) {
	return _Paymentprocessor.Contract.GetMetaInvoiceIdForSub(&_Paymentprocessor.CallOpts, invoiceKey)
}

// GetMetaInvoiceIdForSub is a free data retrieval call binding the contract method 0x5cbc9082.
//
// Solidity: function getMetaInvoiceIdForSub(bytes32 invoiceKey) view returns(bytes32)
func (_Paymentprocessor *PaymentprocessorCallerSession) GetMetaInvoiceIdForSub(invoiceKey [32]byte) ([32]byte, error) {
	return _Paymentprocessor.Contract.GetMetaInvoiceIdForSub(&_Paymentprocessor.CallOpts, invoiceKey)
}

// GetNextInvoiceId is a free data retrieval call binding the contract method 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCaller) GetNextInvoiceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getNextInvoiceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextInvoiceId is a free data retrieval call binding the contract method 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorSession) GetNextInvoiceId() (*big.Int, error) {
	return _Paymentprocessor.Contract.GetNextInvoiceId(&_Paymentprocessor.CallOpts)
}

// GetNextInvoiceId is a free data retrieval call binding the contract method 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCallerSession) GetNextInvoiceId() (*big.Int, error) {
	return _Paymentprocessor.Contract.GetNextInvoiceId(&_Paymentprocessor.CallOpts)
}

// GetNextMetaInvoiceId is a free data retrieval call binding the contract method 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCaller) GetNextMetaInvoiceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getNextMetaInvoiceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextMetaInvoiceId is a free data retrieval call binding the contract method 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorSession) GetNextMetaInvoiceId() (*big.Int, error) {
	return _Paymentprocessor.Contract.GetNextMetaInvoiceId(&_Paymentprocessor.CallOpts)
}

// GetNextMetaInvoiceId is a free data retrieval call binding the contract method 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCallerSession) GetNextMetaInvoiceId() (*big.Int, error) {
	return _Paymentprocessor.Contract.GetNextMetaInvoiceId(&_Paymentprocessor.CallOpts)
}

// GetPredictedAddress is a free data retrieval call binding the contract method 0x4bec0c01.
//
// Solidity: function getPredictedAddress(bytes32 salt) view returns(address)
func (_Paymentprocessor *PaymentprocessorCaller) GetPredictedAddress(opts *bind.CallOpts, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getPredictedAddress", salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPredictedAddress is a free data retrieval call binding the contract method 0x4bec0c01.
//
// Solidity: function getPredictedAddress(bytes32 salt) view returns(address)
func (_Paymentprocessor *PaymentprocessorSession) GetPredictedAddress(salt [32]byte) (common.Address, error) {
	return _Paymentprocessor.Contract.GetPredictedAddress(&_Paymentprocessor.CallOpts, salt)
}

// GetPredictedAddress is a free data retrieval call binding the contract method 0x4bec0c01.
//
// Solidity: function getPredictedAddress(bytes32 salt) view returns(address)
func (_Paymentprocessor *PaymentprocessorCallerSession) GetPredictedAddress(salt [32]byte) (common.Address, error) {
	return _Paymentprocessor.Contract.GetPredictedAddress(&_Paymentprocessor.CallOpts, salt)
}

// GetTokenValueFromUsd is a free data retrieval call binding the contract method 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address paymentToken, uint256 price) view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCaller) GetTokenValueFromUsd(opts *bind.CallOpts, paymentToken common.Address, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "getTokenValueFromUsd", paymentToken, price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenValueFromUsd is a free data retrieval call binding the contract method 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address paymentToken, uint256 price) view returns(uint256)
func (_Paymentprocessor *PaymentprocessorSession) GetTokenValueFromUsd(paymentToken common.Address, price *big.Int) (*big.Int, error) {
	return _Paymentprocessor.Contract.GetTokenValueFromUsd(&_Paymentprocessor.CallOpts, paymentToken, price)
}

// GetTokenValueFromUsd is a free data retrieval call binding the contract method 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address paymentToken, uint256 price) view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCallerSession) GetTokenValueFromUsd(paymentToken common.Address, price *big.Int) (*big.Int, error) {
	return _Paymentprocessor.Contract.GetTokenValueFromUsd(&_Paymentprocessor.CallOpts, paymentToken, price)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Paymentprocessor *PaymentprocessorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Paymentprocessor *PaymentprocessorSession) Owner() (common.Address, error) {
	return _Paymentprocessor.Contract.Owner(&_Paymentprocessor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_Paymentprocessor *PaymentprocessorCallerSession) Owner() (common.Address, error) {
	return _Paymentprocessor.Contract.Owner(&_Paymentprocessor.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Paymentprocessor *PaymentprocessorCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Paymentprocessor *PaymentprocessorSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Paymentprocessor.Contract.OwnershipHandoverExpiresAt(&_Paymentprocessor.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_Paymentprocessor *PaymentprocessorCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _Paymentprocessor.Contract.OwnershipHandoverExpiresAt(&_Paymentprocessor.CallOpts, pendingOwner)
}

// PpStorage is a free data retrieval call binding the contract method 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (_Paymentprocessor *PaymentprocessorCaller) PpStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "ppStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PpStorage is a free data retrieval call binding the contract method 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (_Paymentprocessor *PaymentprocessorSession) PpStorage() (common.Address, error) {
	return _Paymentprocessor.Contract.PpStorage(&_Paymentprocessor.CallOpts)
}

// PpStorage is a free data retrieval call binding the contract method 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (_Paymentprocessor *PaymentprocessorCallerSession) PpStorage() (common.Address, error) {
	return _Paymentprocessor.Contract.PpStorage(&_Paymentprocessor.CallOpts)
}

// TotalMetaInvoiceCreated is a free data retrieval call binding the contract method 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCaller) TotalMetaInvoiceCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "totalMetaInvoiceCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMetaInvoiceCreated is a free data retrieval call binding the contract method 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorSession) TotalMetaInvoiceCreated() (*big.Int, error) {
	return _Paymentprocessor.Contract.TotalMetaInvoiceCreated(&_Paymentprocessor.CallOpts)
}

// TotalMetaInvoiceCreated is a free data retrieval call binding the contract method 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCallerSession) TotalMetaInvoiceCreated() (*big.Int, error) {
	return _Paymentprocessor.Contract.TotalMetaInvoiceCreated(&_Paymentprocessor.CallOpts)
}

// TotalUniqueInvoiceCreated is a free data retrieval call binding the contract method 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCaller) TotalUniqueInvoiceCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paymentprocessor.contract.Call(opts, &out, "totalUniqueInvoiceCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUniqueInvoiceCreated is a free data retrieval call binding the contract method 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorSession) TotalUniqueInvoiceCreated() (*big.Int, error) {
	return _Paymentprocessor.Contract.TotalUniqueInvoiceCreated(&_Paymentprocessor.CallOpts)
}

// TotalUniqueInvoiceCreated is a free data retrieval call binding the contract method 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint256)
func (_Paymentprocessor *PaymentprocessorCallerSession) TotalUniqueInvoiceCreated() (*big.Int, error) {
	return _Paymentprocessor.Contract.TotalUniqueInvoiceCreated(&_Paymentprocessor.CallOpts)
}

// AcceptInvoice is a paid mutator transaction binding the contract method 0x40713eb5.
//
// Solidity: function acceptInvoice(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) AcceptInvoice(opts *bind.TransactOpts, invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "acceptInvoice", invoiceKeys)
}

// AcceptInvoice is a paid mutator transaction binding the contract method 0x40713eb5.
//
// Solidity: function acceptInvoice(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorSession) AcceptInvoice(invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.AcceptInvoice(&_Paymentprocessor.TransactOpts, invoiceKeys)
}

// AcceptInvoice is a paid mutator transaction binding the contract method 0x40713eb5.
//
// Solidity: function acceptInvoice(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) AcceptInvoice(invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.AcceptInvoice(&_Paymentprocessor.TransactOpts, invoiceKeys)
}

// AcceptInvoice0 is a paid mutator transaction binding the contract method 0x47a424b5.
//
// Solidity: function acceptInvoice(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) AcceptInvoice0(opts *bind.TransactOpts, invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "acceptInvoice0", invoiceKey)
}

// AcceptInvoice0 is a paid mutator transaction binding the contract method 0x47a424b5.
//
// Solidity: function acceptInvoice(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorSession) AcceptInvoice0(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.AcceptInvoice0(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// AcceptInvoice0 is a paid mutator transaction binding the contract method 0x47a424b5.
//
// Solidity: function acceptInvoice(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) AcceptInvoice0(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.AcceptInvoice0(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// CancelInvoice is a paid mutator transaction binding the contract method 0xbfe06a2b.
//
// Solidity: function cancelInvoice(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) CancelInvoice(opts *bind.TransactOpts, invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "cancelInvoice", invoiceKeys)
}

// CancelInvoice is a paid mutator transaction binding the contract method 0xbfe06a2b.
//
// Solidity: function cancelInvoice(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorSession) CancelInvoice(invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CancelInvoice(&_Paymentprocessor.TransactOpts, invoiceKeys)
}

// CancelInvoice is a paid mutator transaction binding the contract method 0xbfe06a2b.
//
// Solidity: function cancelInvoice(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) CancelInvoice(invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CancelInvoice(&_Paymentprocessor.TransactOpts, invoiceKeys)
}

// CancelInvoice0 is a paid mutator transaction binding the contract method 0xc0aa7e2e.
//
// Solidity: function cancelInvoice(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) CancelInvoice0(opts *bind.TransactOpts, invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "cancelInvoice0", invoiceKey)
}

// CancelInvoice0 is a paid mutator transaction binding the contract method 0xc0aa7e2e.
//
// Solidity: function cancelInvoice(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorSession) CancelInvoice0(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CancelInvoice0(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// CancelInvoice0 is a paid mutator transaction binding the contract method 0xc0aa7e2e.
//
// Solidity: function cancelInvoice(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) CancelInvoice0(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CancelInvoice0(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Paymentprocessor *PaymentprocessorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CancelOwnershipHandover(&_Paymentprocessor.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CancelOwnershipHandover(&_Paymentprocessor.TransactOpts)
}

// ClaimExpiredInvoiceRefunds is a paid mutator transaction binding the contract method 0xa489897b.
//
// Solidity: function claimExpiredInvoiceRefunds(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) ClaimExpiredInvoiceRefunds(opts *bind.TransactOpts, invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "claimExpiredInvoiceRefunds", invoiceKey)
}

// ClaimExpiredInvoiceRefunds is a paid mutator transaction binding the contract method 0xa489897b.
//
// Solidity: function claimExpiredInvoiceRefunds(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorSession) ClaimExpiredInvoiceRefunds(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.ClaimExpiredInvoiceRefunds(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// ClaimExpiredInvoiceRefunds is a paid mutator transaction binding the contract method 0xa489897b.
//
// Solidity: function claimExpiredInvoiceRefunds(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) ClaimExpiredInvoiceRefunds(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.ClaimExpiredInvoiceRefunds(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Paymentprocessor *PaymentprocessorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CompleteOwnershipHandover(&_Paymentprocessor.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CompleteOwnershipHandover(&_Paymentprocessor.TransactOpts, pendingOwner)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x29092dc0.
//
// Solidity: function createDispute(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) CreateDispute(opts *bind.TransactOpts, invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "createDispute", invoiceKey)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x29092dc0.
//
// Solidity: function createDispute(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorSession) CreateDispute(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CreateDispute(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x29092dc0.
//
// Solidity: function createDispute(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) CreateDispute(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CreateDispute(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// CreateMetaInvoice is a paid mutator transaction binding the contract method 0x5b3fd37e.
//
// Solidity: function createMetaInvoice(address buyer, (address,address,uint32,uint32,uint32,uint256)[] param) returns(bytes32)
func (_Paymentprocessor *PaymentprocessorTransactor) CreateMetaInvoice(opts *bind.TransactOpts, buyer common.Address, param []IAdvancedPaymentProcessorInvoiceCreationParam) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "createMetaInvoice", buyer, param)
}

// CreateMetaInvoice is a paid mutator transaction binding the contract method 0x5b3fd37e.
//
// Solidity: function createMetaInvoice(address buyer, (address,address,uint32,uint32,uint32,uint256)[] param) returns(bytes32)
func (_Paymentprocessor *PaymentprocessorSession) CreateMetaInvoice(buyer common.Address, param []IAdvancedPaymentProcessorInvoiceCreationParam) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CreateMetaInvoice(&_Paymentprocessor.TransactOpts, buyer, param)
}

// CreateMetaInvoice is a paid mutator transaction binding the contract method 0x5b3fd37e.
//
// Solidity: function createMetaInvoice(address buyer, (address,address,uint32,uint32,uint32,uint256)[] param) returns(bytes32)
func (_Paymentprocessor *PaymentprocessorTransactorSession) CreateMetaInvoice(buyer common.Address, param []IAdvancedPaymentProcessorInvoiceCreationParam) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CreateMetaInvoice(&_Paymentprocessor.TransactOpts, buyer, param)
}

// CreateSingleInvoice is a paid mutator transaction binding the contract method 0x8a38c874.
//
// Solidity: function createSingleInvoice((address,address,uint32,uint32,uint32,uint256) param) returns(bytes32)
func (_Paymentprocessor *PaymentprocessorTransactor) CreateSingleInvoice(opts *bind.TransactOpts, param IAdvancedPaymentProcessorInvoiceCreationParam) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "createSingleInvoice", param)
}

// CreateSingleInvoice is a paid mutator transaction binding the contract method 0x8a38c874.
//
// Solidity: function createSingleInvoice((address,address,uint32,uint32,uint32,uint256) param) returns(bytes32)
func (_Paymentprocessor *PaymentprocessorSession) CreateSingleInvoice(param IAdvancedPaymentProcessorInvoiceCreationParam) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CreateSingleInvoice(&_Paymentprocessor.TransactOpts, param)
}

// CreateSingleInvoice is a paid mutator transaction binding the contract method 0x8a38c874.
//
// Solidity: function createSingleInvoice((address,address,uint32,uint32,uint32,uint256) param) returns(bytes32)
func (_Paymentprocessor *PaymentprocessorTransactorSession) CreateSingleInvoice(param IAdvancedPaymentProcessorInvoiceCreationParam) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.CreateSingleInvoice(&_Paymentprocessor.TransactOpts, param)
}

// HandleCancelationRequest is a paid mutator transaction binding the contract method 0xa9a6618a.
//
// Solidity: function handleCancelationRequest(bytes32 invoiceKey, bool accept) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) HandleCancelationRequest(opts *bind.TransactOpts, invoiceKey [32]byte, accept bool) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "handleCancelationRequest", invoiceKey, accept)
}

// HandleCancelationRequest is a paid mutator transaction binding the contract method 0xa9a6618a.
//
// Solidity: function handleCancelationRequest(bytes32 invoiceKey, bool accept) returns()
func (_Paymentprocessor *PaymentprocessorSession) HandleCancelationRequest(invoiceKey [32]byte, accept bool) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.HandleCancelationRequest(&_Paymentprocessor.TransactOpts, invoiceKey, accept)
}

// HandleCancelationRequest is a paid mutator transaction binding the contract method 0xa9a6618a.
//
// Solidity: function handleCancelationRequest(bytes32 invoiceKey, bool accept) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) HandleCancelationRequest(invoiceKey [32]byte, accept bool) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.HandleCancelationRequest(&_Paymentprocessor.TransactOpts, invoiceKey, accept)
}

// PayMetaInvoice is a paid mutator transaction binding the contract method 0x97f2ef98.
//
// Solidity: function payMetaInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) PayMetaInvoice(opts *bind.TransactOpts, invoiceKey [32]byte, paymentToken common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "payMetaInvoice", invoiceKey, paymentToken)
}

// PayMetaInvoice is a paid mutator transaction binding the contract method 0x97f2ef98.
//
// Solidity: function payMetaInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (_Paymentprocessor *PaymentprocessorSession) PayMetaInvoice(invoiceKey [32]byte, paymentToken common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.PayMetaInvoice(&_Paymentprocessor.TransactOpts, invoiceKey, paymentToken)
}

// PayMetaInvoice is a paid mutator transaction binding the contract method 0x97f2ef98.
//
// Solidity: function payMetaInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) PayMetaInvoice(invoiceKey [32]byte, paymentToken common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.PayMetaInvoice(&_Paymentprocessor.TransactOpts, invoiceKey, paymentToken)
}

// PaySingleInvoice is a paid mutator transaction binding the contract method 0x91a61baf.
//
// Solidity: function paySingleInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) PaySingleInvoice(opts *bind.TransactOpts, invoiceKey [32]byte, paymentToken common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "paySingleInvoice", invoiceKey, paymentToken)
}

// PaySingleInvoice is a paid mutator transaction binding the contract method 0x91a61baf.
//
// Solidity: function paySingleInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (_Paymentprocessor *PaymentprocessorSession) PaySingleInvoice(invoiceKey [32]byte, paymentToken common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.PaySingleInvoice(&_Paymentprocessor.TransactOpts, invoiceKey, paymentToken)
}

// PaySingleInvoice is a paid mutator transaction binding the contract method 0x91a61baf.
//
// Solidity: function paySingleInvoice(bytes32 invoiceKey, address paymentToken) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) PaySingleInvoice(invoiceKey [32]byte, paymentToken common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.PaySingleInvoice(&_Paymentprocessor.TransactOpts, invoiceKey, paymentToken)
}

// ReleasePayment is a paid mutator transaction binding the contract method 0x7aa1ed58.
//
// Solidity: function releasePayment(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) ReleasePayment(opts *bind.TransactOpts, invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "releasePayment", invoiceKey)
}

// ReleasePayment is a paid mutator transaction binding the contract method 0x7aa1ed58.
//
// Solidity: function releasePayment(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorSession) ReleasePayment(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.ReleasePayment(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// ReleasePayment is a paid mutator transaction binding the contract method 0x7aa1ed58.
//
// Solidity: function releasePayment(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) ReleasePayment(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.ReleasePayment(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Paymentprocessor *PaymentprocessorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RenounceOwnership(&_Paymentprocessor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RenounceOwnership(&_Paymentprocessor.TransactOpts)
}

// RequestCancelation is a paid mutator transaction binding the contract method 0x0982cb6a.
//
// Solidity: function requestCancelation(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) RequestCancelation(opts *bind.TransactOpts, invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "requestCancelation", invoiceKey)
}

// RequestCancelation is a paid mutator transaction binding the contract method 0x0982cb6a.
//
// Solidity: function requestCancelation(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorSession) RequestCancelation(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RequestCancelation(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// RequestCancelation is a paid mutator transaction binding the contract method 0x0982cb6a.
//
// Solidity: function requestCancelation(bytes32 invoiceKey) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) RequestCancelation(invoiceKey [32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RequestCancelation(&_Paymentprocessor.TransactOpts, invoiceKey)
}

// RequestCancelation0 is a paid mutator transaction binding the contract method 0x7afd65c9.
//
// Solidity: function requestCancelation(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) RequestCancelation0(opts *bind.TransactOpts, invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "requestCancelation0", invoiceKeys)
}

// RequestCancelation0 is a paid mutator transaction binding the contract method 0x7afd65c9.
//
// Solidity: function requestCancelation(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorSession) RequestCancelation0(invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RequestCancelation0(&_Paymentprocessor.TransactOpts, invoiceKeys)
}

// RequestCancelation0 is a paid mutator transaction binding the contract method 0x7afd65c9.
//
// Solidity: function requestCancelation(bytes32[] invoiceKeys) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) RequestCancelation0(invoiceKeys [][32]byte) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RequestCancelation0(&_Paymentprocessor.TransactOpts, invoiceKeys)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Paymentprocessor *PaymentprocessorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RequestOwnershipHandover(&_Paymentprocessor.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _Paymentprocessor.Contract.RequestOwnershipHandover(&_Paymentprocessor.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xff258f33.
//
// Solidity: function resolveDispute(bytes32 invoiceKey, uint8 resolution, uint256 sellerShare) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) ResolveDispute(opts *bind.TransactOpts, invoiceKey [32]byte, resolution uint8, sellerShare *big.Int) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "resolveDispute", invoiceKey, resolution, sellerShare)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xff258f33.
//
// Solidity: function resolveDispute(bytes32 invoiceKey, uint8 resolution, uint256 sellerShare) returns()
func (_Paymentprocessor *PaymentprocessorSession) ResolveDispute(invoiceKey [32]byte, resolution uint8, sellerShare *big.Int) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.ResolveDispute(&_Paymentprocessor.TransactOpts, invoiceKey, resolution, sellerShare)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xff258f33.
//
// Solidity: function resolveDispute(bytes32 invoiceKey, uint8 resolution, uint256 sellerShare) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) ResolveDispute(invoiceKey [32]byte, resolution uint8, sellerShare *big.Int) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.ResolveDispute(&_Paymentprocessor.TransactOpts, invoiceKey, resolution, sellerShare)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddress) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) SetMarketplace(opts *bind.TransactOpts, marketplaceAddress common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "setMarketplace", marketplaceAddress)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddress) returns()
func (_Paymentprocessor *PaymentprocessorSession) SetMarketplace(marketplaceAddress common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.SetMarketplace(&_Paymentprocessor.TransactOpts, marketplaceAddress)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddress) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) SetMarketplace(marketplaceAddress common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.SetMarketplace(&_Paymentprocessor.TransactOpts, marketplaceAddress)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address aggregator) returns()
func (_Paymentprocessor *PaymentprocessorTransactor) SetPriceFeed(opts *bind.TransactOpts, token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "setPriceFeed", token, aggregator)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address aggregator) returns()
func (_Paymentprocessor *PaymentprocessorSession) SetPriceFeed(token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.SetPriceFeed(&_Paymentprocessor.TransactOpts, token, aggregator)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address aggregator) returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) SetPriceFeed(token common.Address, aggregator common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.SetPriceFeed(&_Paymentprocessor.TransactOpts, token, aggregator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Paymentprocessor *PaymentprocessorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.TransferOwnership(&_Paymentprocessor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_Paymentprocessor *PaymentprocessorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paymentprocessor.Contract.TransferOwnership(&_Paymentprocessor.TransactOpts, newOwner)
}

// PaymentprocessorCancelationRequestHandledIterator is returned from FilterCancelationRequestHandled and is used to iterate over the raw logs and unpacked data for CancelationRequestHandled events raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequestHandledIterator struct {
	Event *PaymentprocessorCancelationRequestHandled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorCancelationRequestHandledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorCancelationRequestHandled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorCancelationRequestHandled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorCancelationRequestHandledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorCancelationRequestHandledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorCancelationRequestHandled represents a CancelationRequestHandled event raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequestHandled struct {
	InvoiceKey [32]byte
	Accepted   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelationRequestHandled is a free log retrieval operation binding the contract event 0x6ebf681a47fc200fc83d54a7a5263accb4a601d5db5daa81b681b129d622c40a.
//
// Solidity: event CancelationRequestHandled(bytes32 indexed invoiceKey, bool indexed accepted)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterCancelationRequestHandled(opts *bind.FilterOpts, invoiceKey [][32]byte, accepted []bool) (*PaymentprocessorCancelationRequestHandledIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}
	var acceptedRule []interface{}
	for _, acceptedItem := range accepted {
		acceptedRule = append(acceptedRule, acceptedItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "CancelationRequestHandled", invoiceKeyRule, acceptedRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorCancelationRequestHandledIterator{contract: _Paymentprocessor.contract, event: "CancelationRequestHandled", logs: logs, sub: sub}, nil
}

// WatchCancelationRequestHandled is a free log subscription operation binding the contract event 0x6ebf681a47fc200fc83d54a7a5263accb4a601d5db5daa81b681b129d622c40a.
//
// Solidity: event CancelationRequestHandled(bytes32 indexed invoiceKey, bool indexed accepted)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchCancelationRequestHandled(opts *bind.WatchOpts, sink chan<- *PaymentprocessorCancelationRequestHandled, invoiceKey [][32]byte, accepted []bool) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}
	var acceptedRule []interface{}
	for _, acceptedItem := range accepted {
		acceptedRule = append(acceptedRule, acceptedItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "CancelationRequestHandled", invoiceKeyRule, acceptedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorCancelationRequestHandled)
				if err := _Paymentprocessor.contract.UnpackLog(event, "CancelationRequestHandled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelationRequestHandled is a log parse operation binding the contract event 0x6ebf681a47fc200fc83d54a7a5263accb4a601d5db5daa81b681b129d622c40a.
//
// Solidity: event CancelationRequestHandled(bytes32 indexed invoiceKey, bool indexed accepted)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseCancelationRequestHandled(log types.Log) (*PaymentprocessorCancelationRequestHandled, error) {
	event := new(PaymentprocessorCancelationRequestHandled)
	if err := _Paymentprocessor.contract.UnpackLog(event, "CancelationRequestHandled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorCancelationRequestedIterator is returned from FilterCancelationRequested and is used to iterate over the raw logs and unpacked data for CancelationRequested events raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequestedIterator struct {
	Event *PaymentprocessorCancelationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorCancelationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorCancelationRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorCancelationRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorCancelationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorCancelationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorCancelationRequested represents a CancelationRequested event raised by the Paymentprocessor contract.
type PaymentprocessorCancelationRequested struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelationRequested is a free log retrieval operation binding the contract event 0x491f1dc6f9ef24cc7ec22fee8bec5d729058bcae9a184ded225385c86ea45881.
//
// Solidity: event CancelationRequested(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterCancelationRequested(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorCancelationRequestedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "CancelationRequested", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorCancelationRequestedIterator{contract: _Paymentprocessor.contract, event: "CancelationRequested", logs: logs, sub: sub}, nil
}

// WatchCancelationRequested is a free log subscription operation binding the contract event 0x491f1dc6f9ef24cc7ec22fee8bec5d729058bcae9a184ded225385c86ea45881.
//
// Solidity: event CancelationRequested(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchCancelationRequested(opts *bind.WatchOpts, sink chan<- *PaymentprocessorCancelationRequested, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "CancelationRequested", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorCancelationRequested)
				if err := _Paymentprocessor.contract.UnpackLog(event, "CancelationRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelationRequested is a log parse operation binding the contract event 0x491f1dc6f9ef24cc7ec22fee8bec5d729058bcae9a184ded225385c86ea45881.
//
// Solidity: event CancelationRequested(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseCancelationRequested(log types.Log) (*PaymentprocessorCancelationRequested, error) {
	event := new(PaymentprocessorCancelationRequested)
	if err := _Paymentprocessor.contract.UnpackLog(event, "CancelationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorDisputeCreatedIterator is returned from FilterDisputeCreated and is used to iterate over the raw logs and unpacked data for DisputeCreated events raised by the Paymentprocessor contract.
type PaymentprocessorDisputeCreatedIterator struct {
	Event *PaymentprocessorDisputeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorDisputeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorDisputeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorDisputeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorDisputeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorDisputeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorDisputeCreated represents a DisputeCreated event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeCreated struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDisputeCreated is a free log retrieval operation binding the contract event 0xc5e88e0e17c4064b281ccf73a1bebb1b2c329fd5815939182883f85e7fa67d33.
//
// Solidity: event DisputeCreated(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterDisputeCreated(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorDisputeCreatedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "DisputeCreated", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorDisputeCreatedIterator{contract: _Paymentprocessor.contract, event: "DisputeCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCreated is a free log subscription operation binding the contract event 0xc5e88e0e17c4064b281ccf73a1bebb1b2c329fd5815939182883f85e7fa67d33.
//
// Solidity: event DisputeCreated(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchDisputeCreated(opts *bind.WatchOpts, sink chan<- *PaymentprocessorDisputeCreated, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "DisputeCreated", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorDisputeCreated)
				if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeCreated is a log parse operation binding the contract event 0xc5e88e0e17c4064b281ccf73a1bebb1b2c329fd5815939182883f85e7fa67d33.
//
// Solidity: event DisputeCreated(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseDisputeCreated(log types.Log) (*PaymentprocessorDisputeCreated, error) {
	event := new(PaymentprocessorDisputeCreated)
	if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorDisputeDismissedIterator is returned from FilterDisputeDismissed and is used to iterate over the raw logs and unpacked data for DisputeDismissed events raised by the Paymentprocessor contract.
type PaymentprocessorDisputeDismissedIterator struct {
	Event *PaymentprocessorDisputeDismissed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorDisputeDismissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorDisputeDismissed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorDisputeDismissed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorDisputeDismissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorDisputeDismissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorDisputeDismissed represents a DisputeDismissed event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeDismissed struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDisputeDismissed is a free log retrieval operation binding the contract event 0xbf6c13f05b6f9e66d7144e733c70a2d1cb6460d68f006e4d8084b5ec165faf24.
//
// Solidity: event DisputeDismissed(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterDisputeDismissed(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorDisputeDismissedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "DisputeDismissed", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorDisputeDismissedIterator{contract: _Paymentprocessor.contract, event: "DisputeDismissed", logs: logs, sub: sub}, nil
}

// WatchDisputeDismissed is a free log subscription operation binding the contract event 0xbf6c13f05b6f9e66d7144e733c70a2d1cb6460d68f006e4d8084b5ec165faf24.
//
// Solidity: event DisputeDismissed(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchDisputeDismissed(opts *bind.WatchOpts, sink chan<- *PaymentprocessorDisputeDismissed, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "DisputeDismissed", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorDisputeDismissed)
				if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeDismissed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeDismissed is a log parse operation binding the contract event 0xbf6c13f05b6f9e66d7144e733c70a2d1cb6460d68f006e4d8084b5ec165faf24.
//
// Solidity: event DisputeDismissed(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseDisputeDismissed(log types.Log) (*PaymentprocessorDisputeDismissed, error) {
	event := new(PaymentprocessorDisputeDismissed)
	if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeDismissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the Paymentprocessor contract.
type PaymentprocessorDisputeResolvedIterator struct {
	Event *PaymentprocessorDisputeResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorDisputeResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorDisputeResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorDisputeResolved represents a DisputeResolved event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeResolved struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x65e0c7182ce84cd9087c1b07dc4b65875578877b885848e4be19ee312f2c3d31.
//
// Solidity: event DisputeResolved(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterDisputeResolved(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorDisputeResolvedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "DisputeResolved", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorDisputeResolvedIterator{contract: _Paymentprocessor.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x65e0c7182ce84cd9087c1b07dc4b65875578877b885848e4be19ee312f2c3d31.
//
// Solidity: event DisputeResolved(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *PaymentprocessorDisputeResolved, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "DisputeResolved", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorDisputeResolved)
				if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolved is a log parse operation binding the contract event 0x65e0c7182ce84cd9087c1b07dc4b65875578877b885848e4be19ee312f2c3d31.
//
// Solidity: event DisputeResolved(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseDisputeResolved(log types.Log) (*PaymentprocessorDisputeResolved, error) {
	event := new(PaymentprocessorDisputeResolved)
	if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorDisputeSettledIterator is returned from FilterDisputeSettled and is used to iterate over the raw logs and unpacked data for DisputeSettled events raised by the Paymentprocessor contract.
type PaymentprocessorDisputeSettledIterator struct {
	Event *PaymentprocessorDisputeSettled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorDisputeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorDisputeSettled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorDisputeSettled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorDisputeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorDisputeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorDisputeSettled represents a DisputeSettled event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeSettled struct {
	InvoiceKey   [32]byte
	SellerAmount *big.Int
	BuyerAmount  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDisputeSettled is a free log retrieval operation binding the contract event 0x40c5d1853d4b0104b43d303fb54dbefcd3a41664b94c440aa0a45d96a30a2985.
//
// Solidity: event DisputeSettled(bytes32 indexed invoiceKey, uint256 sellerAmount, uint256 buyerAmount)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterDisputeSettled(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorDisputeSettledIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "DisputeSettled", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorDisputeSettledIterator{contract: _Paymentprocessor.contract, event: "DisputeSettled", logs: logs, sub: sub}, nil
}

// WatchDisputeSettled is a free log subscription operation binding the contract event 0x40c5d1853d4b0104b43d303fb54dbefcd3a41664b94c440aa0a45d96a30a2985.
//
// Solidity: event DisputeSettled(bytes32 indexed invoiceKey, uint256 sellerAmount, uint256 buyerAmount)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchDisputeSettled(opts *bind.WatchOpts, sink chan<- *PaymentprocessorDisputeSettled, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "DisputeSettled", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorDisputeSettled)
				if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeSettled is a log parse operation binding the contract event 0x40c5d1853d4b0104b43d303fb54dbefcd3a41664b94c440aa0a45d96a30a2985.
//
// Solidity: event DisputeSettled(bytes32 indexed invoiceKey, uint256 sellerAmount, uint256 buyerAmount)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseDisputeSettled(log types.Log) (*PaymentprocessorDisputeSettled, error) {
	event := new(PaymentprocessorDisputeSettled)
	if err := _Paymentprocessor.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorEscrowCreatedIterator is returned from FilterEscrowCreated and is used to iterate over the raw logs and unpacked data for EscrowCreated events raised by the Paymentprocessor contract.
type PaymentprocessorEscrowCreatedIterator struct {
	Event *PaymentprocessorEscrowCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorEscrowCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorEscrowCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorEscrowCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorEscrowCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorEscrowCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorEscrowCreated represents a EscrowCreated event raised by the Paymentprocessor contract.
type PaymentprocessorEscrowCreated struct {
	InvoiceKey [32]byte
	Escrow     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEscrowCreated is a free log retrieval operation binding the contract event 0x081d5782a95858ba84888c64fd6c1379fe10f4d1f7e42def7585666088eb3240.
//
// Solidity: event EscrowCreated(bytes32 indexed invoiceKey, address indexed escrow)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterEscrowCreated(opts *bind.FilterOpts, invoiceKey [][32]byte, escrow []common.Address) (*PaymentprocessorEscrowCreatedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}
	var escrowRule []interface{}
	for _, escrowItem := range escrow {
		escrowRule = append(escrowRule, escrowItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "EscrowCreated", invoiceKeyRule, escrowRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorEscrowCreatedIterator{contract: _Paymentprocessor.contract, event: "EscrowCreated", logs: logs, sub: sub}, nil
}

// WatchEscrowCreated is a free log subscription operation binding the contract event 0x081d5782a95858ba84888c64fd6c1379fe10f4d1f7e42def7585666088eb3240.
//
// Solidity: event EscrowCreated(bytes32 indexed invoiceKey, address indexed escrow)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchEscrowCreated(opts *bind.WatchOpts, sink chan<- *PaymentprocessorEscrowCreated, invoiceKey [][32]byte, escrow []common.Address) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}
	var escrowRule []interface{}
	for _, escrowItem := range escrow {
		escrowRule = append(escrowRule, escrowItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "EscrowCreated", invoiceKeyRule, escrowRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorEscrowCreated)
				if err := _Paymentprocessor.contract.UnpackLog(event, "EscrowCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEscrowCreated is a log parse operation binding the contract event 0x081d5782a95858ba84888c64fd6c1379fe10f4d1f7e42def7585666088eb3240.
//
// Solidity: event EscrowCreated(bytes32 indexed invoiceKey, address indexed escrow)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseEscrowCreated(log types.Log) (*PaymentprocessorEscrowCreated, error) {
	event := new(PaymentprocessorEscrowCreated)
	if err := _Paymentprocessor.contract.UnpackLog(event, "EscrowCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorExpiredInvoiceRefundedIterator is returned from FilterExpiredInvoiceRefunded and is used to iterate over the raw logs and unpacked data for ExpiredInvoiceRefunded events raised by the Paymentprocessor contract.
type PaymentprocessorExpiredInvoiceRefundedIterator struct {
	Event *PaymentprocessorExpiredInvoiceRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorExpiredInvoiceRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorExpiredInvoiceRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorExpiredInvoiceRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorExpiredInvoiceRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorExpiredInvoiceRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorExpiredInvoiceRefunded represents a ExpiredInvoiceRefunded event raised by the Paymentprocessor contract.
type PaymentprocessorExpiredInvoiceRefunded struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExpiredInvoiceRefunded is a free log retrieval operation binding the contract event 0x87ae852bba115985a2e1b7f5fba7103a428408e9b37341432485cf04da9206d1.
//
// Solidity: event ExpiredInvoiceRefunded(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterExpiredInvoiceRefunded(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorExpiredInvoiceRefundedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "ExpiredInvoiceRefunded", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorExpiredInvoiceRefundedIterator{contract: _Paymentprocessor.contract, event: "ExpiredInvoiceRefunded", logs: logs, sub: sub}, nil
}

// WatchExpiredInvoiceRefunded is a free log subscription operation binding the contract event 0x87ae852bba115985a2e1b7f5fba7103a428408e9b37341432485cf04da9206d1.
//
// Solidity: event ExpiredInvoiceRefunded(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchExpiredInvoiceRefunded(opts *bind.WatchOpts, sink chan<- *PaymentprocessorExpiredInvoiceRefunded, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "ExpiredInvoiceRefunded", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorExpiredInvoiceRefunded)
				if err := _Paymentprocessor.contract.UnpackLog(event, "ExpiredInvoiceRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExpiredInvoiceRefunded is a log parse operation binding the contract event 0x87ae852bba115985a2e1b7f5fba7103a428408e9b37341432485cf04da9206d1.
//
// Solidity: event ExpiredInvoiceRefunded(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseExpiredInvoiceRefunded(log types.Log) (*PaymentprocessorExpiredInvoiceRefunded, error) {
	event := new(PaymentprocessorExpiredInvoiceRefunded)
	if err := _Paymentprocessor.contract.UnpackLog(event, "ExpiredInvoiceRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorInvoiceAcceptedIterator is returned from FilterInvoiceAccepted and is used to iterate over the raw logs and unpacked data for InvoiceAccepted events raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceAcceptedIterator struct {
	Event *PaymentprocessorInvoiceAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorInvoiceAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorInvoiceAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorInvoiceAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorInvoiceAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorInvoiceAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorInvoiceAccepted represents a InvoiceAccepted event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceAccepted struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvoiceAccepted is a free log retrieval operation binding the contract event 0x1e0b06e7fe90c2cf1faf1136cf60fb577b3df63730281815c5f0d8545e5f03d2.
//
// Solidity: event InvoiceAccepted(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterInvoiceAccepted(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorInvoiceAcceptedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "InvoiceAccepted", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorInvoiceAcceptedIterator{contract: _Paymentprocessor.contract, event: "InvoiceAccepted", logs: logs, sub: sub}, nil
}

// WatchInvoiceAccepted is a free log subscription operation binding the contract event 0x1e0b06e7fe90c2cf1faf1136cf60fb577b3df63730281815c5f0d8545e5f03d2.
//
// Solidity: event InvoiceAccepted(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchInvoiceAccepted(opts *bind.WatchOpts, sink chan<- *PaymentprocessorInvoiceAccepted, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "InvoiceAccepted", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorInvoiceAccepted)
				if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvoiceAccepted is a log parse operation binding the contract event 0x1e0b06e7fe90c2cf1faf1136cf60fb577b3df63730281815c5f0d8545e5f03d2.
//
// Solidity: event InvoiceAccepted(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseInvoiceAccepted(log types.Log) (*PaymentprocessorInvoiceAccepted, error) {
	event := new(PaymentprocessorInvoiceAccepted)
	if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorInvoiceCanceledIterator is returned from FilterInvoiceCanceled and is used to iterate over the raw logs and unpacked data for InvoiceCanceled events raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCanceledIterator struct {
	Event *PaymentprocessorInvoiceCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorInvoiceCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorInvoiceCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorInvoiceCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorInvoiceCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorInvoiceCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCanceled struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvoiceCanceled is a free log retrieval operation binding the contract event 0xfc033046ac827bb58a4208c22eed42c3f50a1cba5925f83820def37179ce0c49.
//
// Solidity: event InvoiceCanceled(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterInvoiceCanceled(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorInvoiceCanceledIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "InvoiceCanceled", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorInvoiceCanceledIterator{contract: _Paymentprocessor.contract, event: "InvoiceCanceled", logs: logs, sub: sub}, nil
}

// WatchInvoiceCanceled is a free log subscription operation binding the contract event 0xfc033046ac827bb58a4208c22eed42c3f50a1cba5925f83820def37179ce0c49.
//
// Solidity: event InvoiceCanceled(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchInvoiceCanceled(opts *bind.WatchOpts, sink chan<- *PaymentprocessorInvoiceCanceled, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "InvoiceCanceled", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorInvoiceCanceled)
				if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvoiceCanceled is a log parse operation binding the contract event 0xfc033046ac827bb58a4208c22eed42c3f50a1cba5925f83820def37179ce0c49.
//
// Solidity: event InvoiceCanceled(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseInvoiceCanceled(log types.Log) (*PaymentprocessorInvoiceCanceled, error) {
	event := new(PaymentprocessorInvoiceCanceled)
	if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorInvoiceCreatedIterator is returned from FilterInvoiceCreated and is used to iterate over the raw logs and unpacked data for InvoiceCreated events raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCreatedIterator struct {
	Event *PaymentprocessorInvoiceCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorInvoiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorInvoiceCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorInvoiceCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorInvoiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorInvoiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorInvoiceCreated represents a InvoiceCreated event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCreated struct {
	InvoiceKey [32]byte
	Invoice    IAdvancedPaymentProcessorInvoice
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvoiceCreated is a free log retrieval operation binding the contract event 0x39193a191e4a22bc3013d1a72e6589a03501b16d14d47daf766689ef9a9e8dad.
//
// Solidity: event InvoiceCreated(bytes32 indexed invoiceKey, (uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32) invoice)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterInvoiceCreated(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorInvoiceCreatedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "InvoiceCreated", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorInvoiceCreatedIterator{contract: _Paymentprocessor.contract, event: "InvoiceCreated", logs: logs, sub: sub}, nil
}

// WatchInvoiceCreated is a free log subscription operation binding the contract event 0x39193a191e4a22bc3013d1a72e6589a03501b16d14d47daf766689ef9a9e8dad.
//
// Solidity: event InvoiceCreated(bytes32 indexed invoiceKey, (uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32) invoice)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchInvoiceCreated(opts *bind.WatchOpts, sink chan<- *PaymentprocessorInvoiceCreated, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "InvoiceCreated", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorInvoiceCreated)
				if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvoiceCreated is a log parse operation binding the contract event 0x39193a191e4a22bc3013d1a72e6589a03501b16d14d47daf766689ef9a9e8dad.
//
// Solidity: event InvoiceCreated(bytes32 indexed invoiceKey, (uint256,address,address,address,address,uint8,uint48,uint48,uint32,uint32,uint32,uint256,uint256,bytes32) invoice)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseInvoiceCreated(log types.Log) (*PaymentprocessorInvoiceCreated, error) {
	event := new(PaymentprocessorInvoiceCreated)
	if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorInvoicePaidIterator is returned from FilterInvoicePaid and is used to iterate over the raw logs and unpacked data for InvoicePaid events raised by the Paymentprocessor contract.
type PaymentprocessorInvoicePaidIterator struct {
	Event *PaymentprocessorInvoicePaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorInvoicePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorInvoicePaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorInvoicePaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorInvoicePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorInvoicePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorInvoicePaid represents a InvoicePaid event raised by the Paymentprocessor contract.
type PaymentprocessorInvoicePaid struct {
	InvoiceKey    [32]byte
	PaymentToken  common.Address
	EscrowAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvoicePaid is a free log retrieval operation binding the contract event 0x787fcac3b50ab9534e1ef2e289dfa2a75eb9481de9f9061d4773232074751582.
//
// Solidity: event InvoicePaid(bytes32 indexed invoiceKey, address paymentToken, address escrowAddress, uint256 amount)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterInvoicePaid(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorInvoicePaidIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "InvoicePaid", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorInvoicePaidIterator{contract: _Paymentprocessor.contract, event: "InvoicePaid", logs: logs, sub: sub}, nil
}

// WatchInvoicePaid is a free log subscription operation binding the contract event 0x787fcac3b50ab9534e1ef2e289dfa2a75eb9481de9f9061d4773232074751582.
//
// Solidity: event InvoicePaid(bytes32 indexed invoiceKey, address paymentToken, address escrowAddress, uint256 amount)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchInvoicePaid(opts *bind.WatchOpts, sink chan<- *PaymentprocessorInvoicePaid, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "InvoicePaid", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorInvoicePaid)
				if err := _Paymentprocessor.contract.UnpackLog(event, "InvoicePaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvoicePaid is a log parse operation binding the contract event 0x787fcac3b50ab9534e1ef2e289dfa2a75eb9481de9f9061d4773232074751582.
//
// Solidity: event InvoicePaid(bytes32 indexed invoiceKey, address paymentToken, address escrowAddress, uint256 amount)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseInvoicePaid(log types.Log) (*PaymentprocessorInvoicePaid, error) {
	event := new(PaymentprocessorInvoicePaid)
	if err := _Paymentprocessor.contract.UnpackLog(event, "InvoicePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorInvoiceRejectedIterator is returned from FilterInvoiceRejected and is used to iterate over the raw logs and unpacked data for InvoiceRejected events raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceRejectedIterator struct {
	Event *PaymentprocessorInvoiceRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorInvoiceRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorInvoiceRejected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorInvoiceRejected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorInvoiceRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorInvoiceRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorInvoiceRejected represents a InvoiceRejected event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceRejected struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvoiceRejected is a free log retrieval operation binding the contract event 0xeef85a536407ac1c854e29c217702ac7a2ca04559011f6eaddb441cb82400f03.
//
// Solidity: event InvoiceRejected(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterInvoiceRejected(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorInvoiceRejectedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "InvoiceRejected", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorInvoiceRejectedIterator{contract: _Paymentprocessor.contract, event: "InvoiceRejected", logs: logs, sub: sub}, nil
}

// WatchInvoiceRejected is a free log subscription operation binding the contract event 0xeef85a536407ac1c854e29c217702ac7a2ca04559011f6eaddb441cb82400f03.
//
// Solidity: event InvoiceRejected(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchInvoiceRejected(opts *bind.WatchOpts, sink chan<- *PaymentprocessorInvoiceRejected, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "InvoiceRejected", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorInvoiceRejected)
				if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceRejected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvoiceRejected is a log parse operation binding the contract event 0xeef85a536407ac1c854e29c217702ac7a2ca04559011f6eaddb441cb82400f03.
//
// Solidity: event InvoiceRejected(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseInvoiceRejected(log types.Log) (*PaymentprocessorInvoiceRejected, error) {
	event := new(PaymentprocessorInvoiceRejected)
	if err := _Paymentprocessor.contract.UnpackLog(event, "InvoiceRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorMetaInvoiceCreatedIterator is returned from FilterMetaInvoiceCreated and is used to iterate over the raw logs and unpacked data for MetaInvoiceCreated events raised by the Paymentprocessor contract.
type PaymentprocessorMetaInvoiceCreatedIterator struct {
	Event *PaymentprocessorMetaInvoiceCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorMetaInvoiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorMetaInvoiceCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorMetaInvoiceCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorMetaInvoiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorMetaInvoiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorMetaInvoiceCreated represents a MetaInvoiceCreated event raised by the Paymentprocessor contract.
type PaymentprocessorMetaInvoiceCreated struct {
	MetaInvoiceKey [32]byte
	MetaInvoice    IAdvancedPaymentProcessorMetaInvoice
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMetaInvoiceCreated is a free log retrieval operation binding the contract event 0xc9cb7eda091b1b9f49202f2e1a7cb20ba63d1af8d2dc0f7c294a3c58a263c381.
//
// Solidity: event MetaInvoiceCreated(bytes32 indexed metaInvoiceKey, (address,uint256,uint256,uint256,address,uint256) metaInvoice)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterMetaInvoiceCreated(opts *bind.FilterOpts, metaInvoiceKey [][32]byte) (*PaymentprocessorMetaInvoiceCreatedIterator, error) {

	var metaInvoiceKeyRule []interface{}
	for _, metaInvoiceKeyItem := range metaInvoiceKey {
		metaInvoiceKeyRule = append(metaInvoiceKeyRule, metaInvoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "MetaInvoiceCreated", metaInvoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorMetaInvoiceCreatedIterator{contract: _Paymentprocessor.contract, event: "MetaInvoiceCreated", logs: logs, sub: sub}, nil
}

// WatchMetaInvoiceCreated is a free log subscription operation binding the contract event 0xc9cb7eda091b1b9f49202f2e1a7cb20ba63d1af8d2dc0f7c294a3c58a263c381.
//
// Solidity: event MetaInvoiceCreated(bytes32 indexed metaInvoiceKey, (address,uint256,uint256,uint256,address,uint256) metaInvoice)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchMetaInvoiceCreated(opts *bind.WatchOpts, sink chan<- *PaymentprocessorMetaInvoiceCreated, metaInvoiceKey [][32]byte) (event.Subscription, error) {

	var metaInvoiceKeyRule []interface{}
	for _, metaInvoiceKeyItem := range metaInvoiceKey {
		metaInvoiceKeyRule = append(metaInvoiceKeyRule, metaInvoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "MetaInvoiceCreated", metaInvoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorMetaInvoiceCreated)
				if err := _Paymentprocessor.contract.UnpackLog(event, "MetaInvoiceCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetaInvoiceCreated is a log parse operation binding the contract event 0xc9cb7eda091b1b9f49202f2e1a7cb20ba63d1af8d2dc0f7c294a3c58a263c381.
//
// Solidity: event MetaInvoiceCreated(bytes32 indexed metaInvoiceKey, (address,uint256,uint256,uint256,address,uint256) metaInvoice)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseMetaInvoiceCreated(log types.Log) (*PaymentprocessorMetaInvoiceCreated, error) {
	event := new(PaymentprocessorMetaInvoiceCreated)
	if err := _Paymentprocessor.contract.UnpackLog(event, "MetaInvoiceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipHandoverCanceledIterator struct {
	Event *PaymentprocessorOwnershipHandoverCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorOwnershipHandoverCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorOwnershipHandoverCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*PaymentprocessorOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorOwnershipHandoverCanceledIterator{contract: _Paymentprocessor.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *PaymentprocessorOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorOwnershipHandoverCanceled)
				if err := _Paymentprocessor.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*PaymentprocessorOwnershipHandoverCanceled, error) {
	event := new(PaymentprocessorOwnershipHandoverCanceled)
	if err := _Paymentprocessor.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipHandoverRequestedIterator struct {
	Event *PaymentprocessorOwnershipHandoverRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorOwnershipHandoverRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorOwnershipHandoverRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*PaymentprocessorOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorOwnershipHandoverRequestedIterator{contract: _Paymentprocessor.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *PaymentprocessorOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorOwnershipHandoverRequested)
				if err := _Paymentprocessor.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseOwnershipHandoverRequested(log types.Log) (*PaymentprocessorOwnershipHandoverRequested, error) {
	event := new(PaymentprocessorOwnershipHandoverRequested)
	if err := _Paymentprocessor.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipTransferredIterator struct {
	Event *PaymentprocessorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorOwnershipTransferred represents a OwnershipTransferred event raised by the Paymentprocessor contract.
type PaymentprocessorOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PaymentprocessorOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorOwnershipTransferredIterator{contract: _Paymentprocessor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentprocessorOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorOwnershipTransferred)
				if err := _Paymentprocessor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Paymentprocessor *PaymentprocessorFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentprocessorOwnershipTransferred, error) {
	event := new(PaymentprocessorOwnershipTransferred)
	if err := _Paymentprocessor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentprocessorPaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the Paymentprocessor contract.
type PaymentprocessorPaymentReleasedIterator struct {
	Event *PaymentprocessorPaymentReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentprocessorPaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentprocessorPaymentReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentprocessorPaymentReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentprocessorPaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentprocessorPaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentprocessorPaymentReleased represents a PaymentReleased event raised by the Paymentprocessor contract.
type PaymentprocessorPaymentReleased struct {
	InvoiceKey [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0xc21bc94c1b5d9d43bbd526118faff1ad8ff9147b010a5308c667a679c4309ea3.
//
// Solidity: event PaymentReleased(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) FilterPaymentReleased(opts *bind.FilterOpts, invoiceKey [][32]byte) (*PaymentprocessorPaymentReleasedIterator, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.FilterLogs(opts, "PaymentReleased", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return &PaymentprocessorPaymentReleasedIterator{contract: _Paymentprocessor.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0xc21bc94c1b5d9d43bbd526118faff1ad8ff9147b010a5308c667a679c4309ea3.
//
// Solidity: event PaymentReleased(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *PaymentprocessorPaymentReleased, invoiceKey [][32]byte) (event.Subscription, error) {

	var invoiceKeyRule []interface{}
	for _, invoiceKeyItem := range invoiceKey {
		invoiceKeyRule = append(invoiceKeyRule, invoiceKeyItem)
	}

	logs, sub, err := _Paymentprocessor.contract.WatchLogs(opts, "PaymentReleased", invoiceKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentprocessorPaymentReleased)
				if err := _Paymentprocessor.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReleased is a log parse operation binding the contract event 0xc21bc94c1b5d9d43bbd526118faff1ad8ff9147b010a5308c667a679c4309ea3.
//
// Solidity: event PaymentReleased(bytes32 indexed invoiceKey)
func (_Paymentprocessor *PaymentprocessorFilterer) ParsePaymentReleased(log types.Log) (*PaymentprocessorPaymentReleased, error) {
	event := new(PaymentprocessorPaymentReleased)
	if err := _Paymentprocessor.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
