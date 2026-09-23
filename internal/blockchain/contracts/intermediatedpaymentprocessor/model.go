package intermediatedpaymentprocessor

import (
	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/callback"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type InvoiceResponse struct {
	Url           string  `json:"url"`
	MetaInvoiceId *string `json:"-"`
	Orders        map[string]struct {
		Seller    string `json:"seller"`
		InvoiceId string `json:"invoiceId"`
	} `json:"orders"`
}

type PaymentProcessor struct {
	explorerURL string
	callbacks   *callback.Client
	address     *common.Address
	instance    *bind.BoundContract
	contract    *gen.Intermediatedpaymentprocessor
	client      *blockchain.Client
}
