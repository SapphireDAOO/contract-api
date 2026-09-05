// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisig

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

// IMultiSigTransaction is an auto generated low-level Go binding around an user-defined struct.
type IMultiSigTransaction struct {
	Target        common.Address
	Value         *big.Int
	Data          []byte
	Nonce         *big.Int
	Status        uint8
	ApprovalCount *big.Int
}

// MultisigMetaData contains all meta data concerning the Multisig contract.
var MultisigMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_initialSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_initialThreshold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyASigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyApprovedByThisSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientSigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotASigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCountBelowThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionNotProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionNotProposedOrApproved\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"name\":\"ApprovalAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"TransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"TransactionProposed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"approveTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"cancelTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"getTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"internalType\":\"structIMultiSig.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"hasApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"proposeTransaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Multisig",
}

// Multisig is an auto generated Go binding around an Ethereum contract.
type Multisig struct {
	abi abi.ABI
}

// NewMultisig creates a new instance of Multisig.
func NewMultisig() *Multisig {
	parsed, err := MultisigMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Multisig{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Multisig) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address[] _initialSigners, uint256 _initialThreshold) returns()
func (multisig *Multisig) PackConstructor(_initialSigners []common.Address, _initialThreshold *big.Int) []byte {
	enc, err := multisig.abi.Pack("", _initialSigners, _initialThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAddSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeb12d61e.
//
// Solidity: function addSigner(address _signer) returns()
func (multisig *Multisig) PackAddSigner(signer common.Address) []byte {
	enc, err := multisig.abi.Pack("addSigner", signer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApproveTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f4fde9f.
//
// Solidity: function approveTransaction(bytes32 _txHash) returns()
func (multisig *Multisig) PackApproveTransaction(txHash [32]byte) []byte {
	enc, err := multisig.abi.Pack("approveTransaction", txHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70ba3339.
//
// Solidity: function cancelTransaction(bytes32 _txHash) returns()
func (multisig *Multisig) PackCancelTransaction(txHash [32]byte) []byte {
	enc, err := multisig.abi.Pack("cancelTransaction", txHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecuteTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc69ed5f2.
//
// Solidity: function executeTransaction(bytes32 _txHash) returns(bytes)
func (multisig *Multisig) PackExecuteTransaction(txHash [32]byte) []byte {
	enc, err := multisig.abi.Pack("executeTransaction", txHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecuteTransaction is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc69ed5f2.
//
// Solidity: function executeTransaction(bytes32 _txHash) returns(bytes)
func (multisig *Multisig) UnpackExecuteTransaction(data []byte) ([]byte, error) {
	out, err := multisig.abi.Unpack("executeTransaction", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, err
}

// PackGetNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (multisig *Multisig) PackGetNonce() []byte {
	enc, err := multisig.abi.Pack("getNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (multisig *Multisig) UnpackGetNonce(data []byte) (*big.Int, error) {
	out, err := multisig.abi.Unpack("getNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetSignerCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (multisig *Multisig) PackGetSignerCount() []byte {
	enc, err := multisig.abi.Pack("getSignerCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSignerCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb715be81.
//
// Solidity: function getSignerCount() view returns(uint256)
func (multisig *Multisig) UnpackGetSignerCount(data []byte) (*big.Int, error) {
	out, err := multisig.abi.Unpack("getSignerCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (multisig *Multisig) PackGetThreshold() []byte {
	enc, err := multisig.abi.Pack("getThreshold")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (multisig *Multisig) UnpackGetThreshold(data []byte) (*big.Int, error) {
	out, err := multisig.abi.Unpack("getThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4aae13ca.
//
// Solidity: function getTransaction(bytes32 _txHash) view returns((address,uint256,bytes,uint256,uint8,uint256))
func (multisig *Multisig) PackGetTransaction(txHash [32]byte) []byte {
	enc, err := multisig.abi.Pack("getTransaction", txHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetTransaction is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4aae13ca.
//
// Solidity: function getTransaction(bytes32 _txHash) view returns((address,uint256,bytes,uint256,uint8,uint256))
func (multisig *Multisig) UnpackGetTransaction(data []byte) (IMultiSigTransaction, error) {
	out, err := multisig.abi.Unpack("getTransaction", data)
	if err != nil {
		return *new(IMultiSigTransaction), err
	}
	out0 := *abi.ConvertType(out[0], new(IMultiSigTransaction)).(*IMultiSigTransaction)
	return out0, err
}

// PackHasApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23cdc3f1.
//
// Solidity: function hasApproved(bytes32 _txHash, address _signer) view returns(bool)
func (multisig *Multisig) PackHasApproved(txHash [32]byte, signer common.Address) []byte {
	enc, err := multisig.abi.Pack("hasApproved", txHash, signer)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23cdc3f1.
//
// Solidity: function hasApproved(bytes32 _txHash, address _signer) view returns(bool)
func (multisig *Multisig) UnpackHasApproved(data []byte) (bool, error) {
	out, err := multisig.abi.Unpack("hasApproved", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7df73e27.
//
// Solidity: function isSigner(address _account) view returns(bool)
func (multisig *Multisig) PackIsSigner(account common.Address) []byte {
	enc, err := multisig.abi.Pack("isSigner", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsSigner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7df73e27.
//
// Solidity: function isSigner(address _account) view returns(bool)
func (multisig *Multisig) UnpackIsSigner(data []byte) (bool, error) {
	out, err := multisig.abi.Unpack("isSigner", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackProposeTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e90852d.
//
// Solidity: function proposeTransaction(address _target, uint256 _value, bytes _data) returns(bytes32 txHash)
func (multisig *Multisig) PackProposeTransaction(target common.Address, value *big.Int, data []byte) []byte {
	enc, err := multisig.abi.Pack("proposeTransaction", target, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProposeTransaction is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e90852d.
//
// Solidity: function proposeTransaction(address _target, uint256 _value, bytes _data) returns(bytes32 txHash)
func (multisig *Multisig) UnpackProposeTransaction(data []byte) ([32]byte, error) {
	out, err := multisig.abi.Unpack("proposeTransaction", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRemoveSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e316ab7.
//
// Solidity: function removeSigner(address _signer) returns()
func (multisig *Multisig) PackRemoveSigner(signer common.Address) []byte {
	enc, err := multisig.abi.Pack("removeSigner", signer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7d7442f.
//
// Solidity: function updateThreshold(uint256 _newThreshold) returns()
func (multisig *Multisig) PackUpdateThreshold(newThreshold *big.Int) []byte {
	enc, err := multisig.abi.Pack("updateThreshold", newThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// MultisigApprovalAdded represents a ApprovalAdded event raised by the Multisig contract.
type MultisigApprovalAdded struct {
	TxHash        [32]byte
	Approver      common.Address
	ApprovalCount *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const MultisigApprovalAddedEventName = "ApprovalAdded"

// ContractEventName returns the user-defined event name.
func (MultisigApprovalAdded) ContractEventName() string {
	return MultisigApprovalAddedEventName
}

// UnpackApprovalAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalAdded(bytes32 indexed txHash, address indexed approver, uint256 approvalCount)
func (multisig *Multisig) UnpackApprovalAddedEvent(log *types.Log) (*MultisigApprovalAdded, error) {
	event := "ApprovalAdded"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigApprovalAdded)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigSignerAdded represents a SignerAdded event raised by the Multisig contract.
type MultisigSignerAdded struct {
	Signer common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const MultisigSignerAddedEventName = "SignerAdded"

// ContractEventName returns the user-defined event name.
func (MultisigSignerAdded) ContractEventName() string {
	return MultisigSignerAddedEventName
}

// UnpackSignerAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignerAdded(address indexed signer)
func (multisig *Multisig) UnpackSignerAddedEvent(log *types.Log) (*MultisigSignerAdded, error) {
	event := "SignerAdded"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigSignerAdded)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigSignerRemoved represents a SignerRemoved event raised by the Multisig contract.
type MultisigSignerRemoved struct {
	Signer common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const MultisigSignerRemovedEventName = "SignerRemoved"

// ContractEventName returns the user-defined event name.
func (MultisigSignerRemoved) ContractEventName() string {
	return MultisigSignerRemovedEventName
}

// UnpackSignerRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignerRemoved(address indexed signer)
func (multisig *Multisig) UnpackSignerRemovedEvent(log *types.Log) (*MultisigSignerRemoved, error) {
	event := "SignerRemoved"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigSignerRemoved)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigThresholdUpdated represents a ThresholdUpdated event raised by the Multisig contract.
type MultisigThresholdUpdated struct {
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const MultisigThresholdUpdatedEventName = "ThresholdUpdated"

// ContractEventName returns the user-defined event name.
func (MultisigThresholdUpdated) ContractEventName() string {
	return MultisigThresholdUpdatedEventName
}

// UnpackThresholdUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ThresholdUpdated(uint256 oldThreshold, uint256 newThreshold)
func (multisig *Multisig) UnpackThresholdUpdatedEvent(log *types.Log) (*MultisigThresholdUpdated, error) {
	event := "ThresholdUpdated"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigThresholdUpdated)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigTransactionApproved represents a TransactionApproved event raised by the Multisig contract.
type MultisigTransactionApproved struct {
	TxHash [32]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const MultisigTransactionApprovedEventName = "TransactionApproved"

// ContractEventName returns the user-defined event name.
func (MultisigTransactionApproved) ContractEventName() string {
	return MultisigTransactionApprovedEventName
}

// UnpackTransactionApprovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransactionApproved(bytes32 indexed txHash)
func (multisig *Multisig) UnpackTransactionApprovedEvent(log *types.Log) (*MultisigTransactionApproved, error) {
	event := "TransactionApproved"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigTransactionApproved)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigTransactionCanceled represents a TransactionCanceled event raised by the Multisig contract.
type MultisigTransactionCanceled struct {
	TxHash [32]byte
	Raw    *types.Log // Blockchain specific contextual infos
}

const MultisigTransactionCanceledEventName = "TransactionCanceled"

// ContractEventName returns the user-defined event name.
func (MultisigTransactionCanceled) ContractEventName() string {
	return MultisigTransactionCanceledEventName
}

// UnpackTransactionCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransactionCanceled(bytes32 indexed txHash)
func (multisig *Multisig) UnpackTransactionCanceledEvent(log *types.Log) (*MultisigTransactionCanceled, error) {
	event := "TransactionCanceled"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigTransactionCanceled)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigTransactionExecuted represents a TransactionExecuted event raised by the Multisig contract.
type MultisigTransactionExecuted struct {
	TxHash   [32]byte
	Executor common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const MultisigTransactionExecutedEventName = "TransactionExecuted"

// ContractEventName returns the user-defined event name.
func (MultisigTransactionExecuted) ContractEventName() string {
	return MultisigTransactionExecutedEventName
}

// UnpackTransactionExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransactionExecuted(bytes32 indexed txHash, address indexed executor)
func (multisig *Multisig) UnpackTransactionExecutedEvent(log *types.Log) (*MultisigTransactionExecuted, error) {
	event := "TransactionExecuted"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigTransactionExecuted)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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

// MultisigTransactionProposed represents a TransactionProposed event raised by the Multisig contract.
type MultisigTransactionProposed struct {
	TxHash   [32]byte
	Target   common.Address
	Value    *big.Int
	Data     []byte
	Nonce    *big.Int
	Proposer common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const MultisigTransactionProposedEventName = "TransactionProposed"

// ContractEventName returns the user-defined event name.
func (MultisigTransactionProposed) ContractEventName() string {
	return MultisigTransactionProposedEventName
}

// UnpackTransactionProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransactionProposed(bytes32 indexed txHash, address indexed target, uint256 value, bytes data, uint256 nonce, address indexed proposer)
func (multisig *Multisig) UnpackTransactionProposedEvent(log *types.Log) (*MultisigTransactionProposed, error) {
	event := "TransactionProposed"
	if log.Topics[0] != multisig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MultisigTransactionProposed)
	if len(log.Data) > 0 {
		if err := multisig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range multisig.abi.Events[event].Inputs {
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
func (multisig *Multisig) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], multisig.abi.Errors["AlreadyASigner"].ID.Bytes()[:4]) {
		return multisig.UnpackAlreadyASignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["AlreadyApproved"].ID.Bytes()[:4]) {
		return multisig.UnpackAlreadyApprovedError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["AlreadyApprovedByThisSigner"].ID.Bytes()[:4]) {
		return multisig.UnpackAlreadyApprovedByThisSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["AlreadyExecuted"].ID.Bytes()[:4]) {
		return multisig.UnpackAlreadyExecutedError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["ExecutionFailed"].ID.Bytes()[:4]) {
		return multisig.UnpackExecutionFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["InsufficientSigners"].ID.Bytes()[:4]) {
		return multisig.UnpackInsufficientSignersError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["InvalidTarget"].ID.Bytes()[:4]) {
		return multisig.UnpackInvalidTargetError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return multisig.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["NotASigner"].ID.Bytes()[:4]) {
		return multisig.UnpackNotASignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["NotSelf"].ID.Bytes()[:4]) {
		return multisig.UnpackNotSelfError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["NotSigner"].ID.Bytes()[:4]) {
		return multisig.UnpackNotSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["SignerCountBelowThreshold"].ID.Bytes()[:4]) {
		return multisig.UnpackSignerCountBelowThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["ThresholdCannotBeZero"].ID.Bytes()[:4]) {
		return multisig.UnpackThresholdCannotBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["TransactionDoesNotExist"].ID.Bytes()[:4]) {
		return multisig.UnpackTransactionDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["TransactionNotApproved"].ID.Bytes()[:4]) {
		return multisig.UnpackTransactionNotApprovedError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["TransactionNotProposed"].ID.Bytes()[:4]) {
		return multisig.UnpackTransactionNotProposedError(raw[4:])
	}
	if bytes.Equal(raw[:4], multisig.abi.Errors["TransactionNotProposedOrApproved"].ID.Bytes()[:4]) {
		return multisig.UnpackTransactionNotProposedOrApprovedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MultisigAlreadyASigner represents a AlreadyASigner error raised by the Multisig contract.
type MultisigAlreadyASigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyASigner()
func MultisigAlreadyASignerErrorID() common.Hash {
	return common.HexToHash("0xcc2fa5583e72317ef7c4a6ba5bc5bb0edf6adc1365c19fcedd0b225f2c8f40e0")
}

// UnpackAlreadyASignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyASigner()
func (multisig *Multisig) UnpackAlreadyASignerError(raw []byte) (*MultisigAlreadyASigner, error) {
	out := new(MultisigAlreadyASigner)
	if err := multisig.abi.UnpackIntoInterface(out, "AlreadyASigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigAlreadyApproved represents a AlreadyApproved error raised by the Multisig contract.
type MultisigAlreadyApproved struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyApproved()
func MultisigAlreadyApprovedErrorID() common.Hash {
	return common.HexToHash("0x101f817a2457b2fcd2728e709f840e8180b3b74a761c9cbd09a4fa605a43f8b1")
}

// UnpackAlreadyApprovedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyApproved()
func (multisig *Multisig) UnpackAlreadyApprovedError(raw []byte) (*MultisigAlreadyApproved, error) {
	out := new(MultisigAlreadyApproved)
	if err := multisig.abi.UnpackIntoInterface(out, "AlreadyApproved", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigAlreadyApprovedByThisSigner represents a AlreadyApprovedByThisSigner error raised by the Multisig contract.
type MultisigAlreadyApprovedByThisSigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyApprovedByThisSigner()
func MultisigAlreadyApprovedByThisSignerErrorID() common.Hash {
	return common.HexToHash("0xefa11217d551b72589899fc9c19d00a363b17848df07ce0ef640eb3510071f4e")
}

// UnpackAlreadyApprovedByThisSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyApprovedByThisSigner()
func (multisig *Multisig) UnpackAlreadyApprovedByThisSignerError(raw []byte) (*MultisigAlreadyApprovedByThisSigner, error) {
	out := new(MultisigAlreadyApprovedByThisSigner)
	if err := multisig.abi.UnpackIntoInterface(out, "AlreadyApprovedByThisSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigAlreadyExecuted represents a AlreadyExecuted error raised by the Multisig contract.
type MultisigAlreadyExecuted struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyExecuted()
func MultisigAlreadyExecutedErrorID() common.Hash {
	return common.HexToHash("0x0dc1019765b80bb1e4e23ed4c31a7dc3f145102e4eff519fd013f6b295eb56f7")
}

// UnpackAlreadyExecutedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyExecuted()
func (multisig *Multisig) UnpackAlreadyExecutedError(raw []byte) (*MultisigAlreadyExecuted, error) {
	out := new(MultisigAlreadyExecuted)
	if err := multisig.abi.UnpackIntoInterface(out, "AlreadyExecuted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigExecutionFailed represents a ExecutionFailed error raised by the Multisig contract.
type MultisigExecutionFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExecutionFailed()
func MultisigExecutionFailedErrorID() common.Hash {
	return common.HexToHash("0xacfdb444727b3b8994850a379f4bfc8a5ca665a55604339199daafa16f687b1a")
}

// UnpackExecutionFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExecutionFailed()
func (multisig *Multisig) UnpackExecutionFailedError(raw []byte) (*MultisigExecutionFailed, error) {
	out := new(MultisigExecutionFailed)
	if err := multisig.abi.UnpackIntoInterface(out, "ExecutionFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigInsufficientSigners represents a InsufficientSigners error raised by the Multisig contract.
type MultisigInsufficientSigners struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientSigners()
func MultisigInsufficientSignersErrorID() common.Hash {
	return common.HexToHash("0xc2ee9b9e218d2e3c08c9c6703b860a5f354eac53450dfd1536ad5aad8ce9993a")
}

// UnpackInsufficientSignersError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientSigners()
func (multisig *Multisig) UnpackInsufficientSignersError(raw []byte) (*MultisigInsufficientSigners, error) {
	out := new(MultisigInsufficientSigners)
	if err := multisig.abi.UnpackIntoInterface(out, "InsufficientSigners", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigInvalidTarget represents a InvalidTarget error raised by the Multisig contract.
type MultisigInvalidTarget struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTarget()
func MultisigInvalidTargetErrorID() common.Hash {
	return common.HexToHash("0x82d5d76a554b1226e0bd53ae5121b9d30262dd775be75ba7d4bd87ba04ede349")
}

// UnpackInvalidTargetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTarget()
func (multisig *Multisig) UnpackInvalidTargetError(raw []byte) (*MultisigInvalidTarget, error) {
	out := new(MultisigInvalidTarget)
	if err := multisig.abi.UnpackIntoInterface(out, "InvalidTarget", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigInvalidThreshold represents a InvalidThreshold error raised by the Multisig contract.
type MultisigInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func MultisigInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (multisig *Multisig) UnpackInvalidThresholdError(raw []byte) (*MultisigInvalidThreshold, error) {
	out := new(MultisigInvalidThreshold)
	if err := multisig.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigNotASigner represents a NotASigner error raised by the Multisig contract.
type MultisigNotASigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotASigner()
func MultisigNotASignerErrorID() common.Hash {
	return common.HexToHash("0xda0357f7ed76d95589df52d07b03ab0b08a3014223be7378b19b3d64a3bd119a")
}

// UnpackNotASignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotASigner()
func (multisig *Multisig) UnpackNotASignerError(raw []byte) (*MultisigNotASigner, error) {
	out := new(MultisigNotASigner)
	if err := multisig.abi.UnpackIntoInterface(out, "NotASigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigNotSelf represents a NotSelf error raised by the Multisig contract.
type MultisigNotSelf struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotSelf()
func MultisigNotSelfErrorID() common.Hash {
	return common.HexToHash("0x29c3b7ee14a7adaa7fba7b043229509fe963f719303620e91e88c61c6b12de34")
}

// UnpackNotSelfError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotSelf()
func (multisig *Multisig) UnpackNotSelfError(raw []byte) (*MultisigNotSelf, error) {
	out := new(MultisigNotSelf)
	if err := multisig.abi.UnpackIntoInterface(out, "NotSelf", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigNotSigner represents a NotSigner error raised by the Multisig contract.
type MultisigNotSigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotSigner()
func MultisigNotSignerErrorID() common.Hash {
	return common.HexToHash("0xa1b035c8a93989f4232e1009d87a1b0b0a298de6594a454477adf8be46e7d892")
}

// UnpackNotSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotSigner()
func (multisig *Multisig) UnpackNotSignerError(raw []byte) (*MultisigNotSigner, error) {
	out := new(MultisigNotSigner)
	if err := multisig.abi.UnpackIntoInterface(out, "NotSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigSignerCountBelowThreshold represents a SignerCountBelowThreshold error raised by the Multisig contract.
type MultisigSignerCountBelowThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignerCountBelowThreshold()
func MultisigSignerCountBelowThresholdErrorID() common.Hash {
	return common.HexToHash("0x082d4dd3b83f15156f44465749f140dfd106577c2af143199a2286ebb4af9a22")
}

// UnpackSignerCountBelowThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignerCountBelowThreshold()
func (multisig *Multisig) UnpackSignerCountBelowThresholdError(raw []byte) (*MultisigSignerCountBelowThreshold, error) {
	out := new(MultisigSignerCountBelowThreshold)
	if err := multisig.abi.UnpackIntoInterface(out, "SignerCountBelowThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigThresholdCannotBeZero represents a ThresholdCannotBeZero error raised by the Multisig contract.
type MultisigThresholdCannotBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ThresholdCannotBeZero()
func MultisigThresholdCannotBeZeroErrorID() common.Hash {
	return common.HexToHash("0xf4124166dd0e538e9ca98fea6e3af8b26855694df19f8dec41711d8e570537d8")
}

// UnpackThresholdCannotBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ThresholdCannotBeZero()
func (multisig *Multisig) UnpackThresholdCannotBeZeroError(raw []byte) (*MultisigThresholdCannotBeZero, error) {
	out := new(MultisigThresholdCannotBeZero)
	if err := multisig.abi.UnpackIntoInterface(out, "ThresholdCannotBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigTransactionDoesNotExist represents a TransactionDoesNotExist error raised by the Multisig contract.
type MultisigTransactionDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransactionDoesNotExist()
func MultisigTransactionDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x5451f114f9c68ef543bbeab205a95eb517ee3f4723b8fed1a7e8075061911e93")
}

// UnpackTransactionDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransactionDoesNotExist()
func (multisig *Multisig) UnpackTransactionDoesNotExistError(raw []byte) (*MultisigTransactionDoesNotExist, error) {
	out := new(MultisigTransactionDoesNotExist)
	if err := multisig.abi.UnpackIntoInterface(out, "TransactionDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigTransactionNotApproved represents a TransactionNotApproved error raised by the Multisig contract.
type MultisigTransactionNotApproved struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransactionNotApproved()
func MultisigTransactionNotApprovedErrorID() common.Hash {
	return common.HexToHash("0x34c68f2b52127ac33f3bf86f0e2854df6425fef5e137c35df0935acb4f8926d0")
}

// UnpackTransactionNotApprovedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransactionNotApproved()
func (multisig *Multisig) UnpackTransactionNotApprovedError(raw []byte) (*MultisigTransactionNotApproved, error) {
	out := new(MultisigTransactionNotApproved)
	if err := multisig.abi.UnpackIntoInterface(out, "TransactionNotApproved", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigTransactionNotProposed represents a TransactionNotProposed error raised by the Multisig contract.
type MultisigTransactionNotProposed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransactionNotProposed()
func MultisigTransactionNotProposedErrorID() common.Hash {
	return common.HexToHash("0xd5a4fdee805422e65507929ed46b5ae710306106913dc29b1ee14e83324a6083")
}

// UnpackTransactionNotProposedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransactionNotProposed()
func (multisig *Multisig) UnpackTransactionNotProposedError(raw []byte) (*MultisigTransactionNotProposed, error) {
	out := new(MultisigTransactionNotProposed)
	if err := multisig.abi.UnpackIntoInterface(out, "TransactionNotProposed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MultisigTransactionNotProposedOrApproved represents a TransactionNotProposedOrApproved error raised by the Multisig contract.
type MultisigTransactionNotProposedOrApproved struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransactionNotProposedOrApproved()
func MultisigTransactionNotProposedOrApprovedErrorID() common.Hash {
	return common.HexToHash("0x76bf04043a21551e2f774864203cc113a7d321db06edadf7fcd307a8a1699102")
}

// UnpackTransactionNotProposedOrApprovedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransactionNotProposedOrApproved()
func (multisig *Multisig) UnpackTransactionNotProposedOrApprovedError(raw []byte) (*MultisigTransactionNotProposedOrApproved, error) {
	out := new(MultisigTransactionNotProposedOrApproved)
	if err := multisig.abi.UnpackIntoInterface(out, "TransactionNotProposedOrApproved", raw); err != nil {
		return nil, err
	}
	return out, nil
}
