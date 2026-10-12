package handler

import (
	"math/big"

	"github.com/SapphireDAOO/contract-api/internal/auth"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/notes"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/oraclemanager"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/simplepaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/callback"
	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
	"github.com/SapphireDAOO/contract-api/internal/query"
)

type ContractHandler struct {
	ExplorerURL             string
	Tokens                  config.Tokens
	Callbacks               *callback.Client
	Subgraph                *query.Client
	PaymentProcessor        *intermediatedpaymentprocessor.PaymentProcessor
	PaymentProcessorStorage *paymentprocessorstorage.PaymentProcessorStorage
	SimplePaymentProcessor  *simplepaymentprocessor.SimplePaymentProcessor
	Oracle                  *oraclemanager.OracleManager
	Notes                   *notes.Notes
	BaseUrl                 string

	FeeReceiver *feereceiver.Client
	ChainID     *big.Int
	Auth        *auth.Service
}

func NewContractHandler(c *ContractHandler) *ContractHandler {
	return &ContractHandler{
		ExplorerURL:             c.ExplorerURL,
		Tokens:                  c.Tokens,
		Callbacks:               c.Callbacks,
		Subgraph:                c.Subgraph,
		PaymentProcessor:        c.PaymentProcessor,
		PaymentProcessorStorage: c.PaymentProcessorStorage,
		SimplePaymentProcessor:  c.SimplePaymentProcessor,
		Oracle:                  c.Oracle,
		Notes:                   c.Notes,
		BaseUrl:                 c.BaseUrl,
		FeeReceiver:             c.FeeReceiver,
		ChainID:                 c.ChainID,
		Auth:                    c.Auth,
	}
}
