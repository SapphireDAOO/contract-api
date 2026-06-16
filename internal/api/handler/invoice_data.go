package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func (h *ContractHandler) GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")
	if orderId == "" {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, errors.New("empty invoice id in path"), "Missing invoiceId parameter")
		return
	}

	var (
		data *query.SmartInvoice
		err  error
	)

	data, err = query.GetInvoiceData(orderId)

	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, err, "failed to fetch invoice data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
