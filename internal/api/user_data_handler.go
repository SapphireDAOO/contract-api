package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) != 2 || pathParts[0] != "invoices" || pathParts[1] == "" {
		http.Error(w, "Invalid or missing invoice ID", http.StatusBadRequest)
		return
	}

	id := pathParts[1]

	if id == "" {
		http.Error(w, "Missing orderId parameter", http.StatusBadRequest)
		return
	}

	var (
		data *query.SmartInvoice
		err  error
	)

	data, err = query.GetInvoiceData(id)

	if data.Buyer.ID == "" {
		orderId, err := utils.Keccak256(id)

		if err != nil {
			http.Error(w, "failed to generate order ID hash", http.StatusInternalServerError)
			return
		}
		data, err = query.GetInvoiceData(orderId.Hex())
	}

	if err != nil {
		http.Error(w, "failed to fetch invoice data"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
