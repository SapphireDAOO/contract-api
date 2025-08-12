package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const TX_URL string = "https://sepolia.etherscan.io/tx/"

type ContractHandler struct {
	Contract *blockchain.Contract
	baseUrl  string
}

func NewContractHandler(contract *blockchain.Contract, baseUrl string) *ContractHandler {
	return &ContractHandler{Contract: contract, baseUrl: baseUrl}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoices []paymentprocessor.IAdvancedPaymentProcessorInvoiceCreationParam

	if err := json.NewDecoder(r.Body).Decode(&invoices); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if len(invoices) == 0 {
		utils.Error(w, http.StatusBadRequest, nil, "no invoice parameters provided")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if len(invoices) == 1 {
		res, err := h.Contract.CreateInvoice(invoices)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err, "error creating invoice")
			return
		}
		token, err := utils.GenerateToken(res.InvoiceId)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err, "token generation failed")
			return
		}

		res.Url = h.baseUrl + token
		json.NewEncoder(w).Encode(res)
		return
	}

	res, err := h.Contract.CreateInvoices(invoices)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "error creating meta invoice")
		return
	}

	token, err := utils.GenerateToken(*res.Key)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "token generation failed")
		return
	}

	res.Url = h.baseUrl + token
	json.NewEncoder(w).Encode(res)
}

func (h *ContractHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvoiceId common.Hash `json:"invoiceId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.Cancel(input.InvoiceId)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"transaction url": TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) Refund(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvoiceId   [32]byte `json:"invoiceId"`
		RefundShare big.Int  `json:"refundShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if input.RefundShare.Cmp(big.NewInt(0)) == 0 {
		utils.Error(w, http.StatusBadRequest, errors.New("share can not be zero"), "invalid request body")
		return
	}

	txHash, err := h.Contract.Refund(input.InvoiceId, &input.RefundShare)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"transaction url": TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) CreateDispute(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvoiceId [32]byte `json:"invoiceId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.CreateDispute(input.InvoiceId)

	if err != nil {
		fmt.Println(err)
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"transaction url": TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) Release(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvoiceId   [32]byte `json:"invoiceId"`
		SellerShare big.Int  `json:"sellerShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.Release(input.InvoiceId, &input.SellerShare)
	if err != nil {
		fmt.Println(err)
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"transaction url": TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) HandleDispute(w http.ResponseWriter, r *http.Request) {
	var input struct {
		InvoiceId   [32]byte                     `json:"invoiceId"`
		Resolution  blockchain.MarketplaceAction `json:"resolution"`
		SellerShare *big.Int                     `json:"sellerShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.Contract.HandleDispute(input.InvoiceId, input.Resolution, input.SellerShare)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"transaction url": TX_URL + txHash.Hex(),
	})
}
