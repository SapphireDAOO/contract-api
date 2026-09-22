package main

import (
	"log/slog"
	"os"

	"github.com/SapphireDAOO/contract-api/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
