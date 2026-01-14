package handler

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"time"

	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	paymentprocesor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessor"
	paymentprocessorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessorStorage"
	"github.com/orgs/SapphireDAOO/contract-api/internal/callback"
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
	var param []utils.CreateInvoiceParam

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	invoices := utils.ConvertParam(param)

	if err := utils.ValidateInvoices(invoices); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	if len(invoices) == 0 {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, nil, "no invoice parameters provided")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	marketplaceAddress, err := h.PaymentProcessorStorage.GetMarketplaceAddress()
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, nil, "error fetching marketplace address: "+err.Error())
		return
	}

	if len(invoices) == 1 {
		res, err := h.PaymentProcessor.CreateInvoice(invoices, *marketplaceAddress)
		if err != nil {
			utils.WriteMappedRevertError(w, err, "error creating invoice")
			return
		}
		id := invoices[0].OrderId
		token, err := utils.GenerateToken(res.Orders[id].OrderId)
		if err != nil {
			utils.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, err, "token generation failed")
			return
		}

		res.Url = h.BaseUrl + token
		json.NewEncoder(w).Encode(res)
		return
	}

	res, err := h.PaymentProcessor.CreateInvoices(invoices, *marketplaceAddress)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "error creating meta invoice")
		return
	}

	token, err := utils.GenerateToken(*res.MetaInvoiceId)

	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, err, "token generation failed")
		return
	}

	res.Url = h.BaseUrl + token
	json.NewEncoder(w).Encode(res)
}

func (h *ContractHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId string `json:"orderId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	n, _ := new(big.Int).SetString(input.OrderId, 10)
	txHash, err := h.PaymentProcessor.Cancel(n)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) Refund(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId     string `json:"orderId"`
		RefundShare string `json:"refundShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if input.OrderId == "" || input.RefundShare == "" {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, nil, "orderId and refundShare is required")
		return
	}

	orderId, _ := new(big.Int).SetString(input.OrderId, 10)
	refundShare, _ := new(big.Int).SetString(input.RefundShare, 10)

	if refundShare.Cmp(big.NewInt(0)) == 0 {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, errors.New("share can not be zero"), "invalid request body")
		return
	}

	transactionTimestamp := time.Now().UTC().UnixMilli()
	txHash, err := h.PaymentProcessor.Refund(orderId, refundShare)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	transactionURL := TX_URL + txHash.Hex()
	go callback.SendRefundCallback(input.OrderId, "", nil, refundShare, transactionURL, transactionTimestamp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": transactionURL,
	})
}

func (h *ContractHandler) CreateDispute(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId string `json:"orderId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	marketplaceAddress, err := h.PaymentProcessorStorage.GetMarketplaceAddress()
	if err != nil && marketplaceAddress != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, nil,
			"error fetching marketplace address: "+err.Error())
		return
	}

	orderId, _ := new(big.Int).SetString(input.OrderId, 10)
	txHash, err := h.PaymentProcessor.CreateDispute(orderId, *marketplaceAddress)

	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": TX_URL + txHash.Hex(),
	})
}

func (h *ContractHandler) Release(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId string `json:"orderId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	orderId, _ := new(big.Int).SetString(input.OrderId, 10)
	transactionTimestamp := time.Now().UTC().UnixMilli()
	txHash, err := h.PaymentProcessor.Release(orderId)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	transactionURL := TX_URL + txHash.Hex()
	go callback.
		SendReleaseCallback(input.OrderId, nil, transactionURL, transactionTimestamp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": transactionURL,
	})
}

func (h *ContractHandler) HandleDispute(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OrderId     string                       `json:"orderId"`
		Resolution  blockchain.MarketplaceAction `json:"resolution"`
		SellerShare string                       `json:"sellerShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	orderId, _ := new(big.Int).SetString(input.OrderId, 10)
	sellerShare, _ := new(big.Int).SetString(input.SellerShare, 10)
	txHash, err := h.PaymentProcessor.HandleDispute(orderId, input.Resolution, sellerShare)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": TX_URL + txHash.Hex(),
	})
}
