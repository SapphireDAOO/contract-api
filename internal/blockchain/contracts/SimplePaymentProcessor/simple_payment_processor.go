package simplepaymentprocessor

import (
	"context"
	"errors"
	"math/big"
	"time"

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
	address := common.HexToAddress(utils.SIMPLE_PAYMENT_PROCESSOR_ADDRESS)
	contract := simpleprocessor.NewSimpleprocessor()
	instance := contract.Instance(client.HTTP, address)

	return &SimplePaymentProcessor{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}

func (c *SimplePaymentProcessor) IsSettlementExpired() (bool, error) {
	data := c.contract.PackGetItems()

	items, err := bind.Call(c.instance, &bind.CallOpts{Pending: false}, data, c.contract.UnpackGetItems)
	if err != nil {
		return false, err
	}

	if len(items) == 0 {
		return false, nil
	}

	head := items[0]

	data = c.contract.PackGetInvoiceData(head)
	result, err := bind.Call(c.instance, &bind.CallOpts{Pending: false}, data, c.contract.UnpackGetInvoiceData)
	if err != nil {
		return false, err
	}

	var exp int64

	switch result.State {
	case 2:
		exp = result.ExpiresAt.Int64()
	default:
		exp = result.ReleaseAt.Int64()
	}

	currentTime := time.Now().Unix() + 60

	return currentTime > exp, nil
}

// GetInvoiceData reads a single invoice. A zero state means the invoice does
// not exist on this processor.
func (c *SimplePaymentProcessor) GetInvoiceData(
	ctx context.Context, invoiceId *big.Int,
) (simpleprocessor.ISimplePaymentProcessorInvoice, error) {
	if c == nil || c.instance == nil {
		return simpleprocessor.ISimplePaymentProcessorInvoice{},
			errors.New("simple payment processor contract is not initialized")
	}

	data := c.contract.PackGetInvoiceData(invoiceId)
	return bind.Call(c.instance, &bind.CallOpts{Context: ctx}, data, c.contract.UnpackGetInvoiceData)
}
