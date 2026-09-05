// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplepaymentprocessor

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

// SimplepaymentprocessorMetaData contains all meta data concerning the Simplepaymentprocessor contract.
var SimplepaymentprocessorMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumInvoicePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_notesAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_wethAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptPayment\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateFee\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"feeValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelInvoice\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createInvoice\",\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_holdPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_storageRef\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_share\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAutomation\",\"inputs\":[],\"outputs\":[{\"name\":\"automationAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecisionWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"decisionWindowValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInvoiceData\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"i\",\"type\":\"tuple\",\"internalType\":\"structISimplePaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"sellerActionDeadline\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getItems\",\"inputs\":[],\"outputs\":[{\"name\":\"items\",\"type\":\"uint216[]\",\"internalType\":\"uint216[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinimumInvoiceValue\",\"inputs\":[],\"outputs\":[{\"name\":\"minimumValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextInvoiceNonceValue\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasDueTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"dueTasksExist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pay\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"_storageRef\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_share\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"escrowAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processDueTasks\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"refundBuyer\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectPayment\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"release\",\"inputs\":[{\"name\":\"_invoiceId\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAutomation\",\"inputs\":[{\"name\":\"_automationAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDecisionWindow\",\"inputs\":[{\"name\":\"_newDecisionWindow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumInvoiceValue\",\"inputs\":[{\"name\":\"_newMinimumInvoiceValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWETH\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AutomationUpdated\",\"inputs\":[{\"name\":\"automation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceAccepted\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCanceled\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceCreated\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"invoice\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structISimplePaymentProcessor.Invoice\",\"components\":[{\"name\":\"invoiceNonce\",\"type\":\"uint216\",\"internalType\":\"uint216\"},{\"name\":\"createdAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"paidAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"releaseAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"expiresAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"sellerActionDeadline\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"escrowHoldPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"withdrawalRetries\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeRate\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"buyer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"escrow\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoicePaid\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountPaid\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sellerActionDeadline\",\"type\":\"uint40\",\"indexed\":false,\"internalType\":\"uint40\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceRefunded\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceRejected\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InvoiceReleased\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"sellerAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentBurned\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFailed\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalRetried\",\"inputs\":[{\"name\":\"invoiceId\",\"type\":\"uint216\",\"indexed\":true,\"internalType\":\"uint216\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"attempt\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AcceptanceWindowExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContractPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DuplicateTask\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowWithdrawFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HoldPeriodHasNotBeenExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectPaymentAmount\",\"inputs\":[{\"name\":\"_sent\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidDecisionWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeAuthorization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHeapPosition\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInvoiceState\",\"inputs\":[{\"name\":\"_invoiceState\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvoiceAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceIsNoLongerValid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvoiceNotEligibleForRefund\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SellerCannotPayOwnedInvoice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaskNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedNativeTransfer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValueIsTooLow\",\"inputs\":[]}]",
	ID:  "Simplepaymentprocessor",
}

// Simplepaymentprocessor is an auto generated Go binding around an Ethereum contract.
type Simplepaymentprocessor struct {
	abi abi.ABI
}

