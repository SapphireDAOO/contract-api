package server

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"github.com/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/multisig"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/notes"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/oraclemanager"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/paymentautomation"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/simplepaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/callback"
	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/SapphireDAOO/contract-api/internal/discord"
	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
	"github.com/SapphireDAOO/contract-api/internal/query"
)

// dependencies are the long-lived collaborators, built once at startup from
// the resolved config and shared by the HTTP handlers and the listeners.
type dependencies struct {
	client      *blockchain.Client
	notifier    *discord.Client
	callbacks   *callback.Client
	subgraph    *query.Client
	feeReceiver *feereceiver.Client

	oracle                  *oraclemanager.OracleManager
	paymentProcessor        *intermediatedpaymentprocessor.PaymentProcessor
	paymentProcessorStorage *paymentprocessorstorage.PaymentProcessorStorage
	simplePaymentProcessor  *simplepaymentprocessor.SimplePaymentProcessor
	multisig                *multisig.Multisig
	paymentAutomation       *paymentautomation.PaymentAutomation
	notes                   *notes.Notes
}

func newDependencies(cfg *config.Config) (*dependencies, error) {
	client, err := blockchain.NewClient(cfg.RPC, cfg.SignerKey)
	if err != nil {
		return nil, err
	}

	addresses := cfg.Contracts.Addresses()

	notifier := discord.NewClient(cfg.URLs.DiscordWebhook)
	callbacks := callback.NewClient(cfg.URLs.Callback, os.Getenv("API_KEY"), cfg.Tokens)

	deps := &dependencies{
		client:    client,
		notifier:  notifier,
		callbacks: callbacks,
		subgraph:  query.NewClient(cfg.URLs.Subgraph),

		paymentProcessor: intermediatedpaymentprocessor.NewPaymentprocessor(
			client, addresses.PaymentProcessor, cfg.URLs.Explorer, callbacks),

		paymentProcessorStorage: paymentprocessorstorage.NewPaymentProcessorStorage(
			client, addresses.PaymentProcessorStorage, cfg.URLs.Explorer, notifier),

		simplePaymentProcessor: simplepaymentprocessor.NewSimplePaymentProcessor(
			client, addresses.SimplePaymentProcessor),

		multisig: multisig.NewMultisig(client, addresses.Multisig, multisig.Peers{
			PaymentProcessor:        addresses.PaymentProcessor,
			SimplePaymentProcessor:  addresses.SimplePaymentProcessor,
			PaymentProcessorStorage: addresses.PaymentProcessorStorage,
		}, cfg.URLs, notifier),

		paymentAutomation: paymentautomation.NewPaymentAutomation(
			client, addresses.PaymentAutomation),
		notes: notes.NewNotes(client, addresses.Notes),
	}

	if cfg.Services.FeeReceiver != "" {
		feeReceiverClient, err := feereceiver.NewClient(cfg.Services.FeeReceiver)
		if err != nil {
			return nil, err
		}
		deps.feeReceiver = feeReceiverClient
	} else {
		log.Print("Fee receivers disabled: services.feeReceiver is not configured")
	}

	if (addresses.OracleManager != common.Address{}) {
		deps.oracle = oraclemanager.NewOracleManager(client, addresses.OracleManager)
	} else {
		log.Print("Exchange rates disabled: contracts.oracleManager is not configured")
	}

	return deps, nil
}

func (d *dependencies) contractHandler(cfg *config.Config) *handler.ContractHandler {
	return handler.NewContractHandler(
		&handler.ContractHandler{
			PaymentProcessor:        d.paymentProcessor,
			PaymentProcessorStorage: d.paymentProcessorStorage,
			SimplePaymentProcessor:  d.simplePaymentProcessor,
			Oracle:                  d.oracle,
			Notes:                   d.notes,
			BaseUrl:                 cfg.URLs.Checkout,
			ExplorerURL:             cfg.URLs.Explorer,
			Tokens:                  cfg.Tokens,
			Callbacks:               d.callbacks,
			Subgraph:                d.subgraph,
			FeeReceiver:             d.feeReceiver,
			ChainID:                 d.client.ChainId,
		},
	)
}

func (d *dependencies) close() {
	if d.feeReceiver != nil {
		if err := d.feeReceiver.Close(); err != nil {
			log.Printf("fee receiver connection close error: %v", err)
		}
	}
}
