package handler

import (
	"context"
	"encoding/json"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SapphireDAOO/contract-api/internal/feereceiver"
	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
)

type stubFeeReceiver struct {
	pb.UnimplementedFeeReceiverServer

	gotPrepare *pb.PrepareAddressRequest
	gotVerify  *pb.VerifyAddressesRequest

	prepareResponse *pb.PrepareAddressResponse
	prepareErr      error
	verifyResponse  *pb.VerifyAddressesResponse
	verifyErr       error
}

func (s *stubFeeReceiver) Generateaddresses(_ context.Context, in *pb.PrepareAddressRequest) (*pb.PrepareAddressResponse, error) {
	s.gotPrepare = in
	if s.prepareErr != nil {
		return nil, s.prepareErr
	}
	return s.prepareResponse, nil
}

func (s *stubFeeReceiver) Send(_ context.Context, in *pb.VerifyAddressesRequest) (*pb.VerifyAddressesResponse, error) {
	s.gotVerify = in
	if s.verifyErr != nil {
		return nil, s.verifyErr
	}
	return s.verifyResponse, nil
}

func newFeeReceiverHandler(t *testing.T, stub *stubFeeReceiver) *ContractHandler {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listening for the stub sidecar: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterFeeReceiverServer(server, stub)
	go server.Serve(listener)
	t.Cleanup(server.Stop)

	client, err := feereceiver.NewClient(listener.Addr().String())
	if err != nil {
		t.Fatalf("dialing the stub sidecar: %v", err)
	}
	t.Cleanup(func() { client.Close() })

	return &ContractHandler{
		Tokens:      testTokens(),
		ChainID:     big.NewInt(8453),
		FeeReceiver: client,
	}
}

func feeReceiverRequest(path, body string) *http.Request {
	return httptest.NewRequest(http.MethodPost, path, strings.NewReader(body))
}

func successBody(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d (body %q)", rec.Code, http.StatusOK, rec.Body.String())
	}
	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v (body %q)", err, rec.Body.String())
	}
	return body
}

