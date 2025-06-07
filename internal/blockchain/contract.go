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

const PAYMENT_PROCESSOR_ADDRESS string = "0xf7B10AD9F2220120af40e974590715856E85E00a"

type MarketplaceAction int

const (
	Pending MarketplaceAction = iota
	Release
	Refund
	DisputeDismissed
)

// relase
// dispute
// dismiss
// settle
// resolve

func NewContract() *Contract {
	address := common.HexToAddress(PAYMENT_PROCESSOR_ADDRESS)
	client := NewClient()

	contract := paymentprocessor.NewPaymentprocessor()

	instance := contract.Instance(client.rpc, address)

	return &Contract{address: address, instance: instance, paymentProcessor: contract, chainId: client.chainId}
}

func (c Contract) CreateInvoice(param []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam) error {

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
