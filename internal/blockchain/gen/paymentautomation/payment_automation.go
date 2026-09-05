// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymentautomation

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

// PaymentautomationMetaData contains all meta data concerning the Paymentautomation contract.
var PaymentautomationMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_processorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checker\",\"inputs\":[],\"outputs\":[{\"name\":\"canExec\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"execPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getForwarder\",\"inputs\":[],\"outputs\":[{\"name\":\"forwarderAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWorkflowOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"workflowOwnerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasDueTasks\",\"inputs\":[],\"outputs\":[{\"name\":\"dueTasksExist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"_metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processDueTasks\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISimplePaymentProcessor\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setForwarderAddress\",\"inputs\":[{\"name\":\"_forwarderAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWorkflowOwner\",\"inputs\":[{\"name\":\"_workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"supported\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"DueTasksProcessed\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"source\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwarderUpdated\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WorkflowOwnerUpdated\",\"inputs\":[{\"name\":\"workflowOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedWorkflowOwner\",\"inputs\":[{\"name\":\"_workflowOwner\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	ID:  "Paymentautomation",
}

// Paymentautomation is an auto generated Go binding around an Ethereum contract.
type Paymentautomation struct {
	abi abi.ABI
}

// NewPaymentautomation creates a new instance of Paymentautomation.
func NewPaymentautomation() *Paymentautomation {
	parsed, err := PaymentautomationMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Paymentautomation{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Paymentautomation) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _processorAddress, address _paymentProcessorStorageAddress) returns()
