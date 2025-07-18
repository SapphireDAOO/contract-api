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
	Seller map[string]struct {
		SubOrderIds map[string]string `json:"hashed_sub_order_ids"`
	} `json:"orders"`
}

type SingleInvoiceResponse struct {
	Url     string  `json:"url"`
	OrderId string  `json:"order_id"`
	Key     *string `json:"hashed_order_id"`
}

const PAYMENT_PROCESSOR_ADDRESS string = "0xb9DD118C880759E62516fa2c88Eb76ba3fd42eae"

func NewContract(client *Client) *Contract {
	address := common.HexToAddress(PAYMENT_PROCESSOR_ADDRESS)

	contract := paymentprocessor.NewPaymentprocessor()

	instance := contract.Instance(client.eth, address)

	return &Contract{address: &address, instance: instance, paymentProcessor: contract, client: *client}
}

func (c *Contract) CreateInvoice(
	param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam,
) (*SingleInvoiceResponse, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	if len(param) != 1 {
		return nil, errors.New("CreateInvoice expects exactly one parameter")
	}

	data := c.paymentProcessor.PackCreateSingleInvoice(param[0])

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.client.eth, tx.Hash())

	if err != nil {
		return nil, err
	}

	invoiceKey := receipt.Logs[0].Topics[1]

	var res SingleInvoiceResponse

	key := invoiceKey.Hex()

	res.Key = &key
	res.OrderId = param[0].OrderId

	return &res, nil
}

func (c *Contract) CreateInvoices(
	param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam,
) (*MetaInvoiceResponse, error) {

	var res MetaInvoiceResponse
	if len(param) == 0 {
		return nil, errors.New("parameter cannot be empty")
	}

	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackCreateMetaInvoice(param)
	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, c.client.eth, tx.Hash())

	if err != nil {
		return nil, err
	}

	metaInvoiceKey := receipt.Logs[len(param)].Topics[1]

	seller := make(map[string]struct {
		SubOrderIds map[string]string `json:"hashed_sub_order_ids"`
	})
	for i := range param {
		sellerAddress := param[i].Seller.Hex()
		id := param[i].OrderId

		hashed, err := utils.Keccak256(id)
		if err != nil {
			return nil, err
		}

		s := seller[sellerAddress]
		if s.SubOrderIds == nil {
			s.SubOrderIds = make(map[string]string)
		}

		s.SubOrderIds[id] = hashed.Hex()

		seller[sellerAddress] = s
	}

	key := metaInvoiceKey.Hex()
	res.Key = &key
	res.Seller = seller

	return &res, nil

}

func (c *Contract) HandleDispute(
	orderId common.Hash, resolver *common.Address, action MarketplaceAction, sellersShare *big.Int,
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

func (c *Contract) getDisputeResolution(action MarketplaceAction) uint8 {
	if action == DismissDispute {
		value, _ := bind.Call(
			c.instance, nil, c.paymentProcessor.PackDISPUTEDISMISSED(), c.paymentProcessor.UnpackDISPUTEDISMISSED)
		return value
	}

	if action == SettleDispute {
		value, _ := bind.Call(
			c.instance, nil, c.paymentProcessor.PackDISPUTESETTLED(), c.paymentProcessor.UnpackDISPUTESETTLED)
		return value
	}

	return 0
}
