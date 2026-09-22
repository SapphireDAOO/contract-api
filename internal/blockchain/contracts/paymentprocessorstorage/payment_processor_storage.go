package paymentprocessorstorage

import (
	"context"
	"errors"
	"math/big"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/discord"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type PaymentProcessorStorage struct {
	explorerURL string
	notifier    *discord.Client
	address     *common.Address
	instance    *bind.BoundContract
	contract    *gen.Paymentprocessorstorage
	client      *blockchain.Client
}

func NewPaymentProcessorStorage(client *blockchain.Client, address common.Address,
	explorerURL string, notifier *discord.Client) *PaymentProcessorStorage {
	contract := gen.NewPaymentprocessorstorage()
	instance := contract.Instance(client.HTTP, address)

	return &PaymentProcessorStorage{
		address:     &address,
		instance:    instance,
		contract:    contract,
		client:      client,
		explorerURL: explorerURL,
		notifier:    notifier,
	}
}

func (c *PaymentProcessorStorage) GetIntermediatedPlatformsOperator() (*common.Address, error) {
	senderData := c.contract.PackGetIntermediatedPlatformsOperator()
	intermediatedOperatorAddress, err := bind.Call(c.instance, &bind.CallOpts{Pending: true},
		senderData, c.contract.UnpackGetIntermediatedPlatformsOperator)

	if err != nil {
		return nil, err
	}

	return &intermediatedOperatorAddress, nil
}

// PauseState reports whether payment processing is halted and when the
// current emergency pause lapses, in one call so the two cannot come from
// different blocks. The contract does not clear the expiry once that moment
// passes, so it is only meaningful compared against the current time.
func (c *PaymentProcessorStorage) PauseState(ctx context.Context) (bool, *big.Int, error) {
	if c == nil || c.instance == nil {
		return false, nil, errors.New("payment processor storage contract is not initialized")
	}

	state, err := bind.Call(c.instance, &bind.CallOpts{Context: ctx},
		c.contract.PackGetPauseState(), c.contract.UnpackGetPauseState)
	if err != nil {
		return false, nil, err
	}

	return state.PausedState, state.Expiry, nil
}
