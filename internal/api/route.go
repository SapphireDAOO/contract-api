package api

import (
	"fmt"
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

type Router struct {
	mux *http.ServeMux
}

func (r *Router) GET(path string, handler http.Handler) {
	pattern := fmt.Sprintf("GET %s", path)
	r.mux.Handle(pattern, handler)
}
func (r *Router) POST(path string, handler http.Handler) {
	pattern := fmt.Sprintf("POST %s", path)
	r.mux.Handle(pattern, handler)
}

func Route(contract *blockchain.Contract, url string) *http.ServeMux {
	router := Router{mux: http.NewServeMux()}
	handler := NewContractHandler(contract, url)

	router.POST("/create", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.CreateInvoice)))
	router.POST("/release", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.Release)))
	router.POST("/handleDispute", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.HandleDispute)))
	router.POST("/cancel", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.Cancel)))
	router.POST("/refund", middleware.AccessControlMiddleWare(http.HandlerFunc(handler.Refund)))

	return router.mux
}
