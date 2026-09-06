// Package server starts the API: it resolves configuration, builds the
// contract bindings, and runs the HTTP server alongside the chain listeners.
package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SapphireDAOO/contract-api/internal/api/routes"
	"github.com/SapphireDAOO/contract-api/internal/config"
)

func Run() error {
	if err := loadEnv(); err != nil {
		return err
	}

	cfg, err := config.Load(config.Path())
	if err != nil {
		return err
	}
	log.Printf("Using %s network", cfg.Network)

	deps, err := newDependencies(cfg)
	if err != nil {
		return err
	}
	defer deps.close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	listeners := deps.startListeners(ctx)

	if err := serve(ctx, newHTTPServer(routes.Route(deps.contractHandler(cfg)))); err != nil {
		return err
	}

	listeners.Wait()
	log.Println("Shutdown complete")
	return nil
}
