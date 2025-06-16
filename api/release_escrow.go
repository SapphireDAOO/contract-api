package handler

import (
	"encoding/json"
	"math/big"
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

func ReleaseEscrow(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://sapphire-dao-website-six.vercel.app")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-KEY")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	contract := blockchain.NewContract(blockchain.NewClient())

	var input struct {
		orderId      [32]byte
		resolution   blockchain.MarketplaceAction
		sellersShare *big.Int
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	txHash, err := contract.ReleaseEscrow(input.orderId, input.resolution, input.sellersShare)
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

func ReleaseEscrowHandler(w http.ResponseWriter, r *http.Request) {
	middleware.AccessControlMiddleWare(CreateInvoice)
}
