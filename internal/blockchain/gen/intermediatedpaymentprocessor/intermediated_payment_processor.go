// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package intermediatedpaymentprocessor

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

// IIntermediatedPaymentProcessorInvoice is an auto generated low-level Go binding around an user-defined struct.
type IIntermediatedPaymentProcessorInvoice struct {
	InvoiceNonce      *big.Int
	PaidAt            *big.Int
	CreatedAt         *big.Int
	ReleaseAt         *big.Int
	ExpiresAt         *big.Int
	State             uint8
	WithdrawalRetries uint8
	EscrowHoldPeriod  uint32
	FeeRate           uint16
	MetaInvoiceId     *big.Int
	Buyer             common.Address
	Seller            common.Address
	Escrow            common.Address
	PaymentToken      common.Address
	FeeReceiver       common.Address
	AmountPaid        *big.Int
	Price             *big.Int
	Balance           *big.Int
}

// IIntermediatedPaymentProcessorInvoiceCreationParam is an auto generated low-level Go binding around an user-defined struct.
type IIntermediatedPaymentProcessorInvoiceCreationParam struct {
	InvoiceId        string
	Seller           common.Address
	Price            *big.Int
	EscrowHoldPeriod uint32
	PaymentTokens    []common.Address
}

// IIntermediatedPaymentProcessorMetaInvoice is an auto generated low-level Go binding around an user-defined struct.
type IIntermediatedPaymentProcessorMetaInvoice struct {
	Price         *big.Int
	SubInvoiceIds []*big.Int
}

// IntermediatedpaymentprocessorMetaData contains all meta data concerning the Intermediatedpaymentprocessor contract.
var IntermediatedpaymentprocessorMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_getDecimals\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeSalt\",\"inputs\":[{\"name\":\"_seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"createDispute\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMetaInvoice\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple[]\",\"internalType\":\"structIIntermediatedPaymentProcessor.InvoiceCreationParam[]\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"paymentTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSingleInvoice\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structIIntermediatedPaymentProcessor.InvoiceCreationParam\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"paymentTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"outputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"i\",\"type\":\"tuple\",\"internalType\":\"structIIntermediatedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetaInvoice\",\"inputs\":[{\"name\":\"_metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"m\",\"type\":\"tuple\",\"internalType\":\"structIIntermediatedPaymentProcessor.MetaInvoice\",\"components\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"subInvoiceIds\",\"type\":\"uint216[]\",\"internalType\":\"uint216[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"currentMinimumPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextInvoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextMetaInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextMetaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPredictedAddress\",\"inputs\":[{\"name\":\"_salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"predictedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenValueFromUsd\",\"inputs\":[{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleDispute\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_resolution\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_sellerShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaymentTokenAllowed\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOracleManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"payMetaInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeReceivers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payMetaInvoiceWithValue\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_feeReceivers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refund\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_refundShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInvoiceReleaseTime\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_holdPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumPrice\",\"inputs\":[{\"name\":\"_newMinimumPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOracle\",\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalMetaInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"totalMetaInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalUniqueInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"totalInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DisputeCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeDismissed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeSettled\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"buyerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"escrow\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCanceled\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"invoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIntermediatedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoicePaid\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"escrowAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedPaymentRecovered\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetaInvoiceCreated\",\"inputs\":[{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"totalPrice\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"inputs\":[{\"name\":\"previousOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReleased\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currency\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentTokensRegistered\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"paymentTokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Refunded\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateReleaseTime\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"newHoldPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BuyerCannotBeSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMetaInvoice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowWithdrawFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeReceiverCountMismatch\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"HoldPeriodCanNotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDisputeResolution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeAuthorization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInvoiceState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMetaInvoicePaymentAmount\",\"inputs\":[{\"name\":\"sent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNativePayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSellersPayoutShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MetaInvoiceAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPaymentTokens\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentTokenNotAllowed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PriceCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceIsTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencerDown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePriceFeed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[]}]",
	ID:  "Intermediatedpaymentprocessor",
}

// Intermediatedpaymentprocessor is an auto generated Go binding around an Ethereum contract.
type Intermediatedpaymentprocessor struct {
	abi abi.ABI
}

// NewIntermediatedpaymentprocessor creates a new instance of Intermediatedpaymentprocessor.
func NewIntermediatedpaymentprocessor() *Intermediatedpaymentprocessor {
	parsed, err := IntermediatedpaymentprocessorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Intermediatedpaymentprocessor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Intermediatedpaymentprocessor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _paymentProcessorStorageAddress, address _oracle) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackConstructor(_paymentProcessorStorageAddress common.Address, _oracle common.Address) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("", _paymentProcessorStorageAddress, _oracle)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ed8bf86.
