package api

import (
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	handler := NewContractHandler()

	mux.Handle("POST /createInvoice", internal.AccessControlMiddleWare(http.HandlerFunc(handler.CreateInvoice)))
	mux.Handle("POST /release", internal.AccessControlMiddleWare(http.HandlerFunc(handler.ReleaseEscrow)))

	return mux
}
