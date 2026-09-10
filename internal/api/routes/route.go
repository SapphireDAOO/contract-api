package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/SapphireDAOO/contract-api/internal/api/middleware"
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

func (r *Router) OPTIONS(path string, handler http.HandlerFunc) {
	pattern := fmt.Sprintf("OPTIONS %s", path)
	r.mux.HandleFunc(pattern, handler)
}

const v1 = "/v1"

const (
	feeReceivers     = v1 + "/fee-receivers"
	feeReceiversAuth = feeReceivers + "/authorization"

	notes     = v1 + "/notes"
	notesOpen = notes + "/open"
)

func Route(contractHandler *handler.ContractHandler) *http.ServeMux {
	router := Router{mux: http.NewServeMux()}

	router.GET("/{$}", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	router.POST(v1+"/invoices", middleware.AccessControlMiddleWare(contractHandler.CreateInvoice))
	router.GET(v1+"/invoices/{invoiceId}", contractHandler.GetInvoiceData)
	router.POST(v1+"/invoices/{invoiceId}/release", middleware.AccessControlMiddleWare(contractHandler.Release))
	router.POST(v1+"/invoices/{invoiceId}/cancel", middleware.AccessControlMiddleWare(contractHandler.Cancel))
	router.POST(v1+"/invoices/{invoiceId}/refund", middleware.AccessControlMiddleWare(contractHandler.Refund))
	router.POST(v1+"/invoices/{invoiceId}/disputes", middleware.AccessControlMiddleWare(contractHandler.CreateDispute))
	router.POST(v1+"/invoices/{invoiceId}/disputes/resolution", middleware.AccessControlMiddleWare(contractHandler.HandleDispute))
	router.GET(v1+"/settlements/status", contractHandler.HandleSettlement)
	router.GET(v1+"/exchangeRate", contractHandler.ExchangeRate)

	router.POST(feeReceivers, middleware.CORS(contractHandler.CreateFeeReceivers))
	router.OPTIONS(feeReceivers, middleware.Preflight)
	router.POST(feeReceiversAuth, middleware.CORS(contractHandler.AuthorizeFeeReceivers))
	router.OPTIONS(feeReceiversAuth, middleware.Preflight)

	router.POST(notes, middleware.AccessControlMiddleWare(contractHandler.WriteNote))
	router.POST(notesOpen, middleware.AccessControlMiddleWare(contractHandler.OpenNote))

	return router.mux
}
