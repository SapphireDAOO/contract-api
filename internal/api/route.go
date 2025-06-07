package api

import (
	"net/http"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewContractHandler()

	mux.Handle("POST /createInvoice", http.HandlerFunc(handler.CreateInvoice))
	mux.Handle("POST /release", http.HandlerFunc(handler.ReleaseEscrow))

	return mux
}
