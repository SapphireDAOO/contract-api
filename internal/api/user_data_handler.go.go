package api

import (
	"encoding/json"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")

	if address == "" {
		http.Error(w, "Missing address parameter", http.StatusBadRequest)
		return
	}

	data, err := query.GetInvoice(address)

	if err != nil {
		http.Error(w, "failed to fetch invoice data"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
