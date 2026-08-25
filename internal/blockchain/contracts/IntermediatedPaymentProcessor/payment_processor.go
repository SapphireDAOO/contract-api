package intermediatedpaymentprocessor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts"
	intermediatedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/IntermediatedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func NewPaymentprocessor(client *blockchain.Client) *PaymentProcessor {
	address := common.HexToAddress(utils.PAYMENT_PROCESSOR_ADDRESS)
	contract := intermediatedprocessor.NewIntermediatedprocessor()
	instance := contract.Instance(client.HTTP, address)

	return &PaymentProcessor{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}

func (c *PaymentProcessor) CreateInvoice(
	param []intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam,
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

	orders := make(map[string]struct {
		Seller  string `json:"seller"`
		OrderId string `json:"orderId"`
	})

	id := param[0].InvoiceId
	o := orders[id]

	o.Seller = param[0].Seller.Hex()
	o.OrderId = utils.OrderIDToUint216(id)

	orders[id] = o

	if err != nil {
		if strings.Contains(utils.Reason(err), "An invoice with this identifier already exists.") {
			return &InvoiceResponse{
				Orders: orders,
			}, nil
		}
		return nil, err
	}

	return &InvoiceResponse{
		Orders: orders,
	}, nil

}

func (c *PaymentProcessor) CreateInvoices(
	param []intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam,
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
		id := param[i].InvoiceId

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
		logs := response.Receipt.Logs
		if len(logs) <= len(param) || len(logs[len(param)].Topics) < 2 {
			return nil, fmt.Errorf(
				"meta invoice id not found in receipt logs for tx %s", response.Receipt.TxHash.Hex())
		}
		result := new(big.Int).SetBytes(logs[len(param)].Topics[1].Bytes()).String()
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

	receipt, err := bind.WaitMined(context.Background(), c.client.HTTP, tx.Hash())
	if err != nil {
		return nil, err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction reverted: %s", tx.Hash().Hex())
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

	receipt, err := bind.WaitMined(context.Background(), c.client.HTTP, tx.Hash())
	if err != nil {
		return nil, err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction reverted: %s", tx.Hash().Hex())
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

	receipt, err := bind.WaitMined(context.Background(), c.client.HTTP, tx.Hash())
	if err != nil {
		return nil, err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction reverted: %s", tx.Hash().Hex())
	}

	hash := tx.Hash()

	return &hash, nil
}

type ReleaseResult struct {
	TxHash         common.Hash
	Seller         common.Address
	PaymentToken   common.Address
	SellerAmount   *big.Int
	BlockTimestamp int64
}

func (c *PaymentProcessor) Release(orderId *big.Int) (*ReleaseResult, error) {
	auth, err := blockchain.Auth(c.client.ChainId)
	if err != nil {
		return nil, err
	}

	data := c.contract.PackRelease(orderId)

	tx, err := bind.Transact(c.instance, auth, data)
	if err != nil {
		return nil, err
	}

	receipt, err := bind.WaitMined(context.Background(), c.client.HTTP, tx.Hash())
	if err != nil {
		return nil, err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction reverted: %s", tx.Hash().Hex())
	}

	result := &ReleaseResult{
		TxHash:         tx.Hash(),
		BlockTimestamp: c.blockTimestampMillis(receipt.BlockHash),
	}
	if event := c.findPaymentReleasedEvent(receipt); event != nil {
		result.Seller = event.Receiver
		result.PaymentToken = event.Currency
		result.SellerAmount = event.SellerAmount
	}
	return result, nil
}

func (c *PaymentProcessor) blockTimestampMillis(blockHash common.Hash) int64 {
	if c.client == nil || c.client.HTTP == nil {
		return 0
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	header, err := c.client.HTTP.HeaderByHash(ctx, blockHash)
	if err != nil || header == nil {
		return 0
	}
	return int64(header.Time) * 1000
}

func (c *PaymentProcessor) findPaymentReleasedEvent(receipt *types.Receipt) *intermediatedprocessor.IntermediatedprocessorPaymentReleased {
	if receipt == nil || c.address == nil {
		return nil
	}
	for _, vlog := range receipt.Logs {
		if vlog == nil || vlog.Address != *c.address {
			continue
		}
		event, err := c.contract.UnpackPaymentReleasedEvent(vlog)
		if err != nil || event == nil {
			continue
		}
		return event
	}
	return nil
}

func (c *PaymentProcessor) GetInvoiceData(orderId *big.Int) (intermediatedprocessor.IIntermediatedPaymentProcessorInvoice, error) {
	if c == nil || c.client == nil || c.client.HTTP == nil {
		return intermediatedprocessor.IIntermediatedPaymentProcessorInvoice{}, errors.New("blockchain client not initialized")
	}
	if c.address == nil {
		return intermediatedprocessor.IIntermediatedPaymentProcessorInvoice{}, errors.New("payment processor address not initialized")
	}

	data := c.contract.PackGetInvoice(orderId)

	out, err := c.client.HTTP.CallContract(context.Background(), ethereum.CallMsg{
		To:   c.address,
		Data: data,
	}, nil)

	if err != nil {
		return intermediatedprocessor.IIntermediatedPaymentProcessorInvoice{}, err
	}

	return c.contract.UnpackGetInvoice(out)
}

const (
	disputeDismissed uint8 = 7
	disputeSettled   uint8 = 8
)

func (c *PaymentProcessor) getDisputeResolution(action blockchain.MarketplaceAction) uint8 {
	switch action {
	case blockchain.DismissDispute:
		return disputeDismissed
	case blockchain.SettleDispute:
		return disputeSettled
	}
	return 0
}
