package routes

import (
	"fmt"
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

type Router struct {
	mux *http.ServeMux
}

func (r *Router) POST(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("POST %s", path)
	r.mux.HandleFunc(pattern, handler)
}

func (r *Router) GET(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("GET %s", path)
	r.mux.Handle(pattern, handler)
}

func Route(contract *blockchain.Contract, url string) *http.ServeMux {
	router := Router{mux: http.NewServeMux()}
	handler := handler.NewContractHandler(contract, url)

	router.POST("/create", middleware.AccessControlMiddleWare(handler.CreateInvoice))
	router.POST("/release", middleware.AccessControlMiddleWare(handler.Release))
	router.POST("/createDispute", middleware.AccessControlMiddleWare(handler.CreateDispute))
	router.POST("/handleDispute", middleware.AccessControlMiddleWare(handler.HandleDispute))
	router.POST("/cancel", middleware.AccessControlMiddleWare(handler.Cancel))
	router.POST("/refund", middleware.AccessControlMiddleWare(handler.Refund))
	router.GET("/invoices/{id}", handler.GetInvoiceData)

	return router.mux
}
