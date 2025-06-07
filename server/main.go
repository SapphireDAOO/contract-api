package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := api.Route()

	server := &http.Server{Addr: port, Handler: mux}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}

}