func (paymentautomation *Paymentautomation) PackConstructor(_processorAddress common.Address, _paymentProcessorStorageAddress common.Address) []byte {
	enc, err := paymentautomation.abi.Pack("", _processorAddress, _paymentProcessorStorageAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChecker is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf5303cf.
//
// Solidity: function checker() view returns(bool canExec, bytes execPayload)
func (paymentautomation *Paymentautomation) PackChecker() []byte {
	enc, err := paymentautomation.abi.Pack("checker")
	if err != nil {
		panic(err)
	}
	return enc
}

// CheckerOutput serves as a container for the return parameters of contract
// method Checker.
type CheckerOutput struct {
	CanExec     bool
	ExecPayload []byte
}

// UnpackChecker is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf5303cf.
//
// Solidity: function checker() view returns(bool canExec, bytes execPayload)
func (paymentautomation *Paymentautomation) UnpackChecker(data []byte) (CheckerOutput, error) {
	out, err := paymentautomation.abi.Unpack("checker", data)
	outstruct := new(CheckerOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.CanExec = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ExecPayload = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, err

}

// PackGetForwarder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa0042526.
//
// Solidity: function getForwarder() view returns(address forwarderAddress)
func (paymentautomation *Paymentautomation) PackGetForwarder() []byte {
	enc, err := paymentautomation.abi.Pack("getForwarder")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetForwarder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa0042526.
//
// Solidity: function getForwarder() view returns(address forwarderAddress)
func (paymentautomation *Paymentautomation) UnpackGetForwarder(data []byte) (common.Address, error) {
	out, err := paymentautomation.abi.Unpack("getForwarder", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetWorkflowOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55139581.
//
// Solidity: function getWorkflowOwner() view returns(address workflowOwnerAddress)
func (paymentautomation *Paymentautomation) PackGetWorkflowOwner() []byte {
	enc, err := paymentautomation.abi.Pack("getWorkflowOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetWorkflowOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x55139581.
//
// Solidity: function getWorkflowOwner() view returns(address workflowOwnerAddress)
func (paymentautomation *Paymentautomation) UnpackGetWorkflowOwner(data []byte) (common.Address, error) {
	out, err := paymentautomation.abi.Unpack("getWorkflowOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackHasDueTasks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a9a0dea.
//
// Solidity: function hasDueTasks() view returns(bool dueTasksExist)
func (paymentautomation *Paymentautomation) PackHasDueTasks() []byte {
	enc, err := paymentautomation.abi.Pack("hasDueTasks")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasDueTasks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0a9a0dea.
//
// Solidity: function hasDueTasks() view returns(bool dueTasksExist)
func (paymentautomation *Paymentautomation) UnpackHasDueTasks(data []byte) (bool, error) {
	out, err := paymentautomation.abi.Unpack("hasDueTasks", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOnReport is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x805f2132.
//
// Solidity: function onReport(bytes _metadata, bytes ) returns()
func (paymentautomation *Paymentautomation) PackOnReport(metadata []byte, arg1 []byte) []byte {
	enc, err := paymentautomation.abi.Pack("onReport", metadata, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (paymentautomation *Paymentautomation) PackPpStorage() []byte {
	enc, err := paymentautomation.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (paymentautomation *Paymentautomation) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := paymentautomation.abi.Unpack("ppStorage", data)
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
func (paymentautomation *Paymentautomation) PackProcessDueTasks() []byte {
	enc, err := paymentautomation.abi.Pack("processDueTasks")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProcessor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce1b1d43.
//
// Solidity: function processor() view returns(address)
func (paymentautomation *Paymentautomation) PackProcessor() []byte {
	enc, err := paymentautomation.abi.Pack("processor")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProcessor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce1b1d43.
//
// Solidity: function processor() view returns(address)
func (paymentautomation *Paymentautomation) UnpackProcessor(data []byte) (common.Address, error) {
	out, err := paymentautomation.abi.Unpack("processor", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSetForwarderAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd777cc6d.
//
// Solidity: function setForwarderAddress(address _forwarderAddress) returns()
func (paymentautomation *Paymentautomation) PackSetForwarderAddress(forwarderAddress common.Address) []byte {
	enc, err := paymentautomation.abi.Pack("setForwarderAddress", forwarderAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetWorkflowOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5ad1c8a.
//
// Solidity: function setWorkflowOwner(address _workflowOwner) returns()
func (paymentautomation *Paymentautomation) PackSetWorkflowOwner(workflowOwner common.Address) []byte {
	enc, err := paymentautomation.abi.Pack("setWorkflowOwner", workflowOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool supported)
func (paymentautomation *Paymentautomation) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := paymentautomation.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool supported)
func (paymentautomation *Paymentautomation) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := paymentautomation.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PaymentautomationDueTasksProcessed represents a DueTasksProcessed event raised by the Paymentautomation contract.
type PaymentautomationDueTasksProcessed struct {
	Caller common.Address
	Source [32]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const PaymentautomationDueTasksProcessedEventName = "DueTasksProcessed"

// ContractEventName returns the user-defined event name.
func (PaymentautomationDueTasksProcessed) ContractEventName() string {
	return PaymentautomationDueTasksProcessedEventName
}

// UnpackDueTasksProcessedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DueTasksProcessed(address indexed caller, bytes32 indexed source)
func (paymentautomation *Paymentautomation) UnpackDueTasksProcessedEvent(log *types.Log) (*PaymentautomationDueTasksProcessed, error) {
	event := "DueTasksProcessed"
	if log.Topics[0] != paymentautomation.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentautomationDueTasksProcessed)
	if len(log.Data) > 0 {
		if err := paymentautomation.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentautomation.abi.Events[event].Inputs {
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

// PaymentautomationForwarderUpdated represents a ForwarderUpdated event raised by the Paymentautomation contract.
type PaymentautomationForwarderUpdated struct {
	Forwarder common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const PaymentautomationForwarderUpdatedEventName = "ForwarderUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentautomationForwarderUpdated) ContractEventName() string {
	return PaymentautomationForwarderUpdatedEventName
}

// UnpackForwarderUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForwarderUpdated(address indexed forwarder)
func (paymentautomation *Paymentautomation) UnpackForwarderUpdatedEvent(log *types.Log) (*PaymentautomationForwarderUpdated, error) {
	event := "ForwarderUpdated"
	if log.Topics[0] != paymentautomation.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentautomationForwarderUpdated)
	if len(log.Data) > 0 {
		if err := paymentautomation.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentautomation.abi.Events[event].Inputs {
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

// PaymentautomationWorkflowOwnerUpdated represents a WorkflowOwnerUpdated event raised by the Paymentautomation contract.
type PaymentautomationWorkflowOwnerUpdated struct {
	WorkflowOwner common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const PaymentautomationWorkflowOwnerUpdatedEventName = "WorkflowOwnerUpdated"

// ContractEventName returns the user-defined event name.
func (PaymentautomationWorkflowOwnerUpdated) ContractEventName() string {
	return PaymentautomationWorkflowOwnerUpdatedEventName
}

// UnpackWorkflowOwnerUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WorkflowOwnerUpdated(address indexed workflowOwner)
func (paymentautomation *Paymentautomation) UnpackWorkflowOwnerUpdatedEvent(log *types.Log) (*PaymentautomationWorkflowOwnerUpdated, error) {
	event := "WorkflowOwnerUpdated"
	if log.Topics[0] != paymentautomation.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PaymentautomationWorkflowOwnerUpdated)
	if len(log.Data) > 0 {
		if err := paymentautomation.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range paymentautomation.abi.Events[event].Inputs {
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
func (paymentautomation *Paymentautomation) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], paymentautomation.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return paymentautomation.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentautomation.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return paymentautomation.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], paymentautomation.abi.Errors["UnauthorizedWorkflowOwner"].ID.Bytes()[:4]) {
		return paymentautomation.UnpackUnauthorizedWorkflowOwnerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PaymentautomationInvalidAddress represents a InvalidAddress error raised by the Paymentautomation contract.
type PaymentautomationInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func PaymentautomationInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (paymentautomation *Paymentautomation) UnpackInvalidAddressError(raw []byte) (*PaymentautomationInvalidAddress, error) {
	out := new(PaymentautomationInvalidAddress)
	if err := paymentautomation.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentautomationNotAuthorized represents a NotAuthorized error raised by the Paymentautomation contract.
type PaymentautomationNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func PaymentautomationNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (paymentautomation *Paymentautomation) UnpackNotAuthorizedError(raw []byte) (*PaymentautomationNotAuthorized, error) {
	out := new(PaymentautomationNotAuthorized)
	if err := paymentautomation.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentautomationUnauthorizedWorkflowOwner represents a UnauthorizedWorkflowOwner error raised by the Paymentautomation contract.
type PaymentautomationUnauthorizedWorkflowOwner struct {
	WorkflowOwner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnauthorizedWorkflowOwner(address _workflowOwner)
func PaymentautomationUnauthorizedWorkflowOwnerErrorID() common.Hash {
	return common.HexToHash("0xbf24162371c5f6f38e8ff64511dfb5fad309cc95c66786a301dcdfe4de62b386")
}

// UnpackUnauthorizedWorkflowOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnauthorizedWorkflowOwner(address _workflowOwner)
func (paymentautomation *Paymentautomation) UnpackUnauthorizedWorkflowOwnerError(raw []byte) (*PaymentautomationUnauthorizedWorkflowOwner, error) {
	out := new(PaymentautomationUnauthorizedWorkflowOwner)
	if err := paymentautomation.abi.UnpackIntoInterface(out, "UnauthorizedWorkflowOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}
