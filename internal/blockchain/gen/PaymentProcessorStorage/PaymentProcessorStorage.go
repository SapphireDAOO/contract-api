// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package processorstorage

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

// IPaymentProcessorStorageConfiguration is an auto generated low-level Go binding around an user-defined struct.
type IPaymentProcessorStorageConfiguration struct {
	Owner             common.Address
	FeeReceiver       common.Address
	Marketplace       common.Address
	FeeRate           *big.Int
	DefaultHoldPeriod *big.Int
	GasThresold       *big.Int
}

// ProcessorstorageMetaData contains all meta data concerning the Processorstorage contract.
var ProcessorstorageMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"defaultHoldPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasThresold\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymentProcessorStorage.Configuration\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HoldPeriodCanNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultHoldPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasThresold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextInvoiceId\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorizedAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"}],\"name\":\"setAuthorizedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDefaultHoldPeriod\",\"type\":\"uint256\"}],\"name\":\"setDefaultHoldPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeReceiverAddress\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasThresold\",\"type\":\"uint256\"}],\"name\":\"setGasThresold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketplaceAddress\",\"type\":\"address\"}],\"name\":\"setMarketplaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInvoiceCreated\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"by\",\"type\":\"uint216\"}],\"name\":\"updateInvoiceId\",\"outputs\":[{\"internalType\":\"uint216\",\"name\":\"\",\"type\":\"uint216\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Processorstorage",
}

// Processorstorage is an auto generated Go binding around an Ethereum contract.
type Processorstorage struct {
	abi abi.ABI
}

