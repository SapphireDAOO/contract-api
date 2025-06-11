package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	contract := blockchain.NewContract(blockchain.NewClient())
	mux := api.Route(contract)

	server := &http.Server{Addr: port, Handler: mux, ReadTimeout: 10 * time.Second, WriteTimeout: 15 * time.Second}

	log.Printf("Listening on port:%s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}

}
