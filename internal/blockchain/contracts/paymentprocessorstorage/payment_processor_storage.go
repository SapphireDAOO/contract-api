package paymentprocessorstorage

import (
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
	marketplaceAddress, err := bind.Call(c.instance, &bind.CallOpts{Pending: true},
		senderData, c.contract.UnpackGetIntermediatedPlatformsOperator)

	if err != nil {
		return nil, err
	}

	return &marketplaceAddress, nil
}