// NewSimplepaymentprocessor creates a new instance of Simplepaymentprocessor.
func NewSimplepaymentprocessor() *Simplepaymentprocessor {
	parsed, err := SimplepaymentprocessorMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Simplepaymentprocessor{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Simplepaymentprocessor) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _paymentProcessorStorageAddress, uint256 _minimumInvoicePrice, address _notesAddress, address _wethAddress) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackConstructor(_paymentProcessorStorageAddress common.Address, _minimumInvoicePrice *big.Int, _notesAddress common.Address, _wethAddress common.Address) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("", _paymentProcessorStorageAddress, _minimumInvoicePrice, _notesAddress, _wethAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAcceptPayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0ecb321d.
//
// Solidity: function acceptPayment(uint216 _invoiceId, address _feeReceiver, bytes _data) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackAcceptPayment(invoiceId *big.Int, feeReceiver common.Address, data []byte) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("acceptPayment", invoiceId, feeReceiver, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCalculateFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x99a5d747.
//
// Solidity: function calculateFee(uint256 _amount) view returns(uint256 feeValue)
func (simplepaymentprocessor *Simplepaymentprocessor) PackCalculateFee(amount *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("calculateFee", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCalculateFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x99a5d747.
//
// Solidity: function calculateFee(uint256 _amount) view returns(uint256 feeValue)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackCalculateFee(data []byte) (*big.Int, error) {
	out, err := simplepaymentprocessor.abi.Unpack("calculateFee", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackCancelInvoice(invoiceId *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("cancelInvoice", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreateInvoice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x580a586c.
//
// Solidity: function createInvoice(uint256 _price, uint32 _holdPeriod, bytes _storageRef, bool _share) returns(uint216 invoiceId)
func (simplepaymentprocessor *Simplepaymentprocessor) PackCreateInvoice(price *big.Int, holdPeriod uint32, storageRef []byte, share bool) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("createInvoice", price, holdPeriod, storageRef, share)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateInvoice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x580a586c.
//
// Solidity: function createInvoice(uint256 _price, uint32 _holdPeriod, bytes _storageRef, bool _share) returns(uint216 invoiceId)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackCreateInvoice(data []byte) (*big.Int, error) {
	out, err := simplepaymentprocessor.abi.Unpack("createInvoice", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackGetAutomation() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("getAutomation")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAutomation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3a7ad1e.
//
// Solidity: function getAutomation() view returns(address automationAddress)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackGetAutomation(data []byte) (common.Address, error) {
	out, err := simplepaymentprocessor.abi.Unpack("getAutomation", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackGetDecisionWindow() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("getDecisionWindow")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDecisionWindow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d86abad.
//
// Solidity: function getDecisionWindow() view returns(uint256 decisionWindowValue)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackGetDecisionWindow(data []byte) (*big.Int, error) {
	out, err := simplepaymentprocessor.abi.Unpack("getDecisionWindow", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackGetInvoiceData(invoiceId *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("getInvoiceData", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetInvoiceData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ec7bbf9.
//
// Solidity: function getInvoiceData(uint216 _invoiceId) view returns((uint216,uint40,uint40,uint40,uint40,uint40,uint32,uint8,uint8,uint16,address,address,address,address,uint256,uint256) i)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackGetInvoiceData(data []byte) (ISimplePaymentProcessorInvoice, error) {
	out, err := simplepaymentprocessor.abi.Unpack("getInvoiceData", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackGetItems() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("getItems")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetItems is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410d59cc.
//
// Solidity: function getItems() view returns(uint216[] items)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackGetItems(data []byte) ([]*big.Int, error) {
	out, err := simplepaymentprocessor.abi.Unpack("getItems", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackGetMinimumInvoiceValue() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("getMinimumInvoiceValue")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMinimumInvoiceValue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1595a882.
//
// Solidity: function getMinimumInvoiceValue() view returns(uint256 minimumValue)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackGetMinimumInvoiceValue(data []byte) (*big.Int, error) {
	out, err := simplepaymentprocessor.abi.Unpack("getMinimumInvoiceValue", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackGetNextInvoiceNonce() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("getNextInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonceValue)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackGetNextInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := simplepaymentprocessor.abi.Unpack("getNextInvoiceNonce", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackHasDueTasks() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("hasDueTasks")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasDueTasks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0a9a0dea.
//
// Solidity: function hasDueTasks() view returns(bool dueTasksExist)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackHasDueTasks(data []byte) (bool, error) {
	out, err := simplepaymentprocessor.abi.Unpack("hasDueTasks", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackPay(invoiceId *big.Int, storageRef []byte, share bool) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("pay", invoiceId, storageRef, share)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9224df4a.
//
// Solidity: function pay(uint216 _invoiceId, bytes _storageRef, bool _share) payable returns(address escrowAddress)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackPay(data []byte) (common.Address, error) {
	out, err := simplepaymentprocessor.abi.Unpack("pay", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackPpStorage() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := simplepaymentprocessor.abi.Unpack("ppStorage", data)
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
func (simplepaymentprocessor *Simplepaymentprocessor) PackProcessDueTasks() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("processDueTasks")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRefundBuyer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x402d5a9b.
//
// Solidity: function refundBuyer(uint216 _invoiceId) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackRefundBuyer(invoiceId *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("refundBuyer", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRejectPayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98201712.
//
// Solidity: function rejectPayment(uint216 _invoiceId) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackRejectPayment(invoiceId *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("rejectPayment", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelease is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb990809.
//
// Solidity: function release(uint216 _invoiceId) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackRelease(invoiceId *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("release", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetAutomation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x730dcc8e.
//
// Solidity: function setAutomation(address _automationAddress) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackSetAutomation(automationAddress common.Address) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("setAutomation", automationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetDecisionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03dd5199.
//
// Solidity: function setDecisionWindow(uint256 _newDecisionWindow) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackSetDecisionWindow(newDecisionWindow *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("setDecisionWindow", newDecisionWindow)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinimumInvoiceValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf82afa07.
//
// Solidity: function setMinimumInvoiceValue(uint256 _newMinimumInvoiceValue) returns()
func (simplepaymentprocessor *Simplepaymentprocessor) PackSetMinimumInvoiceValue(newMinimumInvoiceValue *big.Int) []byte {
	enc, err := simplepaymentprocessor.abi.Pack("setMinimumInvoiceValue", newMinimumInvoiceValue)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWeth is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (simplepaymentprocessor *Simplepaymentprocessor) PackWeth() []byte {
	enc, err := simplepaymentprocessor.abi.Pack("weth")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackWeth is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackWeth(data []byte) (common.Address, error) {
	out, err := simplepaymentprocessor.abi.Unpack("weth", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// SimplepaymentprocessorAutomationUpdated represents a AutomationUpdated event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorAutomationUpdated struct {
	Automation common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorAutomationUpdatedEventName = "AutomationUpdated"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorAutomationUpdated) ContractEventName() string {
	return SimplepaymentprocessorAutomationUpdatedEventName
}

// UnpackAutomationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AutomationUpdated(address indexed automation)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackAutomationUpdatedEvent(log *types.Log) (*SimplepaymentprocessorAutomationUpdated, error) {
	event := "AutomationUpdated"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorAutomationUpdated)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoiceAccepted represents a InvoiceAccepted event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceAccepted struct {
	InvoiceId   *big.Int
	FeeReceiver common.Address
	ReleaseAt   *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoiceAcceptedEventName = "InvoiceAccepted"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoiceAccepted) ContractEventName() string {
	return SimplepaymentprocessorInvoiceAcceptedEventName
}

// UnpackInvoiceAcceptedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceAccepted(uint216 indexed invoiceId, address indexed feeReceiver, uint40 releaseAt)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceAcceptedEvent(log *types.Log) (*SimplepaymentprocessorInvoiceAccepted, error) {
	event := "InvoiceAccepted"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoiceAccepted)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoiceCanceled represents a InvoiceCanceled event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceCanceled struct {
	InvoiceId *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoiceCanceledEventName = "InvoiceCanceled"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoiceCanceled) ContractEventName() string {
	return SimplepaymentprocessorInvoiceCanceledEventName
}

// UnpackInvoiceCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCanceled(uint216 indexed invoiceId)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceCanceledEvent(log *types.Log) (*SimplepaymentprocessorInvoiceCanceled, error) {
	event := "InvoiceCanceled"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoiceCanceled)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoiceCreated represents a InvoiceCreated event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceCreated struct {
	InvoiceId *big.Int
	Invoice   ISimplePaymentProcessorInvoice
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoiceCreatedEventName = "InvoiceCreated"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoiceCreated) ContractEventName() string {
	return SimplepaymentprocessorInvoiceCreatedEventName
}

// UnpackInvoiceCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceCreated(uint216 indexed invoiceId, (uint216,uint40,uint40,uint40,uint40,uint40,uint32,uint8,uint8,uint16,address,address,address,address,uint256,uint256) invoice)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceCreatedEvent(log *types.Log) (*SimplepaymentprocessorInvoiceCreated, error) {
	event := "InvoiceCreated"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoiceCreated)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoicePaid represents a InvoicePaid event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoicePaid struct {
	InvoiceId            *big.Int
	Buyer                common.Address
	AmountPaid           *big.Int
	SellerActionDeadline *big.Int
	Raw                  *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoicePaidEventName = "InvoicePaid"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoicePaid) ContractEventName() string {
	return SimplepaymentprocessorInvoicePaidEventName
}

// UnpackInvoicePaidEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoicePaid(uint216 indexed invoiceId, address indexed buyer, uint256 indexed amountPaid, uint40 sellerActionDeadline)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoicePaidEvent(log *types.Log) (*SimplepaymentprocessorInvoicePaid, error) {
	event := "InvoicePaid"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoicePaid)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoiceRefunded represents a InvoiceRefunded event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceRefunded struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoiceRefundedEventName = "InvoiceRefunded"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoiceRefunded) ContractEventName() string {
	return SimplepaymentprocessorInvoiceRefundedEventName
}

// UnpackInvoiceRefundedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceRefunded(uint216 indexed invoiceId, uint256 amount)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceRefundedEvent(log *types.Log) (*SimplepaymentprocessorInvoiceRefunded, error) {
	event := "InvoiceRefunded"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoiceRefunded)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoiceRejected represents a InvoiceRejected event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceRejected struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoiceRejectedEventName = "InvoiceRejected"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoiceRejected) ContractEventName() string {
	return SimplepaymentprocessorInvoiceRejectedEventName
}

// UnpackInvoiceRejectedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceRejected(uint216 indexed invoiceId, uint256 amount)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceRejectedEvent(log *types.Log) (*SimplepaymentprocessorInvoiceRejected, error) {
	event := "InvoiceRejected"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoiceRejected)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorInvoiceReleased represents a InvoiceReleased event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceReleased struct {
	InvoiceId    *big.Int
	SellerAmount *big.Int
	Fee          *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorInvoiceReleasedEventName = "InvoiceReleased"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorInvoiceReleased) ContractEventName() string {
	return SimplepaymentprocessorInvoiceReleasedEventName
}

// UnpackInvoiceReleasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InvoiceReleased(uint216 indexed invoiceId, uint256 sellerAmount, uint256 fee)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceReleasedEvent(log *types.Log) (*SimplepaymentprocessorInvoiceReleased, error) {
	event := "InvoiceReleased"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorInvoiceReleased)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorPaymentBurned represents a PaymentBurned event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorPaymentBurned struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorPaymentBurnedEventName = "PaymentBurned"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorPaymentBurned) ContractEventName() string {
	return SimplepaymentprocessorPaymentBurnedEventName
}

// UnpackPaymentBurnedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentBurned(uint216 indexed invoiceId, uint256 amount)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackPaymentBurnedEvent(log *types.Log) (*SimplepaymentprocessorPaymentBurned, error) {
	event := "PaymentBurned"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorPaymentBurned)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorTransferFailed represents a TransferFailed event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorTransferFailed struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorTransferFailedEventName = "TransferFailed"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorTransferFailed) ContractEventName() string {
	return SimplepaymentprocessorTransferFailedEventName
}

// UnpackTransferFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferFailed(uint216 indexed invoiceId, address indexed recipient, uint256 amount)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackTransferFailedEvent(log *types.Log) (*SimplepaymentprocessorTransferFailed, error) {
	event := "TransferFailed"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorTransferFailed)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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

// SimplepaymentprocessorWithdrawalRetried represents a WithdrawalRetried event raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorWithdrawalRetried struct {
	InvoiceId *big.Int
	Recipient common.Address
	Amount    *big.Int
	Attempt   uint8
	Raw       *types.Log // Blockchain specific contextual infos
}

const SimplepaymentprocessorWithdrawalRetriedEventName = "WithdrawalRetried"

// ContractEventName returns the user-defined event name.
func (SimplepaymentprocessorWithdrawalRetried) ContractEventName() string {
	return SimplepaymentprocessorWithdrawalRetriedEventName
}

// UnpackWithdrawalRetriedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalRetried(uint216 indexed invoiceId, address indexed recipient, uint256 amount, uint8 attempt)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackWithdrawalRetriedEvent(log *types.Log) (*SimplepaymentprocessorWithdrawalRetried, error) {
	event := "WithdrawalRetried"
	if log.Topics[0] != simplepaymentprocessor.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(SimplepaymentprocessorWithdrawalRetried)
	if len(log.Data) > 0 {
		if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range simplepaymentprocessor.abi.Events[event].Inputs {
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
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["AcceptanceWindowExceeded"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackAcceptanceWindowExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["ContractPaused"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackContractPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["DuplicateTask"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackDuplicateTaskError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["EscrowWithdrawFailed"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackEscrowWithdrawFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["HoldPeriodHasNotBeenExceeded"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackHoldPeriodHasNotBeenExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["IncorrectPaymentAmount"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackIncorrectPaymentAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvalidDecisionWindow"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvalidDecisionWindowError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvalidFeeAuthorization"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvalidFeeAuthorizationError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvalidFeeReceiver"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvalidFeeReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvalidHeapPosition"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvalidHeapPositionError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvalidInvoiceState"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvalidInvoiceStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvoiceAlreadyExists"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvoiceAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvoiceIsNoLongerValid"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvoiceIsNoLongerValidError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["InvoiceNotEligibleForRefund"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackInvoiceNotEligibleForRefundError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["Reentrancy"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["SellerCannotPayOwnedInvoice"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackSellerCannotPayOwnedInvoiceError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["TaskNotFound"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackTaskNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["UnexpectedNativeTransfer"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackUnexpectedNativeTransferError(raw[4:])
	}
	if bytes.Equal(raw[:4], simplepaymentprocessor.abi.Errors["ValueIsTooLow"].ID.Bytes()[:4]) {
		return simplepaymentprocessor.UnpackValueIsTooLowError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// SimplepaymentprocessorAcceptanceWindowExceeded represents a AcceptanceWindowExceeded error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorAcceptanceWindowExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AcceptanceWindowExceeded()
func SimplepaymentprocessorAcceptanceWindowExceededErrorID() common.Hash {
	return common.HexToHash("0x2b8af0bb26858aa842f6aac91213da2376cd6baa62d4c906c6aa0685f42c1a48")
}

// UnpackAcceptanceWindowExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AcceptanceWindowExceeded()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackAcceptanceWindowExceededError(raw []byte) (*SimplepaymentprocessorAcceptanceWindowExceeded, error) {
	out := new(SimplepaymentprocessorAcceptanceWindowExceeded)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "AcceptanceWindowExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorContractPaused represents a ContractPaused error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorContractPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ContractPaused()
func SimplepaymentprocessorContractPausedErrorID() common.Hash {
	return common.HexToHash("0xab35696f06e428ebc5ceba8cd17f8fed287baf43440206d1943af1ee53e6d267")
}

// UnpackContractPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ContractPaused()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackContractPausedError(raw []byte) (*SimplepaymentprocessorContractPaused, error) {
	out := new(SimplepaymentprocessorContractPaused)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "ContractPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorDuplicateTask represents a DuplicateTask error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorDuplicateTask struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateTask()
func SimplepaymentprocessorDuplicateTaskErrorID() common.Hash {
	return common.HexToHash("0x6b22feb9606cb284058f0a2f05a53980401948118942ff234189d20361fe4e93")
}

// UnpackDuplicateTaskError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateTask()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackDuplicateTaskError(raw []byte) (*SimplepaymentprocessorDuplicateTask, error) {
	out := new(SimplepaymentprocessorDuplicateTask)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "DuplicateTask", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorEscrowWithdrawFailed represents a EscrowWithdrawFailed error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorEscrowWithdrawFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowWithdrawFailed()
func SimplepaymentprocessorEscrowWithdrawFailedErrorID() common.Hash {
	return common.HexToHash("0x667ecf9d53e4600a9a128606592ec5e22e0269990439145a2bbc8a983c7af5ac")
}

// UnpackEscrowWithdrawFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowWithdrawFailed()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackEscrowWithdrawFailedError(raw []byte) (*SimplepaymentprocessorEscrowWithdrawFailed, error) {
	out := new(SimplepaymentprocessorEscrowWithdrawFailed)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "EscrowWithdrawFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorHoldPeriodHasNotBeenExceeded represents a HoldPeriodHasNotBeenExceeded error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorHoldPeriodHasNotBeenExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HoldPeriodHasNotBeenExceeded()
func SimplepaymentprocessorHoldPeriodHasNotBeenExceededErrorID() common.Hash {
	return common.HexToHash("0xad2652ac51ebd8cf548af7b7d84853dd5d5efef6e7c1b57e0c81a88fd70d009d")
}

// UnpackHoldPeriodHasNotBeenExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HoldPeriodHasNotBeenExceeded()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackHoldPeriodHasNotBeenExceededError(raw []byte) (*SimplepaymentprocessorHoldPeriodHasNotBeenExceeded, error) {
	out := new(SimplepaymentprocessorHoldPeriodHasNotBeenExceeded)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "HoldPeriodHasNotBeenExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorIncorrectPaymentAmount represents a IncorrectPaymentAmount error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorIncorrectPaymentAmount struct {
	Sent     *big.Int
	Expected *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error IncorrectPaymentAmount(uint256 _sent, uint256 _expected)
func SimplepaymentprocessorIncorrectPaymentAmountErrorID() common.Hash {
	return common.HexToHash("0x47af6acc44a34c9e741a083d95d979d69adf7234f5e86b5e76839bf27fffd8cb")
}

// UnpackIncorrectPaymentAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error IncorrectPaymentAmount(uint256 _sent, uint256 _expected)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackIncorrectPaymentAmountError(raw []byte) (*SimplepaymentprocessorIncorrectPaymentAmount, error) {
	out := new(SimplepaymentprocessorIncorrectPaymentAmount)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "IncorrectPaymentAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvalidDecisionWindow represents a InvalidDecisionWindow error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvalidDecisionWindow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDecisionWindow()
func SimplepaymentprocessorInvalidDecisionWindowErrorID() common.Hash {
	return common.HexToHash("0x39141cc3acd0676e85f3b9927f0f60c5ff570c62f9d8e325c2849504a3f94379")
}

// UnpackInvalidDecisionWindowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDecisionWindow()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvalidDecisionWindowError(raw []byte) (*SimplepaymentprocessorInvalidDecisionWindow, error) {
	out := new(SimplepaymentprocessorInvalidDecisionWindow)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvalidDecisionWindow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvalidFeeAuthorization represents a InvalidFeeAuthorization error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvalidFeeAuthorization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeAuthorization()
func SimplepaymentprocessorInvalidFeeAuthorizationErrorID() common.Hash {
	return common.HexToHash("0x1735eabec15c7395efafdfa0dda5c74faf3517b604ff55660d6ef0e7457f2c1d")
}

// UnpackInvalidFeeAuthorizationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeAuthorization()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvalidFeeAuthorizationError(raw []byte) (*SimplepaymentprocessorInvalidFeeAuthorization, error) {
	out := new(SimplepaymentprocessorInvalidFeeAuthorization)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvalidFeeAuthorization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvalidFeeReceiver represents a InvalidFeeReceiver error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvalidFeeReceiver struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeReceiver()
func SimplepaymentprocessorInvalidFeeReceiverErrorID() common.Hash {
	return common.HexToHash("0xd200485c51caaf66763f8b49c9cfa281a0a10132cb56c8fffc35867701d3fc5f")
}

// UnpackInvalidFeeReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeReceiver()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvalidFeeReceiverError(raw []byte) (*SimplepaymentprocessorInvalidFeeReceiver, error) {
	out := new(SimplepaymentprocessorInvalidFeeReceiver)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvalidFeeReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvalidHeapPosition represents a InvalidHeapPosition error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvalidHeapPosition struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidHeapPosition()
func SimplepaymentprocessorInvalidHeapPositionErrorID() common.Hash {
	return common.HexToHash("0x76f4a2832a11e5a5adb544de785993d5d08fe0ec9c4270e1b0cbc33a9a1f27e8")
}

// UnpackInvalidHeapPositionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidHeapPosition()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvalidHeapPositionError(raw []byte) (*SimplepaymentprocessorInvalidHeapPosition, error) {
	out := new(SimplepaymentprocessorInvalidHeapPosition)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvalidHeapPosition", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvalidInvoiceState represents a InvalidInvoiceState error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvalidInvoiceState struct {
	InvoiceState *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInvoiceState(uint256 _invoiceState)
func SimplepaymentprocessorInvalidInvoiceStateErrorID() common.Hash {
	return common.HexToHash("0x1d5b155656afea98f1415b8e817932b761e3215588ffe226a1d86ff9ac55fe02")
}

// UnpackInvalidInvoiceStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInvoiceState(uint256 _invoiceState)
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvalidInvoiceStateError(raw []byte) (*SimplepaymentprocessorInvalidInvoiceState, error) {
	out := new(SimplepaymentprocessorInvalidInvoiceState)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvalidInvoiceState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvoiceAlreadyExists represents a InvoiceAlreadyExists error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceAlreadyExists()
func SimplepaymentprocessorInvoiceAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x074bc9355c94925fa82ddb49dcc88f1a666f1d1aa24efbddcdbe5f8d98b7ed59")
}

// UnpackInvoiceAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceAlreadyExists()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceAlreadyExistsError(raw []byte) (*SimplepaymentprocessorInvoiceAlreadyExists, error) {
	out := new(SimplepaymentprocessorInvoiceAlreadyExists)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvoiceAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvoiceIsNoLongerValid represents a InvoiceIsNoLongerValid error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceIsNoLongerValid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceIsNoLongerValid()
func SimplepaymentprocessorInvoiceIsNoLongerValidErrorID() common.Hash {
	return common.HexToHash("0xff42dbfc3f50dbd393d8edfcd163d614372311f4609adb53d576d6c3d588f37d")
}

// UnpackInvoiceIsNoLongerValidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceIsNoLongerValid()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceIsNoLongerValidError(raw []byte) (*SimplepaymentprocessorInvoiceIsNoLongerValid, error) {
	out := new(SimplepaymentprocessorInvoiceIsNoLongerValid)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvoiceIsNoLongerValid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorInvoiceNotEligibleForRefund represents a InvoiceNotEligibleForRefund error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorInvoiceNotEligibleForRefund struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvoiceNotEligibleForRefund()
func SimplepaymentprocessorInvoiceNotEligibleForRefundErrorID() common.Hash {
	return common.HexToHash("0xbb126ff1385e3d8206af61fac954f0af8fa8b14efaf24fb5d0f03a32a7f093be")
}

// UnpackInvoiceNotEligibleForRefundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvoiceNotEligibleForRefund()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackInvoiceNotEligibleForRefundError(raw []byte) (*SimplepaymentprocessorInvoiceNotEligibleForRefund, error) {
	out := new(SimplepaymentprocessorInvoiceNotEligibleForRefund)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "InvoiceNotEligibleForRefund", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorNotAuthorized represents a NotAuthorized error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func SimplepaymentprocessorNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackNotAuthorizedError(raw []byte) (*SimplepaymentprocessorNotAuthorized, error) {
	out := new(SimplepaymentprocessorNotAuthorized)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorReentrancy represents a Reentrancy error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Reentrancy()
func SimplepaymentprocessorReentrancyErrorID() common.Hash {
	return common.HexToHash("0xab143c06c9772d69bbbc9f2fe74acd02f810e93b099f3d1dac8448ac9ae35991")
}

// UnpackReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Reentrancy()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackReentrancyError(raw []byte) (*SimplepaymentprocessorReentrancy, error) {
	out := new(SimplepaymentprocessorReentrancy)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "Reentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorSellerCannotPayOwnedInvoice represents a SellerCannotPayOwnedInvoice error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorSellerCannotPayOwnedInvoice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SellerCannotPayOwnedInvoice()
func SimplepaymentprocessorSellerCannotPayOwnedInvoiceErrorID() common.Hash {
	return common.HexToHash("0x020175b17895ca2ebc70bd6da55e4d50ebe6fbc9b1110d0ba1ccb0613bf49691")
}

// UnpackSellerCannotPayOwnedInvoiceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SellerCannotPayOwnedInvoice()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackSellerCannotPayOwnedInvoiceError(raw []byte) (*SimplepaymentprocessorSellerCannotPayOwnedInvoice, error) {
	out := new(SimplepaymentprocessorSellerCannotPayOwnedInvoice)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "SellerCannotPayOwnedInvoice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorTaskNotFound represents a TaskNotFound error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorTaskNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TaskNotFound()
func SimplepaymentprocessorTaskNotFoundErrorID() common.Hash {
	return common.HexToHash("0xc325ae33d18e47931adbda2584c56fef1d3e5e64beab80da59968e1c83c84937")
}

// UnpackTaskNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TaskNotFound()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackTaskNotFoundError(raw []byte) (*SimplepaymentprocessorTaskNotFound, error) {
	out := new(SimplepaymentprocessorTaskNotFound)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "TaskNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorUnexpectedNativeTransfer represents a UnexpectedNativeTransfer error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorUnexpectedNativeTransfer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpectedNativeTransfer()
func SimplepaymentprocessorUnexpectedNativeTransferErrorID() common.Hash {
	return common.HexToHash("0xecb8b30d3aabfab45e153cab7e267d8f707ed670125178fc459f4849e0504fa3")
}

// UnpackUnexpectedNativeTransferError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpectedNativeTransfer()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackUnexpectedNativeTransferError(raw []byte) (*SimplepaymentprocessorUnexpectedNativeTransfer, error) {
	out := new(SimplepaymentprocessorUnexpectedNativeTransfer)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "UnexpectedNativeTransfer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// SimplepaymentprocessorValueIsTooLow represents a ValueIsTooLow error raised by the Simplepaymentprocessor contract.
type SimplepaymentprocessorValueIsTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ValueIsTooLow()
func SimplepaymentprocessorValueIsTooLowErrorID() common.Hash {
	return common.HexToHash("0x5033f274524623c8f2e5518c3a1c9e6345ae1da946950ff4e2153fe3f55d28ec")
}

// UnpackValueIsTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ValueIsTooLow()
func (simplepaymentprocessor *Simplepaymentprocessor) UnpackValueIsTooLowError(raw []byte) (*SimplepaymentprocessorValueIsTooLow, error) {
	out := new(SimplepaymentprocessorValueIsTooLow)
	if err := simplepaymentprocessor.abi.UnpackIntoInterface(out, "ValueIsTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}
