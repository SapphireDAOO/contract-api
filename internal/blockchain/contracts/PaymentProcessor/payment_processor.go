package paymentprocesor

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func NewPaymentprocessor(client *blockchain.Client) *PaymentProcessor {
	address := common.HexToAddress(utils.PAYMENT_PROCESSOR_ADDRESS)
	contract := advancedprocessor.NewAdvancedprocessor()
	instance := contract.Instance(client.HTTP, address)

	return &PaymentProcessor{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}

func (c *PaymentProcessor) CreateInvoice(
	param []advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam,
	marketplaceAddress common.Address,
) (*InvoiceResponse, error) {

	if len(param) != 1 {
		return nil, errors.New("CreateInvoice expects exactly one parameter")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data := c.contract.PackCreateSingleInvoice(param[0])

	_, err := contracts.
		SimulateAndBroadcast(ctx, c.instance, c.client, marketplaceAddress, *c.address, data)

	if err != nil {
		return nil, err
	}

	orders := make(map[string]struct {
		Seller  string `json:"seller"`
		OrderId string `json:"orderId"`
	})

	id := param[0].OrderId
	o := orders[id]

	o.Seller = param[0].Seller.Hex()
	o.OrderId = utils.OrderIDToUint216(id)

	orders[id] = o

	empty := ""
	return &InvoiceResponse{
		MetaInvoiceId: &empty,
		Orders:        orders,
	}, nil

}

func (c *PaymentProcessor) CreateInvoices(
	param []advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam,
	marketplaceAddress common.Address,
) (*InvoiceResponse, error) {
	if len(param) < 2 {
		return nil, errors.New("parameter has to be greater than one")
	}

	orders := make(map[string]struct {
		Seller  string `json:"seller"`
		OrderId string `json:"orderId"`
	})

	for i := range param {
		id := param[i].OrderId

		o := orders[id]
		o.Seller = param[i].Seller.Hex()

		o.OrderId = utils.OrderIDToUint216(id)

		orders[id] = o
	}

	if c.address == nil {
		return nil, errors.New("payment processor contract address is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data := c.contract.PackCreateMetaInvoice(param)

	response, err := contracts.SimulateAndBroadcast(ctx,
		c.instance, c.client, marketplaceAddress, *c.address, data)

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		result := new(big.Int).SetBytes(response.Receipt.Logs[len(param)].Topics[1].Bytes()).String()
		return &InvoiceResponse{
			MetaInvoiceId: &result,
			Orders:        orders,
		}, nil
	}
	return &InvoiceResponse{
		MetaInvoiceId: response.Result,
		Orders:        orders,
	}, nil
}

func (c *PaymentProcessor) CreateDispute(orderId *big.Int, marketplaceAddress common.Address) (*common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := c.contract.PackCreateDispute(orderId)

	response, err := contracts.SimulateAndBroadcast(ctx, c.instance, c.client, marketplaceAddress, *c.address, data)

	if err != nil {
		return nil, err
	}

	return &response.Receipt.TxHash, nil
}

func (c *PaymentProcessor) HandleDispute(
	orderId *big.Int, action blockchain.MarketplaceAction, sellersShare *big.Int,
) (*common.Hash, error) {
	auth, err := blockchain.Auth(c.client.ChainId)

	if sellersShare == nil {
		sellersShare = big.NewInt(0)
	}

	if err != nil {
		return nil, err
	}

	var data []byte

	switch action {
	case blockchain.ResolveDispute:
		data = c.contract.PackResolveDispute(orderId)

	case blockchain.SettleDispute:
		data = c.contract.PackHandleDispute(orderId,
			c.getDisputeResolution(blockchain.SettleDispute), sellersShare)

	case blockchain.DismissDispute:
		data = c.contract.PackHandleDispute(orderId,
			c.getDisputeResolution(blockchain.DismissDispute), sellersShare)

	default:
		return nil, errors.New("unsupported marketplace action")

	}

	tx, err := bind.Transact(c.instance, auth, data)
	if err != nil {
		return nil, err
	}

	hash := tx.Hash()
	return &hash, nil

}

func (c *PaymentProcessor) Cancel(orderId *big.Int) (*common.Hash, error) {
	auth, err := blockchain.Auth(c.client.ChainId)

	if err != nil {
		return nil, err
	}

	data := c.contract.PackCancelInvoice(orderId)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil

}

func (c *PaymentProcessor) Refund(orderId *big.Int, refundShare *big.Int) (*common.Hash, error) {
	auth, err := blockchain.Auth(c.client.ChainId)

	if err != nil {
		return nil, err
	}

	data := c.contract.PackRefund(orderId, refundShare)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}

func (c *PaymentProcessor) Release(orderId *big.Int) (*common.Hash, error) {
	auth, err := blockchain.Auth(c.client.ChainId)

	if err != nil {
		return nil, err
	}

	data := c.contract.PackRelease(orderId)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}

func (c *PaymentProcessor) getDisputeResolution(action blockchain.MarketplaceAction) uint8 {
	if action == blockchain.DismissDispute {
		value, _ := bind.Call(
			c.instance, nil, c.contract.PackDISPUTEDISMISSED(), c.contract.UnpackDISPUTEDISMISSED)
		return value
	}

	if action == blockchain.SettleDispute {
		value, _ := bind.Call(
			c.instance, nil, c.contract.PackDISPUTESETTLED(), c.contract.UnpackDISPUTESETTLED)
		return value
	}

	return 0
}