func TestCreateFeeReceivers(t *testing.T) {
	stub := &stubFeeReceiver{
		prepareResponse: &pb.PrepareAddressResponse{
			EphemeralPublicKey: []string{"0xkey1", "0xkey2"},
		},
	}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"quantity":2,"processor":"intermediated","paymentToken":"USDC"}`))

	body := successBody(t, rec)
	if body["status"] != "success" {
		t.Errorf("status = %v, want success", body["status"])
	}
	keys, ok := body["ephemeralPublicKeys"].([]any)
	if !ok || len(keys) != 2 {
		t.Fatalf("ephemeralPublicKeys = %v, want two keys", body["ephemeralPublicKeys"])
	}

	if stub.gotPrepare.GetQuantity() != 2 {
		t.Errorf("quantity = %d, want 2", stub.gotPrepare.GetQuantity())
	}
	if stub.gotPrepare.GetProcessorKind() != pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED {
		t.Errorf("processorKind = %v, want INTERMEDIATED", stub.gotPrepare.GetProcessorKind())
	}
	if stub.gotPrepare.GetChainId() != 8453 {
		t.Errorf("chainId = %d, want 8453", stub.gotPrepare.GetChainId())
	}
	if stub.gotPrepare.GetPaymentToken() != usdcAddress {
		t.Errorf("paymentToken = %q, want the resolved address %q", stub.gotPrepare.GetPaymentToken(), usdcAddress)
	}
}

func TestCreateFeeReceiversDefaultsQuantityToOne(t *testing.T) {
	stub := &stubFeeReceiver{prepareResponse: &pb.PrepareAddressResponse{}}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers", `{"processor":"simple"}`))

	successBody(t, rec)
	if stub.gotPrepare.GetQuantity() != 1 {
		t.Errorf("quantity = %d, want the default of 1", stub.gotPrepare.GetQuantity())
	}
}

func TestCreateFeeReceiversAllowsAnUnsetPaymentToken(t *testing.T) {
	stub := &stubFeeReceiver{prepareResponse: &pb.PrepareAddressResponse{}}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers", `{"processor":"simple"}`))

	successBody(t, rec)
	if stub.gotPrepare.GetPaymentToken() != "" {
		t.Errorf("paymentToken = %q, want it empty", stub.gotPrepare.GetPaymentToken())
	}
}

func TestCreateFeeReceiversBadRequests(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr string
	}{
		{"malformed json", "not json", "invalid request body"},
		{"missing processor", `{"quantity":1}`, "processor must be simple or intermediated"},
		{"unknown processor", `{"processor":"escrow"}`, "processor must be simple or intermediated"},
		{"unknown payment token", `{"processor":"simple","paymentToken":"DOGE"}`, "unknown token DOGE"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := &stubFeeReceiver{prepareResponse: &pb.PrepareAddressResponse{}}
			h := newFeeReceiverHandler(t, stub)
			rec := httptest.NewRecorder()

			h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers", tt.body))

			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
			}
			if body := decodeError(t, rec); !strings.Contains(body["error"], tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", body["error"], tt.wantErr)
			}

			if stub.gotPrepare != nil {
				t.Error("a rejected request was still forwarded to the sidecar")
			}
		})
	}
}

func TestCreateFeeReceiversMapsSidecarErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{"invalid argument", status.Error(codes.InvalidArgument, "bad chain"), http.StatusBadRequest},
		{"unavailable", status.Error(codes.Unavailable, "no relayer"), http.StatusServiceUnavailable},
		{"deadline exceeded", status.Error(codes.DeadlineExceeded, "slow"), http.StatusGatewayTimeout},
		{"internal", status.Error(codes.Internal, "signer down"), http.StatusBadGateway},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newFeeReceiverHandler(t, &stubFeeReceiver{prepareErr: tt.err})
			rec := httptest.NewRecorder()

			h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers", `{"processor":"simple"}`))

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if body := decodeError(t, rec); !strings.Contains(body["error"], "error creating fee receivers") {
				t.Errorf("error = %q, want the endpoint's message", body["error"])
			}
		})
	}
}

func TestAuthorizeFeeReceivers(t *testing.T) {
	stub := &stubFeeReceiver{
		verifyResponse: &pb.VerifyAddressesResponse{
			Addresses: []string{"0xreceiver1"},
			Signature: "0xsignature",
		},
	}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.AuthorizeFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers/authorization",
		`{"invoiceId":"42","processor":"intermediated","kind":"meta",
		  "paymentToken":"USDC","ephemeralPublicKeys":["0xkey1"]}`))

	body := successBody(t, rec)
	if body["status"] != "success" {
		t.Errorf("status = %v, want success", body["status"])
	}
	if body["signature"] != "0xsignature" {
		t.Errorf("signature = %v, want the signer's authorization", body["signature"])
	}
	receivers, ok := body["feeReceivers"].([]any)
	if !ok || len(receivers) != 1 {
		t.Fatalf("feeReceivers = %v, want one address", body["feeReceivers"])
	}

	if stub.gotVerify.GetInvoiceId() != "42" {
		t.Errorf("invoiceId = %q, want 42", stub.gotVerify.GetInvoiceId())
	}
	if stub.gotVerify.GetInvoiceKind() != pb.InvoiceKind_INVOICE_KIND_META {
		t.Errorf("invoiceKind = %v, want META", stub.gotVerify.GetInvoiceKind())
	}
	if stub.gotVerify.GetProcessorKind() != pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED {
		t.Errorf("processorKind = %v, want INTERMEDIATED", stub.gotVerify.GetProcessorKind())
	}
	if stub.gotVerify.GetChainId() != 8453 {
		t.Errorf("chainId = %d, want 8453", stub.gotVerify.GetChainId())
	}
	if len(stub.gotVerify.GetEphemeralPublicKey()) != 1 {
		t.Errorf("ephemeralPublicKey = %v, want the stored key", stub.gotVerify.GetEphemeralPublicKey())
	}
}

func TestAuthorizeFeeReceiversDefaultsToASingleInvoice(t *testing.T) {
	stub := &stubFeeReceiver{verifyResponse: &pb.VerifyAddressesResponse{}}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.AuthorizeFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers/authorization",
		`{"invoiceId":"42","processor":"simple","ephemeralPublicKeys":["0xkey1"]}`))

	successBody(t, rec)
	if stub.gotVerify.GetInvoiceKind() != pb.InvoiceKind_INVOICE_KIND_SINGLE {
		t.Errorf("invoiceKind = %v, want SINGLE", stub.gotVerify.GetInvoiceKind())
	}
}

func TestAuthorizeFeeReceiversNormalisesTheInvoiceID(t *testing.T) {
	stub := &stubFeeReceiver{verifyResponse: &pb.VerifyAddressesResponse{}}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.AuthorizeFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers/authorization",
		`{"invoiceId":"  0042  ","processor":"simple","ephemeralPublicKeys":["0xkey1"]}`))

	successBody(t, rec)
	if stub.gotVerify.GetInvoiceId() != "42" {
		t.Errorf("invoiceId = %q, want the canonical 42", stub.gotVerify.GetInvoiceId())
	}
}

func TestAuthorizeFeeReceiversBadRequests(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr string
	}{
		{"malformed json", "not json", "invalid request body"},
		{
			"unknown processor",
			`{"invoiceId":"42","processor":"escrow","ephemeralPublicKeys":["0xkey1"]}`,
			"processor must be simple or intermediated",
		},
		{
			"unknown kind",
			`{"invoiceId":"42","processor":"simple","kind":"batch","ephemeralPublicKeys":["0xkey1"]}`,
			"kind must be single or meta",
		},
		{
			"no ephemeral keys",
			`{"invoiceId":"42","processor":"simple","ephemeralPublicKeys":[]}`,
			"ephemeralPublicKeys is required",
		},
		{
			"missing ephemeral keys",
			`{"invoiceId":"42","processor":"simple"}`,
			"ephemeralPublicKeys is required",
		},
		{
			"invoice id is not a number",
			`{"invoiceId":"abc","processor":"simple","ephemeralPublicKeys":["0xkey1"]}`,
			"invoiceId must be a base-10 integer",
		},
		{
			"missing invoice id",
			`{"processor":"simple","ephemeralPublicKeys":["0xkey1"]}`,
			"invoiceId must be a base-10 integer",
		},
		{
			"unknown payment token",
			`{"invoiceId":"42","processor":"simple","paymentToken":"DOGE","ephemeralPublicKeys":["0xkey1"]}`,
			"unknown token DOGE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := &stubFeeReceiver{verifyResponse: &pb.VerifyAddressesResponse{}}
			h := newFeeReceiverHandler(t, stub)
			rec := httptest.NewRecorder()

			h.AuthorizeFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers/authorization", tt.body))

			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
			}
			if body := decodeError(t, rec); !strings.Contains(body["error"], tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", body["error"], tt.wantErr)
			}
			if stub.gotVerify != nil {
				t.Error("a rejected request was still forwarded to the sidecar")
			}
		})
	}
}

func TestAuthorizeFeeReceiversMapsSidecarErrors(t *testing.T) {
	h := newFeeReceiverHandler(t, &stubFeeReceiver{
		verifyErr: status.Error(codes.InvalidArgument, "keys do not match the invoice"),
	})
	rec := httptest.NewRecorder()

	h.AuthorizeFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers/authorization",
		`{"invoiceId":"42","processor":"simple","ephemeralPublicKeys":["0xkey1"]}`))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	body := decodeError(t, rec)
	if !strings.Contains(body["error"], "error authorizing fee receivers") {
		t.Errorf("error = %q, want the endpoint's message", body["error"])
	}
	if !strings.Contains(body["reason"], "keys do not match the invoice") {
		t.Errorf("reason = %q, want the sidecar's message", body["reason"])
	}
}

func TestFeeReceiverEndpointsWithoutTheSidecar(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}

	endpoints := map[string]func(http.ResponseWriter, *http.Request){
		"create":    h.CreateFeeReceivers,
		"authorize": h.AuthorizeFeeReceivers,
	}

	for name, endpoint := range endpoints {
		t.Run(name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			endpoint(rec, feeReceiverRequest("/v1/fee-receivers", `{"processor":"simple"}`))

			if rec.Code != http.StatusServiceUnavailable {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
			}
			if body := decodeError(t, rec); !strings.Contains(body["error"], "fee receivers are unavailable") {
				t.Errorf("error = %q, want it to say the endpoint is unavailable", body["error"])
			}
		})
	}
}
