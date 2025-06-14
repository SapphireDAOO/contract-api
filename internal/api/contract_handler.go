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

func NewContractHandler(contract *blockchain.Contract) *ContractHandler {
	return &ContractHandler{Contract: contract}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {

	var invoice []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.Contract.CreateInvoice(invoice)

	if err != nil {
		http.Error(w, "Error sending transaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "http://www.google.com", http.StatusFound)
}

func (h *ContractHandler) ReleaseEscrow(w http.ResponseWriter, r *http.Request) {
	var input struct {
		orderId      [32]byte
		resolution   blockchain.MarketplaceAction
		sellersShare *big.Int
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	txHash, err := h.Contract.ReleaseEscrow(input.orderId, input.resolution, input.sellersShare)

	if err != nil {
		http.Error(w, "Error sending transaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"hash":   txHash,
	})

}
