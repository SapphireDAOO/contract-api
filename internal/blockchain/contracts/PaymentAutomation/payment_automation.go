package paymentautomation

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/PaymentAutomation"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func NewPaymentAutomation(client *blockchain.Client) *PaymentAutomation {
	address := common.HexToAddress(utils.PAYMENT_AUTOMATION_ADDRESS)
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
		return false, errors.New("payment automation contract is not initialized")
	}

	data := c.contract.PackHasDueTasks()
	return bind.Call(c.instance, &bind.CallOpts{Context: ctx}, data, c.contract.UnpackHasDueTasks)
}

func (c *PaymentAutomation) ProcessDueTasks(ctx context.Context) (*types.Receipt, error) {
	if c == nil || c.instance == nil || c.address == nil {
		return nil, errors.New("payment automation contract is not initialized")
	}

	auth, err := blockchain.Auth(c.client.ChainId)
	if err != nil {
		return nil, err
	}

	data := c.contract.PackProcessDueTasks()

	// Simulate first so a reverting call never costs gas.
	if _, err := c.client.HTTP.CallContract(ctx, ethereum.CallMsg{
		From: auth.From,
		To:   c.address,
		Data: data,
	}, nil); err != nil {
		return nil, err
	}

	tx, err := bind.Transact(c.instance, auth, data)
	if err != nil {
		return nil, err
	}

	return bind.WaitMined(ctx, c.client.HTTP, tx.Hash())
}
