package handler

import (
	"encoding/json"
	"net/http"

	middleware "github.com/orgs/SapphireDAOO/contract-api/internal"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const WEB_URL string = "https://sapphire-dao-website-six.vercel.app/checkout/?data="

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://sapphire-dao-website-six.vercel.app")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-KEY")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	contract := blockchain.NewContract(blockchain.NewClient())

	var invoice []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	invoiceKey, err := contract.CreateInvoice(invoice)

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

func CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	middleware.AccessControlMiddleWare(CreateInvoice)
}
