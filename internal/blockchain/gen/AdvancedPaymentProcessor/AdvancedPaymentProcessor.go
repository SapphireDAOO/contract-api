// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package advancedprocessor

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
	InvoiceId     *big.Int
	PaidAt        *big.Int
	CreatedAt     *big.Int
	ReleaseAt     *big.Int
	State         uint8
	MetaInvoiceId *big.Int
	Buyer         common.Address
	Seller        common.Address
	Escrow        common.Address
	PaymentToken  common.Address
	AmountPaid    *big.Int
	Price         *big.Int
	Balance       *big.Int
}

// IAdvancedPaymentProcessorInvoiceCreationParam is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorInvoiceCreationParam struct {
	OrderId          string
	Seller           common.Address
	Price            *big.Int
	EscrowHoldPeriod *big.Int
}

// IAdvancedPaymentProcessorMetaInvoice is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorMetaInvoice struct {
	Price         *big.Int
	SubInvoiceIds []*big.Int
}

// AdvancedprocessorMetaData contains all meta data concerning the Advancedprocessor contract.
var AdvancedprocessorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentProcessorStorageAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nativeTokenAggregatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BuyerCannotBeSeller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateTask\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDisputeResolution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInvoiceState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativePayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSellersPayoutShare\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvoiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvoiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MetaInvoiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceIsTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TaskNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"DisputeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"DisputeDismissed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerAmount\",\"type\":\"uint256\"}],\"name\":\"DisputeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"}],\"name\":\"EscrowCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"InvoiceCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"components\":[{\"internalType\":\"uint216\",\"name\":\"invoiceId\",\"type\":\"uint216\"},{\"internalType\":\"uint40\",\"name\":\"paidAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"createdAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"releaseAt\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint216\",\"name\":\"metaInvoiceId\",\"type\":\"uint216\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"name\":\"invoice\",\"type\":\"tuple\"}],\"name\":\"InvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"escrowAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvoicePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"metaInvoiceId\",\"type\":\"uint216\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"MetaInvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newHoldPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateReleaseTime\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCELED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_DECIMAL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_DISMISSED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_RESOLVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_SETTLED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIATED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUNDED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"cancelInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"computeSalt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"createDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"escrowHoldPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam[]\",\"name\":\"param\",\"type\":\"tuple[]\"}],\"name\":\"createMetaInvoice\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"escrowHoldPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"createSingleInvoice\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"getInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint216\",\"name\":\"invoiceId\",\"type\":\"uint216\"},{\"internalType\":\"uint40\",\"name\":\"paidAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"createdAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"releaseAt\",\"type\":\"uint40\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint216\",\"name\":\"metaInvoiceId\",\"type\":\"uint216\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getItems\",\"outputs\":[{\"internalType\":\"uint216[]\",\"name\":\"\",\"type\":\"uint216[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"metaInvoiceId\",\"type\":\"uint216\"}],\"name\":\"getMetaInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint216[]\",\"name\":\"subInvoiceIds\",\"type\":\"uint216[]\"}],\"internalType\":\"structIAdvancedPaymentProcessor.MetaInvoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextInvoiceId\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextMetaInvoiceId\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"getPredictedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getTokenValueFromUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"internalType\":\"uint8\",\"name\":\"resolution\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sellerShare\",\"type\":\"uint256\"}],\"name\":\"handleDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"payMetaInvoice\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"paySingleInvoice\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"performUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppStorage\",\"outputs\":[{\"internalType\":\"contractIPaymentProcessorStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"internalType\":\"uint256\",\"name\":\"refundShare\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarderAddress\",\"type\":\"address\"}],\"name\":\"setForwarderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"orderId\",\"type\":\"uint216\"},{\"internalType\":\"uint256\",\"name\":\"holdPeriod\",\"type\":\"uint256\"}],\"name\":\"setInvoiceReleaseTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMetaInvoiceCreated\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUniqueInvoiceCreated\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "Advancedprocessor",
}

// Advancedprocessor is an auto generated Go binding around an Ethereum contract.
type Advancedprocessor struct {
	abi abi.ABI
}

// NewAdvancedprocessor creates a new instance of Advancedprocessor.
func NewAdvancedprocessor() *Advancedprocessor {
	parsed, err := AdvancedprocessorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Advancedprocessor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Advancedprocessor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address paymentProcessorStorageAddress, address nativeTokenAggregatorAddress) returns()
