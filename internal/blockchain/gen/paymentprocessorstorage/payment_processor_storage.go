// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymentprocessorstorage

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
	Owner                          common.Address
	FeeRate                        *big.Int
	FeeReceiver                    common.Address
	IntermediatedPlatformsOperator common.Address
	GasThreshold                   *big.Int
}

// PaymentprocessorstorageMetaData contains all meta data concerning the Paymentprocessorstorage contract.
var PaymentprocessorstorageMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_configuration\",\"type\":\"tuple\",\"internalType\":\"structIPaymentProcessorStorage.Configuration\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeRate\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"intermediatedPlatformsOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_PAYMENT_VALIDITY_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EMERGENCY_PAUSE_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approveEmergencyPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"emergencyPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEmergencyPauseExpiry\",\"inputs\":[],\"outputs\":[{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEmergencyPauser\",\"inputs\":[],\"outputs\":[{\"name\":\"emergencyPauserAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"feeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"feeSignerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGasThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"gasThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntermediatedPlatformsOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"intermediatedPlatformsOperator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextInvoiceNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"nextInvoiceNonceValue\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentValidityDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"validDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"pausedState\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setEmergencyPauser\",\"inputs\":[{\"name\":\"_emergencyPauser\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeRate\",\"inputs\":[{\"name\":\"_newFeeRate\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeReceiver\",\"inputs\":[{\"name\":\"_feeReceiverAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeSigner\",\"inputs\":[{\"name\":\"_feeSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasThreshold\",\"inputs\":[{\"name\":\"_newGasThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntermediatedPlatformsOperator\",\"inputs\":[{\"name\":\"_intermediatedPlatformsOperatorWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentValidityDuration\",\"inputs\":[{\"name\":\"_newValidityDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalInvoiceCreated\",\"inputs\":[],\"outputs\":[{\"name\":\"totalInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateInvoiceNonce\",\"inputs\":[{\"name\":\"_by\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"outputs\":[{\"name\":\"totalInvoices\",\"type\":\"uint216\",\"internalType\":\"uint216\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuthorizationUpdated\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"authorized\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigurationInitialized\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIPaymentProcessorStorage.Configuration\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeRate\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"intermediatedPlatformsOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasThreshold\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPauseApproved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPauserUpdated\",\"inputs\":[{\"name\":\"emergencyPauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeRateUpdated\",\"inputs\":[{\"name\":\"feeRate\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeReceiverUpdated\",\"inputs\":[{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeSignerUpdated\",\"inputs\":[{\"name\":\"feeSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasThresholdUpdated\",\"inputs\":[{\"name\":\"gasThreshold\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntermediatedPlatformsOperatorUpdated\",\"inputs\":[{\"name\":\"intermediatedPlatformsOperator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentValidityDurationUpdated\",\"inputs\":[{\"name\":\"validityDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoActiveEmergencyPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	ID:  "Paymentprocessorstorage",
}

// Paymentprocessorstorage is an auto generated Go binding around an Ethereum contract.
type Paymentprocessorstorage struct {
	abi abi.ABI
}

