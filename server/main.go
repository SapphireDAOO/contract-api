package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
