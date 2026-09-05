package handler

import (
	"errors"
	"net/http"

	"github.com/SapphireDAOO/contract-api/internal/httpx"
)

func (h *ContractHandler) HandleSettlement(w http.ResponseWriter, r *http.Request) {
	exp, err := h.SimplePaymentProcessor.IsSettlementExpired()
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w,
			http.StatusInternalServerError,
			err,
			"failed to determine settlement status",
		)
		return
	}

	if exp {
		httpx.WriteHTTPErrorWithStatus(w,
			http.StatusBadRequest,
			errors.New("settlement window has expired"),
			"settlement time passed",
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}
