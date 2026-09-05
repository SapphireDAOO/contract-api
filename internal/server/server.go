package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/SapphireDAOO/contract-api/internal/api/routes"
	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/multisig"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/notes"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/paymentautomation"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/simplepaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/joho/godotenv"
)

const (
	WEB_URL         string        = "https://sapphire-dao-website-six.vercel.app/checkout/?data="
	BASE_URL        string        = "http://localhost:3000/checkout/?data="
	shutdownTimeout time.Duration = 15 * time.Second
)

func Run() error {
	if err := loadEnv(); err != nil {
		return err
	}

	cfg, err := config.Load(config.Path())
	if err != nil {
		return err
	}
	addresses := cfg.Contracts.Addresses()
	log.Printf("Using %s network", cfg.Network)

	url := checkoutURL()

	client, err := blockchain.NewClient(cfg.RPC)
	if err != nil {
		return err
	}

	pp := intermediatedpaymentprocessor.NewPaymentprocessor(client, addresses.PaymentProcessor)
	pps := paymentprocessorstorage.NewPaymentProcessorStorage(client, addresses.PaymentProcessorStorage)
	spp := simplepaymentprocessor.NewSimplePaymentProcessor(client, addresses.SimplePaymentProcessor)
	ms := multisig.NewMultisig(client, addresses.Multisig, multisig.Peers{
		PaymentProcessor:        addresses.PaymentProcessor,
		SimplePaymentProcessor:  addresses.SimplePaymentProcessor,
		PaymentProcessorStorage: addresses.PaymentProcessorStorage,
	})
	pa := paymentautomation.NewPaymentAutomation(client, addresses.PaymentAutomation)
	notes := notes.NewNotes(client, addresses.Notes)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	listeners := startListeners(ctx, pp, pps, ms, pa)

	contract := handler.NewContractHandler(
		&handler.ContractHandler{
			PaymentProcessor:        pp,
			PaymentProcessorStorage: pps,
			SimplePaymentProcessor:  spp,
			Notes:                   notes,
			BaseUrl:                 url,
		},
	)

	server := newHTTPServer(routes.Route(contract))

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("Server running at port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		log.Println("Shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	listeners.Wait()
	log.Println("Shutdown complete")
	return nil
}

// loadEnv reads .env for local runs. In production the environment is supplied
// by the container, so there is no file to read.
func loadEnv() error {
	if _, ok := os.LookupEnv("PRODUCTION"); ok {
		return nil
	}
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func checkoutURL() string {
	if _, ok := os.LookupEnv("PRODUCTION"); ok {
		return WEB_URL
	}
	return BASE_URL
}

func startListeners(
	ctx context.Context,
	pp *intermediatedpaymentprocessor.PaymentProcessor,
	pps *paymentprocessorstorage.PaymentProcessorStorage,
	ms *multisig.Multisig,
	pa *paymentautomation.PaymentAutomation,
) *sync.WaitGroup {
	var listeners sync.WaitGroup

	for _, listen := range []func(context.Context){
		pp.ListenToPaymentReceivedEvent,
		pp.ListenToReleaseEvent,
		ms.ListenToEvents,
		pps.ListenToPauseEvents,
		pa.PollDueTasks,
	} {
		listeners.Add(1)
		go func() {
			defer listeners.Done()
			listen(ctx)
		}()
	}

	return &listeners
}

func newHTTPServer(mux http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
