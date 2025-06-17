package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
)

func GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	firstParam := r.URL.Query().Get("first")
	skipParam := r.URL.Query().Get("skip")

	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	first := 10
	skip := 0

	if val, err := strconv.Atoi(firstParam); err == nil {
		first = val
	}

	if val, err := strconv.Atoi(skipParam); err == nil {
		skip = val
	}

	data, err := query.GetInvoiceData(id, first, skip)

	if err != nil {
		http.Error(w, "failed to fetch invoice data"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
