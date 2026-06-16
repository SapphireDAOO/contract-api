package simplepaymentprocessor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	simpleprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/SimplePaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

type SimplePaymentProcessor struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *simpleprocessor.Simpleprocessor
	client   *blockchain.Client
}

func NewSimplePaymentProcessor(client *blockchain.Client) *SimplePaymentProcessor {
	address := common.HexToAddress(utils.PAYMENT_PROCESSOR_STORAGE_ADDRESS)
	contract := simpleprocessor.NewSimpleprocessor()
	instance := contract.Instance(client.HTTP, address)

	return &SimplePaymentProcessor{
		address:  &address,
		instance: instance,
		contract: contract,
	}
}
