package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/notes"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/oraclemanager"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/simplepaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/callback"
	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
	"github.com/SapphireDAOO/contract-api/internal/httpx"
	"github.com/SapphireDAOO/contract-api/internal/invoice"
	"github.com/SapphireDAOO/contract-api/internal/query"
)

type ContractHandler struct {
	ExplorerURL             string
	Tokens                  config.Tokens
	Callbacks               *callback.Client
	Subgraph                *query.Client
	PaymentProcessor        *intermediatedpaymentprocessor.PaymentProcessor
	PaymentProcessorStorage *paymentprocessorstorage.PaymentProcessorStorage
	SimplePaymentProcessor  *simplepaymentprocessor.SimplePaymentProcessor
	Oracle                  *oraclemanager.OracleManager
	Notes                   *notes.Notes
	BaseUrl                 string
	// FeeReceiver is nil when services.feeReceiver is not configured, which
	// disables the fee receiver endpoints.
	FeeReceiver *feereceiver.Client
	// ChainID is sent with every fee receiver call so the sidecar refuses a
	// request meant for another network.
	ChainID *big.Int
}

func NewContractHandler(c *ContractHandler) *ContractHandler {
	return &ContractHandler{
		ExplorerURL:             c.ExplorerURL,
		Tokens:                  c.Tokens,
		Callbacks:               c.Callbacks,
		Subgraph:                c.Subgraph,
		PaymentProcessor:        c.PaymentProcessor,
		PaymentProcessorStorage: c.PaymentProcessorStorage,
		SimplePaymentProcessor:  c.SimplePaymentProcessor,
		Oracle:                  c.Oracle,
		Notes:                   c.Notes,
		BaseUrl:                 c.BaseUrl,
		FeeReceiver:             c.FeeReceiver,
		ChainID:                 c.ChainID,
	}
}

func (h *ContractHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var param []invoice.CreateInvoiceParam

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := invoice.ValidateCreateInvoiceParams(param, h.Tokens); err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	invoices, err := invoice.ConvertParam(param, h.Tokens)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	if err := invoice.ValidateInvoices(invoices); err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	intermediatedOperatorAddress, err := h.PaymentProcessorStorage.GetIntermediatedPlatformsOperator()
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, nil, "error fetching marketplace address: "+err.Error())
		return
	}

	if len(invoices) == 1 {
		res, err := h.PaymentProcessor.CreateInvoice(invoices, *intermediatedOperatorAddress)
		if err != nil {
			httpx.WriteMappedRevertError(w, err, "error creating invoice")
			return
		}
		id := invoices[0].InvoiceId
		res.Url = h.BaseUrl + invoice.EncodeIDString(res.Orders[id].OrderId)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			slog.Error("writing the response failed", "error", err)
		}
		return
	}

	res, err := h.PaymentProcessor.CreateInvoices(invoices, *intermediatedOperatorAddress)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "error creating meta invoice")
		return
	}

	res.Url = h.BaseUrl + invoice.EncodeMetaIDString(*res.MetaInvoiceId)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

func (h *ContractHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	orderId, err := parseBigInt("invoiceId", r.PathValue("invoiceId"))
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.PaymentProcessor.Cancel(orderId)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": h.txURL(txHash.Hex()),
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

func (h *ContractHandler) Refund(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RefundShare string `json:"refundShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	invoiceID := r.PathValue("invoiceId")
	if input.RefundShare == "" {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, nil, "refundShare is required")
		return
	}

	orderId, err := parseBigInt("invoiceId", invoiceID)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	refundShare, err := parseBigInt("refundShare", input.RefundShare)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if refundShare.Sign() == 0 {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, errors.New("share can not be zero"), "invalid request body")
		return
	}

	transactionTimestamp := time.Now().UTC().UnixMilli()
	txHash, err := h.PaymentProcessor.Refund(orderId, refundShare)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	transactionURL := h.txURL(txHash.Hex())

	data, err := h.PaymentProcessor.GetInvoiceData(orderId)
	if err != nil {
		// The refund transaction already succeeded; only the callback is skipped.
		slog.Error("refund callback skipped",
			"invoiceId", invoiceID, "reason", "fetching invoice data failed", "error", err)
	} else {
		go h.Callbacks.SendRefundCallback(invoiceID,
			data.PaymentToken.String(), data.AmountPaid, refundShare, transactionURL, transactionTimestamp)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": transactionURL,
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

func (h *ContractHandler) CreateDispute(w http.ResponseWriter, r *http.Request) {
	intermediatedOperatorAddress, err := h.PaymentProcessorStorage.GetIntermediatedPlatformsOperator()
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusInternalServerError, err,
			"error fetching marketplace address")
		return
	}

	orderId, err := parseBigInt("invoiceId", r.PathValue("invoiceId"))
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	txHash, err := h.PaymentProcessor.CreateDispute(orderId, *intermediatedOperatorAddress)

	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": h.txURL(txHash.Hex()),
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

func (h *ContractHandler) Release(w http.ResponseWriter, r *http.Request) {
	invoiceID := r.PathValue("invoiceId")

	orderId, err := parseBigInt("invoiceId", invoiceID)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	res, err := h.PaymentProcessor.Release(orderId)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	transactionURL := h.txURL(res.TxHash.Hex())
	go h.Callbacks.SendReleaseCallback(invoiceID, res.PaymentToken.Hex(), res.Seller.Hex(),
		res.SellerAmount, transactionURL, res.BlockTimestamp)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": transactionURL,
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

func (h *ContractHandler) HandleDispute(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Resolution  blockchain.MarketplaceAction `json:"resolution"`
		SellerShare string                       `json:"sellerShare"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	orderId, err := parseBigInt("invoiceId", r.PathValue("invoiceId"))
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	var sellerShare *big.Int
	if strings.TrimSpace(input.SellerShare) != "" {
		sellerShare, err = parseBigInt("sellerShare", input.SellerShare)
		if err != nil {
			httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
			return
		}
	}

	txHash, err := h.PaymentProcessor.HandleDispute(orderId, input.Resolution, sellerShare)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "Error sending transaction")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{
		"status":         "success",
		"transactionUrl": h.txURL(txHash.Hex()),
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

// txURL links a transaction on the configured explorer. Chains without one
// (a local node) fall back to the bare hash.
func (h *ContractHandler) txURL(txHash string) string {
	return config.Link(h.ExplorerURL, "/tx/", txHash)
}

func parseBigInt(field, value string) (*big.Int, error) {
	n, ok := new(big.Int).SetString(strings.TrimSpace(value), 10)
	if !ok {
		return nil, fmt.Errorf("%s must be a base-10 integer, got %q", field, value)
	}
	return n, nil
}

func (h *ContractHandler) feeReceiverReady(w http.ResponseWriter) bool {
	if h.FeeReceiver == nil || h.ChainID == nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusServiceUnavailable,
			errors.New("fee receiver service is not configured"),
			"fee receivers are unavailable")
		return false
	}
	return true
}
