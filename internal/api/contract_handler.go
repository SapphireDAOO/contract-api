package api

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const TX_URL string = "https://amoy.polygonscan.com/tx/"

type ContractHandler struct {
	Contract *blockchain.Contract
	baseUrl  string
}

func NewContractHandler(contract *blockchain.Contract, baseUrl string) *ContractHandler {
	return &ContractHandler{Contract: contract, baseUrl: baseUrl}
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
		"url": h.baseUrl + token,
	})
}

func (h *ContractHandler) ReleaseEscrow(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId      string                       `json:"orderId"`
		Resolution   blockchain.MarketplaceAction `json:"resolution"`
		SellersShare *big.Int                     `json:"sellersShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var id common.Hash

	if len(input.OrderId) == 66 && strings.HasPrefix(input.OrderId, "0x") {
		id = common.HexToHash(input.OrderId)
	} else {
		hashed, err := utils.Keccak256(input.OrderId)
		if err != nil {
			http.Error(w, "failed to generate order ID hash", http.StatusInternalServerError)
			return
		}
		id = *hashed
	}

	txHash, err := h.Contract.ReleaseEscrow(id, input.Resolution, input.SellersShare)
	if err != nil {
		http.Error(w, "Error sending transaction: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": "success",
		"hash":   TX_URL + txHash.Hex(),
	})

}
