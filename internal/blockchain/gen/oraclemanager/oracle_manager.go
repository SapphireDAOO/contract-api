// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oraclemanager

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

// IOracleManagerPriceFeedConfig is an auto generated low-level Go binding around an user-defined struct.
type IOracleManagerPriceFeedConfig struct {
	Aggregator common.Address
	Heartbeat  *big.Int
	Allowed    bool
}

// OraclemanagerMetaData contains all meta data concerning the Oraclemanager contract.
var OraclemanagerMetaData = bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_paymentProcessorStorageAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sequencerUptimeFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_DECIMAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER_GRACE_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSequencerUptimeFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"feed\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsdPerToken\",\"inputs\":[{\"name\":\"_paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsdPerTokenBatch\",\"inputs\":[{\"name\":\"_paymentTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"prices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"supported\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ppStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractPaymentProcessorStorage\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structIOracleManager.PriceFeedConfig\",\"components\":[{\"name\":\"aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"heartbeat\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSequencerUptimeFeed\",\"inputs\":[{\"name\":\"_sequencerUptimeFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"PriceFeedSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"aggregator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"heartbeat\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SequencerDown\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StalePriceFeed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedToken\",\"inputs\":[]}]",
	ID:  "Oraclemanager",
}

// Oraclemanager is an auto generated Go binding around an Ethereum contract.
type Oraclemanager struct {
	abi abi.ABI
}

// NewOraclemanager creates a new instance of Oraclemanager.
func NewOraclemanager() *Oraclemanager {
	parsed, err := OraclemanagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Oraclemanager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Oraclemanager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _paymentProcessorStorageAddress, address _sequencerUptimeFeed) returns()
