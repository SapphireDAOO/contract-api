package blockchain

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
)

type Contract struct {
	address          common.Address
	instance         *bind.BoundContract
	paymentProcessor *paymentprocessor.Paymentprocessor
	chainId          *big.Int
}

const PAYMENT_PROCESSOR_ADDRESS string = "0x05e6989466427b0bc350DDD897538D8867b1bB58"

func NewContract(client *Client) *Contract {
	address := common.HexToAddress(PAYMENT_PROCESSOR_ADDRESS)

	contract := paymentprocessor.NewPaymentprocessor()

	instance := contract.Instance(client.rpc, address)

	return &Contract{address: address, instance: instance, paymentProcessor: contract, chainId: client.chainId}
}

func (c *Contract) CreateInvoice(param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam) error {

	if len(param) == 0 {
		return errors.New("parameter cannot be empty")
	}

	auth, err := auth(c.chainId)

	if err != nil {
		return err
	}

	if len(param) == 1 {
		_, err := bind.Transact(c.instance, auth, c.paymentProcessor.PackCreateSingleInvoice(param[0]))
		if err != nil {
			return err
		}
	} else {
		data := c.paymentProcessor.PackCreateMetaInvoice(param[0].Buyer, param)
		_, err := bind.Transact(c.instance, auth, data)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Contract) ReleaseEscrow(orderId [32]byte, action MarketplaceAction, sellersShare *big.Int) (*common.Hash, error) {
	auth, err := auth(c.chainId)

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
