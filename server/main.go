package main

import (
	"context"
	"errors"
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
	paymentprocesor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessor"
	paymentprocessorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessorStorage"
)

const (
	WEB_URL         string        = "https://sapphire-dao-website-six.vercel.app/checkout/?data="
	BASE_URL        string        = "http://localhost:3000/checkout/?data="
	shutdownTimeout time.Duration = 15 * time.Second
)

func main() {
	if err := run(); err != nil {
		log.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var url string
	if _, ok := os.LookupEnv("PRODUCTION"); !ok {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("Error loading .env file")
		}
		url = BASE_URL
	} else {
		url = WEB_URL
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	client, err := blockchain.NewClient()
	if err != nil {
		return err
	}

	pp := paymentprocesor.NewPaymentprocessor(client)
	pps := paymentprocessorstorage.NewPaymentProcessorStorage(client)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var listeners sync.WaitGroup
	listeners.Add(2)
	go func() {
		defer listeners.Done()
		pp.ListenToPaymentReceivedEvent(ctx)
	}()
	go func() {
		defer listeners.Done()
		pp.ListenToReleaseEvent(ctx)
	}()

	contract := handler.NewContractHandler(
		&handler.ContractHandler{
			PaymentProcessor:        pp,
			PaymentProcessorStorage: pps,
			BaseUrl:                 url,
		},
	)

	mux := routes.Route(contract)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("Server running at port %s", addr)
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
