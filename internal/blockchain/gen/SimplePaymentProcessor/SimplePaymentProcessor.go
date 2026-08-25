// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleprocessor

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

// ISimplePaymentProcessorInvoice is an auto generated low-level Go binding around an user-defined struct.
type ISimplePaymentProcessorInvoice struct {
	InvoiceNonce         *big.Int
	CreatedAt            *big.Int
	PaidAt               *big.Int
	ReleaseAt            *big.Int
	ExpiresAt            *big.Int
	SellerActionDeadline *big.Int
	EscrowHoldPeriod     uint32
	State                uint8
	WithdrawalRetries    uint8
	FeeRate              uint16
	Seller               common.Address
	Buyer                common.Address
	Escrow               common.Address
	FeeReceiver          common.Address
	Price                *big.Int
	Balance              *big.Int
}

// SimpleprocessorMetaData contains all meta data concerning the Simpleprocessor contract.
var SimpleprocessorMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumInvoicePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_notesAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_wethAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptPayment\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateFee\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"feeValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createInvoice\",\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_holdPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_storageRef\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_share\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAutomation\",\"inputs\":[],\"outputs\":[{\"name\":\"automationAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecisionWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"decisionWindowValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInvoiceData\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"i\",\"type\":\"tuple\",\"internalType\":\"structISimplePaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"sellerActionDeadline\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getItems\",\"inputs\":[],\"outputs\":[{\"name\":\"items\",\"type\":\"uint216[]\",\"internalType\":\"uint216[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumInvoiceValue\",\"inputs\":[],\"outputs\":[{\"name\":\"minimumValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextInvoiceNonceValue\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasDueTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"dueTasksExist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pay\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_storageRef\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_share\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"escrowAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processDueTasks\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"refundBuyer\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectPayment\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAutomation\",\"inputs\":[{\"name\":\"_automationAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDecisionWindow\",\"inputs\":[{\"name\":\"_newDecisionWindow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumInvoiceValue\",\"inputs\":[{\"name\":\"_newMinimumInvoiceValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AutomationUpdated\",\"inputs\":[{\"name\":\"automation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceAccepted\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCanceled\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"invoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISimplePaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"sellerActionDeadline\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoicePaid\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sellerActionDeadline\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceRefunded\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceRejected\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceReleased\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentBurned\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRetried\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"attempt\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AcceptanceWindowExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateTask\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowWithdrawFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HoldPeriodHasNotBeenExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectPaymentAmount\",\"inputs\":[{\"name\":\"_sent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDecisionWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeAuthorization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHeapPosition\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInvoiceState\",\"inputs\":[{\"name\":\"_invoiceState\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvoiceAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceIsNoLongerValid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceNotEligibleForRefund\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SellerCannotPayOwnedInvoice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaskNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedNativeTransfer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValueIsTooLow\",\"inputs\":[]}]",
	ID:  "Simpleprocessor",
}

// Simpleprocessor is an auto generated Go binding around an Ethereum contract.
type Simpleprocessor struct {
	abi abi.ABI
}

// NewSimpleprocessor creates a new instance of Simpleprocessor.
func NewSimpleprocessor() *Simpleprocessor {
	parsed, err := SimpleprocessorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Simpleprocessor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Simpleprocessor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _paymentProcessorStorageAddress, uint256 _minimumInvoicePrice, address _notesAddress, address _wethAddress) returns()