// NewPaymentprocessorstorage creates a new instance of Paymentprocessorstorage.
func NewPaymentprocessorstorage() *Paymentprocessorstorage {
	parsed, err := PaymentprocessorstorageMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Paymentprocessorstorage{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Paymentprocessorstorage) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor((address,uint96,address,address,uint96) _configuration) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackConstructor(_configuration IPaymentProcessorStorageConfiguration) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("", _configuration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBASISPOINTS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (paymentprocessorstorage *Paymentprocessorstorage) PackBASISPOINTS() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("BASIS_POINTS")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBASISPOINTS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackBASISPOINTS(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("BASIS_POINTS", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackDEFAULTPAYMENTVALIDITYPERIOD() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("DEFAULT_PAYMENT_VALIDITY_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTPAYMENTVALIDITYPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd4b332b2.
//
// Solidity: function DEFAULT_PAYMENT_VALIDITY_PERIOD() view returns(uint256)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackDEFAULTPAYMENTVALIDITYPERIOD(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("DEFAULT_PAYMENT_VALIDITY_PERIOD", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackEMERGENCYPAUSEDURATION() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("EMERGENCY_PAUSE_DURATION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEMERGENCYPAUSEDURATION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f9f85ef.
//
// Solidity: function EMERGENCY_PAUSE_DURATION() view returns(uint256)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackEMERGENCYPAUSEDURATION(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("EMERGENCY_PAUSE_DURATION", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackApproveEmergencyPause() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("approveEmergencyPause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackCancelOwnershipHandover() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("cancelOwnershipHandover")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCompleteOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackCompleteOwnershipHandover(pendingOwner common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("completeOwnershipHandover", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEmergencyPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackEmergencyPause() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("emergencyPause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetEmergencyPauseExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x498bbe69.
//
// Solidity: function getEmergencyPauseExpiry() view returns(uint256 expiry)
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetEmergencyPauseExpiry() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getEmergencyPauseExpiry")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEmergencyPauseExpiry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x498bbe69.
//
// Solidity: function getEmergencyPauseExpiry() view returns(uint256 expiry)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetEmergencyPauseExpiry(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getEmergencyPauseExpiry", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetEmergencyPauser() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getEmergencyPauser")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetEmergencyPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x003f0342.
//
// Solidity: function getEmergencyPauser() view returns(address emergencyPauserAddress)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetEmergencyPauser(data []byte) (common.Address, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getEmergencyPauser", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetFeeRate() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getFeeRate")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFeeRate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84e5eed0.
//
// Solidity: function getFeeRate() view returns(uint256 feeRate)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetFeeRate(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getFeeRate", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetFeeReceiver() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getFeeReceiver")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFeeReceiver is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address feeReceiver)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetFeeReceiver(data []byte) (common.Address, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getFeeReceiver", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetFeeSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac3ea0a2.
//
// Solidity: function getFeeSigner() view returns(address feeSignerAddress)
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetFeeSigner() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getFeeSigner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetFeeSigner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac3ea0a2.
//
// Solidity: function getFeeSigner() view returns(address feeSignerAddress)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetFeeSigner(data []byte) (common.Address, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getFeeSigner", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetGasThreshold() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getGasThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetGasThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x55b8245f.
//
// Solidity: function getGasThreshold() view returns(uint256 gasThreshold)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetGasThreshold(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getGasThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetIntermediatedPlatformsOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d2bfd47.
//
// Solidity: function getIntermediatedPlatformsOperator() view returns(address intermediatedPlatformsOperator)
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetIntermediatedPlatformsOperator() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getIntermediatedPlatformsOperator")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetIntermediatedPlatformsOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7d2bfd47.
//
// Solidity: function getIntermediatedPlatformsOperator() view returns(address intermediatedPlatformsOperator)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetIntermediatedPlatformsOperator(data []byte) (common.Address, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getIntermediatedPlatformsOperator", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetNextInvoiceNonce() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getNextInvoiceNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNextInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5614b076.
//
// Solidity: function getNextInvoiceNonce() view returns(uint216 nextInvoiceNonceValue)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetNextInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getNextInvoiceNonce", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackGetPaymentValidityDuration() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("getPaymentValidityDuration")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPaymentValidityDuration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82f0db25.
//
// Solidity: function getPaymentValidityDuration() view returns(uint256 validDuration)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGetPaymentValidityDuration(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("getPaymentValidityDuration", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackIsPaused() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("isPaused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool pausedState)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackIsPaused(data []byte) (bool, error) {
	out, err := paymentprocessorstorage.abi.Unpack("isPaused", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackOwner() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackOwner(data []byte) (common.Address, error) {
	out, err := paymentprocessorstorage.abi.Unpack("owner", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackOwnershipHandoverExpiresAt(pendingOwner common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("ownershipHandoverExpiresAt", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnershipHandoverExpiresAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackOwnershipHandoverExpiresAt(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("ownershipHandoverExpiresAt", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackPause() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackRenounceOwnership() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestOwnershipHandover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackRequestOwnershipHandover() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("requestOwnershipHandover")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetEmergencyPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e70838b.
//
// Solidity: function setEmergencyPauser(address _emergencyPauser) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetEmergencyPauser(emergencyPauser common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setEmergencyPauser", emergencyPauser)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeRate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3775e26.
//
// Solidity: function setFeeRate(uint96 _newFeeRate) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetFeeRate(newFeeRate *big.Int) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setFeeRate", newFeeRate)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeReceiver is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiverAddress) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetFeeReceiver(feeReceiverAddress common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setFeeReceiver", feeReceiverAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetFeeSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7e1aa178.
//
// Solidity: function setFeeSigner(address _feeSigner) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetFeeSigner(feeSigner common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setFeeSigner", feeSigner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetGasThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42d5816f.
//
// Solidity: function setGasThreshold(uint96 _newGasThreshold) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetGasThreshold(newGasThreshold *big.Int) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setGasThreshold", newGasThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetIntermediatedPlatformsOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a329028.
//
// Solidity: function setIntermediatedPlatformsOperator(address _intermediatedPlatformsOperatorWallet) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetIntermediatedPlatformsOperator(intermediatedPlatformsOperatorWallet common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setIntermediatedPlatformsOperator", intermediatedPlatformsOperatorWallet)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetPaymentValidityDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22dceb69.
//
// Solidity: function setPaymentValidityDuration(uint256 _newValidityDuration) returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackSetPaymentValidityDuration(newValidityDuration *big.Int) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("setPaymentValidityDuration", newValidityDuration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTotalInvoiceCreated is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3d1e2ba.
//
// Solidity: function totalInvoiceCreated() view returns(uint216 totalInvoices)
func (paymentprocessorstorage *Paymentprocessorstorage) PackTotalInvoiceCreated() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("totalInvoiceCreated")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalInvoiceCreated is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3d1e2ba.
//
// Solidity: function totalInvoiceCreated() view returns(uint216 totalInvoices)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackTotalInvoiceCreated(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("totalInvoiceCreated", data)
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
func (paymentprocessorstorage *Paymentprocessorstorage) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (paymentprocessorstorage *Paymentprocessorstorage) PackUnpause() []byte {
	enc, err := paymentprocessorstorage.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateInvoiceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9e05db13.
//
// Solidity: function updateInvoiceNonce(uint216 _by) returns(uint216 totalInvoices)
func (paymentprocessorstorage *Paymentprocessorstorage) PackUpdateInvoiceNonce(by *big.Int) []byte {
	enc, err := paymentprocessorstorage.abi.Pack("updateInvoiceNonce", by)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUpdateInvoiceNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9e05db13.
//
// Solidity: function updateInvoiceNonce(uint216 _by) returns(uint216 totalInvoices)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackUpdateInvoiceNonce(data []byte) (*big.Int, error) {
	out, err := paymentprocessorstorage.abi.Unpack("updateInvoiceNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PaymentprocessorstorageAuthorizationUpdated represents a AuthorizationUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageAuthorizationUpdated struct {
	Account    common.Address
	Authorized bool
	Raw        *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageAuthorizationUpdatedEventName = "AuthorizationUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageAuthorizationUpdated) ContractEventName() string {
	return PaymentprocessorstorageAuthorizationUpdatedEventName
}

// UnpackAuthorizationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AuthorizationUpdated(address indexed account, bool authorized)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackAuthorizationUpdatedEvent(log *types.Log) (*PaymentprocessorstorageAuthorizationUpdated, error) {
	event := "AuthorizationUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageAuthorizationUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageConfigurationInitialized represents a ConfigurationInitialized event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageConfigurationInitialized struct {
	Config IPaymentProcessorStorageConfiguration
	Raw    *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageConfigurationInitializedEventName = "ConfigurationInitialized"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageConfigurationInitialized) ContractEventName() string {
	return PaymentprocessorstorageConfigurationInitializedEventName
}

// UnpackConfigurationInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ConfigurationInitialized((address,uint96,address,address,uint96) config)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackConfigurationInitializedEvent(log *types.Log) (*PaymentprocessorstorageConfigurationInitialized, error) {
	event := "ConfigurationInitialized"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageConfigurationInitialized)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageEmergencyPauseApproved represents a EmergencyPauseApproved event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageEmergencyPauseApproved struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageEmergencyPauseApprovedEventName = "EmergencyPauseApproved"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageEmergencyPauseApproved) ContractEventName() string {
	return PaymentprocessorstorageEmergencyPauseApprovedEventName
}

// UnpackEmergencyPauseApprovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyPauseApproved(address indexed account)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackEmergencyPauseApprovedEvent(log *types.Log) (*PaymentprocessorstorageEmergencyPauseApproved, error) {
	event := "EmergencyPauseApproved"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageEmergencyPauseApproved)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageEmergencyPaused represents a EmergencyPaused event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageEmergencyPaused struct {
	Account common.Address
	Expiry  *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageEmergencyPausedEventName = "EmergencyPaused"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageEmergencyPaused) ContractEventName() string {
	return PaymentprocessorstorageEmergencyPausedEventName
}

// UnpackEmergencyPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyPaused(address indexed account, uint256 expiry)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackEmergencyPausedEvent(log *types.Log) (*PaymentprocessorstorageEmergencyPaused, error) {
	event := "EmergencyPaused"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageEmergencyPaused)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageEmergencyPauserUpdated represents a EmergencyPauserUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageEmergencyPauserUpdated struct {
	EmergencyPauser common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageEmergencyPauserUpdatedEventName = "EmergencyPauserUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageEmergencyPauserUpdated) ContractEventName() string {
	return PaymentprocessorstorageEmergencyPauserUpdatedEventName
}

// UnpackEmergencyPauserUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyPauserUpdated(address indexed emergencyPauser)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackEmergencyPauserUpdatedEvent(log *types.Log) (*PaymentprocessorstorageEmergencyPauserUpdated, error) {
	event := "EmergencyPauserUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageEmergencyPauserUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageFeeRateUpdated represents a FeeRateUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageFeeRateUpdated struct {
	FeeRate *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageFeeRateUpdatedEventName = "FeeRateUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageFeeRateUpdated) ContractEventName() string {
	return PaymentprocessorstorageFeeRateUpdatedEventName
}

// UnpackFeeRateUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRateUpdated(uint96 feeRate)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackFeeRateUpdatedEvent(log *types.Log) (*PaymentprocessorstorageFeeRateUpdated, error) {
	event := "FeeRateUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageFeeRateUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageFeeReceiverUpdated represents a FeeReceiverUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageFeeReceiverUpdated struct {
	FeeReceiver common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageFeeReceiverUpdatedEventName = "FeeReceiverUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageFeeReceiverUpdated) ContractEventName() string {
	return PaymentprocessorstorageFeeReceiverUpdatedEventName
}

// UnpackFeeReceiverUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackFeeReceiverUpdatedEvent(log *types.Log) (*PaymentprocessorstorageFeeReceiverUpdated, error) {
	event := "FeeReceiverUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageFeeReceiverUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageFeeSignerUpdated represents a FeeSignerUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageFeeSignerUpdated struct {
	FeeSigner common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageFeeSignerUpdatedEventName = "FeeSignerUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageFeeSignerUpdated) ContractEventName() string {
	return PaymentprocessorstorageFeeSignerUpdatedEventName
}

// UnpackFeeSignerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeSignerUpdated(address indexed feeSigner)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackFeeSignerUpdatedEvent(log *types.Log) (*PaymentprocessorstorageFeeSignerUpdated, error) {
	event := "FeeSignerUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageFeeSignerUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageGasThresholdUpdated represents a GasThresholdUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageGasThresholdUpdated struct {
	GasThreshold *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageGasThresholdUpdatedEventName = "GasThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageGasThresholdUpdated) ContractEventName() string {
	return PaymentprocessorstorageGasThresholdUpdatedEventName
}

// UnpackGasThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GasThresholdUpdated(uint96 gasThreshold)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackGasThresholdUpdatedEvent(log *types.Log) (*PaymentprocessorstorageGasThresholdUpdated, error) {
	event := "GasThresholdUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageGasThresholdUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageIntermediatedPlatformsOperatorUpdated represents a IntermediatedPlatformsOperatorUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageIntermediatedPlatformsOperatorUpdated struct {
	IntermediatedPlatformsOperator common.Address
	Raw                            *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageIntermediatedPlatformsOperatorUpdatedEventName = "IntermediatedPlatformsOperatorUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageIntermediatedPlatformsOperatorUpdated) ContractEventName() string {
	return PaymentprocessorstorageIntermediatedPlatformsOperatorUpdatedEventName
}

// UnpackIntermediatedPlatformsOperatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event IntermediatedPlatformsOperatorUpdated(address indexed intermediatedPlatformsOperator)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackIntermediatedPlatformsOperatorUpdatedEvent(log *types.Log) (*PaymentprocessorstorageIntermediatedPlatformsOperatorUpdated, error) {
	event := "IntermediatedPlatformsOperatorUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageIntermediatedPlatformsOperatorUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageOwnershipHandoverCanceledEventName = "OwnershipHandoverCanceled"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageOwnershipHandoverCanceled) ContractEventName() string {
	return PaymentprocessorstorageOwnershipHandoverCanceledEventName
}

// UnpackOwnershipHandoverCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackOwnershipHandoverCanceledEvent(log *types.Log) (*PaymentprocessorstorageOwnershipHandoverCanceled, error) {
	event := "OwnershipHandoverCanceled"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageOwnershipHandoverCanceled)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageOwnershipHandoverRequestedEventName = "OwnershipHandoverRequested"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageOwnershipHandoverRequested) ContractEventName() string {
	return PaymentprocessorstorageOwnershipHandoverRequestedEventName
}

// UnpackOwnershipHandoverRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackOwnershipHandoverRequestedEvent(log *types.Log) (*PaymentprocessorstorageOwnershipHandoverRequested, error) {
	event := "OwnershipHandoverRequested"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageOwnershipHandoverRequested)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageOwnershipTransferred represents a OwnershipTransferred event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageOwnershipTransferred) ContractEventName() string {
	return PaymentprocessorstorageOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackOwnershipTransferredEvent(log *types.Log) (*PaymentprocessorstorageOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstoragePaused represents a Paused event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstoragePaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstoragePausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstoragePaused) ContractEventName() string {
	return PaymentprocessorstoragePausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address indexed account)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackPausedEvent(log *types.Log) (*PaymentprocessorstoragePaused, error) {
	event := "Paused"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstoragePaused)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstoragePaymentValidityDurationUpdated represents a PaymentValidityDurationUpdated event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstoragePaymentValidityDurationUpdated struct {
	ValidityDuration *big.Int
	Raw              *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstoragePaymentValidityDurationUpdatedEventName = "PaymentValidityDurationUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstoragePaymentValidityDurationUpdated) ContractEventName() string {
	return PaymentprocessorstoragePaymentValidityDurationUpdatedEventName
}

// UnpackPaymentValidityDurationUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentValidityDurationUpdated(uint256 validityDuration)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackPaymentValidityDurationUpdatedEvent(log *types.Log) (*PaymentprocessorstoragePaymentValidityDurationUpdated, error) {
	event := "PaymentValidityDurationUpdated"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstoragePaymentValidityDurationUpdated)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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