func (advancedprocessor *Advancedprocessor) PackConstructor(paymentProcessorStorageAddress common.Address, nativeTokenAggregatorAddress common.Address) []byte {
	enc, err := advancedprocessor.abi.Pack("", paymentProcessorStorageAddress, nativeTokenAggregatorAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBASISPOINTS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (advancedprocessor *Advancedprocessor) PackBASISPOINTS() []byte {
	enc, err := advancedprocessor.abi.Pack("BASIS_POINTS")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBASISPOINTS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (advancedprocessor *Advancedprocessor) UnpackBASISPOINTS(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("BASIS_POINTS", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCANCELED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) PackCANCELED() []byte {
	enc, err := advancedprocessor.abi.Pack("CANCELED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCANCELED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2f4c1da.
//
// Solidity: function CANCELED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackCANCELED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("CANCELED", data)
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
func (advancedprocessor *Advancedprocessor) PackDEFAULTDECIMAL() []byte {
	enc, err := advancedprocessor.abi.Pack("DEFAULT_DECIMAL")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTDECIMAL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackDEFAULTDECIMAL(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("DEFAULT_DECIMAL", data)
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
func (advancedprocessor *Advancedprocessor) PackDISPUTED() []byte {
	enc, err := advancedprocessor.abi.Pack("DISPUTED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x620cac66.
//
// Solidity: function DISPUTED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackDISPUTED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("DISPUTED", data)
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
func (advancedprocessor *Advancedprocessor) PackDISPUTEDISMISSED() []byte {
	enc, err := advancedprocessor.abi.Pack("DISPUTE_DISMISSED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTEDISMISSED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf1201920.
//
// Solidity: function DISPUTE_DISMISSED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackDISPUTEDISMISSED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("DISPUTE_DISMISSED", data)
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
func (advancedprocessor *Advancedprocessor) PackDISPUTERESOLVED() []byte {
	enc, err := advancedprocessor.abi.Pack("DISPUTE_RESOLVED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTERESOLVED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x677a4ece.
//
// Solidity: function DISPUTE_RESOLVED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackDISPUTERESOLVED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("DISPUTE_RESOLVED", data)
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
func (advancedprocessor *Advancedprocessor) PackDISPUTESETTLED() []byte {
	enc, err := advancedprocessor.abi.Pack("DISPUTE_SETTLED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDISPUTESETTLED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x55b9da62.
//
// Solidity: function DISPUTE_SETTLED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackDISPUTESETTLED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("DISPUTE_SETTLED", data)
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
func (advancedprocessor *Advancedprocessor) PackINITIATED() []byte {
	enc, err := advancedprocessor.abi.Pack("INITIATED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackINITIATED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88a03c53.
//
// Solidity: function INITIATED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackINITIATED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("INITIATED", data)
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
func (advancedprocessor *Advancedprocessor) PackPAID() []byte {
	enc, err := advancedprocessor.abi.Pack("PAID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPAID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd31a03c2.
//
// Solidity: function PAID() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackPAID(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("PAID", data)
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
func (advancedprocessor *Advancedprocessor) PackREFUNDED() []byte {
	enc, err := advancedprocessor.abi.Pack("REFUNDED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackREFUNDED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94e15c8f.
//
// Solidity: function REFUNDED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackREFUNDED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("REFUNDED", data)
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
func (advancedprocessor *Advancedprocessor) PackRELEASED() []byte {
	enc, err := advancedprocessor.abi.Pack("RELEASED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRELEASED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b52ea81.
//
// Solidity: function RELEASED() view returns(uint8)
func (advancedprocessor *Advancedprocessor) UnpackRELEASED(data []byte) (uint8, error) {
	out, err := advancedprocessor.abi.Unpack("RELEASED", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackCancelInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde48b793.
//
// Solidity: function cancelInvoice(uint216 orderId) returns()
func (advancedprocessor *Advancedprocessor) PackCancelInvoice(orderId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("cancelInvoice", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCheckUpkeep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool, bytes)
func (advancedprocessor *Advancedprocessor) PackCheckUpkeep(arg0 []byte) []byte {
	enc, err := advancedprocessor.abi.Pack("checkUpkeep", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// CheckUpkeepOutput serves as a container for the return parameters of contract
// method CheckUpkeep.
type CheckUpkeepOutput struct {
	Arg0 bool
	Arg1 []byte
}

// UnpackCheckUpkeep is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool, bytes)
func (advancedprocessor *Advancedprocessor) UnpackCheckUpkeep(data []byte) (CheckUpkeepOutput, error) {
	out, err := advancedprocessor.abi.Unpack("checkUpkeep", data)
	outstruct := new(CheckUpkeepOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, err

}

// PackComputeSalt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a221549.
//
// Solidity: function computeSalt(address seller, address buyer, uint216 orderId) pure returns(bytes32)
func (advancedprocessor *Advancedprocessor) PackComputeSalt(seller common.Address, buyer common.Address, orderId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("computeSalt", seller, buyer, orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackComputeSalt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a221549.
//
// Solidity: function computeSalt(address seller, address buyer, uint216 orderId) pure returns(bytes32)
func (advancedprocessor *Advancedprocessor) UnpackComputeSalt(data []byte) ([32]byte, error) {
	out, err := advancedprocessor.abi.Unpack("computeSalt", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCreateDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6409e54.
//
// Solidity: function createDispute(uint216 orderId) returns()
func (advancedprocessor *Advancedprocessor) PackCreateDispute(orderId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("createDispute", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0edb4521.
//
// Solidity: function createMetaInvoice((string,address,uint256,uint256)[] param) returns(uint216)
func (advancedprocessor *Advancedprocessor) PackCreateMetaInvoice(param []IAdvancedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := advancedprocessor.abi.Pack("createMetaInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0edb4521.
//
// Solidity: function createMetaInvoice((string,address,uint256,uint256)[] param) returns(uint216)
func (advancedprocessor *Advancedprocessor) UnpackCreateMetaInvoice(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("createMetaInvoice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCreateSingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d91ce85.
//
// Solidity: function createSingleInvoice((string,address,uint256,uint256) param) returns(uint216)
func (advancedprocessor *Advancedprocessor) PackCreateSingleInvoice(param IAdvancedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := advancedprocessor.abi.Pack("createSingleInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSingleInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0d91ce85.
//
// Solidity: function createSingleInvoice((string,address,uint256,uint256) param) returns(uint216)
func (advancedprocessor *Advancedprocessor) UnpackCreateSingleInvoice(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("createSingleInvoice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetForwarder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0042526.
//
// Solidity: function getForwarder() view returns(address)
func (advancedprocessor *Advancedprocessor) PackGetForwarder() []byte {
	enc, err := advancedprocessor.abi.Pack("getForwarder")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetForwarder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0042526.
//
// Solidity: function getForwarder() view returns(address)
func (advancedprocessor *Advancedprocessor) UnpackGetForwarder(data []byte) (common.Address, error) {
	out, err := advancedprocessor.abi.Unpack("getForwarder", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cfa3edf.
//
// Solidity: function getInvoice(uint216 orderId) view returns((uint216,uint40,uint40,uint40,uint8,uint216,address,address,address,address,uint256,uint256,uint256))
func (advancedprocessor *Advancedprocessor) PackGetInvoice(orderId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("getInvoice", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4cfa3edf.
//
// Solidity: function getInvoice(uint216 orderId) view returns((uint216,uint40,uint40,uint40,uint8,uint216,address,address,address,address,uint256,uint256,uint256))
func (advancedprocessor *Advancedprocessor) UnpackGetInvoice(data []byte) (IAdvancedPaymentProcessorInvoice, error) {
	out, err := advancedprocessor.abi.Unpack("getInvoice", data)
	if err != nil {
		return *new(IAdvancedPaymentProcessorInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorInvoice)).(*IAdvancedPaymentProcessorInvoice)
	return out0, err
}

// PackGetItems is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410d59cc.
//
// Solidity: function getItems() view returns(uint216[])
func (advancedprocessor *Advancedprocessor) PackGetItems() []byte {
	enc, err := advancedprocessor.abi.Pack("getItems")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetItems is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410d59cc.
//
// Solidity: function getItems() view returns(uint216[])
func (advancedprocessor *Advancedprocessor) UnpackGetItems(data []byte) ([]*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("getItems", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackGetMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33453a3d.
//
// Solidity: function getMetaInvoice(uint216 metaInvoiceId) view returns((uint256,uint216[]))
func (advancedprocessor *Advancedprocessor) PackGetMetaInvoice(metaInvoiceId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("getMetaInvoice", metaInvoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x33453a3d.
//
// Solidity: function getMetaInvoice(uint216 metaInvoiceId) view returns((uint256,uint216[]))
func (advancedprocessor *Advancedprocessor) UnpackGetMetaInvoice(data []byte) (IAdvancedPaymentProcessorMetaInvoice, error) {
	out, err := advancedprocessor.abi.Unpack("getMetaInvoice", data)
	if err != nil {
		return *new(IAdvancedPaymentProcessorMetaInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorMetaInvoice)).(*IAdvancedPaymentProcessorMetaInvoice)
	return out0, err
}

// PackGetNextInvoiceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint216)
func (advancedprocessor *Advancedprocessor) PackGetNextInvoiceId() []byte {
	enc, err := advancedprocessor.abi.Pack("getNextInvoiceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint216)
func (advancedprocessor *Advancedprocessor) UnpackGetNextInvoiceId(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("getNextInvoiceId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetNextMetaInvoiceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint216)
func (advancedprocessor *Advancedprocessor) PackGetNextMetaInvoiceId() []byte {
	enc, err := advancedprocessor.abi.Pack("getNextMetaInvoiceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextMetaInvoiceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x938bad50.
//
// Solidity: function getNextMetaInvoiceId() view returns(uint216)
func (advancedprocessor *Advancedprocessor) UnpackGetNextMetaInvoiceId(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("getNextMetaInvoiceId", data)
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
func (advancedprocessor *Advancedprocessor) PackGetPredictedAddress(salt [32]byte) []byte {
	enc, err := advancedprocessor.abi.Pack("getPredictedAddress", salt)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPredictedAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4bec0c01.
//
// Solidity: function getPredictedAddress(bytes32 salt) view returns(address)
func (advancedprocessor *Advancedprocessor) UnpackGetPredictedAddress(data []byte) (common.Address, error) {
	out, err := advancedprocessor.abi.Unpack("getPredictedAddress", data)
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
func (advancedprocessor *Advancedprocessor) PackGetTokenValueFromUsd(paymentToken common.Address, price *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("getTokenValueFromUsd", paymentToken, price)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTokenValueFromUsd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88516807.
//
// Solidity: function getTokenValueFromUsd(address paymentToken, uint256 price) view returns(uint256)
func (advancedprocessor *Advancedprocessor) UnpackGetTokenValueFromUsd(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("getTokenValueFromUsd", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackHandleDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf00aaec6.
//
// Solidity: function handleDispute(uint216 orderId, uint8 resolution, uint256 sellerShare) returns()
func (advancedprocessor *Advancedprocessor) PackHandleDispute(orderId *big.Int, resolution uint8, sellerShare *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("handleDispute", orderId, resolution, sellerShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPayMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1114e4b.
//
// Solidity: function payMetaInvoice(uint216 orderId, address paymentToken) payable returns()
func (advancedprocessor *Advancedprocessor) PackPayMetaInvoice(orderId *big.Int, paymentToken common.Address) []byte {
	enc, err := advancedprocessor.abi.Pack("payMetaInvoice", orderId, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPaySingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x053d90c2.
//
// Solidity: function paySingleInvoice(uint216 orderId, address paymentToken) payable returns()
func (advancedprocessor *Advancedprocessor) PackPaySingleInvoice(orderId *big.Int, paymentToken common.Address) []byte {
	enc, err := advancedprocessor.abi.Pack("paySingleInvoice", orderId, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPerformUpkeep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (advancedprocessor *Advancedprocessor) PackPerformUpkeep(arg0 []byte) []byte {
	enc, err := advancedprocessor.abi.Pack("performUpkeep", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (advancedprocessor *Advancedprocessor) PackPpStorage() []byte {
	enc, err := advancedprocessor.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (advancedprocessor *Advancedprocessor) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := advancedprocessor.abi.Unpack("ppStorage", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackRefund is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5da97b5e.
//
// Solidity: function refund(uint216 orderId, uint256 refundShare) returns()
func (advancedprocessor *Advancedprocessor) PackRefund(orderId *big.Int, refundShare *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("refund", orderId, refundShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb990809.
//
// Solidity: function release(uint216 orderId) returns()
func (advancedprocessor *Advancedprocessor) PackRelease(orderId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("release", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackResolveDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5168e613.
//
// Solidity: function resolveDispute(uint216 orderId) returns()
func (advancedprocessor *Advancedprocessor) PackResolveDispute(orderId *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("resolveDispute", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForwarderAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd777cc6d.
//
// Solidity: function setForwarderAddress(address forwarderAddress) returns()
func (advancedprocessor *Advancedprocessor) PackSetForwarderAddress(forwarderAddress common.Address) []byte {
	enc, err := advancedprocessor.abi.Pack("setForwarderAddress", forwarderAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetInvoiceReleaseTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50e5d5cf.
//
// Solidity: function setInvoiceReleaseTime(uint216 orderId, uint256 holdPeriod) returns()
func (advancedprocessor *Advancedprocessor) PackSetInvoiceReleaseTime(orderId *big.Int, holdPeriod *big.Int) []byte {
	enc, err := advancedprocessor.abi.Pack("setInvoiceReleaseTime", orderId, holdPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPriceFeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address aggregator) returns()
func (advancedprocessor *Advancedprocessor) PackSetPriceFeed(token common.Address, aggregator common.Address) []byte {
	enc, err := advancedprocessor.abi.Pack("setPriceFeed", token, aggregator)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalMetaInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint216)
func (advancedprocessor *Advancedprocessor) PackTotalMetaInvoiceCreated() []byte {
	enc, err := advancedprocessor.abi.Pack("totalMetaInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalMetaInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66cbb1bd.
//
// Solidity: function totalMetaInvoiceCreated() view returns(uint216)
func (advancedprocessor *Advancedprocessor) UnpackTotalMetaInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("totalMetaInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTotalUniqueInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint216)
func (advancedprocessor *Advancedprocessor) PackTotalUniqueInvoiceCreated() []byte {
	enc, err := advancedprocessor.abi.Pack("totalUniqueInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalUniqueInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x81946608.
//
// Solidity: function totalUniqueInvoiceCreated() view returns(uint216)
func (advancedprocessor *Advancedprocessor) UnpackTotalUniqueInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := advancedprocessor.abi.Unpack("totalUniqueInvoiceCreated", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// AdvancedprocessorDisputeCreated represents a DisputeCreated event raised by the Advancedprocessor contract.
type AdvancedprocessorDisputeCreated struct {
	OrderId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorDisputeCreatedEventName = "DisputeCreated"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorDisputeCreated) ContractEventName() string {
	return AdvancedprocessorDisputeCreatedEventName
}

// UnpackDisputeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeCreated(uint216 indexed orderId)
func (advancedprocessor *Advancedprocessor) UnpackDisputeCreatedEvent(log *types.Log) (*AdvancedprocessorDisputeCreated, error) {
	event := "DisputeCreated"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorDisputeCreated)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorDisputeDismissed represents a DisputeDismissed event raised by the Advancedprocessor contract.
type AdvancedprocessorDisputeDismissed struct {
	OrderId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorDisputeDismissedEventName = "DisputeDismissed"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorDisputeDismissed) ContractEventName() string {
	return AdvancedprocessorDisputeDismissedEventName
}

// UnpackDisputeDismissedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeDismissed(uint216 indexed orderId)
func (advancedprocessor *Advancedprocessor) UnpackDisputeDismissedEvent(log *types.Log) (*AdvancedprocessorDisputeDismissed, error) {
	event := "DisputeDismissed"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorDisputeDismissed)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorDisputeResolved represents a DisputeResolved event raised by the Advancedprocessor contract.
type AdvancedprocessorDisputeResolved struct {
	OrderId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorDisputeResolvedEventName = "DisputeResolved"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorDisputeResolved) ContractEventName() string {
	return AdvancedprocessorDisputeResolvedEventName
}

// UnpackDisputeResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeResolved(uint216 indexed orderId)
func (advancedprocessor *Advancedprocessor) UnpackDisputeResolvedEvent(log *types.Log) (*AdvancedprocessorDisputeResolved, error) {
	event := "DisputeResolved"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorDisputeResolved)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorDisputeSettled represents a DisputeSettled event raised by the Advancedprocessor contract.
type AdvancedprocessorDisputeSettled struct {
	OrderId      *big.Int
	SellerAmount *big.Int
	BuyerAmount  *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorDisputeSettledEventName = "DisputeSettled"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorDisputeSettled) ContractEventName() string {
	return AdvancedprocessorDisputeSettledEventName
}

// UnpackDisputeSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeSettled(uint216 indexed orderId, uint256 sellerAmount, uint256 buyerAmount)
func (advancedprocessor *Advancedprocessor) UnpackDisputeSettledEvent(log *types.Log) (*AdvancedprocessorDisputeSettled, error) {
	event := "DisputeSettled"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorDisputeSettled)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorEscrowCreated represents a EscrowCreated event raised by the Advancedprocessor contract.
type AdvancedprocessorEscrowCreated struct {
	OrderId *big.Int
	Escrow  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorEscrowCreatedEventName = "EscrowCreated"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorEscrowCreated) ContractEventName() string {
	return AdvancedprocessorEscrowCreatedEventName
}

// UnpackEscrowCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EscrowCreated(uint216 indexed orderId, address indexed escrow)
func (advancedprocessor *Advancedprocessor) UnpackEscrowCreatedEvent(log *types.Log) (*AdvancedprocessorEscrowCreated, error) {
	event := "EscrowCreated"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorEscrowCreated)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Advancedprocessor contract.
type AdvancedprocessorInvoiceCanceled struct {
	OrderId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorInvoiceCanceled) ContractEventName() string {
	return AdvancedprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(uint216 indexed orderId)
func (advancedprocessor *Advancedprocessor) UnpackInvoiceCanceledEvent(log *types.Log) (*AdvancedprocessorInvoiceCanceled, error) {
	event := "InvoiceCanceled"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorInvoiceCanceled)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorInvoiceCreated represents a InvoiceCreated event raised by the Advancedprocessor contract.
type AdvancedprocessorInvoiceCreated struct {
	OrderId *big.Int
	Invoice IAdvancedPaymentProcessorInvoice
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorInvoiceCreated) ContractEventName() string {
	return AdvancedprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(uint216 indexed orderId, (uint216,uint40,uint40,uint40,uint8,uint216,address,address,address,address,uint256,uint256,uint256) invoice)
func (advancedprocessor *Advancedprocessor) UnpackInvoiceCreatedEvent(log *types.Log) (*AdvancedprocessorInvoiceCreated, error) {
	event := "InvoiceCreated"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorInvoiceCreated)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorInvoicePaid represents a InvoicePaid event raised by the Advancedprocessor contract.
type AdvancedprocessorInvoicePaid struct {
	OrderId       *big.Int
	PaymentToken  common.Address
	EscrowAddress common.Address
	Amount        *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorInvoicePaidEventName = "InvoicePaid"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorInvoicePaid) ContractEventName() string {
	return AdvancedprocessorInvoicePaidEventName
}

// UnpackInvoicePaidEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoicePaid(uint216 indexed orderId, address paymentToken, address escrowAddress, uint256 amount)
func (advancedprocessor *Advancedprocessor) UnpackInvoicePaidEvent(log *types.Log) (*AdvancedprocessorInvoicePaid, error) {
	event := "InvoicePaid"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorInvoicePaid)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorMetaInvoiceCreated represents a MetaInvoiceCreated event raised by the Advancedprocessor contract.
type AdvancedprocessorMetaInvoiceCreated struct {
	MetaInvoiceId *big.Int
	TotalPrice    *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorMetaInvoiceCreatedEventName = "MetaInvoiceCreated"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorMetaInvoiceCreated) ContractEventName() string {
	return AdvancedprocessorMetaInvoiceCreatedEventName
}

// UnpackMetaInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetaInvoiceCreated(uint216 indexed metaInvoiceId, uint256 indexed totalPrice)
func (advancedprocessor *Advancedprocessor) UnpackMetaInvoiceCreatedEvent(log *types.Log) (*AdvancedprocessorMetaInvoiceCreated, error) {
	event := "MetaInvoiceCreated"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorMetaInvoiceCreated)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorPaymentReleased represents a PaymentReleased event raised by the Advancedprocessor contract.
type AdvancedprocessorPaymentReleased struct {
	OrderId      *big.Int
	SellerAmount *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorPaymentReleasedEventName = "PaymentReleased"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorPaymentReleased) ContractEventName() string {
	return AdvancedprocessorPaymentReleasedEventName
}

// UnpackPaymentReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentReleased(uint216 indexed orderId, uint256 sellerAmount)
func (advancedprocessor *Advancedprocessor) UnpackPaymentReleasedEvent(log *types.Log) (*AdvancedprocessorPaymentReleased, error) {
	event := "PaymentReleased"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorPaymentReleased)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorRefunded represents a Refunded event raised by the Advancedprocessor contract.
type AdvancedprocessorRefunded struct {
	OrderId *big.Int
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorRefundedEventName = "Refunded"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorRefunded) ContractEventName() string {
	return AdvancedprocessorRefundedEventName
}

// UnpackRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Refunded(uint216 indexed orderId, uint256 indexed amount)
func (advancedprocessor *Advancedprocessor) UnpackRefundedEvent(log *types.Log) (*AdvancedprocessorRefunded, error) {
	event := "Refunded"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorRefunded)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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

// AdvancedprocessorUpdateReleaseTime represents a UpdateReleaseTime event raised by the Advancedprocessor contract.
type AdvancedprocessorUpdateReleaseTime struct {
	OrderId       *big.Int
	NewHoldPeriod *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const AdvancedprocessorUpdateReleaseTimeEventName = "UpdateReleaseTime"

// ContractEventName returns the user-defined event name.
func (AdvancedprocessorUpdateReleaseTime) ContractEventName() string {
	return AdvancedprocessorUpdateReleaseTimeEventName
}

// UnpackUpdateReleaseTimeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UpdateReleaseTime(uint216 indexed orderId, uint256 newHoldPeriod)
func (advancedprocessor *Advancedprocessor) UnpackUpdateReleaseTimeEvent(log *types.Log) (*AdvancedprocessorUpdateReleaseTime, error) {
	event := "UpdateReleaseTime"
	if log.Topics[0] != advancedprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AdvancedprocessorUpdateReleaseTime)
	if len(log.Data) > 0 {
		if err := advancedprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range advancedprocessor.abi.Events[event].Inputs {
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
func (advancedprocessor *Advancedprocessor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["BuyerCannotBeSeller"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackBuyerCannotBeSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["DuplicateTask"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackDuplicateTaskError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvalidDisputeResolution"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvalidDisputeResolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvalidInvoiceStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvalidNativePayment"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvalidNativePaymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvalidPaymentToken"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvalidPaymentTokenError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvalidSellersPayoutShare"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvalidSellersPayoutShareError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["InvoiceDoesNotExist"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackInvoiceDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["MetaInvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackMetaInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["PriceCannotBeZero"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackPriceCannotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["PriceIsTooLow"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackPriceIsTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], advancedprocessor.abi.Errors["TaskNotFound"].ID.Bytes()[:4]) {
		return advancedprocessor.UnpackTaskNotFoundError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AdvancedprocessorBuyerCannotBeSeller represents a BuyerCannotBeSeller error raised by the Advancedprocessor contract.
type AdvancedprocessorBuyerCannotBeSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyerCannotBeSeller()
func AdvancedprocessorBuyerCannotBeSellerErrorID() common.Hash {
	return common.HexToHash("0xb12e242105ea4b2bcdc745efefe14be5558f5f16020ec252980cefc86c6a7a77")
}

// UnpackBuyerCannotBeSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyerCannotBeSeller()
func (advancedprocessor *Advancedprocessor) UnpackBuyerCannotBeSellerError(raw []byte) (*AdvancedprocessorBuyerCannotBeSeller, error) {
	out := new(AdvancedprocessorBuyerCannotBeSeller)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "BuyerCannotBeSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorDuplicateTask represents a DuplicateTask error raised by the Advancedprocessor contract.
type AdvancedprocessorDuplicateTask struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateTask()
func AdvancedprocessorDuplicateTaskErrorID() common.Hash {
	return common.HexToHash("0x6b22feb9606cb284058f0a2f05a53980401948118942ff234189d20361fe4e93")
}

// UnpackDuplicateTaskError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateTask()
func (advancedprocessor *Advancedprocessor) UnpackDuplicateTaskError(raw []byte) (*AdvancedprocessorDuplicateTask, error) {
	out := new(AdvancedprocessorDuplicateTask)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "DuplicateTask", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInsufficientBalance represents a InsufficientBalance error raised by the Advancedprocessor contract.
type AdvancedprocessorInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func AdvancedprocessorInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (advancedprocessor *Advancedprocessor) UnpackInsufficientBalanceError(raw []byte) (*AdvancedprocessorInsufficientBalance, error) {
	out := new(AdvancedprocessorInsufficientBalance)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvalidDisputeResolution represents a InvalidDisputeResolution error raised by the Advancedprocessor contract.
type AdvancedprocessorInvalidDisputeResolution struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDisputeResolution()
func AdvancedprocessorInvalidDisputeResolutionErrorID() common.Hash {
	return common.HexToHash("0x34819f908388d0ed594d20c6802439086d46ba2397dca397160a28d8b2bd98b1")
}

// UnpackInvalidDisputeResolutionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDisputeResolution()
func (advancedprocessor *Advancedprocessor) UnpackInvalidDisputeResolutionError(raw []byte) (*AdvancedprocessorInvalidDisputeResolution, error) {
	out := new(AdvancedprocessorInvalidDisputeResolution)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvalidDisputeResolution", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvalidInvoiceState represents a InvalidInvoiceState error raised by the Advancedprocessor contract.
type AdvancedprocessorInvalidInvoiceState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInvoiceState()
func AdvancedprocessorInvalidInvoiceStateErrorID() common.Hash {
	return common.HexToHash("0x487e4409b34dcf5275ed8908061cfcde1e134270e5620e0eaff4d68605de2cbc")
}

// UnpackInvalidInvoiceStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInvoiceState()
func (advancedprocessor *Advancedprocessor) UnpackInvalidInvoiceStateError(raw []byte) (*AdvancedprocessorInvalidInvoiceState, error) {
	out := new(AdvancedprocessorInvalidInvoiceState)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvalidInvoiceState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvalidNativePayment represents a InvalidNativePayment error raised by the Advancedprocessor contract.
type AdvancedprocessorInvalidNativePayment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNativePayment()
func AdvancedprocessorInvalidNativePaymentErrorID() common.Hash {
	return common.HexToHash("0x214510aac5bc5d45b2314d915edc9aa20e9ec869bcb7e6d50d8d068658a871c9")
}

// UnpackInvalidNativePaymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNativePayment()
func (advancedprocessor *Advancedprocessor) UnpackInvalidNativePaymentError(raw []byte) (*AdvancedprocessorInvalidNativePayment, error) {
	out := new(AdvancedprocessorInvalidNativePayment)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvalidNativePayment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvalidPaymentToken represents a InvalidPaymentToken error raised by the Advancedprocessor contract.
type AdvancedprocessorInvalidPaymentToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentToken()
func AdvancedprocessorInvalidPaymentTokenErrorID() common.Hash {
	return common.HexToHash("0x56e7ec5fb32a6b8680de2f9f0c9db77ebacbce2dde5a21c3bd1e84f57d51c79c")
}

// UnpackInvalidPaymentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentToken()
func (advancedprocessor *Advancedprocessor) UnpackInvalidPaymentTokenError(raw []byte) (*AdvancedprocessorInvalidPaymentToken, error) {
	out := new(AdvancedprocessorInvalidPaymentToken)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvalidPaymentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvalidSellersPayoutShare represents a InvalidSellersPayoutShare error raised by the Advancedprocessor contract.
type AdvancedprocessorInvalidSellersPayoutShare struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSellersPayoutShare()
func AdvancedprocessorInvalidSellersPayoutShareErrorID() common.Hash {
	return common.HexToHash("0x453fb42ddfd1ecde870e9bd55d8b7f21b2333b613ee68779ba5f60498951666d")
}

// UnpackInvalidSellersPayoutShareError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSellersPayoutShare()
func (advancedprocessor *Advancedprocessor) UnpackInvalidSellersPayoutShareError(raw []byte) (*AdvancedprocessorInvalidSellersPayoutShare, error) {
	out := new(AdvancedprocessorInvalidSellersPayoutShare)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvalidSellersPayoutShare", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvoiceAlreadyExists represents a InvoiceAlreadyExists error raised by the Advancedprocessor contract.
type AdvancedprocessorInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceAlreadyExists()
func AdvancedprocessorInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x074bc9355c94925fa82ddb49dcc88f1a666f1d1aa24efbddcdbe5f8d98b7ed59")
}

// UnpackInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceAlreadyExists()
func (advancedprocessor *Advancedprocessor) UnpackInvoiceAlreadyExistsError(raw []byte) (*AdvancedprocessorInvoiceAlreadyExists, error) {
	out := new(AdvancedprocessorInvoiceAlreadyExists)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorInvoiceDoesNotExist represents a InvoiceDoesNotExist error raised by the Advancedprocessor contract.
type AdvancedprocessorInvoiceDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceDoesNotExist()
func AdvancedprocessorInvoiceDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x715d9228f420b2c4c07281fb8597619f2ca0c9d8cada84ce032e60ec6407b582")
}

// UnpackInvoiceDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceDoesNotExist()
func (advancedprocessor *Advancedprocessor) UnpackInvoiceDoesNotExistError(raw []byte) (*AdvancedprocessorInvoiceDoesNotExist, error) {
	out := new(AdvancedprocessorInvoiceDoesNotExist)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "InvoiceDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorMetaInvoiceAlreadyExists represents a MetaInvoiceAlreadyExists error raised by the Advancedprocessor contract.
type AdvancedprocessorMetaInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MetaInvoiceAlreadyExists()
func AdvancedprocessorMetaInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xb09960c1d49af4579c96dbfb857f76f135da1b549d276e06491c05ec24747201")
}

// UnpackMetaInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MetaInvoiceAlreadyExists()
func (advancedprocessor *Advancedprocessor) UnpackMetaInvoiceAlreadyExistsError(raw []byte) (*AdvancedprocessorMetaInvoiceAlreadyExists, error) {
	out := new(AdvancedprocessorMetaInvoiceAlreadyExists)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "MetaInvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorNotAuthorized represents a NotAuthorized error raised by the Advancedprocessor contract.
type AdvancedprocessorNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func AdvancedprocessorNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (advancedprocessor *Advancedprocessor) UnpackNotAuthorizedError(raw []byte) (*AdvancedprocessorNotAuthorized, error) {
	out := new(AdvancedprocessorNotAuthorized)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorPriceCannotBeZero represents a PriceCannotBeZero error raised by the Advancedprocessor contract.
type AdvancedprocessorPriceCannotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceCannotBeZero()
func AdvancedprocessorPriceCannotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x2c669f0ac3409adbbadbe16eaad2cac428e45b3cb8de2f47377f30f8b5729e18")
}

// UnpackPriceCannotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceCannotBeZero()
func (advancedprocessor *Advancedprocessor) UnpackPriceCannotBeZeroError(raw []byte) (*AdvancedprocessorPriceCannotBeZero, error) {
	out := new(AdvancedprocessorPriceCannotBeZero)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "PriceCannotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorPriceIsTooLow represents a PriceIsTooLow error raised by the Advancedprocessor contract.
type AdvancedprocessorPriceIsTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceIsTooLow()
func AdvancedprocessorPriceIsTooLowErrorID() common.Hash {
	return common.HexToHash("0xdb8db56995596ab6855aa515b34d8c3549b6c6fd6435c2e3af6e0d886de7e87c")
}

// UnpackPriceIsTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceIsTooLow()
func (advancedprocessor *Advancedprocessor) UnpackPriceIsTooLowError(raw []byte) (*AdvancedprocessorPriceIsTooLow, error) {
	out := new(AdvancedprocessorPriceIsTooLow)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "PriceIsTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AdvancedprocessorTaskNotFound represents a TaskNotFound error raised by the Advancedprocessor contract.
type AdvancedprocessorTaskNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TaskNotFound()
func AdvancedprocessorTaskNotFoundErrorID() common.Hash {
	return common.HexToHash("0xc325ae33d18e47931adbda2584c56fef1d3e5e64beab80da59968e1c83c84937")
}

// UnpackTaskNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TaskNotFound()
func (advancedprocessor *Advancedprocessor) UnpackTaskNotFoundError(raw []byte) (*AdvancedprocessorTaskNotFound, error) {
	out := new(AdvancedprocessorTaskNotFound)
	if err := advancedprocessor.abi.UnpackIntoInterface(out, "TaskNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}
