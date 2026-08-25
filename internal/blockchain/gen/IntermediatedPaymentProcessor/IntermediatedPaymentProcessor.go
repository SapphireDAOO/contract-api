// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package intermediatedprocessor

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
}

// IIntermediatedPaymentProcessorMetaInvoice is an auto generated low-level Go binding around an user-defined struct.
type IIntermediatedPaymentProcessorMetaInvoice struct {
	Price         *big.Int
	SubInvoiceIds []*big.Int
}

// IntermediatedprocessorMetaData contains all meta data concerning the Intermediatedprocessor contract.
var IntermediatedprocessorMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_getDecimals\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"tokenDecimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeSalt\",\"inputs\":[{\"name\":\"_seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"createDispute\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createMetaInvoice\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple[]\",\"internalType\":\"structIIntermediatedPaymentProcessor.InvoiceCreationParam[]\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSingleInvoice\",\"inputs\":[{\"name\":\"_param\",\"type\":\"tuple\",\"internalType\":\"structIIntermediatedPaymentProcessor.InvoiceCreationParam\",\"components\":[{\"name\":\"invoiceId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"i\",\"type\":\"tuple\",\"internalType\":\"structIIntermediatedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMetaInvoice\",\"inputs\":[{\"name\":\"_metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"m\",\"type\":\"tuple\",\"internalType\":\"structIIntermediatedPaymentProcessor.MetaInvoice\",\"components\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"subInvoiceIds\",\"type\":\"uint216[]\",\"internalType\":\"uint216[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"currentMinimumPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextInvoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextMetaInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextMetaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPredictedAddress\",\"inputs\":[{\"name\":\"_salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"predictedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenValueFromUsd\",\"inputs\":[{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleDispute\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_resolution\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_sellerShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOracleManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"payMetaInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"payMetaInvoiceWithValue\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refund\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_refundShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInvoiceReleaseTime\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_holdPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumPrice\",\"inputs\":[{\"name\":\"_newMinimumPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOracle\",\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalMetaInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"totalMetaInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalUniqueInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"totalInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DisputeCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeDismissed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeResolved\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeSettled\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"buyerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EscrowCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"escrow\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCanceled\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"invoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIntermediatedPaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoicePaid\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"escrowAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockedPaymentRecovered\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetaInvoiceCreated\",\"inputs\":[{\"name\":\"metaInvoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"totalPrice\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleUpdated\",\"inputs\":[{\"name\":\"previousOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOracle\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentReleased\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"currency\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Refunded\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateReleaseTime\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"newHoldPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BuyerCannotBeSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Create2EmptyBytecode\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMetaInvoice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowWithdrawFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedDeployment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HoldPeriodCanNotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDisputeResolution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeAuthorization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInvoiceState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMetaInvoicePaymentAmount\",\"inputs\":[{\"name\":\"sent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidNativePayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOracle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSellersPayoutShare\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MetaInvoiceAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PriceIsTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencerDown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePriceFeed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[]}]",
	ID:  "Intermediatedprocessor",
}

// Intermediatedprocessor is an auto generated Go binding around an Ethereum contract.
type Intermediatedprocessor struct {
	abi abi.ABI
}

