package paymentprocesor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
)

type InvoiceResponse struct {
	Url           string  `json:"url"`
	MetaInvoiceId *string `json:"-"`
	Orders        map[string]struct {
		Seller  string `json:"seller"`
		OrderId string `json:"orderId"`
	} `json:"orders"`
}

// type SingleInvoiceResponse struct {
// 	Url     string `json:"url"`
// 	OrderId string `json:"orderId"`
// }

type PaymentProcessor struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *advancedprocessor.Advancedprocessor
	client   *blockchain.Client
}
