package api

import (
	"encoding/json"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("orderId")

	if id == "" {
		http.Error(w, "Missing orderId parameter", http.StatusBadRequest)
		return
	}

	orderId, err := utils.Keccak256(id)

	if err != nil {
		http.Error(w, "failed to generate order ID hash", http.StatusInternalServerError)
		return
	}

	data, err := query.GetInvoiceData(orderId.Hex())

	if err != nil {
		http.Error(w, "failed to fetch invoice data"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
