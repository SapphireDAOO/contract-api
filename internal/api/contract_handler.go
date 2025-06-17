package api

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const WEB_URL string = "https://sapphire-dao-website-six.vercel.app/checkout/?data="

type ContractHandler struct {
	Contract *blockchain.Contract
}

func NewContractHandler(contract *blockchain.Contract) *ContractHandler {
	return &ContractHandler{Contract: contract}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	invoiceKey, err := h.Contract.CreateInvoice(invoice)

	if err != nil {
		http.Error(w, "Error sending transaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(invoiceKey)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"url": WEB_URL + token,
	})
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
