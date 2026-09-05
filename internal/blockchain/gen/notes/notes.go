// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package notes

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

// NotesMetaData contains all meta data concerning the Notes contract.
var NotesMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyContent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"invoiceId\",\"type\":\"uint216\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"share\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedContent\",\"type\":\"bytes\"}],\"name\":\"NoteCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint216\",\"name\":\"invoiceId\",\"type\":\"uint216\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"opened\",\"type\":\"bool\"}],\"name\":\"NoteStateChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ALLOWED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"_invoiceId\",\"type\":\"uint216\"},{\"internalType\":\"address\",\"name\":\"_author\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_encryptedContent\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_share\",\"type\":\"bool\"}],\"name\":\"createNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"_invoiceId\",\"type\":\"uint216\"},{\"internalType\":\"uint256\",\"name\":\"_noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"share\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"openedStatus\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"_invoiceId\",\"type\":\"uint216\"}],\"name\":\"getNoteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalNotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"_invoiceId\",\"type\":\"uint216\"},{\"internalType\":\"uint256\",\"name\":\"_noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ppStorage\",\"outputs\":[{\"internalType\":\"contractIPaymentProcessorStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setAuthorized\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint216\",\"name\":\"_invoiceId\",\"type\":\"uint216\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_noteId\",\"type\":\"uint256\"}],\"name\":\"setOpened\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_newVersion\",\"type\":\"uint8\"}],\"name\":\"updateVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Notes",
}

// Notes is an auto generated Go binding around an Ethereum contract.
type Notes struct {
	abi abi.ABI
}

// NewNotes creates a new instance of Notes.
func NewNotes() *Notes {
	parsed, err := NotesMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Notes{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Notes) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _paymentProcessorStorageAddress) returns()
