package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func (h *ContractHandler) GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	invoiceId := r.PathValue("invoiceId")

	if invoiceId == "" {
		utils.Error(w, http.StatusBadRequest, errors.New("empty invoice id in path"), "Missing invoiceId parameter")
		return
	}

	var (
		data *query.SmartInvoice
		err  error
	)

	data, err = query.GetInvoiceData(invoiceId)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "failed to fetch invoice data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
