package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func (h *ContractHandler) GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		utils.Error(w, http.StatusBadRequest, errors.New("empty invoice id in path"), "Missing orderId parameter")
		return
	}

	var (
		data *query.SmartInvoice
		err  error
	)

	invoiceId, err := handleId(id)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "failed to generate order ID hash")
		return
	}

	data, err = query.GetInvoiceData(invoiceId.Hex())

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "failed to fetch invoice data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
