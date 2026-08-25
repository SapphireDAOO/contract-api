package intermediatedpaymentprocessor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	intermediatedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/IntermediatedPaymentProcessor"
)

type InvoiceResponse struct {
	Url           string  `json:"url"`
	MetaInvoiceId *string `json:"-"`
	Orders        map[string]struct {
		Seller  string `json:"seller"`
		OrderId string `json:"orderId"`
	} `json:"orders"`
}

type PaymentProcessor struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *intermediatedprocessor.Intermediatedprocessor
	client   *blockchain.Client
}
