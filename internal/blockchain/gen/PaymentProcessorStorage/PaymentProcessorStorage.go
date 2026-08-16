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
	FeeRate           *big.Int
	FeeReceiver       common.Address
	DefaultHoldPeriod *big.Int
	Marketplace       common.Address
	GasThreshold      *big.Int
}

// ProcessorstorageMetaData contains all meta data concerning the Processorstorage contract.
var ProcessorstorageMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_configuration\",\"type\":\"tuple\",\"internalType\":\"structIPaymentProcessorStorage.Configuration\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeRate\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultHoldPeriod\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"marketplace\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_PAYMENT_VALIDITY_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMERGENCY_PAUSE_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveEmergencyPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"emergencyPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDefaultHoldPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"defaultHoldPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEmergencyPauseExpiry\",\"inputs\":[],\"outputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEmergencyPauser\",\"inputs\":[],\"outputs\":[{\"name\":\"emergencyPauserAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"feeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGasThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"gasThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMarketplace\",\"inputs\":[],\"outputs\":[{\"name\":\"marketplace\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextInvoiceNonceValue\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentValidityDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"validDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"pausedState\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDefaultHoldPeriod\",\"inputs\":[{\"name\":\"_newDefaultHoldPeriod\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmergencyPauser\",\"inputs\":[{\"name\":\"_emergencyPauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeRate\",\"inputs\":[{\"name\":\"_newFeeRate\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeReceiver\",\"inputs\":[{\"name\":\"_feeReceiverAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasThreshold\",\"inputs\":[{\"name\":\"_newGasThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMarketplaceAddress\",\"inputs\":[{\"name\":\"_marketplaceAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentValidityDuration\",\"inputs\":[{\"name\":\"_newValidityDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"totalInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateInvoiceNonce\",\"inputs\":[{\"name\":\"_by\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"totalInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizationUpdated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"authorized\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigurationInitialized\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentProcessorStorage.Configuration\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeRate\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"defaultHoldPeriod\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"marketplace\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultHoldPeriodUpdated\",\"inputs\":[{\"name\":\"defaultHoldPeriod\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPauseApproved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPauserUpdated\",\"inputs\":[{\"name\":\"emergencyPauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRateUpdated\",\"inputs\":[{\"name\":\"feeRate\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeReceiverUpdated\",\"inputs\":[{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasThresholdUpdated\",\"inputs\":[{\"name\":\"gasThreshold\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MarketplaceUpdated\",\"inputs\":[{\"name\":\"marketplace\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentValidityDurationUpdated\",\"inputs\":[{\"name\":\"validityDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HoldPeriodCanNotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveEmergencyPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
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
// Solidity: constructor((address,uint96,address,uint96,address,uint96) _configuration) returns()
func (processorstorage *Processorstorage) PackConstructor(_configuration IPaymentProcessorStorageConfiguration) []byte {
	enc, err := processorstorage.abi.Pack("", _configuration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBASISPOINTS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (processorstorage *Processorstorage) PackBASISPOINTS() []byte {
	enc, err := processorstorage.abi.Pack("BASIS_POINTS")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBASISPOINTS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (processorstorage *Processorstorage) UnpackBASISPOINTS(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("BASIS_POINTS", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDEFAULTPAYMENTVALIDITYPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4b332b2.
//
// Solidity: function DEFAULT_PAYMENT_VALIDITY_PERIOD() view returns(uint256)
func (processorstorage *Processorstorage) PackDEFAULTPAYMENTVALIDITYPERIOD() []byte {
	enc, err := processorstorage.abi.Pack("DEFAULT_PAYMENT_VALIDITY_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTPAYMENTVALIDITYPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd4b332b2.
//
// Solidity: function DEFAULT_PAYMENT_VALIDITY_PERIOD() view returns(uint256)
func (processorstorage *Processorstorage) UnpackDEFAULTPAYMENTVALIDITYPERIOD(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("DEFAULT_PAYMENT_VALIDITY_PERIOD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackEMERGENCYPAUSEDURATION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f9f85ef.
//
// Solidity: function EMERGENCY_PAUSE_DURATION() view returns(uint256)
func (processorstorage *Processorstorage) PackEMERGENCYPAUSEDURATION() []byte {
	enc, err := processorstorage.abi.Pack("EMERGENCY_PAUSE_DURATION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEMERGENCYPAUSEDURATION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f9f85ef.
//
// Solidity: function EMERGENCY_PAUSE_DURATION() view returns(uint256)
func (processorstorage *Processorstorage) UnpackEMERGENCYPAUSEDURATION(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("EMERGENCY_PAUSE_DURATION", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackApproveEmergencyPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37508c64.
//
// Solidity: function approveEmergencyPause() returns()
func (processorstorage *Processorstorage) PackApproveEmergencyPause() []byte {
	enc, err := processorstorage.abi.Pack("approveEmergencyPause")
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

// PackEmergencyPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (processorstorage *Processorstorage) PackEmergencyPause() []byte {
	enc, err := processorstorage.abi.Pack("emergencyPause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetDefaultHoldPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4026d697.
//
// Solidity: function getDefaultHoldPeriod() view returns(uint256 defaultHoldPeriod)
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
// Solidity: function getDefaultHoldPeriod() view returns(uint256 defaultHoldPeriod)
func (processorstorage *Processorstorage) UnpackGetDefaultHoldPeriod(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getDefaultHoldPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetEmergencyPauseExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x498bbe69.
//
// Solidity: function getEmergencyPauseExpiry() view returns(uint256 expiry)
func (processorstorage *Processorstorage) PackGetEmergencyPauseExpiry() []byte {
	enc, err := processorstorage.abi.Pack("getEmergencyPauseExpiry")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEmergencyPauseExpiry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x498bbe69.
//
// Solidity: function getEmergencyPauseExpiry() view returns(uint256 expiry)
func (processorstorage *Processorstorage) UnpackGetEmergencyPauseExpiry(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getEmergencyPauseExpiry", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetEmergencyPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x003f0342.
//
// Solidity: function getEmergencyPauser() view returns(address emergencyPauserAddress)
func (processorstorage *Processorstorage) PackGetEmergencyPauser() []byte {
	enc, err := processorstorage.abi.Pack("getEmergencyPauser")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEmergencyPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x003f0342.
//
// Solidity: function getEmergencyPauser() view returns(address emergencyPauserAddress)
func (processorstorage *Processorstorage) UnpackGetEmergencyPauser(data []byte) (common.Address, error) {
	out, err := processorstorage.abi.Unpack("getEmergencyPauser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84e5eed0.
//
// Solidity: function getFeeRate() view returns(uint256 feeRate)
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
// Solidity: function getFeeRate() view returns(uint256 feeRate)
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
// Solidity: function getFeeReceiver() view returns(address feeReceiver)
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
// Solidity: function getFeeReceiver() view returns(address feeReceiver)
func (processorstorage *Processorstorage) UnpackGetFeeReceiver(data []byte) (common.Address, error) {
	out, err := processorstorage.abi.Unpack("getFeeReceiver", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetGasThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55b8245f.
//
// Solidity: function getGasThreshold() view returns(uint256 gasThreshold)
func (processorstorage *Processorstorage) PackGetGasThreshold() []byte {
	enc, err := processorstorage.abi.Pack("getGasThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetGasThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x55b8245f.
//
// Solidity: function getGasThreshold() view returns(uint256 gasThreshold)
func (processorstorage *Processorstorage) UnpackGetGasThreshold(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getGasThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetMarketplace is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d21bcd5.
//
// Solidity: function getMarketplace() view returns(address marketplace)
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
// Solidity: function getMarketplace() view returns(address marketplace)
func (processorstorage *Processorstorage) UnpackGetMarketplace(data []byte) (common.Address, error) {
	out, err := processorstorage.abi.Unpack("getMarketplace", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetNextInvoiceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonceValue)
func (processorstorage *Processorstorage) PackGetNextInvoiceNonce() []byte {
	enc, err := processorstorage.abi.Pack("getNextInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonceValue)
func (processorstorage *Processorstorage) UnpackGetNextInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getNextInvoiceNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetPaymentValidityDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82f0db25.
//
// Solidity: function getPaymentValidityDuration() view returns(uint256 validDuration)
func (processorstorage *Processorstorage) PackGetPaymentValidityDuration() []byte {
	enc, err := processorstorage.abi.Pack("getPaymentValidityDuration")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPaymentValidityDuration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82f0db25.
//
// Solidity: function getPaymentValidityDuration() view returns(uint256 validDuration)
func (processorstorage *Processorstorage) UnpackGetPaymentValidityDuration(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("getPaymentValidityDuration", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackIsPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool pausedState)
func (processorstorage *Processorstorage) PackIsPaused() []byte {
	enc, err := processorstorage.abi.Pack("isPaused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool pausedState)
func (processorstorage *Processorstorage) UnpackIsPaused(data []byte) (bool, error) {
	out, err := processorstorage.abi.Unpack("isPaused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
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

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.
//
// Solidity: function pause() returns()
func (processorstorage *Processorstorage) PackPause() []byte {
	enc, err := processorstorage.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
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

// PackSetDefaultHoldPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4ccf11a.
//
// Solidity: function setDefaultHoldPeriod(uint96 _newDefaultHoldPeriod) returns()
func (processorstorage *Processorstorage) PackSetDefaultHoldPeriod(newDefaultHoldPeriod *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("setDefaultHoldPeriod", newDefaultHoldPeriod)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEmergencyPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e70838b.
//
// Solidity: function setEmergencyPauser(address _emergencyPauser) returns()
func (processorstorage *Processorstorage) PackSetEmergencyPauser(emergencyPauser common.Address) []byte {
	enc, err := processorstorage.abi.Pack("setEmergencyPauser", emergencyPauser)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3775e26.
//
// Solidity: function setFeeRate(uint96 _newFeeRate) returns()
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
// Solidity: function setFeeReceiver(address _feeReceiverAddress) returns()
func (processorstorage *Processorstorage) PackSetFeeReceiver(feeReceiverAddress common.Address) []byte {
	enc, err := processorstorage.abi.Pack("setFeeReceiver", feeReceiverAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetGasThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42d5816f.
//
// Solidity: function setGasThreshold(uint96 _newGasThreshold) returns()
func (processorstorage *Processorstorage) PackSetGasThreshold(newGasThreshold *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("setGasThreshold", newGasThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMarketplaceAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb47cc556.
//
// Solidity: function setMarketplaceAddress(address _marketplaceAddress) returns()
func (processorstorage *Processorstorage) PackSetMarketplaceAddress(marketplaceAddress common.Address) []byte {
	enc, err := processorstorage.abi.Pack("setMarketplaceAddress", marketplaceAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPaymentValidityDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22dceb69.
//
// Solidity: function setPaymentValidityDuration(uint256 _newValidityDuration) returns()
func (processorstorage *Processorstorage) PackSetPaymentValidityDuration(newValidityDuration *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("setPaymentValidityDuration", newValidityDuration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3d1e2ba.
//
// Solidity: function totalInvoiceCreated() view returns(uint216 totalInvoices)
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
// Solidity: function totalInvoiceCreated() view returns(uint216 totalInvoices)
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

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (processorstorage *Processorstorage) PackUnpause() []byte {
	enc, err := processorstorage.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateInvoiceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e05db13.
//
// Solidity: function updateInvoiceNonce(uint216 _by) returns(uint216 totalInvoices)
func (processorstorage *Processorstorage) PackUpdateInvoiceNonce(by *big.Int) []byte {
	enc, err := processorstorage.abi.Pack("updateInvoiceNonce", by)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUpdateInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9e05db13.
//
// Solidity: function updateInvoiceNonce(uint216 _by) returns(uint216 totalInvoices)
func (processorstorage *Processorstorage) UnpackUpdateInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := processorstorage.abi.Unpack("updateInvoiceNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// ProcessorstorageAuthorizationUpdated represents a AuthorizationUpdated event raised by the Processorstorage contract.
type ProcessorstorageAuthorizationUpdated struct {
	Account    common.Address
	Authorized bool
	Raw        *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageAuthorizationUpdatedEventName = "AuthorizationUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageAuthorizationUpdated) ContractEventName() string {
	return ProcessorstorageAuthorizationUpdatedEventName
}

// UnpackAuthorizationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorizationUpdated(address indexed account, bool authorized)
func (processorstorage *Processorstorage) UnpackAuthorizationUpdatedEvent(log *types.Log) (*ProcessorstorageAuthorizationUpdated, error) {
	event := "AuthorizationUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageAuthorizationUpdated)
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

// ProcessorstorageConfigurationInitialized represents a ConfigurationInitialized event raised by the Processorstorage contract.
type ProcessorstorageConfigurationInitialized struct {
	Config IPaymentProcessorStorageConfiguration
	Raw    *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageConfigurationInitializedEventName = "ConfigurationInitialized"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageConfigurationInitialized) ContractEventName() string {
	return ProcessorstorageConfigurationInitializedEventName
}

// UnpackConfigurationInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ConfigurationInitialized((address,uint96,address,uint96,address,uint96) config)
func (processorstorage *Processorstorage) UnpackConfigurationInitializedEvent(log *types.Log) (*ProcessorstorageConfigurationInitialized, error) {
	event := "ConfigurationInitialized"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageConfigurationInitialized)
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

// ProcessorstorageDefaultHoldPeriodUpdated represents a DefaultHoldPeriodUpdated event raised by the Processorstorage contract.
type ProcessorstorageDefaultHoldPeriodUpdated struct {
	DefaultHoldPeriod *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageDefaultHoldPeriodUpdatedEventName = "DefaultHoldPeriodUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageDefaultHoldPeriodUpdated) ContractEventName() string {
	return ProcessorstorageDefaultHoldPeriodUpdatedEventName
}

// UnpackDefaultHoldPeriodUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultHoldPeriodUpdated(uint96 defaultHoldPeriod)
func (processorstorage *Processorstorage) UnpackDefaultHoldPeriodUpdatedEvent(log *types.Log) (*ProcessorstorageDefaultHoldPeriodUpdated, error) {
	event := "DefaultHoldPeriodUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageDefaultHoldPeriodUpdated)
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

// ProcessorstorageEmergencyPauseApproved represents a EmergencyPauseApproved event raised by the Processorstorage contract.
type ProcessorstorageEmergencyPauseApproved struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageEmergencyPauseApprovedEventName = "EmergencyPauseApproved"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageEmergencyPauseApproved) ContractEventName() string {
	return ProcessorstorageEmergencyPauseApprovedEventName
}

// UnpackEmergencyPauseApprovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyPauseApproved(address indexed account)
func (processorstorage *Processorstorage) UnpackEmergencyPauseApprovedEvent(log *types.Log) (*ProcessorstorageEmergencyPauseApproved, error) {
	event := "EmergencyPauseApproved"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageEmergencyPauseApproved)
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

// ProcessorstorageEmergencyPaused represents a EmergencyPaused event raised by the Processorstorage contract.
type ProcessorstorageEmergencyPaused struct {
	Account common.Address
	Expiry  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageEmergencyPausedEventName = "EmergencyPaused"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageEmergencyPaused) ContractEventName() string {
	return ProcessorstorageEmergencyPausedEventName
}

// UnpackEmergencyPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyPaused(address indexed account, uint256 expiry)
func (processorstorage *Processorstorage) UnpackEmergencyPausedEvent(log *types.Log) (*ProcessorstorageEmergencyPaused, error) {
	event := "EmergencyPaused"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageEmergencyPaused)
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

// ProcessorstorageEmergencyPauserUpdated represents a EmergencyPauserUpdated event raised by the Processorstorage contract.
type ProcessorstorageEmergencyPauserUpdated struct {
	EmergencyPauser common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageEmergencyPauserUpdatedEventName = "EmergencyPauserUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageEmergencyPauserUpdated) ContractEventName() string {
	return ProcessorstorageEmergencyPauserUpdatedEventName
}

// UnpackEmergencyPauserUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyPauserUpdated(address indexed emergencyPauser)
func (processorstorage *Processorstorage) UnpackEmergencyPauserUpdatedEvent(log *types.Log) (*ProcessorstorageEmergencyPauserUpdated, error) {
	event := "EmergencyPauserUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageEmergencyPauserUpdated)
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

// ProcessorstorageFeeRateUpdated represents a FeeRateUpdated event raised by the Processorstorage contract.
type ProcessorstorageFeeRateUpdated struct {
	FeeRate *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageFeeRateUpdatedEventName = "FeeRateUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageFeeRateUpdated) ContractEventName() string {
	return ProcessorstorageFeeRateUpdatedEventName
}

// UnpackFeeRateUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRateUpdated(uint96 feeRate)
func (processorstorage *Processorstorage) UnpackFeeRateUpdatedEvent(log *types.Log) (*ProcessorstorageFeeRateUpdated, error) {
	event := "FeeRateUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageFeeRateUpdated)
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

// ProcessorstorageFeeReceiverUpdated represents a FeeReceiverUpdated event raised by the Processorstorage contract.
type ProcessorstorageFeeReceiverUpdated struct {
	FeeReceiver common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageFeeReceiverUpdatedEventName = "FeeReceiverUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageFeeReceiverUpdated) ContractEventName() string {
	return ProcessorstorageFeeReceiverUpdatedEventName
}

// UnpackFeeReceiverUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (processorstorage *Processorstorage) UnpackFeeReceiverUpdatedEvent(log *types.Log) (*ProcessorstorageFeeReceiverUpdated, error) {
	event := "FeeReceiverUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageFeeReceiverUpdated)
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

// ProcessorstorageGasThresholdUpdated represents a GasThresholdUpdated event raised by the Processorstorage contract.
type ProcessorstorageGasThresholdUpdated struct {
	GasThreshold *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageGasThresholdUpdatedEventName = "GasThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageGasThresholdUpdated) ContractEventName() string {
	return ProcessorstorageGasThresholdUpdatedEventName
}

// UnpackGasThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GasThresholdUpdated(uint96 gasThreshold)
func (processorstorage *Processorstorage) UnpackGasThresholdUpdatedEvent(log *types.Log) (*ProcessorstorageGasThresholdUpdated, error) {
	event := "GasThresholdUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageGasThresholdUpdated)
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

// ProcessorstorageMarketplaceUpdated represents a MarketplaceUpdated event raised by the Processorstorage contract.
type ProcessorstorageMarketplaceUpdated struct {
	Marketplace common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageMarketplaceUpdatedEventName = "MarketplaceUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageMarketplaceUpdated) ContractEventName() string {
	return ProcessorstorageMarketplaceUpdatedEventName
}

// UnpackMarketplaceUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MarketplaceUpdated(address indexed marketplace)
func (processorstorage *Processorstorage) UnpackMarketplaceUpdatedEvent(log *types.Log) (*ProcessorstorageMarketplaceUpdated, error) {
	event := "MarketplaceUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageMarketplaceUpdated)
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

// ProcessorstoragePaused represents a Paused event raised by the Processorstorage contract.
type ProcessorstoragePaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ProcessorstoragePausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (ProcessorstoragePaused) ContractEventName() string {
	return ProcessorstoragePausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address indexed account)
func (processorstorage *Processorstorage) UnpackPausedEvent(log *types.Log) (*ProcessorstoragePaused, error) {
	event := "Paused"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstoragePaused)
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

// ProcessorstoragePaymentValidityDurationUpdated represents a PaymentValidityDurationUpdated event raised by the Processorstorage contract.
type ProcessorstoragePaymentValidityDurationUpdated struct {
	ValidityDuration *big.Int
	Raw              *types.Log // Blockchain specific contextual infos
}

const ProcessorstoragePaymentValidityDurationUpdatedEventName = "PaymentValidityDurationUpdated"

// ContractEventName returns the user-defined event name.
func (ProcessorstoragePaymentValidityDurationUpdated) ContractEventName() string {
	return ProcessorstoragePaymentValidityDurationUpdatedEventName
}

// UnpackPaymentValidityDurationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentValidityDurationUpdated(uint256 validityDuration)
func (processorstorage *Processorstorage) UnpackPaymentValidityDurationUpdatedEvent(log *types.Log) (*ProcessorstoragePaymentValidityDurationUpdated, error) {
	event := "PaymentValidityDurationUpdated"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstoragePaymentValidityDurationUpdated)
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

// ProcessorstorageUnpaused represents a Unpaused event raised by the Processorstorage contract.
type ProcessorstorageUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ProcessorstorageUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (ProcessorstorageUnpaused) ContractEventName() string {
	return ProcessorstorageUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address indexed account)
func (processorstorage *Processorstorage) UnpackUnpausedEvent(log *types.Log) (*ProcessorstorageUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != processorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ProcessorstorageUnpaused)
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
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["AlreadyPaused"].ID.Bytes()[:4]) {
		return processorstorage.UnpackAlreadyPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["HoldPeriodCanNotBeZero"].ID.Bytes()[:4]) {
		return processorstorage.UnpackHoldPeriodCanNotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["InvalidFeeRate"].ID.Bytes()[:4]) {
		return processorstorage.UnpackInvalidFeeRateError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNewOwnerIsZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NoActiveEmergencyPause"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNoActiveEmergencyPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NoHandoverRequest"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNoHandoverRequestError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], processorstorage.abi.Errors["NotPaused"].ID.Bytes()[:4]) {
		return processorstorage.UnpackNotPausedError(raw[4:])
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

// ProcessorstorageAlreadyPaused represents a AlreadyPaused error raised by the Processorstorage contract.
type ProcessorstorageAlreadyPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyPaused()
func ProcessorstorageAlreadyPausedErrorID() common.Hash {
	return common.HexToHash("0x1785c68176ff5ca26e02299a48022fe13a267aed4ebbbf517400769c3e8e8df7")
}

// UnpackAlreadyPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyPaused()
func (processorstorage *Processorstorage) UnpackAlreadyPausedError(raw []byte) (*ProcessorstorageAlreadyPaused, error) {
	out := new(ProcessorstorageAlreadyPaused)
	if err := processorstorage.abi.UnpackIntoInterface(out, "AlreadyPaused", raw); err != nil {
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

// ProcessorstorageInvalidFeeRate represents a InvalidFeeRate error raised by the Processorstorage contract.
type ProcessorstorageInvalidFeeRate struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeRate()
func ProcessorstorageInvalidFeeRateErrorID() common.Hash {
	return common.HexToHash("0x56d69198c50c349b33dac636e06a8847667e835557d137a3943dab95f3d5ce59")
}

// UnpackInvalidFeeRateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeRate()
func (processorstorage *Processorstorage) UnpackInvalidFeeRateError(raw []byte) (*ProcessorstorageInvalidFeeRate, error) {
	out := new(ProcessorstorageInvalidFeeRate)
	if err := processorstorage.abi.UnpackIntoInterface(out, "InvalidFeeRate", raw); err != nil {
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

// ProcessorstorageNoActiveEmergencyPause represents a NoActiveEmergencyPause error raised by the Processorstorage contract.
type ProcessorstorageNoActiveEmergencyPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoActiveEmergencyPause()
func ProcessorstorageNoActiveEmergencyPauseErrorID() common.Hash {
	return common.HexToHash("0xdb469296406d35a0b1b7813ae476c3491d1c8f00379ed4f980870b27cbd368b2")
}

// UnpackNoActiveEmergencyPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoActiveEmergencyPause()
func (processorstorage *Processorstorage) UnpackNoActiveEmergencyPauseError(raw []byte) (*ProcessorstorageNoActiveEmergencyPause, error) {
	out := new(ProcessorstorageNoActiveEmergencyPause)
	if err := processorstorage.abi.UnpackIntoInterface(out, "NoActiveEmergencyPause", raw); err != nil {
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

// ProcessorstorageNotPaused represents a NotPaused error raised by the Processorstorage contract.
type ProcessorstorageNotPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotPaused()
func ProcessorstorageNotPausedErrorID() common.Hash {
	return common.HexToHash("0x6cd602013233635730773e15e89b8a778034d859147e8f706bcd1aa42e228e06")
}

// UnpackNotPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotPaused()
func (processorstorage *Processorstorage) UnpackNotPausedError(raw []byte) (*ProcessorstorageNotPaused, error) {
	out := new(ProcessorstorageNotPaused)
	if err := processorstorage.abi.UnpackIntoInterface(out, "NotPaused", raw); err != nil {
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