func (simpleprocessor *Simpleprocessor) PackConstructor(_paymentProcessorStorageAddress common.Address, _minimumInvoicePrice *big.Int, _notesAddress common.Address, _wethAddress common.Address) []byte {
	enc, err := simpleprocessor.abi.Pack("", _paymentProcessorStorageAddress, _minimumInvoicePrice, _notesAddress, _wethAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAcceptPayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ecb321d.
//
// Solidity: function acceptPayment(uint216 _invoiceId, address _feeReceiver, bytes _data) returns()
func (simpleprocessor *Simpleprocessor) PackAcceptPayment(invoiceId *big.Int, feeReceiver common.Address, data []byte) []byte {
	enc, err := simpleprocessor.abi.Pack("acceptPayment", invoiceId, feeReceiver, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCalculateFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x99a5d747.
//
// Solidity: function calculateFee(uint256 _amount) view returns(uint256 feeValue)
func (simpleprocessor *Simpleprocessor) PackCalculateFee(amount *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("calculateFee", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCalculateFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x99a5d747.
//
// Solidity: function calculateFee(uint256 _amount) view returns(uint256 feeValue)
func (simpleprocessor *Simpleprocessor) UnpackCalculateFee(data []byte) (*big.Int, error) {
	out, err := simpleprocessor.abi.Unpack("calculateFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCancelInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde48b793.
//
// Solidity: function cancelInvoice(uint216 _invoiceId) returns()
func (simpleprocessor *Simpleprocessor) PackCancelInvoice(invoiceId *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("cancelInvoice", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x580a586c.
//
// Solidity: function createInvoice(uint256 _price, uint32 _holdPeriod, bytes _storageRef, bool _share) returns(uint216 invoiceId)
func (simpleprocessor *Simpleprocessor) PackCreateInvoice(price *big.Int, holdPeriod uint32, storageRef []byte, share bool) []byte {
	enc, err := simpleprocessor.abi.Pack("createInvoice", price, holdPeriod, storageRef, share)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x580a586c.
//
// Solidity: function createInvoice(uint256 _price, uint32 _holdPeriod, bytes _storageRef, bool _share) returns(uint216 invoiceId)
func (simpleprocessor *Simpleprocessor) UnpackCreateInvoice(data []byte) (*big.Int, error) {
	out, err := simpleprocessor.abi.Unpack("createInvoice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetAutomation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3a7ad1e.
//
// Solidity: function getAutomation() view returns(address automationAddress)
func (simpleprocessor *Simpleprocessor) PackGetAutomation() []byte {
	enc, err := simpleprocessor.abi.Pack("getAutomation")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAutomation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3a7ad1e.
//
// Solidity: function getAutomation() view returns(address automationAddress)
func (simpleprocessor *Simpleprocessor) UnpackGetAutomation(data []byte) (common.Address, error) {
	out, err := simpleprocessor.abi.Unpack("getAutomation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetDecisionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d86abad.
//
// Solidity: function getDecisionWindow() view returns(uint256 decisionWindowValue)
func (simpleprocessor *Simpleprocessor) PackGetDecisionWindow() []byte {
	enc, err := simpleprocessor.abi.Pack("getDecisionWindow")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDecisionWindow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d86abad.
//
// Solidity: function getDecisionWindow() view returns(uint256 decisionWindowValue)
func (simpleprocessor *Simpleprocessor) UnpackGetDecisionWindow(data []byte) (*big.Int, error) {
	out, err := simpleprocessor.abi.Unpack("getDecisionWindow", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetInvoiceData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ec7bbf9.
//
// Solidity: function getInvoiceData(uint216 _invoiceId) view returns((uint216,uint40,uint40,uint40,uint40,uint40,uint32,uint8,uint8,uint16,address,address,address,address,uint256,uint256) i)
func (simpleprocessor *Simpleprocessor) PackGetInvoiceData(invoiceId *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("getInvoiceData", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoiceData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ec7bbf9.
//
// Solidity: function getInvoiceData(uint216 _invoiceId) view returns((uint216,uint40,uint40,uint40,uint40,uint40,uint32,uint8,uint8,uint16,address,address,address,address,uint256,uint256) i)
func (simpleprocessor *Simpleprocessor) UnpackGetInvoiceData(data []byte) (ISimplePaymentProcessorInvoice, error) {
	out, err := simpleprocessor.abi.Unpack("getInvoiceData", data)
	if err != nil {
		return *new(ISimplePaymentProcessorInvoice), err
	}
	out0 := *abi.ConvertType(out[0], new(ISimplePaymentProcessorInvoice)).(*ISimplePaymentProcessorInvoice)
	return out0, err
}

// PackGetItems is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410d59cc.
//
// Solidity: function getItems() view returns(uint216[] items)
func (simpleprocessor *Simpleprocessor) PackGetItems() []byte {
	enc, err := simpleprocessor.abi.Pack("getItems")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetItems is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410d59cc.
//
// Solidity: function getItems() view returns(uint216[] items)
func (simpleprocessor *Simpleprocessor) UnpackGetItems(data []byte) ([]*big.Int, error) {
	out, err := simpleprocessor.abi.Unpack("getItems", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackGetMinimumInvoiceValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1595a882.
//
// Solidity: function getMinimumInvoiceValue() view returns(uint256 minimumValue)
func (simpleprocessor *Simpleprocessor) PackGetMinimumInvoiceValue() []byte {
	enc, err := simpleprocessor.abi.Pack("getMinimumInvoiceValue")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMinimumInvoiceValue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1595a882.
//
// Solidity: function getMinimumInvoiceValue() view returns(uint256 minimumValue)
func (simpleprocessor *Simpleprocessor) UnpackGetMinimumInvoiceValue(data []byte) (*big.Int, error) {
	out, err := simpleprocessor.abi.Unpack("getMinimumInvoiceValue", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetNextInvoiceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonceValue)
func (simpleprocessor *Simpleprocessor) PackGetNextInvoiceNonce() []byte {
	enc, err := simpleprocessor.abi.Pack("getNextInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonceValue)
func (simpleprocessor *Simpleprocessor) UnpackGetNextInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := simpleprocessor.abi.Unpack("getNextInvoiceNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackHasDueTasks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a9a0dea.
//
// Solidity: function hasDueTasks() view returns(bool dueTasksExist)
func (simpleprocessor *Simpleprocessor) PackHasDueTasks() []byte {
	enc, err := simpleprocessor.abi.Pack("hasDueTasks")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasDueTasks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0a9a0dea.
//
// Solidity: function hasDueTasks() view returns(bool dueTasksExist)
func (simpleprocessor *Simpleprocessor) UnpackHasDueTasks(data []byte) (bool, error) {
	out, err := simpleprocessor.abi.Unpack("hasDueTasks", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackPay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9224df4a.
//
// Solidity: function pay(uint216 _invoiceId, bytes _storageRef, bool _share) payable returns(address escrowAddress)
func (simpleprocessor *Simpleprocessor) PackPay(invoiceId *big.Int, storageRef []byte, share bool) []byte {
	enc, err := simpleprocessor.abi.Pack("pay", invoiceId, storageRef, share)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9224df4a.
//
// Solidity: function pay(uint216 _invoiceId, bytes _storageRef, bool _share) payable returns(address escrowAddress)
func (simpleprocessor *Simpleprocessor) UnpackPay(data []byte) (common.Address, error) {
	out, err := simpleprocessor.abi.Unpack("pay", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (simpleprocessor *Simpleprocessor) PackPpStorage() []byte {
	enc, err := simpleprocessor.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (simpleprocessor *Simpleprocessor) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := simpleprocessor.abi.Unpack("ppStorage", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProcessDueTasks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb138aa8a.
//
// Solidity: function processDueTasks() returns()
func (simpleprocessor *Simpleprocessor) PackProcessDueTasks() []byte {
	enc, err := simpleprocessor.abi.Pack("processDueTasks")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRefundBuyer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x402d5a9b.
//
// Solidity: function refundBuyer(uint216 _invoiceId) returns()
func (simpleprocessor *Simpleprocessor) PackRefundBuyer(invoiceId *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("refundBuyer", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRejectPayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98201712.
//
// Solidity: function rejectPayment(uint216 _invoiceId) returns()
func (simpleprocessor *Simpleprocessor) PackRejectPayment(invoiceId *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("rejectPayment", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb990809.
//
// Solidity: function release(uint216 _invoiceId) returns()
func (simpleprocessor *Simpleprocessor) PackRelease(invoiceId *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("release", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetAutomation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x730dcc8e.
//
// Solidity: function setAutomation(address _automationAddress) returns()
func (simpleprocessor *Simpleprocessor) PackSetAutomation(automationAddress common.Address) []byte {
	enc, err := simpleprocessor.abi.Pack("setAutomation", automationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetDecisionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03dd5199.
//
// Solidity: function setDecisionWindow(uint256 _newDecisionWindow) returns()
func (simpleprocessor *Simpleprocessor) PackSetDecisionWindow(newDecisionWindow *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("setDecisionWindow", newDecisionWindow)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinimumInvoiceValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf82afa07.
//
// Solidity: function setMinimumInvoiceValue(uint256 _newMinimumInvoiceValue) returns()
func (simpleprocessor *Simpleprocessor) PackSetMinimumInvoiceValue(newMinimumInvoiceValue *big.Int) []byte {
	enc, err := simpleprocessor.abi.Pack("setMinimumInvoiceValue", newMinimumInvoiceValue)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWeth is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (simpleprocessor *Simpleprocessor) PackWeth() []byte {
	enc, err := simpleprocessor.abi.Pack("weth")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackWeth is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (simpleprocessor *Simpleprocessor) UnpackWeth(data []byte) (common.Address, error) {
	out, err := simpleprocessor.abi.Unpack("weth", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// SimpleprocessorAutomationUpdated represents a AutomationUpdated event raised by the Simpleprocessor contract.
type SimpleprocessorAutomationUpdated struct {
	Automation common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorAutomationUpdatedEventName = "AutomationUpdated"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorAutomationUpdated) ContractEventName() string {
	return SimpleprocessorAutomationUpdatedEventName
}

// UnpackAutomationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AutomationUpdated(address indexed automation)
func (simpleprocessor *Simpleprocessor) UnpackAutomationUpdatedEvent(log *types.Log) (*SimpleprocessorAutomationUpdated, error) {
	event := "AutomationUpdated"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorAutomationUpdated)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoiceAccepted represents a InvoiceAccepted event raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceAccepted struct {
	InvoiceId   *big.Int
	FeeReceiver common.Address
	ReleaseAt   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoiceAcceptedEventName = "InvoiceAccepted"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoiceAccepted) ContractEventName() string {
	return SimpleprocessorInvoiceAcceptedEventName
}

// UnpackInvoiceAcceptedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceAccepted(uint216 indexed invoiceId, address indexed feeReceiver, uint40 releaseAt)
func (simpleprocessor *Simpleprocessor) UnpackInvoiceAcceptedEvent(log *types.Log) (*SimpleprocessorInvoiceAccepted, error) {
	event := "InvoiceAccepted"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoiceAccepted)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceCanceled struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoiceCanceled) ContractEventName() string {
	return SimpleprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(uint216 indexed invoiceId)
func (simpleprocessor *Simpleprocessor) UnpackInvoiceCanceledEvent(log *types.Log) (*SimpleprocessorInvoiceCanceled, error) {
	event := "InvoiceCanceled"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoiceCanceled)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoiceCreated represents a InvoiceCreated event raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceCreated struct {
	InvoiceId *big.Int
	Invoice   ISimplePaymentProcessorInvoice
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoiceCreated) ContractEventName() string {
	return SimpleprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(uint216 indexed invoiceId, (uint216,uint40,uint40,uint40,uint40,uint40,uint32,uint8,uint8,uint16,address,address,address,address,uint256,uint256) invoice)
func (simpleprocessor *Simpleprocessor) UnpackInvoiceCreatedEvent(log *types.Log) (*SimpleprocessorInvoiceCreated, error) {
	event := "InvoiceCreated"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoiceCreated)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoicePaid represents a InvoicePaid event raised by the Simpleprocessor contract.
type SimpleprocessorInvoicePaid struct {
	InvoiceId            *big.Int
	Buyer                common.Address
	AmountPaid           *big.Int
	SellerActionDeadline *big.Int
	Raw                  *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoicePaidEventName = "InvoicePaid"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoicePaid) ContractEventName() string {
	return SimpleprocessorInvoicePaidEventName
}

// UnpackInvoicePaidEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoicePaid(uint216 indexed invoiceId, address indexed buyer, uint256 indexed amountPaid, uint40 sellerActionDeadline)
func (simpleprocessor *Simpleprocessor) UnpackInvoicePaidEvent(log *types.Log) (*SimpleprocessorInvoicePaid, error) {
	event := "InvoicePaid"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoicePaid)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoiceRefunded represents a InvoiceRefunded event raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceRefunded struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoiceRefundedEventName = "InvoiceRefunded"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoiceRefunded) ContractEventName() string {
	return SimpleprocessorInvoiceRefundedEventName
}

// UnpackInvoiceRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceRefunded(uint216 indexed invoiceId, uint256 amount)
func (simpleprocessor *Simpleprocessor) UnpackInvoiceRefundedEvent(log *types.Log) (*SimpleprocessorInvoiceRefunded, error) {
	event := "InvoiceRefunded"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoiceRefunded)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoiceRejected represents a InvoiceRejected event raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceRejected struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoiceRejectedEventName = "InvoiceRejected"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoiceRejected) ContractEventName() string {
	return SimpleprocessorInvoiceRejectedEventName
}

// UnpackInvoiceRejectedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceRejected(uint216 indexed invoiceId, uint256 amount)
func (simpleprocessor *Simpleprocessor) UnpackInvoiceRejectedEvent(log *types.Log) (*SimpleprocessorInvoiceRejected, error) {
	event := "InvoiceRejected"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoiceRejected)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorInvoiceReleased represents a InvoiceReleased event raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceReleased struct {
	InvoiceId    *big.Int
	SellerAmount *big.Int
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorInvoiceReleasedEventName = "InvoiceReleased"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorInvoiceReleased) ContractEventName() string {
	return SimpleprocessorInvoiceReleasedEventName
}

// UnpackInvoiceReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceReleased(uint216 indexed invoiceId, uint256 sellerAmount, uint256 fee)
func (simpleprocessor *Simpleprocessor) UnpackInvoiceReleasedEvent(log *types.Log) (*SimpleprocessorInvoiceReleased, error) {
	event := "InvoiceReleased"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorInvoiceReleased)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorPaymentBurned represents a PaymentBurned event raised by the Simpleprocessor contract.
type SimpleprocessorPaymentBurned struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorPaymentBurnedEventName = "PaymentBurned"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorPaymentBurned) ContractEventName() string {
	return SimpleprocessorPaymentBurnedEventName
}

// UnpackPaymentBurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentBurned(uint216 indexed invoiceId, uint256 amount)
func (simpleprocessor *Simpleprocessor) UnpackPaymentBurnedEvent(log *types.Log) (*SimpleprocessorPaymentBurned, error) {
	event := "PaymentBurned"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorPaymentBurned)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorTransferFailed represents a TransferFailed event raised by the Simpleprocessor contract.
type SimpleprocessorTransferFailed struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorTransferFailedEventName = "TransferFailed"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorTransferFailed) ContractEventName() string {
	return SimpleprocessorTransferFailedEventName
}

// UnpackTransferFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferFailed(uint216 indexed invoiceId, address indexed recipient, uint256 amount)
func (simpleprocessor *Simpleprocessor) UnpackTransferFailedEvent(log *types.Log) (*SimpleprocessorTransferFailed, error) {
	event := "TransferFailed"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorTransferFailed)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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

// SimpleprocessorWithdrawalRetried represents a WithdrawalRetried event raised by the Simpleprocessor contract.
type SimpleprocessorWithdrawalRetried struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Attempt   uint8
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimpleprocessorWithdrawalRetriedEventName = "WithdrawalRetried"

// ContractEventName returns the user-defined event name.
func (SimpleprocessorWithdrawalRetried) ContractEventName() string {
	return SimpleprocessorWithdrawalRetriedEventName
}

// UnpackWithdrawalRetriedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalRetried(uint216 indexed invoiceId, address indexed recipient, uint256 amount, uint8 attempt)
func (simpleprocessor *Simpleprocessor) UnpackWithdrawalRetriedEvent(log *types.Log) (*SimpleprocessorWithdrawalRetried, error) {
	event := "WithdrawalRetried"
	if log.Topics[0] != simpleprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimpleprocessorWithdrawalRetried)
	if len(log.Data) > 0 {
		if err := simpleprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simpleprocessor.abi.Events[event].Inputs {
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
func (simpleprocessor *Simpleprocessor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["AcceptanceWindowExceeded"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackAcceptanceWindowExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["ContractPaused"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["DuplicateTask"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackDuplicateTaskError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["EscrowWithdrawFailed"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackEscrowWithdrawFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["HoldPeriodHasNotBeenExceeded"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackHoldPeriodHasNotBeenExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["IncorrectPaymentAmount"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackIncorrectPaymentAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvalidDecisionWindow"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvalidDecisionWindowError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvalidFeeAuthorization"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvalidFeeAuthorizationError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvalidFeeReceiver"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvalidFeeReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvalidHeapPosition"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvalidHeapPositionError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvalidInvoiceStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvoiceIsNoLongerValid"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvoiceIsNoLongerValidError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["InvoiceNotEligibleForRefund"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackInvoiceNotEligibleForRefundError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["Reentrancy"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["SellerCannotPayOwnedInvoice"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackSellerCannotPayOwnedInvoiceError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["TaskNotFound"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackTaskNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["UnexpectedNativeTransfer"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackUnexpectedNativeTransferError(raw[4:])
	}
	if bytes.Equal(raw[:4], simpleprocessor.abi.Errors["ValueIsTooLow"].ID.Bytes()[:4]) {
		return simpleprocessor.UnpackValueIsTooLowError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SimpleprocessorAcceptanceWindowExceeded represents a AcceptanceWindowExceeded error raised by the Simpleprocessor contract.
type SimpleprocessorAcceptanceWindowExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AcceptanceWindowExceeded()
func SimpleprocessorAcceptanceWindowExceededErrorID() common.Hash {
	return common.HexToHash("0x2b8af0bb26858aa842f6aac91213da2376cd6baa62d4c906c6aa0685f42c1a48")
}

// UnpackAcceptanceWindowExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AcceptanceWindowExceeded()
func (simpleprocessor *Simpleprocessor) UnpackAcceptanceWindowExceededError(raw []byte) (*SimpleprocessorAcceptanceWindowExceeded, error) {
	out := new(SimpleprocessorAcceptanceWindowExceeded)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "AcceptanceWindowExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorContractPaused represents a ContractPaused error raised by the Simpleprocessor contract.
type SimpleprocessorContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ContractPaused()
func SimpleprocessorContractPausedErrorID() common.Hash {
	return common.HexToHash("0xab35696f06e428ebc5ceba8cd17f8fed287baf43440206d1943af1ee53e6d267")
}

// UnpackContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ContractPaused()
func (simpleprocessor *Simpleprocessor) UnpackContractPausedError(raw []byte) (*SimpleprocessorContractPaused, error) {
	out := new(SimpleprocessorContractPaused)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "ContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorDuplicateTask represents a DuplicateTask error raised by the Simpleprocessor contract.
type SimpleprocessorDuplicateTask struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateTask()
func SimpleprocessorDuplicateTaskErrorID() common.Hash {
	return common.HexToHash("0x6b22feb9606cb284058f0a2f05a53980401948118942ff234189d20361fe4e93")
}

// UnpackDuplicateTaskError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateTask()
func (simpleprocessor *Simpleprocessor) UnpackDuplicateTaskError(raw []byte) (*SimpleprocessorDuplicateTask, error) {
	out := new(SimpleprocessorDuplicateTask)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "DuplicateTask", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorEscrowWithdrawFailed represents a EscrowWithdrawFailed error raised by the Simpleprocessor contract.
type SimpleprocessorEscrowWithdrawFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowWithdrawFailed()
func SimpleprocessorEscrowWithdrawFailedErrorID() common.Hash {
	return common.HexToHash("0x667ecf9d53e4600a9a128606592ec5e22e0269990439145a2bbc8a983c7af5ac")
}

// UnpackEscrowWithdrawFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowWithdrawFailed()
func (simpleprocessor *Simpleprocessor) UnpackEscrowWithdrawFailedError(raw []byte) (*SimpleprocessorEscrowWithdrawFailed, error) {
	out := new(SimpleprocessorEscrowWithdrawFailed)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "EscrowWithdrawFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorHoldPeriodHasNotBeenExceeded represents a HoldPeriodHasNotBeenExceeded error raised by the Simpleprocessor contract.
type SimpleprocessorHoldPeriodHasNotBeenExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HoldPeriodHasNotBeenExceeded()
func SimpleprocessorHoldPeriodHasNotBeenExceededErrorID() common.Hash {
	return common.HexToHash("0xad2652ac51ebd8cf548af7b7d84853dd5d5efef6e7c1b57e0c81a88fd70d009d")
}

// UnpackHoldPeriodHasNotBeenExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HoldPeriodHasNotBeenExceeded()
func (simpleprocessor *Simpleprocessor) UnpackHoldPeriodHasNotBeenExceededError(raw []byte) (*SimpleprocessorHoldPeriodHasNotBeenExceeded, error) {
	out := new(SimpleprocessorHoldPeriodHasNotBeenExceeded)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "HoldPeriodHasNotBeenExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorIncorrectPaymentAmount represents a IncorrectPaymentAmount error raised by the Simpleprocessor contract.
type SimpleprocessorIncorrectPaymentAmount struct {
	Sent     *big.Int
	Expected *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error IncorrectPaymentAmount(uint256 _sent, uint256 _expected)
func SimpleprocessorIncorrectPaymentAmountErrorID() common.Hash {
	return common.HexToHash("0x47af6acc44a34c9e741a083d95d979d69adf7234f5e86b5e76839bf27fffd8cb")
}

// UnpackIncorrectPaymentAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error IncorrectPaymentAmount(uint256 _sent, uint256 _expected)
func (simpleprocessor *Simpleprocessor) UnpackIncorrectPaymentAmountError(raw []byte) (*SimpleprocessorIncorrectPaymentAmount, error) {
	out := new(SimpleprocessorIncorrectPaymentAmount)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "IncorrectPaymentAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvalidDecisionWindow represents a InvalidDecisionWindow error raised by the Simpleprocessor contract.
type SimpleprocessorInvalidDecisionWindow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDecisionWindow()
func SimpleprocessorInvalidDecisionWindowErrorID() common.Hash {
	return common.HexToHash("0x39141cc3acd0676e85f3b9927f0f60c5ff570c62f9d8e325c2849504a3f94379")
}

// UnpackInvalidDecisionWindowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDecisionWindow()
func (simpleprocessor *Simpleprocessor) UnpackInvalidDecisionWindowError(raw []byte) (*SimpleprocessorInvalidDecisionWindow, error) {
	out := new(SimpleprocessorInvalidDecisionWindow)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvalidDecisionWindow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvalidFeeAuthorization represents a InvalidFeeAuthorization error raised by the Simpleprocessor contract.
type SimpleprocessorInvalidFeeAuthorization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeAuthorization()
func SimpleprocessorInvalidFeeAuthorizationErrorID() common.Hash {
	return common.HexToHash("0x1735eabec15c7395efafdfa0dda5c74faf3517b604ff55660d6ef0e7457f2c1d")
}

// UnpackInvalidFeeAuthorizationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeAuthorization()
func (simpleprocessor *Simpleprocessor) UnpackInvalidFeeAuthorizationError(raw []byte) (*SimpleprocessorInvalidFeeAuthorization, error) {
	out := new(SimpleprocessorInvalidFeeAuthorization)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvalidFeeAuthorization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvalidFeeReceiver represents a InvalidFeeReceiver error raised by the Simpleprocessor contract.
type SimpleprocessorInvalidFeeReceiver struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeReceiver()
func SimpleprocessorInvalidFeeReceiverErrorID() common.Hash {
	return common.HexToHash("0xd200485c51caaf66763f8b49c9cfa281a0a10132cb56c8fffc35867701d3fc5f")
}

// UnpackInvalidFeeReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeReceiver()
func (simpleprocessor *Simpleprocessor) UnpackInvalidFeeReceiverError(raw []byte) (*SimpleprocessorInvalidFeeReceiver, error) {
	out := new(SimpleprocessorInvalidFeeReceiver)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvalidFeeReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvalidHeapPosition represents a InvalidHeapPosition error raised by the Simpleprocessor contract.
type SimpleprocessorInvalidHeapPosition struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidHeapPosition()
func SimpleprocessorInvalidHeapPositionErrorID() common.Hash {
	return common.HexToHash("0x76f4a2832a11e5a5adb544de785993d5d08fe0ec9c4270e1b0cbc33a9a1f27e8")
}

// UnpackInvalidHeapPositionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidHeapPosition()
func (simpleprocessor *Simpleprocessor) UnpackInvalidHeapPositionError(raw []byte) (*SimpleprocessorInvalidHeapPosition, error) {
	out := new(SimpleprocessorInvalidHeapPosition)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvalidHeapPosition", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvalidInvoiceState represents a InvalidInvoiceState error raised by the Simpleprocessor contract.
type SimpleprocessorInvalidInvoiceState struct {
	InvoiceState *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInvoiceState(uint256 _invoiceState)
func SimpleprocessorInvalidInvoiceStateErrorID() common.Hash {
	return common.HexToHash("0x1d5b155656afea98f1415b8e817932b761e3215588ffe226a1d86ff9ac55fe02")
}

// UnpackInvalidInvoiceStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInvoiceState(uint256 _invoiceState)
func (simpleprocessor *Simpleprocessor) UnpackInvalidInvoiceStateError(raw []byte) (*SimpleprocessorInvalidInvoiceState, error) {
	out := new(SimpleprocessorInvalidInvoiceState)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvalidInvoiceState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvoiceAlreadyExists represents a InvoiceAlreadyExists error raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceAlreadyExists()
func SimpleprocessorInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x074bc9355c94925fa82ddb49dcc88f1a666f1d1aa24efbddcdbe5f8d98b7ed59")
}

// UnpackInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceAlreadyExists()
func (simpleprocessor *Simpleprocessor) UnpackInvoiceAlreadyExistsError(raw []byte) (*SimpleprocessorInvoiceAlreadyExists, error) {
	out := new(SimpleprocessorInvoiceAlreadyExists)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvoiceIsNoLongerValid represents a InvoiceIsNoLongerValid error raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceIsNoLongerValid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceIsNoLongerValid()
func SimpleprocessorInvoiceIsNoLongerValidErrorID() common.Hash {
	return common.HexToHash("0xff42dbfc3f50dbd393d8edfcd163d614372311f4609adb53d576d6c3d588f37d")
}

// UnpackInvoiceIsNoLongerValidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceIsNoLongerValid()
func (simpleprocessor *Simpleprocessor) UnpackInvoiceIsNoLongerValidError(raw []byte) (*SimpleprocessorInvoiceIsNoLongerValid, error) {
	out := new(SimpleprocessorInvoiceIsNoLongerValid)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvoiceIsNoLongerValid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorInvoiceNotEligibleForRefund represents a InvoiceNotEligibleForRefund error raised by the Simpleprocessor contract.
type SimpleprocessorInvoiceNotEligibleForRefund struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceNotEligibleForRefund()
func SimpleprocessorInvoiceNotEligibleForRefundErrorID() common.Hash {
	return common.HexToHash("0xbb126ff1385e3d8206af61fac954f0af8fa8b14efaf24fb5d0f03a32a7f093be")
}

// UnpackInvoiceNotEligibleForRefundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceNotEligibleForRefund()
func (simpleprocessor *Simpleprocessor) UnpackInvoiceNotEligibleForRefundError(raw []byte) (*SimpleprocessorInvoiceNotEligibleForRefund, error) {
	out := new(SimpleprocessorInvoiceNotEligibleForRefund)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "InvoiceNotEligibleForRefund", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorNotAuthorized represents a NotAuthorized error raised by the Simpleprocessor contract.
type SimpleprocessorNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func SimpleprocessorNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (simpleprocessor *Simpleprocessor) UnpackNotAuthorizedError(raw []byte) (*SimpleprocessorNotAuthorized, error) {
	out := new(SimpleprocessorNotAuthorized)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorReentrancy represents a Reentrancy error raised by the Simpleprocessor contract.
type SimpleprocessorReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Reentrancy()
func SimpleprocessorReentrancyErrorID() common.Hash {
	return common.HexToHash("0xab143c06c9772d69bbbc9f2fe74acd02f810e93b099f3d1dac8448ac9ae35991")
}

// UnpackReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Reentrancy()
func (simpleprocessor *Simpleprocessor) UnpackReentrancyError(raw []byte) (*SimpleprocessorReentrancy, error) {
	out := new(SimpleprocessorReentrancy)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "Reentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorSellerCannotPayOwnedInvoice represents a SellerCannotPayOwnedInvoice error raised by the Simpleprocessor contract.
type SimpleprocessorSellerCannotPayOwnedInvoice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SellerCannotPayOwnedInvoice()
func SimpleprocessorSellerCannotPayOwnedInvoiceErrorID() common.Hash {
	return common.HexToHash("0x020175b17895ca2ebc70bd6da55e4d50ebe6fbc9b1110d0ba1ccb0613bf49691")
}

// UnpackSellerCannotPayOwnedInvoiceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SellerCannotPayOwnedInvoice()
func (simpleprocessor *Simpleprocessor) UnpackSellerCannotPayOwnedInvoiceError(raw []byte) (*SimpleprocessorSellerCannotPayOwnedInvoice, error) {
	out := new(SimpleprocessorSellerCannotPayOwnedInvoice)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "SellerCannotPayOwnedInvoice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorTaskNotFound represents a TaskNotFound error raised by the Simpleprocessor contract.
type SimpleprocessorTaskNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TaskNotFound()
func SimpleprocessorTaskNotFoundErrorID() common.Hash {
	return common.HexToHash("0xc325ae33d18e47931adbda2584c56fef1d3e5e64beab80da59968e1c83c84937")
}

// UnpackTaskNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TaskNotFound()
func (simpleprocessor *Simpleprocessor) UnpackTaskNotFoundError(raw []byte) (*SimpleprocessorTaskNotFound, error) {
	out := new(SimpleprocessorTaskNotFound)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "TaskNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorUnexpectedNativeTransfer represents a UnexpectedNativeTransfer error raised by the Simpleprocessor contract.
type SimpleprocessorUnexpectedNativeTransfer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpectedNativeTransfer()
func SimpleprocessorUnexpectedNativeTransferErrorID() common.Hash {
	return common.HexToHash("0xecb8b30d3aabfab45e153cab7e267d8f707ed670125178fc459f4849e0504fa3")
}

// UnpackUnexpectedNativeTransferError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpectedNativeTransfer()
func (simpleprocessor *Simpleprocessor) UnpackUnexpectedNativeTransferError(raw []byte) (*SimpleprocessorUnexpectedNativeTransfer, error) {
	out := new(SimpleprocessorUnexpectedNativeTransfer)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "UnexpectedNativeTransfer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprocessorValueIsTooLow represents a ValueIsTooLow error raised by the Simpleprocessor contract.
type SimpleprocessorValueIsTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ValueIsTooLow()
func SimpleprocessorValueIsTooLowErrorID() common.Hash {
	return common.HexToHash("0x5033f274524623c8f2e5518c3a1c9e6345ae1da946950ff4e2153fe3f55d28ec")
}

// UnpackValueIsTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ValueIsTooLow()
func (simpleprocessor *Simpleprocessor) UnpackValueIsTooLowError(raw []byte) (*SimpleprocessorValueIsTooLow, error) {
	out := new(SimpleprocessorValueIsTooLow)
	if err := simpleprocessor.abi.UnpackIntoInterface(out, "ValueIsTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}
