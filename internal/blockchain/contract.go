package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

const PAYMENT_PROCESSOR_ADDRESS string = "0x1A1b771B7e6cE617d22A148d08d0395Ca29f208a"

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

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
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
	} else {
		return &SingleInvoiceResponse{
			OrderId:   param[0].OrderId,
			InvoiceId: *response.result,
		}, nil
	}
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
	} else {
		return &MetaInvoiceResponse{
			Key:    response.result,
			Orders: orders,
		}, nil
	}

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

func (c *Contract) Refund(orderId [32]byte, amount *big.Int) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackRefund(orderId, amount)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}

func (c *Contract) Release(orderId [32]byte, sellerShare *big.Int) (*common.Hash, error) {
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

func waitForReceipt(ctx context.Context, client *ethclient.Client, hash common.Hash) (chan *types.Receipt, chan error) {
	receiptChan := make(chan *types.Receipt, 1)
	errorChan := make(chan error, 1)

	go func() {
		receipt, err := bind.WaitMined(ctx, client, hash)
		if err != nil {
			errorChan <- err
			return
		}
		receiptChan <- receipt
	}()
	return receiptChan, errorChan
}

func anotherWaitForReceipt(ctx context.Context, orderId string, result []byte,
	errorChan <-chan error, receiptChan <-chan *types.Receipt) (any, error) {
	select {
	case <-ctx.Done():
		fmt.Println("Timeout occured while sending transaction to be mined", ctx.Err())
		return &SingleInvoiceResponse{
			OrderId:   orderId,
			InvoiceId: "0x" + common.Bytes2Hex(result),
		}, nil

	case err := <-errorChan:
		return nil, err

	case receipt := <-receiptChan:
		if len(receipt.Logs) == 0 {
			return nil, errors.New("no logs in receipt")
		}
		return &SingleInvoiceResponse{
			OrderId:   orderId,
			InvoiceId: receipt.Logs[0].Topics[1].Hex(),
		}, nil
	}

	// select {
	// case <-ctx.Done():
	// 	fmt.Println("Timeout occured while sending transaction to be mined", ctx.Err())
	// 	key := "0x" + common.Bytes2Hex(result)
	// 	return &MetaInvoiceResponse{
	// 		Key:    &key,
	// 		Orders: orders,
	// 	}, nil

	// case err = <-errorChan:
	// 	return nil, err

	// case receipt := <-receiptChan:
	// 	if len(result) == 0 {
	// 		return nil, errors.New("no logs in receipt")
	// 	}
	// 	key := receipt.Logs[len(param)].Topics[1].Hex()
	// 	return &MetaInvoiceResponse{
	// 		Key:    &key,
	// 		Orders: orders,
	// 	}, nil
	// }

}