// NewProcessorstorage creates a new instance of Processorstorage.
func NewProcessorstorage() *Processorstorage {
	parsed, err := ProcessorstorageMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Processorstorage{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Processorstorage) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor((address,address,address,uint256,uint256,uint256) configuration) returns()
func (processorstorage *Processorstorage) PackConstructor(configuration IPaymentProcessorStorageConfiguration) []byte {
	enc, err := processorstorage.abi.Pack("", configuration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (processorstorage *Processorstorage) PackCancelOwnershipHandover() []byte {
	enc, err := processorstorage.abi.Pack("cancelOwnershipHandover")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCompleteOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (processorstorage *Processorstorage) PackCompleteOwnershipHandover(pendingOwner common.Address) []byte {
	enc, err := processorstorage.abi.Pack("completeOwnershipHandover", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (processorstorage *Processorstorage) PackExecute(target common.Address, data []byte) []byte {
	enc, err := processorstorage.abi.Pack("execute", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecute is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1cff79cd.
//
// Solidity: function execute(address target, bytes data) returns(bytes)
func (processorstorage *Processorstorage) UnpackExecute(data []byte) ([]byte, error) {
	out, err := processorstorage.abi.Unpack("execute", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// PackGetDefaultHoldPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4026d697.
//
// Solidity: function getDefaultHoldPeriod() view returns(uint256)
func (processorstorage *Processorstorage) PackGetDefaultHoldPeriod() []byte {
	enc, err := processorstorage.abi.Pack("getDefaultHoldPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetDefaultHoldPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4026d697.
//
// Solidity: function getDefaultHoldPeriod() view returns(uint256)
func (processorstorage *Processorstorage) UnpackGetDefaultHoldPeriod(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getDefaultHoldPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84e5eed0.
//
// Solidity: function getFeeRate() view returns(uint256)
func (processorstorage *Processorstorage) PackGetFeeRate() []byte {
	enc, err := processorstorage.abi.Pack("getFeeRate")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFeeRate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84e5eed0.
//
// Solidity: function getFeeRate() view returns(uint256)
func (processorstorage *Processorstorage) UnpackGetFeeRate(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getFeeRate", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetFeeReceiver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (processorstorage *Processorstorage) PackGetFeeReceiver() []byte {
	enc, err := processorstorage.abi.Pack("getFeeReceiver")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFeeReceiver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (processorstorage *Processorstorage) UnpackGetFeeReceiver(data []byte) (common.Address, error) {
	out, err := processorstorage.abi.Unpack("getFeeReceiver", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetGasThresold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x243da283.
//
// Solidity: function getGasThresold() view returns(uint256)
func (processorstorage *Processorstorage) PackGetGasThresold() []byte {
	enc, err := processorstorage.abi.Pack("getGasThresold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetGasThresold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x243da283.
//
// Solidity: function getGasThresold() view returns(uint256)
func (processorstorage *Processorstorage) UnpackGetGasThresold(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getGasThresold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetMarketplace is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (processorstorage *Processorstorage) PackGetMarketplace() []byte {
	enc, err := processorstorage.abi.Pack("getMarketplace")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetMarketplace is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address)
func (processorstorage *Processorstorage) UnpackGetMarketplace(data []byte) (common.Address, error) {
	out, err := processorstorage.abi.Unpack("getMarketplace", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetNextInvoiceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint216)
func (processorstorage *Processorstorage) PackGetNextInvoiceId() []byte {
	enc, err := processorstorage.abi.Pack("getNextInvoiceId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1471dcb3.
//
// Solidity: function getNextInvoiceId() view returns(uint216)
func (processorstorage *Processorstorage) UnpackGetNextInvoiceId(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getNextInvoiceId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (processorstorage *Processorstorage) PackOwner() []byte {
	enc, err := processorstorage.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (processorstorage *Processorstorage) UnpackOwner(data []byte) (common.Address, error) {
	out, err := processorstorage.abi.Unpack("owner", data)
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
func (processorstorage *Processorstorage) PackOwnershipHandoverExpiresAt(pendingOwner common.Address) []byte {
	enc, err := processorstorage.abi.Pack("ownershipHandoverExpiresAt", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnershipHandoverExpiresAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (processorstorage *Processorstorage) UnpackOwnershipHandoverExpiresAt(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("ownershipHandoverExpiresAt", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (processorstorage *Processorstorage) PackRenounceOwnership() []byte {
	enc, err := processorstorage.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (processorstorage *Processorstorage) PackRequestOwnershipHandover() []byte {
	enc, err := processorstorage.abi.Pack("requestOwnershipHandover")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetAuthorizedAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1351cf51.
//
// Solidity: function setAuthorizedAddress(address authorizedAddress, bool authorized) returns()
func (processorstorage *Processorstorage) PackSetAuthorizedAddress(authorizedAddress common.Address, authorized bool) []byte {
	enc, err := processorstorage.abi.Pack("setAuthorizedAddress", authorizedAddress, authorized)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetDefaultHoldPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5aabfe3.
//
// Solidity: function setDefaultHoldPeriod(uint256 newDefaultHoldPeriod) returns()
func (processorstorage *Processorstorage) PackSetDefaultHoldPeriod(newDefaultHoldPeriod *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("setDefaultHoldPeriod", newDefaultHoldPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 newFeeRate) returns()
func (processorstorage *Processorstorage) PackSetFeeRate(newFeeRate *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("setFeeRate", newFeeRate)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeReceiver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefdcd974.
//
// Solidity: function setFeeReceiver(address feeReceiverAddress) returns()
func (processorstorage *Processorstorage) PackSetFeeReceiver(feeReceiverAddress common.Address) []byte {
	enc, err := processorstorage.abi.Pack("setFeeReceiver", feeReceiverAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetGasThresold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da43bec.
//
// Solidity: function setGasThresold(uint256 newGasThresold) returns()
func (processorstorage *Processorstorage) PackSetGasThresold(newGasThresold *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("setGasThresold", newGasThresold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMarketplaceAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb47cc556.
//
// Solidity: function setMarketplaceAddress(address marketplaceAddress) returns()
func (processorstorage *Processorstorage) PackSetMarketplaceAddress(marketplaceAddress common.Address) []byte {
	enc, err := processorstorage.abi.Pack("setMarketplaceAddress", marketplaceAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3d1e2ba.
//
// Solidity: function totalInvoiceCreated() view returns(uint216)
func (processorstorage *Processorstorage) PackTotalInvoiceCreated() []byte {
	enc, err := processorstorage.abi.Pack("totalInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3d1e2ba.
//
// Solidity: function totalInvoiceCreated() view returns(uint216)
func (processorstorage *Processorstorage) UnpackTotalInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("totalInvoiceCreated", data)
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
func (processorstorage *Processorstorage) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := processorstorage.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateInvoiceId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e9384ef.
//
// Solidity: function updateInvoiceId(uint216 by) returns(uint216)
func (processorstorage *Processorstorage) PackUpdateInvoiceId(by *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("updateInvoiceId", by)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUpdateInvoiceId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e9384ef.
//
// Solidity: function updateInvoiceId(uint216 by) returns(uint216)
func (processorstorage *Processorstorage) UnpackUpdateInvoiceId(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("updateInvoiceId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// ProcessorstorageOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Processorstorage contract.
type ProcessorstorageOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageOwnershipHandoverCanceledEventName = "OwnershipHandoverCanceled"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageOwnershipHandoverCanceled) ContractEventName() string {
	return ProcessorstorageOwnershipHandoverCanceledEventName
}

// UnpackOwnershipHandoverCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (processorstorage *Processorstorage) UnpackOwnershipHandoverCanceledEvent(log *types.Log) (*ProcessorstorageOwnershipHandoverCanceled, error) {
	event := "OwnershipHandoverCanceled"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageOwnershipHandoverCanceled)
	if len(log.Data) > 0 {
		if err := processorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range processorstorage.abi.Events[event].Inputs {
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

// ProcessorstorageOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Processorstorage contract.
type ProcessorstorageOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageOwnershipHandoverRequestedEventName = "OwnershipHandoverRequested"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageOwnershipHandoverRequested) ContractEventName() string {
	return ProcessorstorageOwnershipHandoverRequestedEventName
}

// UnpackOwnershipHandoverRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (processorstorage *Processorstorage) UnpackOwnershipHandoverRequestedEvent(log *types.Log) (*ProcessorstorageOwnershipHandoverRequested, error) {
	event := "OwnershipHandoverRequested"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageOwnershipHandoverRequested)
	if len(log.Data) > 0 {
		if err := processorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range processorstorage.abi.Events[event].Inputs {
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

// ProcessorstorageOwnershipTransferred represents a OwnershipTransferred event raised by the Processorstorage contract.
type ProcessorstorageOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageOwnershipTransferred) ContractEventName() string {
	return ProcessorstorageOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (processorstorage *Processorstorage) UnpackOwnershipTransferredEvent(log *types.Log) (*ProcessorstorageOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := processorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range processorstorage.abi.Events[event].Inputs {
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
func (processorstorage *Processorstorage) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["AlreadyInitialized"].ID.Bytes()[:4]) {
		return processorstorage.UnpackAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["CallFailed"].ID.Bytes()[:4]) {
		return processorstorage.UnpackCallFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["HoldPeriodCanNotBeZero"].ID.Bytes()[:4]) {
		return processorstorage.UnpackHoldPeriodCanNotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNewOwnerIsZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NoHandoverRequest"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNoHandoverRequestError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return processorstorage.UnpackUnauthorizedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ProcessorstorageAlreadyInitialized represents a AlreadyInitialized error raised by the Processorstorage contract.
type ProcessorstorageAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInitialized()
func ProcessorstorageAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0x0dc149f07762891dbcea3fe72770f3d63a1863fc54b2f084e8c59ec476996927")
}

// UnpackAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInitialized()
func (processorstorage *Processorstorage) UnpackAlreadyInitializedError(raw []byte) (*ProcessorstorageAlreadyInitialized, error) {
	out := new(ProcessorstorageAlreadyInitialized)
	if err := processorstorage.abi.UnpackIntoInterface(out, "AlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorstorageCallFailed represents a CallFailed error raised by the Processorstorage contract.
type ProcessorstorageCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CallFailed()
func ProcessorstorageCallFailedErrorID() common.Hash {
	return common.HexToHash("0x3204506f6a813959cffc5fd04d7ea98039e9c7444048856baf6c10ae312eb843")
}

// UnpackCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CallFailed()
func (processorstorage *Processorstorage) UnpackCallFailedError(raw []byte) (*ProcessorstorageCallFailed, error) {
	out := new(ProcessorstorageCallFailed)
	if err := processorstorage.abi.UnpackIntoInterface(out, "CallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorstorageHoldPeriodCanNotBeZero represents a HoldPeriodCanNotBeZero error raised by the Processorstorage contract.
type ProcessorstorageHoldPeriodCanNotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error HoldPeriodCanNotBeZero()
func ProcessorstorageHoldPeriodCanNotBeZeroErrorID() common.Hash {
	return common.HexToHash("0x705a71532da8bae84d5c54245bfd200d9655b2c961da65ccc7fcf54a50ad44b4")
}

// UnpackHoldPeriodCanNotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error HoldPeriodCanNotBeZero()
func (processorstorage *Processorstorage) UnpackHoldPeriodCanNotBeZeroError(raw []byte) (*ProcessorstorageHoldPeriodCanNotBeZero, error) {
	out := new(ProcessorstorageHoldPeriodCanNotBeZero)
	if err := processorstorage.abi.UnpackIntoInterface(out, "HoldPeriodCanNotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorstorageNewOwnerIsZeroAddress represents a NewOwnerIsZeroAddress error raised by the Processorstorage contract.
type ProcessorstorageNewOwnerIsZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NewOwnerIsZeroAddress()
func ProcessorstorageNewOwnerIsZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x7448fbae245b5163a637f61fac94c5376c3e155928452ce47ee52d8c1b99587a")
}

// UnpackNewOwnerIsZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NewOwnerIsZeroAddress()
func (processorstorage *Processorstorage) UnpackNewOwnerIsZeroAddressError(raw []byte) (*ProcessorstorageNewOwnerIsZeroAddress, error) {
	out := new(ProcessorstorageNewOwnerIsZeroAddress)
	if err := processorstorage.abi.UnpackIntoInterface(out, "NewOwnerIsZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorstorageNoHandoverRequest represents a NoHandoverRequest error raised by the Processorstorage contract.
type ProcessorstorageNoHandoverRequest struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoHandoverRequest()
func ProcessorstorageNoHandoverRequestErrorID() common.Hash {
	return common.HexToHash("0x6f5e8818469c73d5be4a0d17c371cde64695907022629c1d064c895f98d466a6")
}

// UnpackNoHandoverRequestError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoHandoverRequest()
func (processorstorage *Processorstorage) UnpackNoHandoverRequestError(raw []byte) (*ProcessorstorageNoHandoverRequest, error) {
	out := new(ProcessorstorageNoHandoverRequest)
	if err := processorstorage.abi.UnpackIntoInterface(out, "NoHandoverRequest", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorstorageNotAuthorized represents a NotAuthorized error raised by the Processorstorage contract.
type ProcessorstorageNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func ProcessorstorageNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (processorstorage *Processorstorage) UnpackNotAuthorizedError(raw []byte) (*ProcessorstorageNotAuthorized, error) {
	out := new(ProcessorstorageNotAuthorized)
	if err := processorstorage.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorstorageUnauthorized represents a Unauthorized error raised by the Processorstorage contract.
type ProcessorstorageUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func ProcessorstorageUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (processorstorage *Processorstorage) UnpackUnauthorizedError(raw []byte) (*ProcessorstorageUnauthorized, error) {
	out := new(ProcessorstorageUnauthorized)
	if err := processorstorage.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}
