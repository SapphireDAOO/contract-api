package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"

	intermediatedpaymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/IntermediatedPaymentProcessor"
	notes "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/Notes"
	paymentprocessorstorage "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/PaymentProcessorStorage"
	simplepaymentprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/contracts/SimplePaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/callback"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const TX_URL string = "https://sepolia.basescan.org/tx/"

func parseBigInt(field, value string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(strings.TrimSpace(value), 10)
	if !ok {
		return nil, fmt.Errorf("%s must be a base-10 integer, got %q", field, value)
	}
	return n, nil
}

type ContractHandler struct {
	PaymentProcessor        *intermediatedpaymentprocessor.PaymentProcessor
	PaymentProcessorStorage *paymentprocessorstorage.PaymentProcessorStorage
	SimplePaymentProcessor  *simplepaymentprocessor.SimplePaymentProcessor
	Notes                   *notes.Notes
	BaseUrl                 string
}

func NewContractHandler(c *ContractHandler) *ContractHandler {
	return &ContractHandler{
		PaymentProcessor:        c.PaymentProcessor,
		PaymentProcessorStorage: c.PaymentProcessorStorage,
		SimplePaymentProcessor:  c.SimplePaymentProcessor,
		Notes:                   c.Notes,
		BaseUrl:                 c.BaseUrl,
	}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var param []utils.CreateInvoiceParam

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := utils.ValidateCreateInvoiceParams(param); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	invoices := utils.ConvertParam(param)

	if err := utils.ValidateInvoices(invoices); err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
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
		id := invoices[0].InvoiceId
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

	orderId, err := parseBigInt("orderId", input.OrderId)
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.PaymentProcessor.Cancel(orderId)
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

	orderId, err := parseBigInt("orderId", input.OrderId)
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	refundShare, err := parseBigInt("refundShare", input.RefundShare)
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if refundShare.Sign() == 0 {
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

	data, err := h.PaymentProcessor.GetInvoiceData(orderId)
	if err != nil {
		// The refund transaction already succeeded; only the callback is skipped.
		log.Printf("refund callback skipped for orderId %s: fetching invoice data failed: %v", input.OrderId, err)
	} else {
		go callback.SendRefundCallback(input.OrderId,
			data.PaymentToken.String(), data.AmountPaid, refundShare, transactionURL, transactionTimestamp)
	}

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
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, err,
			"error fetching marketplace address")
		return
	}

	orderId, err := parseBigInt("orderId", input.OrderId)
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

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

	orderId, err := parseBigInt("orderId", input.OrderId)
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	res, err := h.PaymentProcessor.Release(orderId)
	if err != nil {
		utils.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	transactionURL := TX_URL + res.TxHash.Hex()
	go callback.
		SendReleaseCallback(input.OrderId, res.PaymentToken.Hex(), res.Seller.Hex(),
			res.SellerAmount, transactionURL, res.BlockTimestamp)

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

	orderId, err := parseBigInt("orderId", input.OrderId)
	if err != nil {
		utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	var sellerShare *big.Int
	if strings.TrimSpace(input.SellerShare) != "" {
		sellerShare, err = parseBigInt("sellerShare", input.SellerShare)
		if err != nil {
			utils.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
			return
		}
	}

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
