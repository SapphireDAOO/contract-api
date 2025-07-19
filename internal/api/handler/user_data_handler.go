package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(pathParts) != 2 || pathParts[0] != "invoices" || pathParts[1] == "" {
		utils.Error(w, http.StatusBadRequest, errors.New("invalid URL path structure"), "Invalid or missing invoice ID")
		return
	}

	id := pathParts[1]

	if id == "" {
		utils.Error(w, http.StatusBadRequest, errors.New("empty invoice id in path"), "Missing orderId parameter")
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
			utils.Error(w, http.StatusInternalServerError, err, "failed to generate order ID hash")
			return
		}
		data, err = query.GetInvoiceData(orderId.Hex())
	}

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "failed to fetch invoice data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
