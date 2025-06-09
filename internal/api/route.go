package api

import (
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

func Route(contract *blockchain.Contract) *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewContractHandler(contract)

	mux.Handle("POST /create-invoice", internal.AccessControlMiddleWare(http.HandlerFunc(handler.CreateInvoice)))
	mux.Handle("POST /release", internal.AccessControlMiddleWare(http.HandlerFunc(handler.ReleaseEscrow)))

	return mux
}
