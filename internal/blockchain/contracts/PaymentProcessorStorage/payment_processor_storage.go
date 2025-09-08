package paymentprocessorstorage

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	processorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/PaymentProcessorStorage"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

type PaymentProcessorStorage struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *processorstorage.Processorstorage
	client   *blockchain.Client
}

func NewPaymentProcessorStorage(client *blockchain.Client) *PaymentProcessorStorage {
	address := common.HexToAddress(utils.PAYMENT_PROCESSOR_STORAGE_ADDRESS)
	contract := processorstorage.NewProcessorstorage()
	instance := contract.Instance(client.HTTP, address)

	return &PaymentProcessorStorage{
		address:  &address,
		instance: instance,
		contract: contract,
	}
}

func (c *PaymentProcessorStorage) GetMarketplaceAddress() (*common.Address, error) {
	senderData := c.contract.PackGetMarketplace()
	marketplaceAddress, err := bind.Call(c.instance, &bind.CallOpts{Pending: true},
		senderData, c.contract.UnpackGetMarketplace)

	if err != nil {
		return nil, err
	}

	return &marketplaceAddress, nil
}
