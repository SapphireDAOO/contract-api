package oraclemanager

import (
	"fmt"

	"math/big"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/oraclemanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type OracleManager struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Oraclemanager
	client   *blockchain.Client
}

func NewOracleManager(client *blockchain.Client, address common.Address) *OracleManager {
	contract := gen.NewOraclemanager()
	instance := contract.Instance(client.HTTP, address)

	return &OracleManager{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}

func (c *OracleManager) Address() common.Address {
	if c == nil || c.address == nil {
		return common.Address{}
	}
	return *c.address
}

func (c *OracleManager) UsdPerTokenBatch(tokens []common.Address) ([]*big.Int, error) {
	if c == nil || c.instance == nil {
		return nil, ErrNotInitialized
	}
	if len(tokens) == 0 {
		return nil, nil
	}

	data := c.contract.PackGetUsdPerTokenBatch(tokens)

	prices, err := bind.Call(c.instance, &bind.CallOpts{Pending: false}, data, c.contract.UnpackGetUsdPerTokenBatch)
	if err != nil {
		return nil, err
	}
	if len(prices) != len(tokens) {
		return nil, fmt.Errorf("%w: %d prices for %d tokens", ErrPriceCountMismatch, len(prices), len(tokens))
	}

	return prices, nil
}

func (c *OracleManager) Decimals() (uint8, error) {
	if c == nil || c.instance == nil {
		return 0, ErrNotInitialized
	}

	data := c.contract.PackDEFAULTDECIMAL()
	return bind.Call(c.instance, &bind.CallOpts{Pending: false}, data, c.contract.UnpackDEFAULTDECIMAL)
}
