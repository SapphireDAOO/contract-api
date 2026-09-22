package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
	"github.com/SapphireDAOO/contract-api/internal/httpx"
	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
)

// Fee receivers are issued in two steps the caller drives separately, so a
// cut-off between them strands nothing it could have recorded:
//
// POST /v1/fee-receivers derives the addresses, delegates each to the relayer
// and approves the sweeper, returning only the ephemeral public keys. Those
// keys are the sole way to recover the addresses, so the caller must persist
// them before doing anything else.
//
// POST /v1/fee-receivers/authorization hands them back and receives the
// addresses with the fee signer's signature, to pass on-chain as the
// processor's _feeReceiver(s) / _data arguments.

// CreateFeeReceivers derives fee receivers and puts each on-chain. It returns
// the ephemeral public keys, which the caller must store: they cannot be
// recovered, and the sponsored gas is spent whether or not they are kept.
func (h *ContractHandler) CreateFeeReceivers(w http.ResponseWriter, r *http.Request) {
	if !h.feeReceiverReady(w) {
		return
	}

	var input struct {
		Quantity     int32  `json:"quantity"`
		Processor    string `json:"processor"`
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

	res, err := h.FeeReceiver.Generate(r.Context(), &pb.PrepareAddressRequest{
		Quantity:      input.Quantity,
		ProcessorKind: processor,
		ChainId:       h.ChainID.Uint64(),
		PaymentToken:  paymentToken,
	})
	if err != nil {
		httpx.WriteGRPCError(w, err, "error creating fee receivers")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]any{
		"status":              "success",
		"ephemeralPublicKeys": res.GetEphemeralPublicKey(),
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}

// AuthorizeFeeReceivers turns stored ephemeral public keys back into fee
// receiver addresses and returns the fee signer's authorization over them.
func (h *ContractHandler) AuthorizeFeeReceivers(w http.ResponseWriter, r *http.Request) {
	if !h.feeReceiverReady(w) {
		return
	}

	var input struct {
		InvoiceID           string   `json:"invoiceId"`
		Processor           string   `json:"processor"`
		Kind                string   `json:"kind"`
		PaymentToken        string   `json:"paymentToken"`
		EphemeralPublicKeys []string `json:"ephemeralPublicKeys"`
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

	if len(input.EphemeralPublicKeys) == 0 {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest,
			errors.New("no ephemeral public keys"),
			"ephemeralPublicKeys is required")
		return
	}

	invoiceID, err := parseBigInt("invoiceId", input.InvoiceID)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	paymentToken, err := feereceiver.ResolveFeeToken(h.Tokens, input.PaymentToken)
	if err != nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest, err, err.Error())
		return
	}

	res, err := h.FeeReceiver.Authorize(r.Context(), &pb.VerifyAddressesRequest{
		EphemeralPublicKey: input.EphemeralPublicKeys,
		ProcessorKind:      processor,
		InvoiceKind:        kind,
		ChainId:            h.ChainID.Uint64(),
		InvoiceId:          invoiceID.String(),
		PaymentToken:       paymentToken,
	})
	if err != nil {
		httpx.WriteGRPCError(w, err, "error authorizing fee receivers")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]any{
		"status":       "success",
		"feeReceivers": res.GetAddresses(),
		"signature":    res.GetSignature(),
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}
