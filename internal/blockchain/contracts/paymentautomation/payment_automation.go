package paymentautomation

import (
	"context"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/paymentautomation"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const minProcessDueTasksGas uint64 = 5_000_000

func NewPaymentAutomation(client *blockchain.Client, address common.Address) *PaymentAutomation {
	contract := gen.NewPaymentautomation()
	instance := contract.Instance(client.HTTP, address)

	return &PaymentAutomation{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}

func (c *PaymentAutomation) HasDueTasks(ctx context.Context) (bool, error) {
	if c == nil || c.instance == nil {
		return false, ErrNotInitialized
	}

	data := c.contract.PackHasDueTasks()
	return bind.Call(c.instance, &bind.CallOpts{Context: ctx}, data, c.contract.UnpackHasDueTasks)
}

func (c *PaymentAutomation) ProcessDueTasks(ctx context.Context) (*types.Receipt, error) {
	if c == nil || c.instance == nil || c.address == nil {
		return nil, ErrNotInitialized
	}

	auth, err := c.client.Auth()
	if err != nil {
		return nil, err
	}

	data := c.contract.PackProcessDueTasks()

	estimate, err := c.client.HTTP.EstimateGas(ctx, ethereum.CallMsg{
		From: auth.From,
		To:   c.address,
		Data: data,
	})
	if err != nil {
		return nil, err
	}
	auth.GasLimit = max(estimate, minProcessDueTasksGas)

	tx, err := bind.Transact(c.instance, auth, data)
	if err != nil {
		return nil, err
	}

	return bind.WaitMined(ctx, c.client.HTTP, tx.Hash())
}