func (oraclemanager *Oraclemanager) PackConstructor(_paymentProcessorStorageAddress common.Address, _sequencerUptimeFeed common.Address) []byte {
	enc, err := oraclemanager.abi.Pack("", _paymentProcessorStorageAddress, _sequencerUptimeFeed)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTDECIMAL is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (oraclemanager *Oraclemanager) PackDEFAULTDECIMAL() []byte {
	enc, err := oraclemanager.abi.Pack("DEFAULT_DECIMAL")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTDECIMAL is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x26c5eaea.
//
// Solidity: function DEFAULT_DECIMAL() view returns(uint8)
func (oraclemanager *Oraclemanager) UnpackDEFAULTDECIMAL(data []byte) (uint8, error) {
	out, err := oraclemanager.abi.Unpack("DEFAULT_DECIMAL", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackSEQUENCERGRACEPERIOD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd307b6db.
//
// Solidity: function SEQUENCER_GRACE_PERIOD() view returns(uint256)
func (oraclemanager *Oraclemanager) PackSEQUENCERGRACEPERIOD() []byte {
	enc, err := oraclemanager.abi.Pack("SEQUENCER_GRACE_PERIOD")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSEQUENCERGRACEPERIOD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd307b6db.
//
// Solidity: function SEQUENCER_GRACE_PERIOD() view returns(uint256)
func (oraclemanager *Oraclemanager) UnpackSEQUENCERGRACEPERIOD(data []byte) (*big.Int, error) {
	out, err := oraclemanager.abi.Unpack("SEQUENCER_GRACE_PERIOD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetSequencerUptimeFeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa056076c.
//
// Solidity: function getSequencerUptimeFeed() view returns(address feed)
func (oraclemanager *Oraclemanager) PackGetSequencerUptimeFeed() []byte {
	enc, err := oraclemanager.abi.Pack("getSequencerUptimeFeed")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetSequencerUptimeFeed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa056076c.
//
// Solidity: function getSequencerUptimeFeed() view returns(address feed)
func (oraclemanager *Oraclemanager) UnpackGetSequencerUptimeFeed(data []byte) (common.Address, error) {
	out, err := oraclemanager.abi.Unpack("getSequencerUptimeFeed", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetUsdPerToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9866981f.
//
// Solidity: function getUsdPerToken(address _paymentToken) view returns(uint256)
func (oraclemanager *Oraclemanager) PackGetUsdPerToken(paymentToken common.Address) []byte {
	enc, err := oraclemanager.abi.Pack("getUsdPerToken", paymentToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUsdPerToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9866981f.
//
// Solidity: function getUsdPerToken(address _paymentToken) view returns(uint256)
func (oraclemanager *Oraclemanager) UnpackGetUsdPerToken(data []byte) (*big.Int, error) {
	out, err := oraclemanager.abi.Unpack("getUsdPerToken", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetUsdPerTokenBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea73f40c.
//
// Solidity: function getUsdPerTokenBatch(address[] _paymentTokens) view returns(uint256[] prices)
func (oraclemanager *Oraclemanager) PackGetUsdPerTokenBatch(paymentTokens []common.Address) []byte {
	enc, err := oraclemanager.abi.Pack("getUsdPerTokenBatch", paymentTokens)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetUsdPerTokenBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xea73f40c.
//
// Solidity: function getUsdPerTokenBatch(address[] _paymentTokens) view returns(uint256[] prices)
func (oraclemanager *Oraclemanager) UnpackGetUsdPerTokenBatch(data []byte) ([]*big.Int, error) {
	out, err := oraclemanager.abi.Unpack("getUsdPerTokenBatch", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackIsSupportedToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x240028e8.
//
// Solidity: function isSupportedToken(address _token) view returns(bool supported)
func (oraclemanager *Oraclemanager) PackIsSupportedToken(token common.Address) []byte {
	enc, err := oraclemanager.abi.Pack("isSupportedToken", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsSupportedToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x240028e8.
//
// Solidity: function isSupportedToken(address _token) view returns(bool supported)
func (oraclemanager *Oraclemanager) UnpackIsSupportedToken(data []byte) (bool, error) {
	out, err := oraclemanager.abi.Unpack("isSupportedToken", data)
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
func (oraclemanager *Oraclemanager) PackPpStorage() []byte {
	enc, err := oraclemanager.abi.Pack("ppStorage")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPpStorage is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x49f97927.
//
// Solidity: function ppStorage() view returns(address)
func (oraclemanager *Oraclemanager) UnpackPpStorage(data []byte) (common.Address, error) {
	out, err := oraclemanager.abi.Unpack("ppStorage", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSetPriceFeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c83920.
//
// Solidity: function setPriceFeed(address _token, (address,uint96,bool) _config) returns()
func (oraclemanager *Oraclemanager) PackSetPriceFeed(token common.Address, config IOracleManagerPriceFeedConfig) []byte {
	enc, err := oraclemanager.abi.Pack("setPriceFeed", token, config)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetSequencerUptimeFeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb8f44963.
//
// Solidity: function setSequencerUptimeFeed(address _sequencerUptimeFeed) returns()
func (oraclemanager *Oraclemanager) PackSetSequencerUptimeFeed(sequencerUptimeFeed common.Address) []byte {
	enc, err := oraclemanager.abi.Pack("setSequencerUptimeFeed", sequencerUptimeFeed)
	if err != nil {
		panic(err)
	}
	return enc
}

// OraclemanagerPriceFeedSet represents a PriceFeedSet event raised by the Oraclemanager contract.
type OraclemanagerPriceFeedSet struct {
	Token      common.Address
	Aggregator common.Address
	Heartbeat  *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const OraclemanagerPriceFeedSetEventName = "PriceFeedSet"

// ContractEventName returns the user-defined event name.
func (OraclemanagerPriceFeedSet) ContractEventName() string {
	return OraclemanagerPriceFeedSetEventName
}

// UnpackPriceFeedSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PriceFeedSet(address indexed token, address indexed aggregator, uint96 heartbeat)
func (oraclemanager *Oraclemanager) UnpackPriceFeedSetEvent(log *types.Log) (*OraclemanagerPriceFeedSet, error) {
	event := "PriceFeedSet"
	if log.Topics[0] != oraclemanager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OraclemanagerPriceFeedSet)
	if len(log.Data) > 0 {
		if err := oraclemanager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range oraclemanager.abi.Events[event].Inputs {
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
func (oraclemanager *Oraclemanager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], oraclemanager.abi.Errors["InvalidPrice"].ID.Bytes()[:4]) {
		return oraclemanager.UnpackInvalidPriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], oraclemanager.abi.Errors["NotAuthorized"].ID.Bytes()[:4]) {
		return oraclemanager.UnpackNotAuthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], oraclemanager.abi.Errors["SequencerDown"].ID.Bytes()[:4]) {
		return oraclemanager.UnpackSequencerDownError(raw[4:])
	}
	if bytes.Equal(raw[:4], oraclemanager.abi.Errors["StalePrice"].ID.Bytes()[:4]) {
		return oraclemanager.UnpackStalePriceError(raw[4:])
	}
	if bytes.Equal(raw[:4], oraclemanager.abi.Errors["StalePriceFeed"].ID.Bytes()[:4]) {
		return oraclemanager.UnpackStalePriceFeedError(raw[4:])
	}
	if bytes.Equal(raw[:4], oraclemanager.abi.Errors["UnsupportedToken"].ID.Bytes()[:4]) {
		return oraclemanager.UnpackUnsupportedTokenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// OraclemanagerInvalidPrice represents a InvalidPrice error raised by the Oraclemanager contract.
type OraclemanagerInvalidPrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPrice()
func OraclemanagerInvalidPriceErrorID() common.Hash {
	return common.HexToHash("0x00bfc9219afe7e8e3b9f14a4708e4cd3d8acb04e325ce992b2a60a58a519683a")
}

// UnpackInvalidPriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPrice()
func (oraclemanager *Oraclemanager) UnpackInvalidPriceError(raw []byte) (*OraclemanagerInvalidPrice, error) {
	out := new(OraclemanagerInvalidPrice)
	if err := oraclemanager.abi.UnpackIntoInterface(out, "InvalidPrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OraclemanagerNotAuthorized represents a NotAuthorized error raised by the Oraclemanager contract.
type OraclemanagerNotAuthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAuthorized()
func OraclemanagerNotAuthorizedErrorID() common.Hash {
	return common.HexToHash("0xea8e4eb51685727b38a21cb154eb3ebd023f607c62908e0f6f0b645d782af2a4")
}

// UnpackNotAuthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAuthorized()
func (oraclemanager *Oraclemanager) UnpackNotAuthorizedError(raw []byte) (*OraclemanagerNotAuthorized, error) {
	out := new(OraclemanagerNotAuthorized)
	if err := oraclemanager.abi.UnpackIntoInterface(out, "NotAuthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OraclemanagerSequencerDown represents a SequencerDown error raised by the Oraclemanager contract.
type OraclemanagerSequencerDown struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SequencerDown()
func OraclemanagerSequencerDownErrorID() common.Hash {
	return common.HexToHash("0x032b3d00cfb14fdf4eecb317aaf61db9dd7331083f0db9baa2eae06ec3e15ecb")
}

// UnpackSequencerDownError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SequencerDown()
func (oraclemanager *Oraclemanager) UnpackSequencerDownError(raw []byte) (*OraclemanagerSequencerDown, error) {
	out := new(OraclemanagerSequencerDown)
	if err := oraclemanager.abi.UnpackIntoInterface(out, "SequencerDown", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OraclemanagerStalePrice represents a StalePrice error raised by the Oraclemanager contract.
type OraclemanagerStalePrice struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StalePrice()
func OraclemanagerStalePriceErrorID() common.Hash {
	return common.HexToHash("0x19abf40e7c2e0280d6137a5d95d9f3793d913552f00d0e25e4ab4388bcc0d573")
}

// UnpackStalePriceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StalePrice()
func (oraclemanager *Oraclemanager) UnpackStalePriceError(raw []byte) (*OraclemanagerStalePrice, error) {
	out := new(OraclemanagerStalePrice)
	if err := oraclemanager.abi.UnpackIntoInterface(out, "StalePrice", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OraclemanagerStalePriceFeed represents a StalePriceFeed error raised by the Oraclemanager contract.
type OraclemanagerStalePriceFeed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StalePriceFeed()
func OraclemanagerStalePriceFeedErrorID() common.Hash {
	return common.HexToHash("0x1087e109db85b72cf66a8dbc341a9e5601a49c3f12e82151b3eb6e742d4a766e")
}

// UnpackStalePriceFeedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StalePriceFeed()
func (oraclemanager *Oraclemanager) UnpackStalePriceFeedError(raw []byte) (*OraclemanagerStalePriceFeed, error) {
	out := new(OraclemanagerStalePriceFeed)
	if err := oraclemanager.abi.UnpackIntoInterface(out, "StalePriceFeed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OraclemanagerUnsupportedToken represents a UnsupportedToken error raised by the Oraclemanager contract.
type OraclemanagerUnsupportedToken struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedToken()
func OraclemanagerUnsupportedTokenErrorID() common.Hash {
	return common.HexToHash("0x6a1728823cfcc894fe1dcf37bfe71f201fb66b0b61862091f422023e22ea5ab9")
}

// UnpackUnsupportedTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedToken()
func (oraclemanager *Oraclemanager) UnpackUnsupportedTokenError(raw []byte) (*OraclemanagerUnsupportedToken, error) {
	out := new(OraclemanagerUnsupportedToken)
	if err := oraclemanager.abi.UnpackIntoInterface(out, "UnsupportedToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}
