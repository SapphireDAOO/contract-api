package blockchain

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
)

type Contract struct {
	address          *common.Address
	instance         *bind.BoundContract
	paymentProcessor *paymentprocessor.Paymentprocessor
	client           Client
}

const PAYMENT_PROCESSOR_ADDRESS string = "0x7f98897854Ab2B2370C3BF2694514f05375Ea606"

func NewContract(client *Client) *Contract {
	address := common.HexToAddress(PAYMENT_PROCESSOR_ADDRESS)

	contract := paymentprocessor.NewPaymentprocessor()

	instance := contract.Instance(client.eth, address)

	return &Contract{address: &address, instance: instance, paymentProcessor: contract, client: *client}
}

func (c *Contract) CreateInvoice(param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam) (*common.Hash, error) {
	if len(param) == 0 {
		return nil, errors.New("parameter cannot be empty")
	}

	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	if len(param) == 1 {
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

		return &invoiceKey, nil
	} else {

		data := c.paymentProcessor.PackCreateMetaInvoice(param[0].Buyer, param)

		tx, err := bind.Transact(c.instance, auth, data)

		if err != nil {
			return nil, err
		}

		receipt, err := bind.WaitMined(context.Background(), c.client.eth, tx.Hash())

		if err != nil {
			return nil, err
		}

		invoiceKey := receipt.Logs[2].Topics[1]

		return &invoiceKey, nil
	}

}

func (c *Contract) ReleaseEscrow(orderId [32]byte, action MarketplaceAction, sellersShare *big.Int) (*common.Hash, error) {
	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	var data []byte

	switch action {
	case Release:
		data = c.paymentProcessor.PackReleasePayment(orderId)

	case SettleDispute:
		data = c.paymentProcessor.PackHandleDispute(orderId, c.getDisputeResolution(SettleDispute), sellersShare)

	case DismissDispute:
		data = c.paymentProcessor.PackHandleDispute(orderId, c.getDisputeResolution(DismissDispute), nil)

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
