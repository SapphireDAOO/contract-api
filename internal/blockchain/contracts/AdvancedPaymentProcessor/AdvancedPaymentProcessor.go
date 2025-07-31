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
	InvoiceId     *big.Int
	Buyer         common.Address
	Seller        common.Address
	Escrow        common.Address
	PaymentToken  common.Address
	State         uint8
	PaidAt        *big.Int
	CreatedAt     *big.Int
	AmountPaid    *big.Int
	Price         *big.Int
	Balance       *big.Int
	MetaInvoiceId [32]byte
}

// IAdvancedPaymentProcessorInvoiceCreationParam is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorInvoiceCreationParam struct {
	OrderId string
	Seller  common.Address
	Price   *big.Int
}

// IAdvancedPaymentProcessorMetaInvoice is an auto generated low-level Go binding around an user-defined struct.
type IAdvancedPaymentProcessorMetaInvoice struct {
	Price         *big.Int
	SubInvoiceIds [][32]byte
}

// PaymentprocessorMetaData contains all meta data concerning the Paymentprocessor contract.
var PaymentprocessorMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentProcessorStorageAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplaceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nativeTokenAggregatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyerCannotBeSeller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDisputeResolution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInvoiceState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativePayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSellersPayoutShare\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvoiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvoiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MetaInvoiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriceCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"DisputeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"DisputeDismissed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerAmount\",\"type\":\"uint256\"}],\"name\":\"DisputeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"}],\"name\":\"EscrowCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"InvoiceCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint48\",\"name\":\"paidAt\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"createdAt\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metaInvoiceId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"name\":\"invoice\",\"type\":\"tuple\"}],\"name\":\"InvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"escrowAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvoicePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"metaInvoiceId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"MetaInvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellerAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerAmount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCELED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_DECIMAL\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_DISMISSED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_RESOLVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_SETTLED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIATED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUNDED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"cancelInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"computeSalt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"createDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam[]\",\"name\":\"param\",\"type\":\"tuple[]\"}],\"name\":\"createMetaInvoice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structIAdvancedPaymentProcessor.InvoiceCreationParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"createSingleInvoice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"getInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint48\",\"name\":\"paidAt\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"createdAt\",\"type\":\"uint48\"},{\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metaInvoiceId\",\"type\":\"bytes32\"}],\"internalType\":\"structIAdvancedPaymentProcessor.Invoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"metaInvoiceId\",\"type\":\"bytes32\"}],\"name\":\"getMetaInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"subInvoiceIds\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIAdvancedPaymentProcessor.MetaInvoice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextInvoiceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextMetaInvoiceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"getPredictedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getTokenValueFromUsd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"resolution\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sellerShare\",\"type\":\"uint256\"}],\"name\":\"handleDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"payMetaInvoice\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"paySingleInvoice\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppStorage\",\"outputs\":[{\"internalType\":\"contractIPaymentProcessorStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refundShare\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sellerShare\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketplaceAddress\",\"type\":\"address\"}],\"name\":\"setMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMetaInvoiceCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUniqueInvoiceCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// PackCancelInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0aa7e2e.
//
// Solidity: function cancelInvoice(bytes32 orderId) returns()
func (paymentprocessor *Paymentprocessor) PackCancelInvoice(orderId [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("cancelInvoice", orderId)
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
// Solidity: function computeSalt(address seller, address buyer, bytes32 orderId) pure returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackComputeSalt(seller common.Address, buyer common.Address, orderId [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("computeSalt", seller, buyer, orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackComputeSalt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c074054.
//
// Solidity: function computeSalt(address seller, address buyer, bytes32 orderId) pure returns(bytes32)
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
// Solidity: function createDispute(bytes32 orderId) returns()
func (paymentprocessor *Paymentprocessor) PackCreateDispute(orderId [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("createDispute", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateMetaInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fc8236a.
//
// Solidity: function createMetaInvoice((string,address,uint256)[] param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackCreateMetaInvoice(param []IAdvancedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := paymentprocessor.abi.Pack("createMetaInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7fc8236a.
//
// Solidity: function createMetaInvoice((string,address,uint256)[] param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) UnpackCreateMetaInvoice(data []byte) ([32]byte, error) {
	out, err := paymentprocessor.abi.Unpack("createMetaInvoice", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackCreateSingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x458e0be9.
//
// Solidity: function createSingleInvoice((string,address,uint256) param) returns(bytes32)
func (paymentprocessor *Paymentprocessor) PackCreateSingleInvoice(param IAdvancedPaymentProcessorInvoiceCreationParam) []byte {
	enc, err := paymentprocessor.abi.Pack("createSingleInvoice", param)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateSingleInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x458e0be9.
//
// Solidity: function createSingleInvoice((string,address,uint256) param) returns(bytes32)
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
// Solidity: function getInvoice(bytes32 orderId) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint256,uint256,uint256,bytes32))
func (paymentprocessor *Paymentprocessor) PackGetInvoice(orderId [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("getInvoice", orderId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcb802c8b.
//
// Solidity: function getInvoice(bytes32 orderId) view returns((uint256,address,address,address,address,uint8,uint48,uint48,uint256,uint256,uint256,bytes32))
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
// Solidity: function getMetaInvoice(bytes32 metaInvoiceId) view returns((uint256,bytes32[]))
func (paymentprocessor *Paymentprocessor) PackGetMetaInvoice(metaInvoiceId [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("getMetaInvoice", metaInvoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMetaInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ccef592.
//
// Solidity: function getMetaInvoice(bytes32 metaInvoiceId) view returns((uint256,bytes32[]))
func (paymentprocessor *Paymentprocessor) UnpackGetMetaInvoice(data []byte) (IAdvancedPaymentProcessorMetaInvoice, error) {
	out, err := paymentprocessor.abi.Unpack("getMetaInvoice", data)
	if err != nil {
		return *new(IAdvancedPaymentProcessorMetaInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(IAdvancedPaymentProcessorMetaInvoice)).(*IAdvancedPaymentProcessorMetaInvoice)
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

// PackHandleDispute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc768634f.
//
// Solidity: function handleDispute(bytes32 orderId, uint8 resolution, uint256 sellerShare) returns()
func (paymentprocessor *Paymentprocessor) PackHandleDispute(orderId [32]byte, resolution uint8, sellerShare *big.Int) []byte {
	enc, err := paymentprocessor.abi.Pack("handleDispute", orderId, resolution, sellerShare)
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
// Solidity: function payMetaInvoice(bytes32 orderId, address paymentToken) payable returns()
func (paymentprocessor *Paymentprocessor) PackPayMetaInvoice(orderId [32]byte, paymentToken common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("payMetaInvoice", orderId, paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPaySingleInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91a61baf.
//
// Solidity: function paySingleInvoice(bytes32 orderId, address paymentToken) payable returns()
func (paymentprocessor *Paymentprocessor) PackPaySingleInvoice(orderId [32]byte, paymentToken common.Address) []byte {
	enc, err := paymentprocessor.abi.Pack("paySingleInvoice", orderId, paymentToken)
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

// PackRefund is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x695eda19.
//
// Solidity: function refund(bytes32 orderId, uint256 refundShare) returns()
func (paymentprocessor *Paymentprocessor) PackRefund(orderId [32]byte, refundShare *big.Int) []byte {
	enc, err := paymentprocessor.abi.Pack("refund", orderId, refundShare)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66afd8ef.
//
// Solidity: function release(bytes32 orderId, uint256 sellerShare) returns()
func (paymentprocessor *Paymentprocessor) PackRelease(orderId [32]byte, sellerShare *big.Int) []byte {
	enc, err := paymentprocessor.abi.Pack("release", orderId, sellerShare)
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
// the contract method with ID 0xc8218e3e.
//
// Solidity: function resolveDispute(bytes32 orderId) returns()
func (paymentprocessor *Paymentprocessor) PackResolveDispute(orderId [32]byte) []byte {
	enc, err := paymentprocessor.abi.Pack("resolveDispute", orderId)
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

// PaymentprocessorDisputeCreated represents a DisputeCreated event raised by the Paymentprocessor contract.
type PaymentprocessorDisputeCreated struct {
	OrderId [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeCreatedEventName = "DisputeCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeCreated) ContractEventName() string {
	return PaymentprocessorDisputeCreatedEventName
}

// UnpackDisputeCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeCreated(bytes32 indexed orderId)
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
	OrderId [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeDismissedEventName = "DisputeDismissed"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeDismissed) ContractEventName() string {
	return PaymentprocessorDisputeDismissedEventName
}

// UnpackDisputeDismissedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeDismissed(bytes32 indexed orderId)
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
	OrderId [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorDisputeResolvedEventName = "DisputeResolved"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorDisputeResolved) ContractEventName() string {
	return PaymentprocessorDisputeResolvedEventName
}

// UnpackDisputeResolvedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeResolved(bytes32 indexed orderId)
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
	OrderId      [32]byte
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
// Solidity: event DisputeSettled(bytes32 indexed orderId, uint256 sellerAmount, uint256 buyerAmount)
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
	OrderId [32]byte
	Escrow  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorEscrowCreatedEventName = "EscrowCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorEscrowCreated) ContractEventName() string {
	return PaymentprocessorEscrowCreatedEventName
}

// UnpackEscrowCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EscrowCreated(bytes32 indexed orderId, address indexed escrow)
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

// PaymentprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceCanceled struct {
	OrderId [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoiceCanceled) ContractEventName() string {
	return PaymentprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(bytes32 indexed orderId)
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
	OrderId [32]byte
	Invoice IAdvancedPaymentProcessorInvoice
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorInvoiceCreated) ContractEventName() string {
	return PaymentprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(bytes32 indexed orderId, (uint256,address,address,address,address,uint8,uint48,uint48,uint256,uint256,uint256,bytes32) invoice)
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
	OrderId       [32]byte
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
// Solidity: event InvoicePaid(bytes32 indexed orderId, address paymentToken, address escrowAddress, uint256 amount)
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

// PaymentprocessorMetaInvoiceCreated represents a MetaInvoiceCreated event raised by the Paymentprocessor contract.
type PaymentprocessorMetaInvoiceCreated struct {
	MetaInvoiceId [32]byte
	TotalPrice    *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorMetaInvoiceCreatedEventName = "MetaInvoiceCreated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorMetaInvoiceCreated) ContractEventName() string {
	return PaymentprocessorMetaInvoiceCreatedEventName
}

// UnpackMetaInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MetaInvoiceCreated(bytes32 indexed metaInvoiceId, uint256 indexed totalPrice)
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
	OrderId      [32]byte
	SellerAmount *big.Int
	BuyerAmount  *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorPaymentReleasedEventName = "PaymentReleased"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorPaymentReleased) ContractEventName() string {
	return PaymentprocessorPaymentReleasedEventName
}

// UnpackPaymentReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentReleased(bytes32 indexed orderId, uint256 sellerAmount, uint256 buyerAmount)
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

// PaymentprocessorRefunded represents a Refunded event raised by the Paymentprocessor contract.
type PaymentprocessorRefunded struct {
	OrderId [32]byte
	Amount  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorRefundedEventName = "Refunded"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorRefunded) ContractEventName() string {
	return PaymentprocessorRefundedEventName
}

// UnpackRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Refunded(bytes32 indexed orderId, uint256 indexed amount)
func (paymentprocessor *Paymentprocessor) UnpackRefundedEvent(log *types.Log) (*PaymentprocessorRefunded, error) {
	event := "Refunded"
	if log.Topics[0] != paymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorRefunded)
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
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["BuyerCannotBeSeller"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackBuyerCannotBeSellerError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidDisputeResolution"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidDisputeResolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvalidInvoiceStateError(raw[4:])
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
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["InvoiceDoesNotExist"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackInvoiceDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["MetaInvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackMetaInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNewOwnerIsZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NoHandoverRequest"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNoHandoverRequestError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["PriceCannotBeZero"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackPriceCannotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessor.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return paymentprocessor.UnpackUnauthorizedError(raw[4:])
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

// PaymentprocessorBuyerCannotBeSeller represents a BuyerCannotBeSeller error raised by the Paymentprocessor contract.
type PaymentprocessorBuyerCannotBeSeller struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BuyerCannotBeSeller()
func PaymentprocessorBuyerCannotBeSellerErrorID() common.Hash {
	return common.HexToHash("0xb12e242105ea4b2bcdc745efefe14be5558f5f16020ec252980cefc86c6a7a77")
}

// UnpackBuyerCannotBeSellerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BuyerCannotBeSeller()
func (paymentprocessor *Paymentprocessor) UnpackBuyerCannotBeSellerError(raw []byte) (*PaymentprocessorBuyerCannotBeSeller, error) {
	out := new(PaymentprocessorBuyerCannotBeSeller)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "BuyerCannotBeSeller", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorInsufficientBalance represents a InsufficientBalance error raised by the Paymentprocessor contract.
type PaymentprocessorInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func PaymentprocessorInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (paymentprocessor *Paymentprocessor) UnpackInsufficientBalanceError(raw []byte) (*PaymentprocessorInsufficientBalance, error) {
	out := new(PaymentprocessorInsufficientBalance)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
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

// PaymentprocessorInvoiceAlreadyExists represents a InvoiceAlreadyExists error raised by the Paymentprocessor contract.
type PaymentprocessorInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceAlreadyExists()
func PaymentprocessorInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x074bc9355c94925fa82ddb49dcc88f1a666f1d1aa24efbddcdbe5f8d98b7ed59")
}

// UnpackInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceAlreadyExists()
func (paymentprocessor *Paymentprocessor) UnpackInvoiceAlreadyExistsError(raw []byte) (*PaymentprocessorInvoiceAlreadyExists, error) {
	out := new(PaymentprocessorInvoiceAlreadyExists)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "InvoiceAlreadyExists", raw); err != nil {
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

// PaymentprocessorMetaInvoiceAlreadyExists represents a MetaInvoiceAlreadyExists error raised by the Paymentprocessor contract.
type PaymentprocessorMetaInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MetaInvoiceAlreadyExists()
func PaymentprocessorMetaInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xb09960c1d49af4579c96dbfb857f76f135da1b549d276e06491c05ec24747201")
}

// UnpackMetaInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MetaInvoiceAlreadyExists()
func (paymentprocessor *Paymentprocessor) UnpackMetaInvoiceAlreadyExistsError(raw []byte) (*PaymentprocessorMetaInvoiceAlreadyExists, error) {
	out := new(PaymentprocessorMetaInvoiceAlreadyExists)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "MetaInvoiceAlreadyExists", raw); err != nil {
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

// PaymentprocessorPriceCannotBeZero represents a PriceCannotBeZero error raised by the Paymentprocessor contract.
type PaymentprocessorPriceCannotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PriceCannotBeZero()
func PaymentprocessorPriceCannotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x2c669f0ac3409adbbadbe16eaad2cac428e45b3cb8de2f47377f30f8b5729e18")
}

// UnpackPriceCannotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PriceCannotBeZero()
func (paymentprocessor *Paymentprocessor) UnpackPriceCannotBeZeroError(raw []byte) (*PaymentprocessorPriceCannotBeZero, error) {
	out := new(PaymentprocessorPriceCannotBeZero)
	if err := paymentprocessor.abi.UnpackIntoInterface(out, "PriceCannotBeZero", raw); err != nil {
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
