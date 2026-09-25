package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/SapphireDAOO/contract-api/internal/httpx"
	"github.com/SapphireDAOO/contract-api/internal/query"
)

func (h *ContractHandler) GetInvoiceData(w http.ResponseWriter, r *http.Request) {
	invoiceId := r.PathValue("invoiceId")
	if invoiceId == "" {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, errors.New("empty invoice id in path"), "Missing invoiceId parameter")
		return
	}

	var (
		data *query.SmartInvoice
		err  error
	)

	data, err = h.Subgraph.GetInvoiceData(invoiceId)

	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, err, "failed to fetch invoice data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}
