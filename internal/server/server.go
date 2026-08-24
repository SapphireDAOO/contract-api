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

	"github.com/joho/godotenv"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/routes"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	multisig "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/Multisig"
	paymentautomation "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentAutomation"
	paymentprocesor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessor"
	paymentprocessorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessorStorage"
	simplepaymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/SimplePaymentProcessor"
)

const (
	WEB_URL         string        = "https://sapphire-dao-website-six.vercel.app/checkout/?data="
	BASE_URL        string        = "http://localhost:3000/checkout/?data="
	shutdownTimeout time.Duration = 15 * time.Second
)

func Run() error {
	url, err := checkoutURL()
	if err != nil {
		return err
	}

	client, err := blockchain.NewClient()
	if err != nil {
		return err
	}

	pp := paymentprocesor.NewPaymentprocessor(client)
	pps := paymentprocessorstorage.NewPaymentProcessorStorage(client)
	spp := simplepaymentprocessor.NewSimplePaymentProcessor(client)
	ms := multisig.NewMultisig(client)
	pa := paymentautomation.NewPaymentAutomation(client)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	listeners := startListeners(ctx, pp, pps, ms, pa)

	contract := handler.NewContractHandler(
		&handler.ContractHandler{
			PaymentProcessor:        pp,
			PaymentProcessorStorage: pps,
			SimplePaymentProcessor:  spp,
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

func checkoutURL() (string, error) {
	if _, ok := os.LookupEnv("PRODUCTION"); ok {
		return WEB_URL, nil
	}
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}
	return BASE_URL, nil
}

func startListeners(
	ctx context.Context,
	pp *paymentprocesor.PaymentProcessor,
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
