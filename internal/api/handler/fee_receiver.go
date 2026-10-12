package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/intermediatedpaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/blockchain/contracts/simplepaymentprocessor"
	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
	"github.com/SapphireDAOO/contract-api/internal/httpx"
	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
)

func (h *ContractHandler) CreateFeeReceivers(w http.ResponseWriter, r *http.Request) {
	if !h.feeReceiverReady(w) {
		return
	}

	var input struct {
		InvoiceID    string `json:"invoiceId"`
		Processor    string `json:"processor"`
		Kind         string `json:"kind"`
		Quantity     int32  `json:"quantity"`
		PaymentToken string `json:"paymentToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	processor, err := feereceiver.ParseProcessorKind(input.Processor)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	kind, err := feereceiver.ParseInvoiceKind(input.Kind)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	invoiceID, err := parseBigInt("invoiceId", input.InvoiceID)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	if processor == pb.ProcessorKind_PROCESSOR_KIND_SIMPLE && kind == pb.InvoiceKind_INVOICE_KIND_META {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, ErrMetaNotIntermediated, ErrMetaNotIntermediated.Error())
		return
	}

	// One receiver covers a single invoice; a meta invoice asks for one per
	// sub-invoice.
	if input.Quantity == 0 {
		input.Quantity = 1
	}

	paymentToken, err := feereceiver.ResolveFeeToken(h.Tokens, input.PaymentToken)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	allowed, err := h.invoiceAllowsFeeReceivers(r.Context(), processor, kind, invoiceID)
	if err != nil {
		slog.Error("invoice state read failed", "invoiceId", invoiceID.String(), "error", err)
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadGateway, nil, "unable to verify the invoice")
		return
	}

	if !allowed {
		err := fmt.Errorf("invoice %s %w", invoiceID, ErrNoFeeReceiverState)
		httpx.WriteHTTPErrorWithStatus(w, http.StatusConflict, err, "invoice cannot take fee receivers")
		return
	}

	res, err := h.FeeReceiver.Create(r.Context(), &pb.CreateFeeReceiversRequest{
		ProcessorKind: processor,
		InvoiceKind:   kind,
		ChainId:       h.ChainID.Uint64(),
		InvoiceId:     invoiceID.String(),
		PaymentToken:  paymentToken,
		Quantity:      input.Quantity,
	})
	if err != nil {
		httpx.WriteGRPCError(w, err, "error creating fee receivers")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]any{
		"status":              "success",
		"feeReceivers":        res.GetAddresses(),
		"signature":           res.GetSignature(),
		"ephemeralPublicKeys": res.GetEphemeralPublicKey(),
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

func (h *ContractHandler) invoiceAllowsFeeReceivers(ctx context.Context, processor pb.ProcessorKind,
	kind pb.InvoiceKind, invoiceID *big.Int) (bool, error) {
	switch {
	case processor == pb.ProcessorKind_PROCESSOR_KIND_SIMPLE:
		invoice, err := h.SimplePaymentProcessor.GetInvoiceData(ctx, invoiceID)
		if err != nil {
			return false, err
		}
		return invoice.State == simplepaymentprocessor.StatePaid, nil

	case kind == pb.InvoiceKind_INVOICE_KIND_META:
		subInvoiceIDs, err := h.PaymentProcessor.SubInvoiceIDs(ctx, invoiceID)
		if err != nil {
			return false, err
		}
		for _, id := range subInvoiceIDs {
			invoice, err := h.PaymentProcessor.GetInvoiceData(id)
			if err != nil {
				return false, err
			}
			if invoice.State == intermediatedpaymentprocessor.StatePaid {
				return true, nil
			}
		}
		return false, nil

	default:
		invoice, err := h.PaymentProcessor.GetInvoiceData(invoiceID)
		if err != nil {
			return false, err
		}
		return invoice.State == intermediatedpaymentprocessor.StateCreated, nil
	}
}