// NewIntermediatedprocessor creates a new instance of Intermediatedprocessor.
func NewIntermediatedprocessor() *Intermediatedprocessor {
	parsed, err := IntermediatedprocessorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Intermediatedprocessor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Intermediatedprocessor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _paymentProcessorStorageAddress, address _oracle) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackConstructor(_paymentProcessorStorageAddress common.Address, _oracle common.Address) []byte {
	enc, err := intermediatedprocessor.abi.Pack("", _paymentProcessorStorageAddress, _oracle)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ed8bf86.
//
// Solidity: function _getDecimals(address _token) view returns(uint8 tokenDecimals)
func (intermediatedprocessor *Intermediatedprocessor) PackGetDecimals(token common.Address) []byte {
	enc, err := intermediatedprocessor.abi.Pack("_getDecimals", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2ed8bf86.
//
// Solidity: function _getDecimals(address _token) view returns(uint8 tokenDecimals)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetDecimals(data []byte) (uint8, error) {
	out, err := intermediatedprocessor.abi.Unpack("_getDecimals", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackCancelInvoice(invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("cancelInvoice", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackComputeSalt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a221549.
//
// Solidity: function computeSalt(address _seller, address _buyer, uint216 _invoiceId) pure returns(bytes32 salt)
func (intermediatedprocessor *Intermediatedprocessor) PackComputeSalt(seller common.Address, buyer common.Address, invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("computeSalt", seller, buyer, invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackComputeSalt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a221549.
//
// Solidity: function computeSalt(address _seller, address _buyer, uint216 _invoiceId) pure returns(bytes32 salt)
func (intermediatedprocessor *Intermediatedprocessor) UnpackComputeSalt(data []byte) ([32]byte, error) {
	out, err := intermediatedprocessor.abi.Unpack("computeSalt", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackCreateDispute(invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("createDispute", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd05d9f1d.
//
// Solidity: function createMetaInvoice((string,address,uint256,uint32)[] _param) returns(uint216 metaInvoiceId)
func (intermediatedprocessor *Intermediatedprocessor) PackCreateMetaInvoice(param []IIntermediatedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := intermediatedprocessor.abi.Pack("createMetaInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd05d9f1d.
//
// Solidity: function createMetaInvoice((string,address,uint256,uint32)[] _param) returns(uint216 metaInvoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackCreateMetaInvoice(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("createMetaInvoice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCreateSingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x180bb2c1.
//
// Solidity: function createSingleInvoice((string,address,uint256,uint32) _param) returns(uint216 invoiceId)
func (intermediatedprocessor *Intermediatedprocessor) PackCreateSingleInvoice(param IIntermediatedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := intermediatedprocessor.abi.Pack("createSingleInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSingleInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x180bb2c1.
//
// Solidity: function createSingleInvoice((string,address,uint256,uint32) _param) returns(uint216 invoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackCreateSingleInvoice(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("createSingleInvoice", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetInvoice(invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("getInvoice", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4cfa3edf.
//
// Solidity: function getInvoice(uint216 _invoiceId) view returns((uint216,uint40,uint40,uint40,uint40,uint8,uint8,uint32,uint16,uint216,address,address,address,address,address,uint256,uint256,uint256) i)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetInvoice(data []byte) (IIntermediatedPaymentProcessorInvoice, error) {
	out, err := intermediatedprocessor.abi.Unpack("getInvoice", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetMetaInvoice(metaInvoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("getMetaInvoice", metaInvoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x33453a3d.
//
// Solidity: function getMetaInvoice(uint216 _metaInvoiceId) view returns((uint256,uint216[]) m)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetMetaInvoice(data []byte) (IIntermediatedPaymentProcessorMetaInvoice, error) {
	out, err := intermediatedprocessor.abi.Unpack("getMetaInvoice", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetMinimumPrice() []byte {
	enc, err := intermediatedprocessor.abi.Pack("getMinimumPrice")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMinimumPrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac4316cb.
//
// Solidity: function getMinimumPrice() view returns(uint256 currentMinimumPrice)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetMinimumPrice(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("getMinimumPrice", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetNextInvoiceNonce() []byte {
	enc, err := intermediatedprocessor.abi.Pack("getNextInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonce)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetNextInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("getNextInvoiceNonce", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetNextMetaInvoiceNonce() []byte {
	enc, err := intermediatedprocessor.abi.Pack("getNextMetaInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextMetaInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x540f0f3c.
//
// Solidity: function getNextMetaInvoiceNonce() view returns(uint216 nextMetaInvoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetNextMetaInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("getNextMetaInvoiceNonce", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetPredictedAddress(salt [32]byte, invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("getPredictedAddress", salt, invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPredictedAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc4ed09d9.
//
// Solidity: function getPredictedAddress(bytes32 _salt, uint216 _invoiceId) view returns(address predictedAddress)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetPredictedAddress(data []byte) (common.Address, error) {
	out, err := intermediatedprocessor.abi.Unpack("getPredictedAddress", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackGetTokenValueFromUsd(paymentToken common.Address, usdAmount *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("getTokenValueFromUsd", paymentToken, usdAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenValueFromUsd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address _paymentToken, uint256 _usdAmount) view returns(uint256 tokenValue)
func (intermediatedprocessor *Intermediatedprocessor) UnpackGetTokenValueFromUsd(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("getTokenValueFromUsd", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackHandleDispute(invoiceId *big.Int, resolution uint8, sellerShare *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("handleDispute", invoiceId, resolution, sellerShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (intermediatedprocessor *Intermediatedprocessor) PackOracle() []byte {
	enc, err := intermediatedprocessor.abi.Pack("oracle")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOracle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (intermediatedprocessor *Intermediatedprocessor) UnpackOracle(data []byte) (common.Address, error) {
	out, err := intermediatedprocessor.abi.Unpack("oracle", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackPayInvoice(invoiceId *big.Int, paymentToken common.Address, feeReceiver common.Address, data []byte) []byte {
	enc, err := intermediatedprocessor.abi.Pack("payInvoice", invoiceId, paymentToken, feeReceiver, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPayMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1114e4b.
//
// Solidity: function payMetaInvoice(uint216 _invoiceId, address _paymentToken) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackPayMetaInvoice(invoiceId *big.Int, paymentToken common.Address) []byte {
	enc, err := intermediatedprocessor.abi.Pack("payMetaInvoice", invoiceId, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPayMetaInvoiceWithValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4859734.
//
// Solidity: function payMetaInvoiceWithValue(uint216 _invoiceId) payable returns()
func (intermediatedprocessor *Intermediatedprocessor) PackPayMetaInvoiceWithValue(invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("payMetaInvoiceWithValue", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (intermediatedprocessor *Intermediatedprocessor) PackPpStorage() []byte {
	enc, err := intermediatedprocessor.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (intermediatedprocessor *Intermediatedprocessor) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := intermediatedprocessor.abi.Unpack("ppStorage", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackRefund(invoiceId *big.Int, refundShare *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("refund", invoiceId, refundShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb990809.
//
// Solidity: function release(uint216 _invoiceId) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackRelease(invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("release", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackResolveDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5168e613.
//
// Solidity: function resolveDispute(uint216 _invoiceId) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackResolveDispute(invoiceId *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("resolveDispute", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetInvoiceReleaseTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50e5d5cf.
//
// Solidity: function setInvoiceReleaseTime(uint216 _invoiceId, uint256 _holdPeriod) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackSetInvoiceReleaseTime(invoiceId *big.Int, holdPeriod *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("setInvoiceReleaseTime", invoiceId, holdPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinimumPrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25cea976.
//
// Solidity: function setMinimumPrice(uint256 _newMinimumPrice) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackSetMinimumPrice(newMinimumPrice *big.Int) []byte {
	enc, err := intermediatedprocessor.abi.Pack("setMinimumPrice", newMinimumPrice)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetOracle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (intermediatedprocessor *Intermediatedprocessor) PackSetOracle(oracle common.Address) []byte {
	enc, err := intermediatedprocessor.abi.Pack("setOracle", oracle)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalMetaInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint216 totalMetaInvoices)
func (intermediatedprocessor *Intermediatedprocessor) PackTotalMetaInvoiceCreated() []byte {
	enc, err := intermediatedprocessor.abi.Pack("totalMetaInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalMetaInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint216 totalMetaInvoices)
func (intermediatedprocessor *Intermediatedprocessor) UnpackTotalMetaInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("totalMetaInvoiceCreated", data)
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
func (intermediatedprocessor *Intermediatedprocessor) PackTotalUniqueInvoiceCreated() []byte {
	enc, err := intermediatedprocessor.abi.Pack("totalUniqueInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalUniqueInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint216 totalInvoices)
func (intermediatedprocessor *Intermediatedprocessor) UnpackTotalUniqueInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := intermediatedprocessor.abi.Unpack("totalUniqueInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// IntermediatedprocessorDisputeCreated represents a DisputeCreated event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorDisputeCreated struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorDisputeCreatedEventName = "DisputeCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorDisputeCreated) ContractEventName() string {
	return IntermediatedprocessorDisputeCreatedEventName
}

// UnpackDisputeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeCreated(uint216 indexed invoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackDisputeCreatedEvent(log *types.Log) (*IntermediatedprocessorDisputeCreated, error) {
	event := "DisputeCreated"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorDisputeCreated)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorDisputeDismissed represents a DisputeDismissed event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorDisputeDismissed struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorDisputeDismissedEventName = "DisputeDismissed"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorDisputeDismissed) ContractEventName() string {
	return IntermediatedprocessorDisputeDismissedEventName
}

// UnpackDisputeDismissedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeDismissed(uint216 indexed invoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackDisputeDismissedEvent(log *types.Log) (*IntermediatedprocessorDisputeDismissed, error) {
	event := "DisputeDismissed"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorDisputeDismissed)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorDisputeResolved represents a DisputeResolved event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorDisputeResolved struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorDisputeResolvedEventName = "DisputeResolved"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorDisputeResolved) ContractEventName() string {
	return IntermediatedprocessorDisputeResolvedEventName
}

// UnpackDisputeResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeResolved(uint216 indexed invoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackDisputeResolvedEvent(log *types.Log) (*IntermediatedprocessorDisputeResolved, error) {
	event := "DisputeResolved"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorDisputeResolved)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorDisputeSettled represents a DisputeSettled event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorDisputeSettled struct {
	InvoiceId    *big.Int
	SellerAmount *big.Int
	BuyerAmount  *big.Int
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorDisputeSettledEventName = "DisputeSettled"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorDisputeSettled) ContractEventName() string {
	return IntermediatedprocessorDisputeSettledEventName
}

// UnpackDisputeSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeSettled(uint216 indexed invoiceId, uint256 sellerAmount, uint256 buyerAmount, uint256 fee)
func (intermediatedprocessor *Intermediatedprocessor) UnpackDisputeSettledEvent(log *types.Log) (*IntermediatedprocessorDisputeSettled, error) {
	event := "DisputeSettled"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorDisputeSettled)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorEscrowCreated represents a EscrowCreated event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorEscrowCreated struct {
	InvoiceId *big.Int
	Escrow    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorEscrowCreatedEventName = "EscrowCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorEscrowCreated) ContractEventName() string {
	return IntermediatedprocessorEscrowCreatedEventName
}

// UnpackEscrowCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EscrowCreated(uint216 indexed invoiceId, address indexed escrow)
func (intermediatedprocessor *Intermediatedprocessor) UnpackEscrowCreatedEvent(log *types.Log) (*IntermediatedprocessorEscrowCreated, error) {
	event := "EscrowCreated"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorEscrowCreated)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvoiceCanceled struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorInvoiceCanceled) ContractEventName() string {
	return IntermediatedprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(uint216 indexed invoiceId)
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvoiceCanceledEvent(log *types.Log) (*IntermediatedprocessorInvoiceCanceled, error) {
	event := "InvoiceCanceled"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorInvoiceCanceled)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorInvoiceCreated represents a InvoiceCreated event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvoiceCreated struct {
	InvoiceId *big.Int
	Invoice   IIntermediatedPaymentProcessorInvoice
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorInvoiceCreated) ContractEventName() string {
	return IntermediatedprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(uint216 indexed invoiceId, (uint216,uint40,uint40,uint40,uint40,uint8,uint8,uint32,uint16,uint216,address,address,address,address,address,uint256,uint256,uint256) invoice)
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvoiceCreatedEvent(log *types.Log) (*IntermediatedprocessorInvoiceCreated, error) {
	event := "InvoiceCreated"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorInvoiceCreated)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorInvoicePaid represents a InvoicePaid event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvoicePaid struct {
	InvoiceId     *big.Int
	PaymentToken  common.Address
	EscrowAddress common.Address
	Amount        *big.Int
	ReleaseAt     *big.Int
	FeeReceiver   common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorInvoicePaidEventName = "InvoicePaid"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorInvoicePaid) ContractEventName() string {
	return IntermediatedprocessorInvoicePaidEventName
}

// UnpackInvoicePaidEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoicePaid(uint216 indexed invoiceId, address paymentToken, address escrowAddress, uint256 amount, uint40 releaseAt, address feeReceiver)
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvoicePaidEvent(log *types.Log) (*IntermediatedprocessorInvoicePaid, error) {
	event := "InvoicePaid"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorInvoicePaid)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorLockedPaymentRecovered represents a LockedPaymentRecovered event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorLockedPaymentRecovered struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorLockedPaymentRecoveredEventName = "LockedPaymentRecovered"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorLockedPaymentRecovered) ContractEventName() string {
	return IntermediatedprocessorLockedPaymentRecoveredEventName
}

// UnpackLockedPaymentRecoveredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LockedPaymentRecovered(uint216 indexed invoiceId, address indexed recipient, uint256 amount)
func (intermediatedprocessor *Intermediatedprocessor) UnpackLockedPaymentRecoveredEvent(log *types.Log) (*IntermediatedprocessorLockedPaymentRecovered, error) {
	event := "LockedPaymentRecovered"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorLockedPaymentRecovered)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorMetaInvoiceCreated represents a MetaInvoiceCreated event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorMetaInvoiceCreated struct {
	MetaInvoiceId *big.Int
	TotalPrice    *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorMetaInvoiceCreatedEventName = "MetaInvoiceCreated"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorMetaInvoiceCreated) ContractEventName() string {
	return IntermediatedprocessorMetaInvoiceCreatedEventName
}

// UnpackMetaInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetaInvoiceCreated(uint216 indexed metaInvoiceId, uint256 indexed totalPrice)
func (intermediatedprocessor *Intermediatedprocessor) UnpackMetaInvoiceCreatedEvent(log *types.Log) (*IntermediatedprocessorMetaInvoiceCreated, error) {
	event := "MetaInvoiceCreated"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorMetaInvoiceCreated)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorOracleUpdated represents a OracleUpdated event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorOracleUpdated struct {
	PreviousOracle common.Address
	NewOracle      common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorOracleUpdatedEventName = "OracleUpdated"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorOracleUpdated) ContractEventName() string {
	return IntermediatedprocessorOracleUpdatedEventName
}

// UnpackOracleUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OracleUpdated(address indexed previousOracle, address indexed newOracle)
func (intermediatedprocessor *Intermediatedprocessor) UnpackOracleUpdatedEvent(log *types.Log) (*IntermediatedprocessorOracleUpdated, error) {
	event := "OracleUpdated"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorOracleUpdated)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorPaymentReleased represents a PaymentReleased event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorPaymentReleased struct {
	InvoiceId    *big.Int
	Receiver     common.Address
	Currency     common.Address
	SellerAmount *big.Int
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorPaymentReleasedEventName = "PaymentReleased"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorPaymentReleased) ContractEventName() string {
	return IntermediatedprocessorPaymentReleasedEventName
}

// UnpackPaymentReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentReleased(uint216 indexed invoiceId, address receiver, address currency, uint256 sellerAmount, uint256 fee)
func (intermediatedprocessor *Intermediatedprocessor) UnpackPaymentReleasedEvent(log *types.Log) (*IntermediatedprocessorPaymentReleased, error) {
	event := "PaymentReleased"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorPaymentReleased)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorRefunded represents a Refunded event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorRefunded struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorRefundedEventName = "Refunded"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorRefunded) ContractEventName() string {
	return IntermediatedprocessorRefundedEventName
}

// UnpackRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Refunded(uint216 indexed invoiceId, uint256 indexed amount)
func (intermediatedprocessor *Intermediatedprocessor) UnpackRefundedEvent(log *types.Log) (*IntermediatedprocessorRefunded, error) {
	event := "Refunded"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorRefunded)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorTransferFailed represents a TransferFailed event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorTransferFailed struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorTransferFailedEventName = "TransferFailed"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorTransferFailed) ContractEventName() string {
	return IntermediatedprocessorTransferFailedEventName
}

// UnpackTransferFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferFailed(uint216 indexed invoiceId, address indexed recipient, uint256 amount)
func (intermediatedprocessor *Intermediatedprocessor) UnpackTransferFailedEvent(log *types.Log) (*IntermediatedprocessorTransferFailed, error) {
	event := "TransferFailed"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorTransferFailed)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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

// IntermediatedprocessorUpdateReleaseTime represents a UpdateReleaseTime event raised by the Intermediatedprocessor contract.
type IntermediatedprocessorUpdateReleaseTime struct {
	InvoiceId     *big.Int
	NewHoldPeriod *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const IntermediatedprocessorUpdateReleaseTimeEventName = "UpdateReleaseTime"

// ContractEventName returns the user-defined event name.
func (IntermediatedprocessorUpdateReleaseTime) ContractEventName() string {
	return IntermediatedprocessorUpdateReleaseTimeEventName
}

// UnpackUpdateReleaseTimeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UpdateReleaseTime(uint216 indexed invoiceId, uint256 newHoldPeriod)
func (intermediatedprocessor *Intermediatedprocessor) UnpackUpdateReleaseTimeEvent(log *types.Log) (*IntermediatedprocessorUpdateReleaseTime, error) {
	event := "UpdateReleaseTime"
	if log.Topics[0] != intermediatedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IntermediatedprocessorUpdateReleaseTime)
	if len(log.Data) > 0 {
		if err := intermediatedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range intermediatedprocessor.abi.Events[event].Inputs {
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
func (intermediatedprocessor *Intermediatedprocessor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["BuyerCannotBeSeller"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackBuyerCannotBeSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["ContractPaused"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["Create2EmptyBytecode"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackCreate2EmptyBytecodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["EmptyMetaInvoice"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackEmptyMetaInvoiceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["EscrowWithdrawFailed"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackEscrowWithdrawFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["FailedDeployment"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackFailedDeploymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["HoldPeriodCanNotBeZero"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackHoldPeriodCanNotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidDisputeResolution"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidDisputeResolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidFeeAuthorization"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidFeeAuthorizationError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidFeeReceiver"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidFeeReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidInvoiceStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidMetaInvoicePaymentAmount"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidMetaInvoicePaymentAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidNativePayment"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidNativePaymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidOracle"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidOracleError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidPrice"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidSeller"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvalidSellersPayoutShare"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvalidSellersPayoutShareError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvoiceDoesNotExist"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvoiceDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["InvoiceExpired"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackInvoiceExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["MetaInvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackMetaInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["PriceCannotBeZero"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackPriceCannotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["PriceIsTooLow"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackPriceIsTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["Reentrancy"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["SequencerDown"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackSequencerDownError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["StalePrice"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackStalePriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["StalePriceFeed"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackStalePriceFeedError(raw[4:])
	}
	if bytes.Equal(raw[:4], intermediatedprocessor.abi.Errors["UnsupportedToken"].ID.Bytes()[:4]) {
		return intermediatedprocessor.UnpackUnsupportedTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IntermediatedprocessorBuyerCannotBeSeller represents a BuyerCannotBeSeller error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorBuyerCannotBeSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyerCannotBeSeller()
func IntermediatedprocessorBuyerCannotBeSellerErrorID() common.Hash {
	return common.HexToHash("0xb12e242105ea4b2bcdc745efefe14be5558f5f16020ec252980cefc86c6a7a77")
}

// UnpackBuyerCannotBeSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyerCannotBeSeller()
func (intermediatedprocessor *Intermediatedprocessor) UnpackBuyerCannotBeSellerError(raw []byte) (*IntermediatedprocessorBuyerCannotBeSeller, error) {
	out := new(IntermediatedprocessorBuyerCannotBeSeller)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "BuyerCannotBeSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorContractPaused represents a ContractPaused error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ContractPaused()
func IntermediatedprocessorContractPausedErrorID() common.Hash {
	return common.HexToHash("0xab35696f06e428ebc5ceba8cd17f8fed287baf43440206d1943af1ee53e6d267")
}

// UnpackContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ContractPaused()
func (intermediatedprocessor *Intermediatedprocessor) UnpackContractPausedError(raw []byte) (*IntermediatedprocessorContractPaused, error) {
	out := new(IntermediatedprocessorContractPaused)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "ContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorCreate2EmptyBytecode represents a Create2EmptyBytecode error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorCreate2EmptyBytecode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Create2EmptyBytecode()
func IntermediatedprocessorCreate2EmptyBytecodeErrorID() common.Hash {
	return common.HexToHash("0x4ca249dcffe41558ef8b961d71c905e4fa4317a1663f377b9610642e4e0abdb6")
}

// UnpackCreate2EmptyBytecodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Create2EmptyBytecode()
func (intermediatedprocessor *Intermediatedprocessor) UnpackCreate2EmptyBytecodeError(raw []byte) (*IntermediatedprocessorCreate2EmptyBytecode, error) {
	out := new(IntermediatedprocessorCreate2EmptyBytecode)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "Create2EmptyBytecode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorEmptyMetaInvoice represents a EmptyMetaInvoice error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorEmptyMetaInvoice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyMetaInvoice()
func IntermediatedprocessorEmptyMetaInvoiceErrorID() common.Hash {
	return common.HexToHash("0x815ba404f0d3eea5259f820bd75186cf6e09fe9a2e3f59f2f7a517f382abfd35")
}

// UnpackEmptyMetaInvoiceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyMetaInvoice()
func (intermediatedprocessor *Intermediatedprocessor) UnpackEmptyMetaInvoiceError(raw []byte) (*IntermediatedprocessorEmptyMetaInvoice, error) {
	out := new(IntermediatedprocessorEmptyMetaInvoice)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "EmptyMetaInvoice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorEscrowWithdrawFailed represents a EscrowWithdrawFailed error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorEscrowWithdrawFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowWithdrawFailed()
func IntermediatedprocessorEscrowWithdrawFailedErrorID() common.Hash {
	return common.HexToHash("0x667ecf9d53e4600a9a128606592ec5e22e0269990439145a2bbc8a983c7af5ac")
}

// UnpackEscrowWithdrawFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowWithdrawFailed()
func (intermediatedprocessor *Intermediatedprocessor) UnpackEscrowWithdrawFailedError(raw []byte) (*IntermediatedprocessorEscrowWithdrawFailed, error) {
	out := new(IntermediatedprocessorEscrowWithdrawFailed)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "EscrowWithdrawFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorFailedDeployment represents a FailedDeployment error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorFailedDeployment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedDeployment()
func IntermediatedprocessorFailedDeploymentErrorID() common.Hash {
	return common.HexToHash("0xb06ebf3d5067824a3fe5d5ba19471e035a7de6c88dac362c77b162830a5b9093")
}

// UnpackFailedDeploymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedDeployment()
func (intermediatedprocessor *Intermediatedprocessor) UnpackFailedDeploymentError(raw []byte) (*IntermediatedprocessorFailedDeployment, error) {
	out := new(IntermediatedprocessorFailedDeployment)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "FailedDeployment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorHoldPeriodCanNotBeZero represents a HoldPeriodCanNotBeZero error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorHoldPeriodCanNotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HoldPeriodCanNotBeZero()
func IntermediatedprocessorHoldPeriodCanNotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x705a71532da8bae84d5c54245bfd200d9655b2c961da65ccc7fcf54a50ad44b4")
}

// UnpackHoldPeriodCanNotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HoldPeriodCanNotBeZero()
func (intermediatedprocessor *Intermediatedprocessor) UnpackHoldPeriodCanNotBeZeroError(raw []byte) (*IntermediatedprocessorHoldPeriodCanNotBeZero, error) {
	out := new(IntermediatedprocessorHoldPeriodCanNotBeZero)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "HoldPeriodCanNotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInsufficientBalance represents a InsufficientBalance error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func IntermediatedprocessorInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInsufficientBalanceError(raw []byte) (*IntermediatedprocessorInsufficientBalance, error) {
	out := new(IntermediatedprocessorInsufficientBalance)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidDisputeResolution represents a InvalidDisputeResolution error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidDisputeResolution struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDisputeResolution()
func IntermediatedprocessorInvalidDisputeResolutionErrorID() common.Hash {
	return common.HexToHash("0x34819f908388d0ed594d20c6802439086d46ba2397dca397160a28d8b2bd98b1")
}

// UnpackInvalidDisputeResolutionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDisputeResolution()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidDisputeResolutionError(raw []byte) (*IntermediatedprocessorInvalidDisputeResolution, error) {
	out := new(IntermediatedprocessorInvalidDisputeResolution)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidDisputeResolution", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidFeeAuthorization represents a InvalidFeeAuthorization error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidFeeAuthorization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeAuthorization()
func IntermediatedprocessorInvalidFeeAuthorizationErrorID() common.Hash {
	return common.HexToHash("0x1735eabec15c7395efafdfa0dda5c74faf3517b604ff55660d6ef0e7457f2c1d")
}

// UnpackInvalidFeeAuthorizationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeAuthorization()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidFeeAuthorizationError(raw []byte) (*IntermediatedprocessorInvalidFeeAuthorization, error) {
	out := new(IntermediatedprocessorInvalidFeeAuthorization)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidFeeAuthorization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidFeeReceiver represents a InvalidFeeReceiver error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidFeeReceiver struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeReceiver()
func IntermediatedprocessorInvalidFeeReceiverErrorID() common.Hash {
	return common.HexToHash("0xd200485c51caaf66763f8b49c9cfa281a0a10132cb56c8fffc35867701d3fc5f")
}

// UnpackInvalidFeeReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeReceiver()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidFeeReceiverError(raw []byte) (*IntermediatedprocessorInvalidFeeReceiver, error) {
	out := new(IntermediatedprocessorInvalidFeeReceiver)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidFeeReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidInvoiceState represents a InvalidInvoiceState error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidInvoiceState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInvoiceState()
func IntermediatedprocessorInvalidInvoiceStateErrorID() common.Hash {
	return common.HexToHash("0x487e4409b34dcf5275ed8908061cfcde1e134270e5620e0eaff4d68605de2cbc")
}

// UnpackInvalidInvoiceStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInvoiceState()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidInvoiceStateError(raw []byte) (*IntermediatedprocessorInvalidInvoiceState, error) {
	out := new(IntermediatedprocessorInvalidInvoiceState)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidInvoiceState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidMetaInvoicePaymentAmount represents a InvalidMetaInvoicePaymentAmount error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidMetaInvoicePaymentAmount struct {
	Sent     *big.Int
	Expected *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMetaInvoicePaymentAmount(uint256 sent, uint256 expected)
func IntermediatedprocessorInvalidMetaInvoicePaymentAmountErrorID() common.Hash {
	return common.HexToHash("0xc7632c7d819e4ec9dbca7ec79df876ac7d5ca98ab46cf285c8f8a7ff52ea72a3")
}

// UnpackInvalidMetaInvoicePaymentAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMetaInvoicePaymentAmount(uint256 sent, uint256 expected)
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidMetaInvoicePaymentAmountError(raw []byte) (*IntermediatedprocessorInvalidMetaInvoicePaymentAmount, error) {
	out := new(IntermediatedprocessorInvalidMetaInvoicePaymentAmount)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidMetaInvoicePaymentAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidNativePayment represents a InvalidNativePayment error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidNativePayment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNativePayment()
func IntermediatedprocessorInvalidNativePaymentErrorID() common.Hash {
	return common.HexToHash("0x214510aac5bc5d45b2314d915edc9aa20e9ec869bcb7e6d50d8d068658a871c9")
}

// UnpackInvalidNativePaymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNativePayment()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidNativePaymentError(raw []byte) (*IntermediatedprocessorInvalidNativePayment, error) {
	out := new(IntermediatedprocessorInvalidNativePayment)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidNativePayment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidOracle represents a InvalidOracle error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidOracle struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidOracle()
func IntermediatedprocessorInvalidOracleErrorID() common.Hash {
	return common.HexToHash("0x9589a27d464cce309224596a505cbfd22e5fde1f0f420cecf8a6b6c1d65791b6")
}

// UnpackInvalidOracleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidOracle()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidOracleError(raw []byte) (*IntermediatedprocessorInvalidOracle, error) {
	out := new(IntermediatedprocessorInvalidOracle)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidOracle", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidPrice represents a InvalidPrice error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPrice()
func IntermediatedprocessorInvalidPriceErrorID() common.Hash {
	return common.HexToHash("0x00bfc9219afe7e8e3b9f14a4708e4cd3d8acb04e325ce992b2a60a58a519683a")
}

// UnpackInvalidPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPrice()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidPriceError(raw []byte) (*IntermediatedprocessorInvalidPrice, error) {
	out := new(IntermediatedprocessorInvalidPrice)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidSeller represents a InvalidSeller error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSeller()
func IntermediatedprocessorInvalidSellerErrorID() common.Hash {
	return common.HexToHash("0xbab7ca35fcde13672ca7744c85f31cdd0a5c3f882f4b4992269c1e7dc56732e9")
}

// UnpackInvalidSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSeller()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidSellerError(raw []byte) (*IntermediatedprocessorInvalidSeller, error) {
	out := new(IntermediatedprocessorInvalidSeller)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvalidSellersPayoutShare represents a InvalidSellersPayoutShare error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvalidSellersPayoutShare struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSellersPayoutShare()
func IntermediatedprocessorInvalidSellersPayoutShareErrorID() common.Hash {
	return common.HexToHash("0x453fb42ddfd1ecde870e9bd55d8b7f21b2333b613ee68779ba5f60498951666d")
}

// UnpackInvalidSellersPayoutShareError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSellersPayoutShare()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvalidSellersPayoutShareError(raw []byte) (*IntermediatedprocessorInvalidSellersPayoutShare, error) {
	out := new(IntermediatedprocessorInvalidSellersPayoutShare)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvalidSellersPayoutShare", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvoiceAlreadyExists represents a InvoiceAlreadyExists error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceAlreadyExists()
func IntermediatedprocessorInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x074bc9355c94925fa82ddb49dcc88f1a666f1d1aa24efbddcdbe5f8d98b7ed59")
}

// UnpackInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceAlreadyExists()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvoiceAlreadyExistsError(raw []byte) (*IntermediatedprocessorInvoiceAlreadyExists, error) {
	out := new(IntermediatedprocessorInvoiceAlreadyExists)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvoiceDoesNotExist represents a InvoiceDoesNotExist error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvoiceDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceDoesNotExist()
func IntermediatedprocessorInvoiceDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x715d9228f420b2c4c07281fb8597619f2ca0c9d8cada84ce032e60ec6407b582")
}

// UnpackInvoiceDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceDoesNotExist()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvoiceDoesNotExistError(raw []byte) (*IntermediatedprocessorInvoiceDoesNotExist, error) {
	out := new(IntermediatedprocessorInvoiceDoesNotExist)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvoiceDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorInvoiceExpired represents a InvoiceExpired error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorInvoiceExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceExpired()
func IntermediatedprocessorInvoiceExpiredErrorID() common.Hash {
	return common.HexToHash("0xf04e9cf09371be6ef375f7f016c1ac94b6c5c6a4d247bec67ae9568dd6b911b6")
}

// UnpackInvoiceExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceExpired()
func (intermediatedprocessor *Intermediatedprocessor) UnpackInvoiceExpiredError(raw []byte) (*IntermediatedprocessorInvoiceExpired, error) {
	out := new(IntermediatedprocessorInvoiceExpired)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "InvoiceExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorMetaInvoiceAlreadyExists represents a MetaInvoiceAlreadyExists error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorMetaInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MetaInvoiceAlreadyExists()
func IntermediatedprocessorMetaInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xb09960c1d49af4579c96dbfb857f76f135da1b549d276e06491c05ec24747201")
}

// UnpackMetaInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MetaInvoiceAlreadyExists()
func (intermediatedprocessor *Intermediatedprocessor) UnpackMetaInvoiceAlreadyExistsError(raw []byte) (*IntermediatedprocessorMetaInvoiceAlreadyExists, error) {
	out := new(IntermediatedprocessorMetaInvoiceAlreadyExists)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "MetaInvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorNotAuthorized represents a NotAuthorized error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func IntermediatedprocessorNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (intermediatedprocessor *Intermediatedprocessor) UnpackNotAuthorizedError(raw []byte) (*IntermediatedprocessorNotAuthorized, error) {
	out := new(IntermediatedprocessorNotAuthorized)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorPriceCannotBeZero represents a PriceCannotBeZero error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorPriceCannotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceCannotBeZero()
func IntermediatedprocessorPriceCannotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x2c669f0ac3409adbbadbe16eaad2cac428e45b3cb8de2f47377f30f8b5729e18")
}

// UnpackPriceCannotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceCannotBeZero()
func (intermediatedprocessor *Intermediatedprocessor) UnpackPriceCannotBeZeroError(raw []byte) (*IntermediatedprocessorPriceCannotBeZero, error) {
	out := new(IntermediatedprocessorPriceCannotBeZero)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "PriceCannotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorPriceIsTooLow represents a PriceIsTooLow error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorPriceIsTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceIsTooLow()
func IntermediatedprocessorPriceIsTooLowErrorID() common.Hash {
	return common.HexToHash("0xdb8db56995596ab6855aa515b34d8c3549b6c6fd6435c2e3af6e0d886de7e87c")
}

// UnpackPriceIsTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceIsTooLow()
func (intermediatedprocessor *Intermediatedprocessor) UnpackPriceIsTooLowError(raw []byte) (*IntermediatedprocessorPriceIsTooLow, error) {
	out := new(IntermediatedprocessorPriceIsTooLow)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "PriceIsTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorReentrancy represents a Reentrancy error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Reentrancy()
func IntermediatedprocessorReentrancyErrorID() common.Hash {
	return common.HexToHash("0xab143c06c9772d69bbbc9f2fe74acd02f810e93b099f3d1dac8448ac9ae35991")
}

// UnpackReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Reentrancy()
func (intermediatedprocessor *Intermediatedprocessor) UnpackReentrancyError(raw []byte) (*IntermediatedprocessorReentrancy, error) {
	out := new(IntermediatedprocessorReentrancy)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "Reentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorSequencerDown represents a SequencerDown error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorSequencerDown struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SequencerDown()
func IntermediatedprocessorSequencerDownErrorID() common.Hash {
	return common.HexToHash("0x032b3d00cfb14fdf4eecb317aaf61db9dd7331083f0db9baa2eae06ec3e15ecb")
}

// UnpackSequencerDownError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SequencerDown()
func (intermediatedprocessor *Intermediatedprocessor) UnpackSequencerDownError(raw []byte) (*IntermediatedprocessorSequencerDown, error) {
	out := new(IntermediatedprocessorSequencerDown)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "SequencerDown", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorStalePrice represents a StalePrice error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorStalePrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StalePrice()
func IntermediatedprocessorStalePriceErrorID() common.Hash {
	return common.HexToHash("0x19abf40e7c2e0280d6137a5d95d9f3793d913552f00d0e25e4ab4388bcc0d573")
}

// UnpackStalePriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StalePrice()
func (intermediatedprocessor *Intermediatedprocessor) UnpackStalePriceError(raw []byte) (*IntermediatedprocessorStalePrice, error) {
	out := new(IntermediatedprocessorStalePrice)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "StalePrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorStalePriceFeed represents a StalePriceFeed error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorStalePriceFeed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StalePriceFeed()
func IntermediatedprocessorStalePriceFeedErrorID() common.Hash {
	return common.HexToHash("0x1087e109db85b72cf66a8dbc341a9e5601a49c3f12e82151b3eb6e742d4a766e")
}

// UnpackStalePriceFeedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StalePriceFeed()
func (intermediatedprocessor *Intermediatedprocessor) UnpackStalePriceFeedError(raw []byte) (*IntermediatedprocessorStalePriceFeed, error) {
	out := new(IntermediatedprocessorStalePriceFeed)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "StalePriceFeed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IntermediatedprocessorUnsupportedToken represents a UnsupportedToken error raised by the Intermediatedprocessor contract.
type IntermediatedprocessorUnsupportedToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedToken()
func IntermediatedprocessorUnsupportedTokenErrorID() common.Hash {
	return common.HexToHash("0x6a1728823cfcc894fe1dcf37bfe71f201fb66b0b61862091f422023e22ea5ab9")
}

// UnpackUnsupportedTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedToken()
func (intermediatedprocessor *Intermediatedprocessor) UnpackUnsupportedTokenError(raw []byte) (*IntermediatedprocessorUnsupportedToken, error) {
	out := new(IntermediatedprocessorUnsupportedToken)
	if err := intermediatedprocessor.abi.UnpackIntoInterface(out, "UnsupportedToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}
