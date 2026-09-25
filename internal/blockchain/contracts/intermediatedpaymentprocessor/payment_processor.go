package intermediatedpaymentprocessor

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/tx"
	"github.com/SapphireDAOO/contract-api/internal/callback"
	"github.com/SapphireDAOO/contract-api/internal/invoice"
	"github.com/SapphireDAOO/contract-api/internal/revert"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewPaymentprocessor(client *blockchain.Client, address common.Address,
	explorerURL string, callbacks *callback.Client) *PaymentProcessor {
	contract := gen.NewIntermediatedpaymentprocessor()
	instance := contract.Instance(client.HTTP, address)

	return &PaymentProcessor{
		address:     &address,
		instance:    instance,
		contract:    contract,
		client:      client,
		explorerURL: explorerURL,
		callbacks:   callbacks,
	}
}

func (c *PaymentProcessor) CreateInvoice(
	param []gen.IIntermediatedPaymentProcessorInvoiceCreationParam,
	intermediatedOperatorAddress common.Address,
) (*InvoiceResponse, error) {

	if len(param) != 1 {
		return nil, errors.New("CreateInvoice expects exactly one parameter")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data := c.contract.PackCreateSingleInvoice(param[0])

	_, err := tx.
		SimulateAndBroadcast(ctx, c.instance, c.client, intermediatedOperatorAddress, *c.address, data)

	invoices := make(map[string]struct {
		Seller    string `json:"seller"`
		InvoiceId string `json:"invoiceId"`
	})

	id := param[0].InvoiceId
	inv := invoices[id]

	inv.Seller = param[0].Seller.Hex()
	inv.InvoiceId = invoice.InvoiceIDToUint216(id)

	invoices[id] = inv

	if err != nil {
		if strings.Contains(revert.Reason(err), "An invoice with this identifier already exists.") {
			return &InvoiceResponse{
				Invoices: invoices,
			}, nil
		}
		return nil, err
	}

	return &InvoiceResponse{
		Invoices: invoices,
	}, nil

}

func (c *PaymentProcessor) CreateInvoices(
	param []gen.IIntermediatedPaymentProcessorInvoiceCreationParam,
	intermediatedOperatorAddress common.Address,
) (*InvoiceResponse, error) {
	if len(param) < 2 {
		return nil, errors.New("parameter has to be greater than one")
	}

	invoices := make(map[string]struct {
		Seller    string `json:"seller"`
		InvoiceId string `json:"invoiceId"`
	})

	for i := range param {
		id := param[i].InvoiceId

		inv := invoices[id]
		inv.Seller = param[i].Seller.Hex()

		inv.InvoiceId = invoice.InvoiceIDToUint216(id)

		invoices[id] = inv
	}

	if c.address == nil {
		return nil, errors.New("payment processor contract address is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	data := c.contract.PackCreateMetaInvoice(param)

	response, err := tx.SimulateAndBroadcast(ctx,
		c.instance, c.client, intermediatedOperatorAddress, *c.address, data)

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		result, err := c.metaInvoiceIDFromLogs(response.Receipt)
		if err != nil {
			return nil, err
		}
		return &InvoiceResponse{
			MetaInvoiceId: result,
			Invoices:      invoices,
		}, nil
	}
	return &InvoiceResponse{
		MetaInvoiceId: response.Result,
		Invoices:      invoices,
	}, nil
}

// metaInvoiceIDFromLogs finds the MetaInvoiceCreated event by its signature.
// Picking a log by position is not safe: each sub-invoice emits several logs,
// and how many is the contract's business, not this code's.
func (c *PaymentProcessor) metaInvoiceIDFromLogs(receipt *types.Receipt) (*string, error) {
	if receipt == nil {
		return nil, errors.New("no receipt returned for meta invoice creation")
	}

	for _, log := range receipt.Logs {
		if log == nil || len(log.Topics) == 0 || log.Address != *c.address {
			continue
		}

		event, err := c.contract.UnpackMetaInvoiceCreatedEvent(log)
		if err != nil {
			continue
		}

		id := event.MetaInvoiceId.String()
		return &id, nil
	}

	return nil, fmt.Errorf(
		"MetaInvoiceCreated event not found in receipt logs for tx %s", receipt.TxHash.Hex())
}

func (c *PaymentProcessor) CreateDispute(invoiceId *big.Int, intermediatedOperatorAddress common.Address) (*common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data := c.contract.PackCreateDispute(invoiceId)

	response, err := tx.SimulateAndBroadcast(ctx, c.instance, c.client, intermediatedOperatorAddress, *c.address, data)

	if err != nil {
		return nil, err
	}

	return &response.Receipt.TxHash, nil
}

func (c *PaymentProcessor) HandleDispute(
	invoiceId *big.Int, action blockchain.MarketplaceAction, sellersShare *big.Int,
) (*common.Hash, error) {
	auth, err := c.client.Auth()

	if sellersShare == nil {
		sellersShare = big.NewInt(0)
	}

	if err != nil {
		return nil, err
	}

	var data []byte

	switch action {
	case blockchain.ResolveDispute:
		data = c.contract.PackResolveDispute(invoiceId)

	case blockchain.SettleDispute:
		data = c.contract.PackHandleDispute(invoiceId,
			c.getDisputeResolution(blockchain.SettleDispute), sellersShare)

	case blockchain.DismissDispute:
		data = c.contract.PackHandleDispute(invoiceId,
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

func (c *PaymentProcessor) Cancel(invoiceId *big.Int) (*common.Hash, error) {
	auth, err := c.client.Auth()

	if err != nil {
		return nil, err
	}

	data := c.contract.PackCancelInvoice(invoiceId)

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

func (c *PaymentProcessor) Refund(invoiceId *big.Int, refundShare *big.Int) (*common.Hash, error) {
	auth, err := c.client.Auth()

	if err != nil {
		return nil, err
	}

	data := c.contract.PackRefund(invoiceId, refundShare)

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

func (c *PaymentProcessor) Release(invoiceId *big.Int) (*ReleaseResult, error) {
	auth, err := c.client.Auth()
	if err != nil {
		return nil, err
	}

	data := c.contract.PackRelease(invoiceId)

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

func (c *PaymentProcessor) findPaymentReleasedEvent(receipt *types.Receipt) *gen.IntermediatedpaymentprocessorPaymentReleased {
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

func (c *PaymentProcessor) GetInvoiceData(invoiceId *big.Int) (gen.IIntermediatedPaymentProcessorInvoice, error) {
	if c == nil || c.client == nil || c.client.HTTP == nil {
		return gen.IIntermediatedPaymentProcessorInvoice{}, errors.New("blockchain client not initialized")
	}
	if c.address == nil {
		return gen.IIntermediatedPaymentProcessorInvoice{}, errors.New("payment processor address not initialized")
	}

	data := c.contract.PackGetInvoice(invoiceId)

	out, err := c.client.HTTP.CallContract(context.Background(), ethereum.CallMsg{
		To:   c.address,
		Data: data,
	}, nil)

	if err != nil {
		return gen.IIntermediatedPaymentProcessorInvoice{}, err
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
