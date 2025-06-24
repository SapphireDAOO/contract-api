package api

import (
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

func Route(contract *blockchain.Contract, url string) *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewContractHandler(contract, url)

	mux.Handle("POST /create", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.CreateInvoice)))
	mux.Handle("POST /release", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.ReleaseEscrow)))
	mux.Handle("GET /invoices/{id}", http.HandlerFunc(GetInvoiceData))

	return mux
}
