package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

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
