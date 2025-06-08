package api

import (
	"encoding/json"
	"math/big"
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

	http.Redirect(w, r, "", http.StatusFound)
}

func (h *ContractHandler) ReleaseEscrow(w http.ResponseWriter, r *http.Request) {
	var input struct {
		orderId      [32]byte
		resolution   blockchain.MarketplaceAction
		sellersShare *big.Int
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	txHash, err := h.Contract.ReleaseEscrow(input.orderId, input.resolution, input.sellersShare)

	if err != nil {
		http.Error(w, "Error sending transaction: "+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"hash":   txHash,
	})

}