//
// Solidity: function _getDecimals(address _token) view returns(uint8 tokenDecimals)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetDecimals(token common.Address) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("_getDecimals", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ed8bf86.
//
// Solidity: function _getDecimals(address _token) view returns(uint8 tokenDecimals)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetDecimals(data []byte) (uint8, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("_getDecimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackCancelInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde48b793.
//
// Solidity: function cancelInvoice(uint216 _invoiceId) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackCancelInvoice(invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("cancelInvoice", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackComputeSalt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a221549.
//
// Solidity: function computeSalt(address _seller, address _buyer, uint216 _invoiceId) pure returns(bytes32 salt)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackComputeSalt(seller common.Address, buyer common.Address, invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("computeSalt", seller, buyer, invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackComputeSalt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a221549.
//
// Solidity: function computeSalt(address _seller, address _buyer, uint216 _invoiceId) pure returns(bytes32 salt)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackComputeSalt(data []byte) ([32]byte, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("computeSalt", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCreateDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6409e54.
//
// Solidity: function createDispute(uint216 _invoiceId) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackCreateDispute(invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("createDispute", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda9f2cad.
//
// Solidity: function createMetaInvoice((string,address,uint256,uint32,address[])[] _param) returns(uint216 metaInvoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackCreateMetaInvoice(param []IIntermediatedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("createMetaInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda9f2cad.
//
// Solidity: function createMetaInvoice((string,address,uint256,uint32,address[])[] _param) returns(uint216 metaInvoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackCreateMetaInvoice(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("createMetaInvoice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCreateSingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69855c9d.
//
// Solidity: function createSingleInvoice((string,address,uint256,uint32,address[]) _param) returns(uint216 invoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackCreateSingleInvoice(param IIntermediatedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("createSingleInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSingleInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x69855c9d.
//
// Solidity: function createSingleInvoice((string,address,uint256,uint32,address[]) _param) returns(uint216 invoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackCreateSingleInvoice(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("createSingleInvoice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cfa3edf.
//
// Solidity: function getInvoice(uint216 _invoiceId) view returns((uint216,uint40,uint40,uint40,uint40,uint8,uint8,uint32,uint16,uint216,address,address,address,address,address,uint256,uint256,uint256) i)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetInvoice(invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getInvoice", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4cfa3edf.
//
// Solidity: function getInvoice(uint216 _invoiceId) view returns((uint216,uint40,uint40,uint40,uint40,uint8,uint8,uint32,uint16,uint216,address,address,address,address,address,uint256,uint256,uint256) i)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetInvoice(data []byte) (IIntermediatedPaymentProcessorInvoice, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getInvoice", data)
	if err != nil {
		return *new(IIntermediatedPaymentProcessorInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IIntermediatedPaymentProcessorInvoice)).(*IIntermediatedPaymentProcessorInvoice)
	return out0, err
}

// PackGetMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33453a3d.
//
// Solidity: function getMetaInvoice(uint216 _metaInvoiceId) view returns((uint256,uint216[]) m)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetMetaInvoice(metaInvoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getMetaInvoice", metaInvoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x33453a3d.
//
// Solidity: function getMetaInvoice(uint216 _metaInvoiceId) view returns((uint256,uint216[]) m)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetMetaInvoice(data []byte) (IIntermediatedPaymentProcessorMetaInvoice, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getMetaInvoice", data)
	if err != nil {
		return *new(IIntermediatedPaymentProcessorMetaInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IIntermediatedPaymentProcessorMetaInvoice)).(*IIntermediatedPaymentProcessorMetaInvoice)
	return out0, err
}

// PackGetMinimumPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac4316cb.
//
// Solidity: function getMinimumPrice() view returns(uint256 currentMinimumPrice)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetMinimumPrice() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getMinimumPrice")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMinimumPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac4316cb.
//
// Solidity: function getMinimumPrice() view returns(uint256 currentMinimumPrice)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetMinimumPrice(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getMinimumPrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetNextInvoiceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonce)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetNextInvoiceNonce() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getNextInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonce)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetNextInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getNextInvoiceNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetNextMetaInvoiceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x540f0f3c.
//
// Solidity: function getNextMetaInvoiceNonce() view returns(uint216 nextMetaInvoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetNextMetaInvoiceNonce() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getNextMetaInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextMetaInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x540f0f3c.
//
// Solidity: function getNextMetaInvoiceNonce() view returns(uint216 nextMetaInvoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetNextMetaInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getNextMetaInvoiceNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetPredictedAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4ed09d9.
//
// Solidity: function getPredictedAddress(bytes32 _salt, uint216 _invoiceId) view returns(address predictedAddress)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetPredictedAddress(salt [32]byte, invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getPredictedAddress", salt, invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPredictedAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc4ed09d9.
//
// Solidity: function getPredictedAddress(bytes32 _salt, uint216 _invoiceId) view returns(address predictedAddress)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetPredictedAddress(data []byte) (common.Address, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getPredictedAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetTokenValueFromUsd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address _paymentToken, uint256 _usdAmount) view returns(uint256 tokenValue)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackGetTokenValueFromUsd(paymentToken common.Address, usdAmount *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("getTokenValueFromUsd", paymentToken, usdAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenValueFromUsd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address _paymentToken, uint256 _usdAmount) view returns(uint256 tokenValue)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackGetTokenValueFromUsd(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("getTokenValueFromUsd", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackHandleDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf00aaec6.
//
// Solidity: function handleDispute(uint216 _invoiceId, uint8 _resolution, uint256 _sellerShare) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackHandleDispute(invoiceId *big.Int, resolution uint8, sellerShare *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("handleDispute", invoiceId, resolution, sellerShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsPaymentTokenAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9da39ea9.
//
// Solidity: function isPaymentTokenAllowed(uint216 _invoiceId, address _paymentToken) view returns(bool allowed)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackIsPaymentTokenAllowed(invoiceId *big.Int, paymentToken common.Address) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("isPaymentTokenAllowed", invoiceId, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsPaymentTokenAllowed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9da39ea9.
//
// Solidity: function isPaymentTokenAllowed(uint216 _invoiceId, address _paymentToken) view returns(bool allowed)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackIsPaymentTokenAllowed(data []byte) (bool, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("isPaymentTokenAllowed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackOracle() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("oracle")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOracle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackOracle(data []byte) (common.Address, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("oracle", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPayInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40b3d7ab.
//
// Solidity: function payInvoice(uint216 _invoiceId, address _paymentToken, address _feeReceiver, bytes _data) payable returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackPayInvoice(invoiceId *big.Int, paymentToken common.Address, feeReceiver common.Address, data []byte) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("payInvoice", invoiceId, paymentToken, feeReceiver, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPayMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7db1208d.
//
// Solidity: function payMetaInvoice(uint216 _invoiceId, address _paymentToken, address[] _feeReceivers, bytes _data) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackPayMetaInvoice(invoiceId *big.Int, paymentToken common.Address, feeReceivers []common.Address, data []byte) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("payMetaInvoice", invoiceId, paymentToken, feeReceivers, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPayMetaInvoiceWithValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a8bf1ba.
//
// Solidity: function payMetaInvoiceWithValue(uint216 _invoiceId, address[] _feeReceivers, bytes _data) payable returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackPayMetaInvoiceWithValue(invoiceId *big.Int, feeReceivers []common.Address, data []byte) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("payMetaInvoiceWithValue", invoiceId, feeReceivers, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackPpStorage() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("ppStorage", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackRefund is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5da97b5e.
//
// Solidity: function refund(uint216 _invoiceId, uint256 _refundShare) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackRefund(invoiceId *big.Int, refundShare *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("refund", invoiceId, refundShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb990809.
//
// Solidity: function release(uint216 _invoiceId) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackRelease(invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("release", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackResolveDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5168e613.
//
// Solidity: function resolveDispute(uint216 _invoiceId) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackResolveDispute(invoiceId *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("resolveDispute", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetInvoiceReleaseTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50e5d5cf.
//
// Solidity: function setInvoiceReleaseTime(uint216 _invoiceId, uint256 _holdPeriod) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackSetInvoiceReleaseTime(invoiceId *big.Int, holdPeriod *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("setInvoiceReleaseTime", invoiceId, holdPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinimumPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25cea976.
//
// Solidity: function setMinimumPrice(uint256 _newMinimumPrice) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackSetMinimumPrice(newMinimumPrice *big.Int) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("setMinimumPrice", newMinimumPrice)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackSetOracle(oracle common.Address) []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("setOracle", oracle)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalMetaInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint216 totalMetaInvoices)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackTotalMetaInvoiceCreated() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("totalMetaInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalMetaInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint216 totalMetaInvoices)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackTotalMetaInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("totalMetaInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTotalUniqueInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint216 totalInvoices)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) PackTotalUniqueInvoiceCreated() []byte {
	enc, err := intermediatedpaymentprocessor.abi.Pack("totalUniqueInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalUniqueInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint216 totalInvoices)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackTotalUniqueInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := intermediatedpaymentprocessor.abi.Unpack("totalUniqueInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// IntermediatedpaymentprocessorDisputeCreated represents a DisputeCreated event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorDisputeCreated struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorDisputeCreatedEventName = "DisputeCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorDisputeCreated) ContractEventName() string {
	return IntermediatedpaymentprocessorDisputeCreatedEventName
}

// UnpackDisputeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeCreated(uint216 indexed invoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackDisputeCreatedEvent(log *types.Log) (*IntermediatedpaymentprocessorDisputeCreated, error) {
	event := "DisputeCreated"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorDisputeCreated)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorDisputeDismissed represents a DisputeDismissed event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorDisputeDismissed struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorDisputeDismissedEventName = "DisputeDismissed"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorDisputeDismissed) ContractEventName() string {
	return IntermediatedpaymentprocessorDisputeDismissedEventName
}

// UnpackDisputeDismissedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeDismissed(uint216 indexed invoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackDisputeDismissedEvent(log *types.Log) (*IntermediatedpaymentprocessorDisputeDismissed, error) {
	event := "DisputeDismissed"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorDisputeDismissed)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorDisputeResolved represents a DisputeResolved event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorDisputeResolved struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorDisputeResolvedEventName = "DisputeResolved"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorDisputeResolved) ContractEventName() string {
	return IntermediatedpaymentprocessorDisputeResolvedEventName
}

// UnpackDisputeResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeResolved(uint216 indexed invoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackDisputeResolvedEvent(log *types.Log) (*IntermediatedpaymentprocessorDisputeResolved, error) {
	event := "DisputeResolved"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorDisputeResolved)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorDisputeSettled represents a DisputeSettled event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorDisputeSettled struct {
	InvoiceId    *big.Int
	SellerAmount *big.Int
	BuyerAmount  *big.Int
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorDisputeSettledEventName = "DisputeSettled"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorDisputeSettled) ContractEventName() string {
	return IntermediatedpaymentprocessorDisputeSettledEventName
}

// UnpackDisputeSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeSettled(uint216 indexed invoiceId, uint256 sellerAmount, uint256 buyerAmount, uint256 fee)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackDisputeSettledEvent(log *types.Log) (*IntermediatedpaymentprocessorDisputeSettled, error) {
	event := "DisputeSettled"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorDisputeSettled)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorEscrowCreated represents a EscrowCreated event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorEscrowCreated struct {
	InvoiceId *big.Int
	Escrow    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorEscrowCreatedEventName = "EscrowCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorEscrowCreated) ContractEventName() string {
	return IntermediatedpaymentprocessorEscrowCreatedEventName
}

// UnpackEscrowCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EscrowCreated(uint216 indexed invoiceId, address indexed escrow)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackEscrowCreatedEvent(log *types.Log) (*IntermediatedpaymentprocessorEscrowCreated, error) {
	event := "EscrowCreated"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorEscrowCreated)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvoiceCanceled struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorInvoiceCanceled) ContractEventName() string {
	return IntermediatedpaymentprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(uint216 indexed invoiceId)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvoiceCanceledEvent(log *types.Log) (*IntermediatedpaymentprocessorInvoiceCanceled, error) {
	event := "InvoiceCanceled"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorInvoiceCanceled)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorInvoiceCreated represents a InvoiceCreated event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvoiceCreated struct {
	InvoiceId *big.Int
	Invoice   IIntermediatedPaymentProcessorInvoice
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorInvoiceCreated) ContractEventName() string {
	return IntermediatedpaymentprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(uint216 indexed invoiceId, (uint216,uint40,uint40,uint40,uint40,uint8,uint8,uint32,uint16,uint216,address,address,address,address,address,uint256,uint256,uint256) invoice)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvoiceCreatedEvent(log *types.Log) (*IntermediatedpaymentprocessorInvoiceCreated, error) {
	event := "InvoiceCreated"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorInvoiceCreated)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorInvoicePaid represents a InvoicePaid event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvoicePaid struct {
	InvoiceId     *big.Int
	PaymentToken  common.Address
	EscrowAddress common.Address
	Amount        *big.Int
	ReleaseAt     *big.Int
	FeeReceiver   common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorInvoicePaidEventName = "InvoicePaid"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorInvoicePaid) ContractEventName() string {
	return IntermediatedpaymentprocessorInvoicePaidEventName
}

// UnpackInvoicePaidEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoicePaid(uint216 indexed invoiceId, address paymentToken, address escrowAddress, uint256 amount, uint40 releaseAt, address feeReceiver)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvoicePaidEvent(log *types.Log) (*IntermediatedpaymentprocessorInvoicePaid, error) {
	event := "InvoicePaid"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorInvoicePaid)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorLockedPaymentRecovered represents a LockedPaymentRecovered event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorLockedPaymentRecovered struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorLockedPaymentRecoveredEventName = "LockedPaymentRecovered"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorLockedPaymentRecovered) ContractEventName() string {
	return IntermediatedpaymentprocessorLockedPaymentRecoveredEventName
}

// UnpackLockedPaymentRecoveredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LockedPaymentRecovered(uint216 indexed invoiceId, address indexed recipient, uint256 amount)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackLockedPaymentRecoveredEvent(log *types.Log) (*IntermediatedpaymentprocessorLockedPaymentRecovered, error) {
	event := "LockedPaymentRecovered"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorLockedPaymentRecovered)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorMetaInvoiceCreated represents a MetaInvoiceCreated event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorMetaInvoiceCreated struct {
	MetaInvoiceId *big.Int
	TotalPrice    *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorMetaInvoiceCreatedEventName = "MetaInvoiceCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorMetaInvoiceCreated) ContractEventName() string {
	return IntermediatedpaymentprocessorMetaInvoiceCreatedEventName
}

// UnpackMetaInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetaInvoiceCreated(uint216 indexed metaInvoiceId, uint256 indexed totalPrice)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackMetaInvoiceCreatedEvent(log *types.Log) (*IntermediatedpaymentprocessorMetaInvoiceCreated, error) {
	event := "MetaInvoiceCreated"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorMetaInvoiceCreated)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorOracleUpdated represents a OracleUpdated event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorOracleUpdated struct {
	PreviousOracle common.Address
	NewOracle      common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorOracleUpdatedEventName = "OracleUpdated"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorOracleUpdated) ContractEventName() string {
	return IntermediatedpaymentprocessorOracleUpdatedEventName
}

// UnpackOracleUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OracleUpdated(address indexed previousOracle, address indexed newOracle)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackOracleUpdatedEvent(log *types.Log) (*IntermediatedpaymentprocessorOracleUpdated, error) {
	event := "OracleUpdated"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorOracleUpdated)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorPaymentReleased represents a PaymentReleased event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorPaymentReleased struct {
	InvoiceId    *big.Int
	Receiver     common.Address
	Currency     common.Address
	SellerAmount *big.Int
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorPaymentReleasedEventName = "PaymentReleased"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorPaymentReleased) ContractEventName() string {
	return IntermediatedpaymentprocessorPaymentReleasedEventName
}

// UnpackPaymentReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentReleased(uint216 indexed invoiceId, address receiver, address currency, uint256 sellerAmount, uint256 fee)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackPaymentReleasedEvent(log *types.Log) (*IntermediatedpaymentprocessorPaymentReleased, error) {
	event := "PaymentReleased"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorPaymentReleased)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorPaymentTokensRegistered represents a PaymentTokensRegistered event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorPaymentTokensRegistered struct {
	InvoiceId     *big.Int
	PaymentTokens []common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorPaymentTokensRegisteredEventName = "PaymentTokensRegistered"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorPaymentTokensRegistered) ContractEventName() string {
	return IntermediatedpaymentprocessorPaymentTokensRegisteredEventName
}

// UnpackPaymentTokensRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentTokensRegistered(uint216 indexed invoiceId, address[] paymentTokens)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackPaymentTokensRegisteredEvent(log *types.Log) (*IntermediatedpaymentprocessorPaymentTokensRegistered, error) {
	event := "PaymentTokensRegistered"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorPaymentTokensRegistered)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorRefunded represents a Refunded event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorRefunded struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorRefundedEventName = "Refunded"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorRefunded) ContractEventName() string {
	return IntermediatedpaymentprocessorRefundedEventName
}

// UnpackRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Refunded(uint216 indexed invoiceId, uint256 indexed amount)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackRefundedEvent(log *types.Log) (*IntermediatedpaymentprocessorRefunded, error) {
	event := "Refunded"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorRefunded)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorTransferFailed represents a TransferFailed event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorTransferFailed struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorTransferFailedEventName = "TransferFailed"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorTransferFailed) ContractEventName() string {
	return IntermediatedpaymentprocessorTransferFailedEventName
}

// UnpackTransferFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferFailed(uint216 indexed invoiceId, address indexed recipient, uint256 amount)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackTransferFailedEvent(log *types.Log) (*IntermediatedpaymentprocessorTransferFailed, error) {
	event := "TransferFailed"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorTransferFailed)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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

// IntermediatedpaymentprocessorUpdateReleaseTime represents a UpdateReleaseTime event raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorUpdateReleaseTime struct {
	InvoiceId     *big.Int
	NewHoldPeriod *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedpaymentprocessorUpdateReleaseTimeEventName = "UpdateReleaseTime"

// ContractEventName returns the user-defined event name.
func (IntermediatedpaymentprocessorUpdateReleaseTime) ContractEventName() string {
	return IntermediatedpaymentprocessorUpdateReleaseTimeEventName
}

// UnpackUpdateReleaseTimeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UpdateReleaseTime(uint216 indexed invoiceId, uint256 newHoldPeriod)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackUpdateReleaseTimeEvent(log *types.Log) (*IntermediatedpaymentprocessorUpdateReleaseTime, error) {
	event := "UpdateReleaseTime"
	if log.Topics[0] != intermediatedpaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedpaymentprocessorUpdateReleaseTime)
	if len(log.Data) > 0 {
		if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedpaymentprocessor.abi.Events[event].Inputs {
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
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["BuyerCannotBeSeller"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackBuyerCannotBeSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["ContractPaused"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["EmptyMetaInvoice"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackEmptyMetaInvoiceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["EscrowWithdrawFailed"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackEscrowWithdrawFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["FeeReceiverCountMismatch"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackFeeReceiverCountMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["HoldPeriodCanNotBeZero"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackHoldPeriodCanNotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidDisputeResolution"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidDisputeResolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidFeeAuthorization"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidFeeAuthorizationError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidFeeReceiver"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidFeeReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidInvoiceStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidMetaInvoicePaymentAmount"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidMetaInvoicePaymentAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidNativePayment"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidNativePaymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidOracle"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidOracleError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidPrice"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidSeller"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvalidSellersPayoutShare"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvalidSellersPayoutShareError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvoiceDoesNotExist"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvoiceDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["InvoiceExpired"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackInvoiceExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["MetaInvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackMetaInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["NoPaymentTokens"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackNoPaymentTokensError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["PaymentTokenNotAllowed"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackPaymentTokenNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["PriceCannotBeZero"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackPriceCannotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["PriceIsTooLow"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackPriceIsTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["Reentrancy"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["SequencerDown"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackSequencerDownError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["StalePrice"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackStalePriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["StalePriceFeed"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackStalePriceFeedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedpaymentprocessor.abi.Errors["UnsupportedToken"].ID.Bytes()[:4]) {
		return intermediatedpaymentprocessor.UnpackUnsupportedTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IntermediatedpaymentprocessorBuyerCannotBeSeller represents a BuyerCannotBeSeller error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorBuyerCannotBeSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyerCannotBeSeller()
func IntermediatedpaymentprocessorBuyerCannotBeSellerErrorID() common.Hash {
	return common.HexToHash("0xb12e242105ea4b2bcdc745efefe14be5558f5f16020ec252980cefc86c6a7a77")
}

// UnpackBuyerCannotBeSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyerCannotBeSeller()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackBuyerCannotBeSellerError(raw []byte) (*IntermediatedpaymentprocessorBuyerCannotBeSeller, error) {
	out := new(IntermediatedpaymentprocessorBuyerCannotBeSeller)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "BuyerCannotBeSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorContractPaused represents a ContractPaused error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ContractPaused()
func IntermediatedpaymentprocessorContractPausedErrorID() common.Hash {
	return common.HexToHash("0xab35696f06e428ebc5ceba8cd17f8fed287baf43440206d1943af1ee53e6d267")
}

// UnpackContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ContractPaused()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackContractPausedError(raw []byte) (*IntermediatedpaymentprocessorContractPaused, error) {
	out := new(IntermediatedpaymentprocessorContractPaused)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "ContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorCreate2EmptyBytecode represents a Create2EmptyBytecode error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorCreate2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func IntermediatedpaymentprocessorCreate2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackCreate2EmptyBytecodeError(raw []byte) (*IntermediatedpaymentprocessorCreate2EmptyBytecode, error) {
	out := new(IntermediatedpaymentprocessorCreate2EmptyBytecode)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorEmptyMetaInvoice represents a EmptyMetaInvoice error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorEmptyMetaInvoice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyMetaInvoice()
func IntermediatedpaymentprocessorEmptyMetaInvoiceErrorID() common.Hash {
	return common.HexToHash("0x815ba404f0d3eea5259f820bd75186cf6e09fe9a2e3f59f2f7a517f382abfd35")
}

// UnpackEmptyMetaInvoiceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyMetaInvoice()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackEmptyMetaInvoiceError(raw []byte) (*IntermediatedpaymentprocessorEmptyMetaInvoice, error) {
	out := new(IntermediatedpaymentprocessorEmptyMetaInvoice)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "EmptyMetaInvoice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorEscrowWithdrawFailed represents a EscrowWithdrawFailed error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorEscrowWithdrawFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowWithdrawFailed()
func IntermediatedpaymentprocessorEscrowWithdrawFailedErrorID() common.Hash {
	return common.HexToHash("0x667ecf9d53e4600a9a128606592ec5e22e0269990439145a2bbc8a983c7af5ac")
}

// UnpackEscrowWithdrawFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowWithdrawFailed()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackEscrowWithdrawFailedError(raw []byte) (*IntermediatedpaymentprocessorEscrowWithdrawFailed, error) {
	out := new(IntermediatedpaymentprocessorEscrowWithdrawFailed)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "EscrowWithdrawFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorFailedDeployment represents a FailedDeployment error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func IntermediatedpaymentprocessorFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackFailedDeploymentError(raw []byte) (*IntermediatedpaymentprocessorFailedDeployment, error) {
	out := new(IntermediatedpaymentprocessorFailedDeployment)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorFeeReceiverCountMismatch represents a FeeReceiverCountMismatch error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorFeeReceiverCountMismatch struct {
	Provided *big.Int
	Expected *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeReceiverCountMismatch(uint256 provided, uint256 expected)
func IntermediatedpaymentprocessorFeeReceiverCountMismatchErrorID() common.Hash {
	return common.HexToHash("0xa0c3f1201017b0a80caaeacc6d8996ecf119f17592f113f1fc0e894fe826c3ee")
}

// UnpackFeeReceiverCountMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeReceiverCountMismatch(uint256 provided, uint256 expected)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackFeeReceiverCountMismatchError(raw []byte) (*IntermediatedpaymentprocessorFeeReceiverCountMismatch, error) {
	out := new(IntermediatedpaymentprocessorFeeReceiverCountMismatch)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "FeeReceiverCountMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorHoldPeriodCanNotBeZero represents a HoldPeriodCanNotBeZero error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorHoldPeriodCanNotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HoldPeriodCanNotBeZero()
func IntermediatedpaymentprocessorHoldPeriodCanNotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x705a71532da8bae84d5c54245bfd200d9655b2c961da65ccc7fcf54a50ad44b4")
}

// UnpackHoldPeriodCanNotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HoldPeriodCanNotBeZero()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackHoldPeriodCanNotBeZeroError(raw []byte) (*IntermediatedpaymentprocessorHoldPeriodCanNotBeZero, error) {
	out := new(IntermediatedpaymentprocessorHoldPeriodCanNotBeZero)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "HoldPeriodCanNotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInsufficientBalance represents a InsufficientBalance error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func IntermediatedpaymentprocessorInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInsufficientBalanceError(raw []byte) (*IntermediatedpaymentprocessorInsufficientBalance, error) {
	out := new(IntermediatedpaymentprocessorInsufficientBalance)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidDisputeResolution represents a InvalidDisputeResolution error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidDisputeResolution struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDisputeResolution()
func IntermediatedpaymentprocessorInvalidDisputeResolutionErrorID() common.Hash {
	return common.HexToHash("0x34819f908388d0ed594d20c6802439086d46ba2397dca397160a28d8b2bd98b1")
}

// UnpackInvalidDisputeResolutionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDisputeResolution()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidDisputeResolutionError(raw []byte) (*IntermediatedpaymentprocessorInvalidDisputeResolution, error) {
	out := new(IntermediatedpaymentprocessorInvalidDisputeResolution)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidDisputeResolution", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidFeeAuthorization represents a InvalidFeeAuthorization error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidFeeAuthorization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeAuthorization()
func IntermediatedpaymentprocessorInvalidFeeAuthorizationErrorID() common.Hash {
	return common.HexToHash("0x1735eabec15c7395efafdfa0dda5c74faf3517b604ff55660d6ef0e7457f2c1d")
}

// UnpackInvalidFeeAuthorizationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeAuthorization()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidFeeAuthorizationError(raw []byte) (*IntermediatedpaymentprocessorInvalidFeeAuthorization, error) {
	out := new(IntermediatedpaymentprocessorInvalidFeeAuthorization)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidFeeAuthorization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidFeeReceiver represents a InvalidFeeReceiver error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidFeeReceiver struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeReceiver()
func IntermediatedpaymentprocessorInvalidFeeReceiverErrorID() common.Hash {
	return common.HexToHash("0xd200485c51caaf66763f8b49c9cfa281a0a10132cb56c8fffc35867701d3fc5f")
}

// UnpackInvalidFeeReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeReceiver()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidFeeReceiverError(raw []byte) (*IntermediatedpaymentprocessorInvalidFeeReceiver, error) {
	out := new(IntermediatedpaymentprocessorInvalidFeeReceiver)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidFeeReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidInvoiceState represents a InvalidInvoiceState error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidInvoiceState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInvoiceState()
func IntermediatedpaymentprocessorInvalidInvoiceStateErrorID() common.Hash {
	return common.HexToHash("0x487e4409b34dcf5275ed8908061cfcde1e134270e5620e0eaff4d68605de2cbc")
}

// UnpackInvalidInvoiceStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInvoiceState()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidInvoiceStateError(raw []byte) (*IntermediatedpaymentprocessorInvalidInvoiceState, error) {
	out := new(IntermediatedpaymentprocessorInvalidInvoiceState)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidInvoiceState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidMetaInvoicePaymentAmount represents a InvalidMetaInvoicePaymentAmount error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidMetaInvoicePaymentAmount struct {
	Sent     *big.Int
	Expected *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMetaInvoicePaymentAmount(uint256 sent, uint256 expected)
func IntermediatedpaymentprocessorInvalidMetaInvoicePaymentAmountErrorID() common.Hash {
	return common.HexToHash("0xc7632c7d819e4ec9dbca7ec79df876ac7d5ca98ab46cf285c8f8a7ff52ea72a3")
}

// UnpackInvalidMetaInvoicePaymentAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMetaInvoicePaymentAmount(uint256 sent, uint256 expected)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidMetaInvoicePaymentAmountError(raw []byte) (*IntermediatedpaymentprocessorInvalidMetaInvoicePaymentAmount, error) {
	out := new(IntermediatedpaymentprocessorInvalidMetaInvoicePaymentAmount)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidMetaInvoicePaymentAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidNativePayment represents a InvalidNativePayment error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidNativePayment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNativePayment()
func IntermediatedpaymentprocessorInvalidNativePaymentErrorID() common.Hash {
	return common.HexToHash("0x214510aac5bc5d45b2314d915edc9aa20e9ec869bcb7e6d50d8d068658a871c9")
}

// UnpackInvalidNativePaymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNativePayment()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidNativePaymentError(raw []byte) (*IntermediatedpaymentprocessorInvalidNativePayment, error) {
	out := new(IntermediatedpaymentprocessorInvalidNativePayment)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidNativePayment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidOracle represents a InvalidOracle error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidOracle struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidOracle()
func IntermediatedpaymentprocessorInvalidOracleErrorID() common.Hash {
	return common.HexToHash("0x9589a27d464cce309224596a505cbfd22e5fde1f0f420cecf8a6b6c1d65791b6")
}

// UnpackInvalidOracleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidOracle()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidOracleError(raw []byte) (*IntermediatedpaymentprocessorInvalidOracle, error) {
	out := new(IntermediatedpaymentprocessorInvalidOracle)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidOracle", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidPrice represents a InvalidPrice error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPrice()
func IntermediatedpaymentprocessorInvalidPriceErrorID() common.Hash {
	return common.HexToHash("0x00bfc9219afe7e8e3b9f14a4708e4cd3d8acb04e325ce992b2a60a58a519683a")
}

// UnpackInvalidPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPrice()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidPriceError(raw []byte) (*IntermediatedpaymentprocessorInvalidPrice, error) {
	out := new(IntermediatedpaymentprocessorInvalidPrice)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidSeller represents a InvalidSeller error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSeller()
func IntermediatedpaymentprocessorInvalidSellerErrorID() common.Hash {
	return common.HexToHash("0xbab7ca35fcde13672ca7744c85f31cdd0a5c3f882f4b4992269c1e7dc56732e9")
}

// UnpackInvalidSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSeller()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidSellerError(raw []byte) (*IntermediatedpaymentprocessorInvalidSeller, error) {
	out := new(IntermediatedpaymentprocessorInvalidSeller)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvalidSellersPayoutShare represents a InvalidSellersPayoutShare error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvalidSellersPayoutShare struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSellersPayoutShare()
func IntermediatedpaymentprocessorInvalidSellersPayoutShareErrorID() common.Hash {
	return common.HexToHash("0x453fb42ddfd1ecde870e9bd55d8b7f21b2333b613ee68779ba5f60498951666d")
}

// UnpackInvalidSellersPayoutShareError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSellersPayoutShare()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvalidSellersPayoutShareError(raw []byte) (*IntermediatedpaymentprocessorInvalidSellersPayoutShare, error) {
	out := new(IntermediatedpaymentprocessorInvalidSellersPayoutShare)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvalidSellersPayoutShare", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvoiceAlreadyExists represents a InvoiceAlreadyExists error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceAlreadyExists()
func IntermediatedpaymentprocessorInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x074bc9355c94925fa82ddb49dcc88f1a666f1d1aa24efbddcdbe5f8d98b7ed59")
}

// UnpackInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceAlreadyExists()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvoiceAlreadyExistsError(raw []byte) (*IntermediatedpaymentprocessorInvoiceAlreadyExists, error) {
	out := new(IntermediatedpaymentprocessorInvoiceAlreadyExists)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvoiceDoesNotExist represents a InvoiceDoesNotExist error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvoiceDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceDoesNotExist()
func IntermediatedpaymentprocessorInvoiceDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x715d9228f420b2c4c07281fb8597619f2ca0c9d8cada84ce032e60ec6407b582")
}

// UnpackInvoiceDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceDoesNotExist()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvoiceDoesNotExistError(raw []byte) (*IntermediatedpaymentprocessorInvoiceDoesNotExist, error) {
	out := new(IntermediatedpaymentprocessorInvoiceDoesNotExist)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvoiceDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorInvoiceExpired represents a InvoiceExpired error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorInvoiceExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceExpired()
func IntermediatedpaymentprocessorInvoiceExpiredErrorID() common.Hash {
	return common.HexToHash("0xf04e9cf09371be6ef375f7f016c1ac94b6c5c6a4d247bec67ae9568dd6b911b6")
}

// UnpackInvoiceExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceExpired()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackInvoiceExpiredError(raw []byte) (*IntermediatedpaymentprocessorInvoiceExpired, error) {
	out := new(IntermediatedpaymentprocessorInvoiceExpired)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "InvoiceExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorMetaInvoiceAlreadyExists represents a MetaInvoiceAlreadyExists error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorMetaInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MetaInvoiceAlreadyExists()
func IntermediatedpaymentprocessorMetaInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xb09960c1d49af4579c96dbfb857f76f135da1b549d276e06491c05ec24747201")
}

// UnpackMetaInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MetaInvoiceAlreadyExists()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackMetaInvoiceAlreadyExistsError(raw []byte) (*IntermediatedpaymentprocessorMetaInvoiceAlreadyExists, error) {
	out := new(IntermediatedpaymentprocessorMetaInvoiceAlreadyExists)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "MetaInvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorNoPaymentTokens represents a NoPaymentTokens error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorNoPaymentTokens struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPaymentTokens()
func IntermediatedpaymentprocessorNoPaymentTokensErrorID() common.Hash {
	return common.HexToHash("0xb883eab0fbc18a684735ee65aa188bdbbd75f51ef9aeda0e59a1c2b7d4fa55c5")
}

// UnpackNoPaymentTokensError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPaymentTokens()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackNoPaymentTokensError(raw []byte) (*IntermediatedpaymentprocessorNoPaymentTokens, error) {
	out := new(IntermediatedpaymentprocessorNoPaymentTokens)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "NoPaymentTokens", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorNotAuthorized represents a NotAuthorized error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func IntermediatedpaymentprocessorNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackNotAuthorizedError(raw []byte) (*IntermediatedpaymentprocessorNotAuthorized, error) {
	out := new(IntermediatedpaymentprocessorNotAuthorized)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorPaymentTokenNotAllowed represents a PaymentTokenNotAllowed error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorPaymentTokenNotAllowed struct {
	InvoiceId    *big.Int
	PaymentToken common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentTokenNotAllowed(uint216 invoiceId, address paymentToken)
func IntermediatedpaymentprocessorPaymentTokenNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x8e6656fbc0fc74fe6d8f2df365687f5e04719673f75fda2eeadaf139389c43ae")
}

// UnpackPaymentTokenNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentTokenNotAllowed(uint216 invoiceId, address paymentToken)
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackPaymentTokenNotAllowedError(raw []byte) (*IntermediatedpaymentprocessorPaymentTokenNotAllowed, error) {
	out := new(IntermediatedpaymentprocessorPaymentTokenNotAllowed)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "PaymentTokenNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorPriceCannotBeZero represents a PriceCannotBeZero error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorPriceCannotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceCannotBeZero()
func IntermediatedpaymentprocessorPriceCannotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x2c669f0ac3409adbbadbe16eaad2cac428e45b3cb8de2f47377f30f8b5729e18")
}

// UnpackPriceCannotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceCannotBeZero()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackPriceCannotBeZeroError(raw []byte) (*IntermediatedpaymentprocessorPriceCannotBeZero, error) {
	out := new(IntermediatedpaymentprocessorPriceCannotBeZero)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "PriceCannotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorPriceIsTooLow represents a PriceIsTooLow error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorPriceIsTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceIsTooLow()
func IntermediatedpaymentprocessorPriceIsTooLowErrorID() common.Hash {
	return common.HexToHash("0xdb8db56995596ab6855aa515b34d8c3549b6c6fd6435c2e3af6e0d886de7e87c")
}

// UnpackPriceIsTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceIsTooLow()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackPriceIsTooLowError(raw []byte) (*IntermediatedpaymentprocessorPriceIsTooLow, error) {
	out := new(IntermediatedpaymentprocessorPriceIsTooLow)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "PriceIsTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorReentrancy represents a Reentrancy error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Reentrancy()
func IntermediatedpaymentprocessorReentrancyErrorID() common.Hash {
	return common.HexToHash("0xab143c06c9772d69bbbc9f2fe74acd02f810e93b099f3d1dac8448ac9ae35991")
}

// UnpackReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Reentrancy()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackReentrancyError(raw []byte) (*IntermediatedpaymentprocessorReentrancy, error) {
	out := new(IntermediatedpaymentprocessorReentrancy)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "Reentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorSequencerDown represents a SequencerDown error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorSequencerDown struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SequencerDown()
func IntermediatedpaymentprocessorSequencerDownErrorID() common.Hash {
	return common.HexToHash("0x032b3d00cfb14fdf4eecb317aaf61db9dd7331083f0db9baa2eae06ec3e15ecb")
}

// UnpackSequencerDownError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SequencerDown()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackSequencerDownError(raw []byte) (*IntermediatedpaymentprocessorSequencerDown, error) {
	out := new(IntermediatedpaymentprocessorSequencerDown)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "SequencerDown", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorStalePrice represents a StalePrice error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorStalePrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StalePrice()
func IntermediatedpaymentprocessorStalePriceErrorID() common.Hash {
	return common.HexToHash("0x19abf40e7c2e0280d6137a5d95d9f3793d913552f00d0e25e4ab4388bcc0d573")
}

// UnpackStalePriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StalePrice()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackStalePriceError(raw []byte) (*IntermediatedpaymentprocessorStalePrice, error) {
	out := new(IntermediatedpaymentprocessorStalePrice)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "StalePrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorStalePriceFeed represents a StalePriceFeed error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorStalePriceFeed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StalePriceFeed()
func IntermediatedpaymentprocessorStalePriceFeedErrorID() common.Hash {
	return common.HexToHash("0x1087e109db85b72cf66a8dbc341a9e5601a49c3f12e82151b3eb6e742d4a766e")
}

// UnpackStalePriceFeedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StalePriceFeed()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackStalePriceFeedError(raw []byte) (*IntermediatedpaymentprocessorStalePriceFeed, error) {
	out := new(IntermediatedpaymentprocessorStalePriceFeed)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "StalePriceFeed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedpaymentprocessorUnsupportedToken represents a UnsupportedToken error raised by the Intermediatedpaymentprocessor contract.
type IntermediatedpaymentprocessorUnsupportedToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedToken()
func IntermediatedpaymentprocessorUnsupportedTokenErrorID() common.Hash {
	return common.HexToHash("0x6a1728823cfcc894fe1dcf37bfe71f201fb66b0b61862091f422023e22ea5ab9")
}

// UnpackUnsupportedTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedToken()
func (intermediatedpaymentprocessor *Intermediatedpaymentprocessor) UnpackUnsupportedTokenError(raw []byte) (*IntermediatedpaymentprocessorUnsupportedToken, error) {
	out := new(IntermediatedpaymentprocessorUnsupportedToken)
	if err := intermediatedpaymentprocessor.abi.UnpackIntoInterface(out, "UnsupportedToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}
