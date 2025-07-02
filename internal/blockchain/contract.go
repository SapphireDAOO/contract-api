package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

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
	Key    *common.Hash `json:"key,omitempty"`
	Orders map[string]struct {
		Seller      string            `json:"seller"`
		SubOrderIds map[string]string `json:"sub_order_ids"`
	} `json:"orders"`
}

type SingleInvoiceResponse struct {
	Key     *common.Hash `json:"key,omitempty"`
	OrderId string       `json:"order_id"`
}

const PAYMENT_PROCESSOR_ADDRESS string = "0x953ff222255730544c8e118a2ccb5dfb856bfbad"

func NewContract(client *Client) *Contract {
	address := common.HexToAddress(PAYMENT_PROCESSOR_ADDRESS)

	contract := paymentprocessor.NewPaymentprocessor()

	instance := contract.Instance(client.eth, address)

	return &Contract{address: &address, instance: instance, paymentProcessor: contract, client: *client}
}

func (c *Contract) CreateInvoice(param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam) (*SingleInvoiceResponse, error) {
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

	receipt, err := bind.WaitMined(context.Background(), c.client.eth, tx.Hash())

	if err != nil {
		return nil, err
	}

	invoiceKey := receipt.Logs[0].Topics[1]

	var res SingleInvoiceResponse

	orderId, err := utils.Keccak256(param[0].OrderId)

	if err != nil {
		return nil, err
	}

	res.Key = &invoiceKey
	res.OrderId = orderId.Hex()

	return &res, nil
}

func (c *Contract) CreateInvoices(param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam) (*MetaInvoiceResponse, error) {

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

	receipt, err := bind.WaitMined(context.Background(), c.client.eth, tx.Hash())

	if err != nil {
		return nil, err
	}

	metaInvoiceKey := receipt.Logs[len(param)-1].Topics[1]

	orders := make(map[string]struct {
		Seller      string            `json:"seller"`
		SubOrderIds map[string]string `json:"sub_order_ids"`
	})
	for i := range param {
		seller := param[i].Seller.Hex()
		id := param[i].OrderId

		hashed, err := utils.Keccak256(id)
		if err != nil {
			return nil, err
		}

		orderId, err := computeOrderId(seller, metaInvoiceKey.Hex())
		if err != nil {
			return nil, err
		}

		order := orders[orderId]
		if order.SubOrderIds == nil {
			order.SubOrderIds = make(map[string]string)
		}
		order.Seller = seller
		order.SubOrderIds[id] = hashed.Hex()

		orders[orderId] = order
	}

	res.Key = &metaInvoiceKey
	res.Orders = orders

	return &res, nil

}

func (c *Contract) ReleaseEscrow(orderId common.Hash, resolver *common.Address, action MarketplaceAction, sellersShare *big.Int) (*common.Hash, error) {
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
		data = c.paymentProcessor.PackResolveDispute(orderId, *resolver)

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

func (c *Contract) Cancel(itemId common.Hash) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackCancelInvoice(itemId)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil

}

func (c *Contract) Refund(itemId [32]byte) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackClaimExpiredInvoiceRefunds(itemId)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}

func (c *Contract) Release(itemId [32]byte) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	data := c.paymentProcessor.PackReleasePayment(itemId)

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	hash := tx.Hash()

	return &hash, nil
}

func (c *Contract) getDisputeResolution(action MarketplaceAction) uint8 {
	if action == DismissDispute {
		value, _ := bind.Call(c.instance, nil, c.paymentProcessor.PackDISPUTEDISMISSED(), c.paymentProcessor.UnpackDISPUTEDISMISSED)
		return value
	}

	if action == SettleDispute {
		value, _ := bind.Call(c.instance, nil, c.paymentProcessor.PackDISPUTESETTLED(), c.paymentProcessor.UnpackDISPUTESETTLED)
		return value
	}

	return 0
}

func computeOrderId(seller, metaInvoiceId string) (string, error) {
	data := fmt.Sprintf("%s%s", seller, strings.TrimPrefix(metaInvoiceId, "0x"))

	hash, err := utils.Keccak256(data)

	if err != nil {
		return "", err
	}

	return hash.Hex(), err
}
