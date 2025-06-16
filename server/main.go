package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	handler "github.com/orgs/SapphireDAOO/contract-api/api"
	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	mux := http.NewServeMux()

	mux.Handle("POST /create-invoice", middleware.AccessControlMiddleWare(handler.CreateInvoiceHandler))
	mux.Handle("POST /release", middleware.AccessControlMiddleWare(handler.ReleaseEscrowHandler))
	mux.Handle("GET /user-data", http.HandlerFunc(handler.GetUserDataHandler))

	server := &http.Server{Addr: port, Handler: mux, ReadTimeout: 10 * time.Second, WriteTimeout: 15 * time.Second}

	log.Printf("Listening on port %s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}

}
