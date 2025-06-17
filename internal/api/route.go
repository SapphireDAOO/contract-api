package api

import (
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

func Route(contract *blockchain.Contract) *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewContractHandler(contract)

	mux.Handle("POST /create-invoice", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.CreateInvoice)))
	mux.Handle("POST /release", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.ReleaseEscrow)))
	mux.Handle("GET /invoice-data", http.HandlerFunc(GetInvoiceData))

	return mux
}
