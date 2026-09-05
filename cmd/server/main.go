package main

import (
	"log"
	"os"

	"github.com/SapphireDAOO/contract-api/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Printf("Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
