package paymentprocesor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
)

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

type PaymentProcessor struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *advancedprocessor.Advancedprocessor
	client   *blockchain.Client
}
