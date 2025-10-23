package routes

import (
	"fmt"
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/api/handler"
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

func Route(h *handler.ContractHandler) *http.ServeMux {
	router := Router{mux: http.NewServeMux()}
	contractHandler := handler.NewContractHandler(h)

	router.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))
	router.POST("/create", middleware.AccessControlMiddleWare(contractHandler.CreateInvoice))
	router.POST("/release", middleware.AccessControlMiddleWare(contractHandler.Release))
	router.POST("/createDispute", middleware.AccessControlMiddleWare(contractHandler.CreateDispute))
	router.POST("/handleDispute", middleware.AccessControlMiddleWare(contractHandler.HandleDispute))
	router.POST("/cancel", middleware.AccessControlMiddleWare(contractHandler.Cancel))
	router.POST("/refund", middleware.AccessControlMiddleWare(contractHandler.Refund))
	router.GET("/invoices/{orderId}", contractHandler.GetInvoiceData)

	return router.mux
}