func (notes *Notes) PackConstructor(_paymentProcessorStorageAddress common.Address) []byte {
	enc, err := notes.abi.Pack("", _paymentProcessorStorageAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackALLOWED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5f768bf.
//
// Solidity: function ALLOWED() view returns(uint256)
func (notes *Notes) PackALLOWED() []byte {
	enc, err := notes.abi.Pack("ALLOWED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackALLOWED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5f768bf.
//
// Solidity: function ALLOWED() view returns(uint256)
func (notes *Notes) UnpackALLOWED(data []byte) (*big.Int, error) {
	out, err := notes.abi.Unpack("ALLOWED", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackNOTALLOWED is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95c31a56.
//
// Solidity: function NOT_ALLOWED() view returns(uint256)
func (notes *Notes) PackNOTALLOWED() []byte {
	enc, err := notes.abi.Pack("NOT_ALLOWED")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNOTALLOWED is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95c31a56.
//
// Solidity: function NOT_ALLOWED() view returns(uint256)
func (notes *Notes) UnpackNOTALLOWED(data []byte) (*big.Int, error) {
	out, err := notes.abi.Unpack("NOT_ALLOWED", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackCreateNote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdee3cb23.
//
// Solidity: function createNote(uint216 _invoiceId, address _author, bytes _encryptedContent, bool _share) returns(uint256 noteId)
func (notes *Notes) PackCreateNote(invoiceId *big.Int, author common.Address, encryptedContent []byte, share bool) []byte {
	enc, err := notes.abi.Pack("createNote", invoiceId, author, encryptedContent, share)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCreateNote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdee3cb23.
//
// Solidity: function createNote(uint216 _invoiceId, address _author, bytes _encryptedContent, bool _share) returns(uint256 noteId)
func (notes *Notes) UnpackCreateNote(data []byte) (*big.Int, error) {
	out, err := notes.abi.Unpack("createNote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetCurrentVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfabec44a.
//
// Solidity: function getCurrentVersion() view returns(uint8 v)
func (notes *Notes) PackGetCurrentVersion() []byte {
	enc, err := notes.abi.Pack("getCurrentVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetCurrentVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfabec44a.
//
// Solidity: function getCurrentVersion() view returns(uint8 v)
func (notes *Notes) UnpackGetCurrentVersion(data []byte) (uint8, error) {
	out, err := notes.abi.Unpack("getCurrentVersion", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackGetNote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe735995.
//
// Solidity: function getNote(uint216 _invoiceId, uint256 _noteId) view returns(address author, bool share, bytes content, bool openedStatus, uint8 version)
func (notes *Notes) PackGetNote(invoiceId *big.Int, noteId *big.Int) []byte {
	enc, err := notes.abi.Pack("getNote", invoiceId, noteId)
	if err != nil {
		panic(err)
	}
	return enc
}

// GetNoteOutput serves as a container for the return parameters of contract
// method GetNote.
type GetNoteOutput struct {
	Author       common.Address
	Share        bool
	Content      []byte
	OpenedStatus bool
	Version      uint8
}

// UnpackGetNote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfe735995.
//
// Solidity: function getNote(uint216 _invoiceId, uint256 _noteId) view returns(address author, bool share, bytes content, bool openedStatus, uint8 version)
func (notes *Notes) UnpackGetNote(data []byte) (GetNoteOutput, error) {
	out, err := notes.abi.Unpack("getNote", data)
	outstruct := new(GetNoteOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Author = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Share = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Content = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.OpenedStatus = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Version = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	return *outstruct, err

}

// PackGetNoteCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b0c22b1.
//
// Solidity: function getNoteCount(uint216 _invoiceId) view returns(uint256 totalNotes)
func (notes *Notes) PackGetNoteCount(invoiceId *big.Int) []byte {
	enc, err := notes.abi.Pack("getNoteCount", invoiceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetNoteCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1b0c22b1.
//
// Solidity: function getNoteCount(uint216 _invoiceId) view returns(uint256 totalNotes)
func (notes *Notes) UnpackGetNoteCount(data []byte) (*big.Int, error) {
	out, err := notes.abi.Unpack("getNoteCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackIsOpened is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ec0a68d.
//
// Solidity: function isOpened(uint216 _invoiceId, uint256 _noteId, address _user) view returns(bool isOpen)
func (notes *Notes) PackIsOpened(invoiceId *big.Int, noteId *big.Int, user common.Address) []byte {
	enc, err := notes.abi.Pack("isOpened", invoiceId, noteId, user)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsOpened is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3ec0a68d.
//
// Solidity: function isOpened(uint216 _invoiceId, uint256 _noteId, address _user) view returns(bool isOpen)
func (notes *Notes) UnpackIsOpened(data []byte) (bool, error) {
	out, err := notes.abi.Unpack("isOpened", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackPpStorage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (notes *Notes) PackPpStorage() []byte {
	enc, err := notes.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (notes *Notes) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := notes.abi.Unpack("ppStorage", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSetAuthorized is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x711bf9b2.
//
// Solidity: function setAuthorized(address _user, bool _enabled) returns()
func (notes *Notes) PackSetAuthorized(user common.Address, enabled bool) []byte {
	enc, err := notes.abi.Pack("setAuthorized", user, enabled)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetOpened is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcc99d73.
//
// Solidity: function setOpened(uint216 _invoiceId, address _account, uint256 _noteId) returns()
func (notes *Notes) PackSetOpened(invoiceId *big.Int, account common.Address, noteId *big.Int) []byte {
	enc, err := notes.abi.Pack("setOpened", invoiceId, account, noteId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdateVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2e14495.
//
// Solidity: function updateVersion(uint8 _newVersion) returns()
func (notes *Notes) PackUpdateVersion(newVersion uint8) []byte {
	enc, err := notes.abi.Pack("updateVersion", newVersion)
	if err != nil {
		panic(err)
	}
	return enc
}

// NotesNoteCreated represents a NoteCreated event raised by the Notes contract.
type NotesNoteCreated struct {
	InvoiceId        *big.Int
	NoteId           *big.Int
	Author           common.Address
	Share            bool
	EncryptedContent []byte
	Raw              *types.Log // Blockchain specific contextual infos
}

const NotesNoteCreatedEventName = "NoteCreated"

// ContractEventName returns the user-defined event name.
func (NotesNoteCreated) ContractEventName() string {
	return NotesNoteCreatedEventName
}

// UnpackNoteCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NoteCreated(uint216 indexed invoiceId, uint256 indexed noteId, address indexed author, bool share, bytes encryptedContent)
func (notes *Notes) UnpackNoteCreatedEvent(log *types.Log) (*NotesNoteCreated, error) {
	event := "NoteCreated"
	if log.Topics[0] != notes.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NotesNoteCreated)
	if len(log.Data) > 0 {
		if err := notes.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range notes.abi.Events[event].Inputs {
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

// NotesNoteStateChanged represents a NoteStateChanged event raised by the Notes contract.
type NotesNoteStateChanged struct {
	InvoiceId *big.Int
	NoteId    *big.Int
	User      common.Address
	Opened    bool
	Raw       *types.Log // Blockchain specific contextual infos
}

const NotesNoteStateChangedEventName = "NoteStateChanged"

// ContractEventName returns the user-defined event name.
func (NotesNoteStateChanged) ContractEventName() string {
	return NotesNoteStateChangedEventName
}

// UnpackNoteStateChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NoteStateChanged(uint216 indexed invoiceId, uint256 indexed noteId, address indexed user, bool opened)
func (notes *Notes) UnpackNoteStateChangedEvent(log *types.Log) (*NotesNoteStateChanged, error) {
	event := "NoteStateChanged"
	if log.Topics[0] != notes.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(NotesNoteStateChanged)
	if len(log.Data) > 0 {
		if err := notes.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range notes.abi.Events[event].Inputs {
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
func (notes *Notes) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], notes.abi.Errors["EmptyContent"].ID.Bytes()[:4]) {
		return notes.UnpackEmptyContentError(raw[4:])
	}
	if bytes.Equal(raw[:4], notes.abi.Errors["NoteNotFound"].ID.Bytes()[:4]) {
		return notes.UnpackNoteNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], notes.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return notes.UnpackUnauthorizedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// NotesEmptyContent represents a EmptyContent error raised by the Notes contract.
type NotesEmptyContent struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyContent()
func NotesEmptyContentErrorID() common.Hash {
	return common.HexToHash("0x68b3703615f3995ce649856cbb6890fc4faae4c4115369e35365f18349b3af6a")
}

// UnpackEmptyContentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyContent()
func (notes *Notes) UnpackEmptyContentError(raw []byte) (*NotesEmptyContent, error) {
	out := new(NotesEmptyContent)
	if err := notes.abi.UnpackIntoInterface(out, "EmptyContent", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NotesNoteNotFound represents a NoteNotFound error raised by the Notes contract.
type NotesNoteNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoteNotFound()
func NotesNoteNotFoundErrorID() common.Hash {
	return common.HexToHash("0xdc0b7363ceb6a2b5ab15883bc0fc655de5fbef5fe415d47d09c615e438a8fab5")
}

// UnpackNoteNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoteNotFound()
func (notes *Notes) UnpackNoteNotFoundError(raw []byte) (*NotesNoteNotFound, error) {
	out := new(NotesNoteNotFound)
	if err := notes.abi.UnpackIntoInterface(out, "NoteNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// NotesUnauthorized represents a Unauthorized error raised by the Notes contract.
type NotesUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func NotesUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (notes *Notes) UnpackUnauthorizedError(raw []byte) (*NotesUnauthorized, error) {
	out := new(NotesUnauthorized)
	if err := notes.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}
