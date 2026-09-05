package simplepaymentprocessor

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/simplepaymentprocessor"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type SimplePaymentProcessor struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Simplepaymentprocessor
	client   *blockchain.Client
}

func NewSimplePaymentProcessor(client *blockchain.Client, address common.Address) *SimplePaymentProcessor {
	contract := gen.NewSimplepaymentprocessor()
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
) (gen.ISimplePaymentProcessorInvoice, error) {
	if c == nil || c.instance == nil {
		return gen.ISimplePaymentProcessorInvoice{},
			errors.New("simple payment processor contract is not initialized")
	}

	data := c.contract.PackGetInvoiceData(invoiceId)
	return bind.Call(c.instance, &bind.CallOpts{Context: ctx}, data, c.contract.UnpackGetInvoiceData)
}
