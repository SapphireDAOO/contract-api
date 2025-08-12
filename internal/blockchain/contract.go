package blockchain

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

type Contract struct {
	address          *common.Address
	instance         *bind.BoundContract
	paymentProcessor *paymentprocessor.Paymentprocessor
	client           Client
}

type MetaInvoiceResponse struct {
	Url    string  `json:"url"`
	Key    *string `json:"-"`
	Orders map[string]struct {
		Seller    string `json:"seller"`
		InvoiceId string `json:"invoiceId"`
	} `json:"orders"`
}

type SingleInvoiceResponse struct {
	Url       string `json:"url"`
	OrderId   string `json:"orderId"`
	InvoiceId string `json:"invoiceId"`
}

const PAYMENT_PROCESSOR_ADDRESS string = "0x8f73c398ECcd94874752c1dFa48F20A092C8Cf86"

func NewContract(client *Client) *Contract {
	address := common.HexToAddress(PAYMENT_PROCESSOR_ADDRESS)
	contract := paymentprocessor.NewPaymentprocessor()
	instance := contract.Instance(client.eth, address)

	return &Contract{address: &address, instance: instance, paymentProcessor: contract, client: *client}
}

func (c *Contract) CreateInvoice(
	param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam,
) (*SingleInvoiceResponse, error) {
	if len(param) != 1 {
		return nil, errors.New("CreateInvoice expects exactly one parameter")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := c.paymentProcessor.PackCreateSingleInvoice(param[0])

	response, err := c.simulateAndBroadcast(ctx, data)

	if err != nil {
		return nil, err
	}

	if response.result == nil {
		return &SingleInvoiceResponse{
			OrderId:   param[0].OrderId,
			InvoiceId: response.receipt.Logs[0].Topics[1].Hex(),
		}, nil
	}
	return &SingleInvoiceResponse{
		OrderId:   param[0].OrderId,
		InvoiceId: *response.result,
	}, nil

}

func (c *Contract) CreateInvoices(
	param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam,
) (*MetaInvoiceResponse, error) {
	if len(param) < 2 {
		return nil, errors.New("parameter has to be greater than one")
	}

	orders := make(map[string]struct {
		Seller    string `json:"seller"`
		InvoiceId string `json:"invoiceId"`
	})

	for i := range param {
		sellerAddress := param[i].Seller.Hex()
		orderId := param[i].OrderId

		invoiceId, err := utils.Keccak256(orderId)
		if err != nil {
			return nil, err
		}

		o := orders[orderId]
		o.Seller = sellerAddress
		o.InvoiceId = invoiceId.Hex()

		orders[orderId] = o
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := c.paymentProcessor.PackCreateMetaInvoice(param)

	response, err := c.simulateAndBroadcast(ctx, data)

	if err != nil {
		return nil, err
	}

	if response.result == nil {
		key := response.receipt.Logs[len(param)].Topics[1].Hex()
		return &MetaInvoiceResponse{
			Key:    &key,
			Orders: orders,
		}, nil
	}
	return &MetaInvoiceResponse{
		Key:    response.result,
		Orders: orders,
	}, nil
}

func (c *Contract) CreateDispute(orderId common.Hash) (*common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := c.paymentProcessor.PackCreateDispute(orderId)

	response, err := c.simulateAndBroadcast(ctx, data)

	if err != nil {
		return nil, err
	}

	return &response.receipt.TxHash, nil
}

func (c *Contract) HandleDispute(
	orderId common.Hash, action MarketplaceAction, sellersShare *big.Int,
) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if sellersShare == nil {
		sellersShare = big.NewInt(0)
	}

	if err != nil {
		return nil, err
	}

	var data []byte

	switch action {
	case ResolveDispute:
		data = c.paymentProcessor.PackResolveDispute(orderId)

	case SettleDispute:
		data = c.paymentProcessor.PackHandleDispute(orderId, c.getDisputeResolution(SettleDispute), sellersShare)

	case DismissDispute:
		data = c.paymentProcessor.PackHandleDispute(orderId, c.getDisputeResolution(DismissDispute), sellersShare)

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

func (c *Contract) Cancel(orderId common.Hash) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackCancelInvoice(orderId)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil

}

func (c *Contract) Refund(orderId common.Hash, refundShare *big.Int) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackRefund(orderId, refundShare)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}

func (c *Contract) Release(orderId common.Hash, sellerShare *big.Int) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackRelease(orderId, sellerShare)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}
