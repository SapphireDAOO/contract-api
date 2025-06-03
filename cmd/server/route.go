package main

import (
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/api"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /createInvoice", http.HandlerFunc(api.CreateInvoice))
	mux.Handle("POST /release", http.HandlerFunc(api.ReleaseEscrow))

	return mux
}
