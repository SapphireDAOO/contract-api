package handler

import (
	"errors"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func (h *ContractHandler) HandleSettlement(w http.ResponseWriter, r *http.Request) {
	exp, err := h.SimplePaymentProcessor.IsSettlementExpired()
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w,
			http.StatusInternalServerError,
			err,
			"failed to determine settlement statu",
		)
		return
	}

	if exp {
		utils.WriteHTTPErrorWithStatus(w,
			http.StatusBadRequest,
			errors.New("settlement window has expired"),
			"settlement time passed",
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}
