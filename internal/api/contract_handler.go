package api

import (
	"encoding/json"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
)

type ContractHandler struct {
	Contract *blockchain.Contract
}

func NewContractHandler() *ContractHandler {
	return &ContractHandler{Contract: blockchain.NewContract()}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {

	var invoice []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam

	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	err := h.Contract.CreateInvoice(invoice)

	if err != nil {
		http.Error(w, "Error sending transaction: "+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"State": "success",
	})

}

func (h *ContractHandler) ReleaseEscrow(w http.ResponseWriter, r *http.Request) {}
