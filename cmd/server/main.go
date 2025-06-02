package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := Route()

	server := &http.Server{Addr: port, Handler: mux}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}

}
