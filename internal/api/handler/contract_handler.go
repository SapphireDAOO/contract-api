package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocesor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessor"
	paymentprocessorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessorStorage"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const TX_URL string = "https://sepolia.etherscan.io/tx/"

type ContractHandler struct {
	PaymentProcessor        *paymentprocesor.PaymentProcessor
	PaymentProcessorStorage *paymentprocessorstorage.PaymentProcessorStorage
	BaseUrl                 string
}

func NewContractHandler(c *ContractHandler) *ContractHandler {
	return &ContractHandler{
		PaymentProcessor:        c.PaymentProcessor,
		PaymentProcessorStorage: c.PaymentProcessorStorage,
		BaseUrl:                 c.BaseUrl,
	}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoices []advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam

	if err := json.NewDecoder(r.Body).Decode(&invoices); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := utils.ValidateInvoices(invoices); err != nil {
		utils.Error(w, http.StatusBadRequest, err, err.Error())
		return
	}

	if len(invoices) == 0 {
		utils.Error(w, http.StatusBadRequest, nil, "no invoice parameters provided")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	marketplaceAddress, err := h.PaymentProcessorStorage.GetMarketplaceAddress()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, nil, "error fetching marketplace address: "+err.Error())
		return
	}

	if len(invoices) == 1 {
		res, err := h.PaymentProcessor.CreateInvoice(invoices, *marketplaceAddress)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err, "error creating invoice")
			return
		}
		token, err := utils.GenerateToken(res.OrderId)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, err, "token generation failed")
			return
		}

		res.Url = h.BaseUrl + token
		json.NewEncoder(w).Encode(res)
		return
	}

	res, err := h.PaymentProcessor.CreateInvoices(invoices, *marketplaceAddress)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "error creating meta invoice")
		return
	}

	token, err := utils.GenerateToken(*res.MetaInvoiceId)

	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err, "token generation failed")
		return
	}

	res.Url = h.BaseUrl + token
	json.NewEncoder(w).Encode(res)
}

func (h *ContractHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId *big.Int `json:"OrderId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.PaymentProcessor.Cancel(input.OrderId)
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
		OrderId     *big.Int `json:"invoiceId"`
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

	txHash, err := h.PaymentProcessor.Refund(input.OrderId, &input.RefundShare)
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
		OrderId *big.Int `json:"invoiceId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	marketplaceAddress, err := h.PaymentProcessorStorage.GetMarketplaceAddress()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, nil, "error fetching marketplace address: "+err.Error())
		return
	}

	txHash, err := h.PaymentProcessor.CreateDispute(input.OrderId, *marketplaceAddress)

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
		OrderId *big.Int `json:"invoiceId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.PaymentProcessor.Release(input.OrderId)
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
		OrderId     *big.Int                     `json:"orderId"`
		Resolution  blockchain.MarketplaceAction `json:"resolution"`
		SellerShare *big.Int                     `json:"sellerShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.Error(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.PaymentProcessor.HandleDispute(input.OrderId, input.Resolution, input.SellerShare)
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
