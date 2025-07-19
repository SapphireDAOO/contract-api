package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/routes"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

const (
	WEB_URL  string = "https://sapphire-dao-website-six.vercel.app/checkout/?data="
	BASE_URL string = "http://localhost:3000/checkout/?data="
)

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

	contract := blockchain.NewContract(blockchain.NewClient())

	mux := routes.Route(contract, url)

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("Server running at port %s", addr)

	return server.ListenAndServe()
}
