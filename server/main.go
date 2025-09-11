package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/routes"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocesor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessor"
	paymentprocessorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessorStorage"
)

const (
	WEB_URL  string = "https://sapphire-dao-website-six.vercel.app/checkout/?data="
	BASE_URL string = "http://localhost:3000/checkout/?data="
)

func main() {

	if err := run(); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
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

	// go pp.ListenToReleaseEvent()

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

	log.Printf("Server running at port %s", addr)

	return server.ListenAndServe()
}