// PaymentprocessorstorageUnpaused represents a Unpaused event raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PaymentprocessorstorageUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (PaymentprocessorstorageUnpaused) ContractEventName() string {
	return PaymentprocessorstorageUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address indexed account)
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackUnpausedEvent(log *types.Log) (*PaymentprocessorstorageUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != paymentprocessorstorage.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentprocessorstorageUnpaused)
	if len(log.Data) > 0 {
		if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentprocessorstorage.abi.Events[event].Inputs {
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
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["AlreadyInitialized"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["AlreadyPaused"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackAlreadyPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["InvalidFeeRate"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackInvalidFeeRateError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["InvalidFeeSigner"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackInvalidFeeSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["NewOwnerIsZeroAddress"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackNewOwnerIsZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["NoActiveEmergencyPause"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackNoActiveEmergencyPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["NoHandoverRequest"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackNoHandoverRequestError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["NotPaused"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackNotPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentprocessorstorage.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return paymentprocessorstorage.UnpackUnauthorizedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PaymentprocessorstorageAlreadyInitialized represents a AlreadyInitialized error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInitialized()
func PaymentprocessorstorageAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0x0dc149f07762891dbcea3fe72770f3d63a1863fc54b2f084e8c59ec476996927")
}

// UnpackAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInitialized()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackAlreadyInitializedError(raw []byte) (*PaymentprocessorstorageAlreadyInitialized, error) {
	out := new(PaymentprocessorstorageAlreadyInitialized)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "AlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageAlreadyPaused represents a AlreadyPaused error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageAlreadyPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyPaused()
func PaymentprocessorstorageAlreadyPausedErrorID() common.Hash {
	return common.HexToHash("0x1785c68176ff5ca26e02299a48022fe13a267aed4ebbbf517400769c3e8e8df7")
}

// UnpackAlreadyPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyPaused()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackAlreadyPausedError(raw []byte) (*PaymentprocessorstorageAlreadyPaused, error) {
	out := new(PaymentprocessorstorageAlreadyPaused)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "AlreadyPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageInvalidFeeRate represents a InvalidFeeRate error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageInvalidFeeRate struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeRate()
func PaymentprocessorstorageInvalidFeeRateErrorID() common.Hash {
	return common.HexToHash("0x56d69198c50c349b33dac636e06a8847667e835557d137a3943dab95f3d5ce59")
}

// UnpackInvalidFeeRateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeRate()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackInvalidFeeRateError(raw []byte) (*PaymentprocessorstorageInvalidFeeRate, error) {
	out := new(PaymentprocessorstorageInvalidFeeRate)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "InvalidFeeRate", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageInvalidFeeSigner represents a InvalidFeeSigner error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageInvalidFeeSigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeSigner()
func PaymentprocessorstorageInvalidFeeSignerErrorID() common.Hash {
	return common.HexToHash("0x20d80102f2c2e94add93a29fd4abca50beaac174ffcd1d339ade41326ba6323d")
}

// UnpackInvalidFeeSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeSigner()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackInvalidFeeSignerError(raw []byte) (*PaymentprocessorstorageInvalidFeeSigner, error) {
	out := new(PaymentprocessorstorageInvalidFeeSigner)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "InvalidFeeSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageNewOwnerIsZeroAddress represents a NewOwnerIsZeroAddress error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageNewOwnerIsZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NewOwnerIsZeroAddress()
func PaymentprocessorstorageNewOwnerIsZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x7448fbae245b5163a637f61fac94c5376c3e155928452ce47ee52d8c1b99587a")
}

// UnpackNewOwnerIsZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NewOwnerIsZeroAddress()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackNewOwnerIsZeroAddressError(raw []byte) (*PaymentprocessorstorageNewOwnerIsZeroAddress, error) {
	out := new(PaymentprocessorstorageNewOwnerIsZeroAddress)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "NewOwnerIsZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageNoActiveEmergencyPause represents a NoActiveEmergencyPause error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageNoActiveEmergencyPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoActiveEmergencyPause()
func PaymentprocessorstorageNoActiveEmergencyPauseErrorID() common.Hash {
	return common.HexToHash("0xdb469296406d35a0b1b7813ae476c3491d1c8f00379ed4f980870b27cbd368b2")
}

// UnpackNoActiveEmergencyPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoActiveEmergencyPause()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackNoActiveEmergencyPauseError(raw []byte) (*PaymentprocessorstorageNoActiveEmergencyPause, error) {
	out := new(PaymentprocessorstorageNoActiveEmergencyPause)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "NoActiveEmergencyPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageNoHandoverRequest represents a NoHandoverRequest error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageNoHandoverRequest struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoHandoverRequest()
func PaymentprocessorstorageNoHandoverRequestErrorID() common.Hash {
	return common.HexToHash("0x6f5e8818469c73d5be4a0d17c371cde64695907022629c1d064c895f98d466a6")
}

// UnpackNoHandoverRequestError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoHandoverRequest()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackNoHandoverRequestError(raw []byte) (*PaymentprocessorstorageNoHandoverRequest, error) {
	out := new(PaymentprocessorstorageNoHandoverRequest)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "NoHandoverRequest", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageNotAuthorized represents a NotAuthorized error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func PaymentprocessorstorageNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackNotAuthorizedError(raw []byte) (*PaymentprocessorstorageNotAuthorized, error) {
	out := new(PaymentprocessorstorageNotAuthorized)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageNotPaused represents a NotPaused error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageNotPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotPaused()
func PaymentprocessorstorageNotPausedErrorID() common.Hash {
	return common.HexToHash("0x6cd602013233635730773e15e89b8a778034d859147e8f706bcd1aa42e228e06")
}

// UnpackNotPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotPaused()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackNotPausedError(raw []byte) (*PaymentprocessorstorageNotPaused, error) {
	out := new(PaymentprocessorstorageNotPaused)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "NotPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentprocessorstorageUnauthorized represents a Unauthorized error raised by the Paymentprocessorstorage contract.
type PaymentprocessorstorageUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func PaymentprocessorstorageUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (paymentprocessorstorage *Paymentprocessorstorage) UnpackUnauthorizedError(raw []byte) (*PaymentprocessorstorageUnauthorized, error) {
	out := new(PaymentprocessorstorageUnauthorized)
	if err := paymentprocessorstorage.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}
