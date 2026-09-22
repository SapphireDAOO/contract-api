package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort     = "8080"
	requestTimeout  = 30 * time.Second
	shutdownTimeout = 15 * time.Second
)

func newHTTPServer(mux http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	return &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  requestTimeout,
		WriteTimeout: requestTimeout,
	}
}

// serve runs srv until it fails or ctx is cancelled, then drains it within
// shutdownTimeout. A serve error is returned; a clean shutdown is not.
func serve(ctx context.Context, srv *http.Server) error {
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("http server listening", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
		close(serverErr)
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		slog.Info("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("http server shutdown failed", "error", err)
	}
	return nil
}
