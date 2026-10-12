package handler

import (
	"context"
	"encoding/json"
	"fmt"
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

	got *pb.CreateFeeReceiversRequest

	response *pb.CreateFeeReceiversResponse
	err      error
}

func (s *stubFeeReceiver) CreateFeeReceivers(_ context.Context, in *pb.CreateFeeReceiversRequest) (*pb.CreateFeeReceiversResponse, error) {
	s.got = in
	if s.err != nil {
		return nil, s.err
	}
	return s.response, nil
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

	h := &ContractHandler{
		Tokens:      testTokens(),
		ChainID:     big.NewInt(8453),
		FeeReceiver: client,
	}
	readyChain().install(t, h)
	return h
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
		response: &pb.CreateFeeReceiversResponse{
			Addresses:          []string{"0xreceiver1", "0xreceiver2"},
			Signature:          "0xsignature",
			EphemeralPublicKey: []string{"0xkey1", "0xkey2"},
		},
	}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"invoiceId":"42","processor":"intermediated","kind":"meta",
		  "quantity":2,"paymentToken":"USDC"}`))

	body := successBody(t, rec)
	if body["status"] != "success" {
		t.Errorf("status = %v, want success", body["status"])
	}
	if body["signature"] != "0xsignature" {
		t.Errorf("signature = %v, want the signer's authorization", body["signature"])
	}
	receivers, ok := body["feeReceivers"].([]any)
	if !ok || len(receivers) != 2 {
		t.Fatalf("feeReceivers = %v, want two addresses", body["feeReceivers"])
	}
	keys, ok := body["ephemeralPublicKeys"].([]any)
	if !ok || len(keys) != 2 {
		t.Fatalf("ephemeralPublicKeys = %v, want two keys", body["ephemeralPublicKeys"])
	}

	if stub.got.GetInvoiceId() != "42" {
		t.Errorf("invoiceId = %q, want 42", stub.got.GetInvoiceId())
	}
	if stub.got.GetInvoiceKind() != pb.InvoiceKind_INVOICE_KIND_META {
		t.Errorf("invoiceKind = %v, want META", stub.got.GetInvoiceKind())
	}
	if stub.got.GetProcessorKind() != pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED {
		t.Errorf("processorKind = %v, want INTERMEDIATED", stub.got.GetProcessorKind())
	}
	if stub.got.GetQuantity() != 2 {
		t.Errorf("quantity = %d, want 2", stub.got.GetQuantity())
	}
	if stub.got.GetChainId() != 8453 {
		t.Errorf("chainId = %d, want 8453", stub.got.GetChainId())
	}
	if stub.got.GetPaymentToken() != usdcAddress {
		t.Errorf("paymentToken = %q, want the resolved address %q", stub.got.GetPaymentToken(), usdcAddress)
	}
}

func TestCreateFeeReceiversDefaults(t *testing.T) {
	stub := &stubFeeReceiver{response: &pb.CreateFeeReceiversResponse{}}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"invoiceId":"42","processor":"simple"}`))

	successBody(t, rec)
	if stub.got.GetQuantity() != 1 {
		t.Errorf("quantity = %d, want the default of 1", stub.got.GetQuantity())
	}
	if stub.got.GetInvoiceKind() != pb.InvoiceKind_INVOICE_KIND_SINGLE {
		t.Errorf("invoiceKind = %v, want SINGLE", stub.got.GetInvoiceKind())
	}
	if stub.got.GetPaymentToken() != "" {
		t.Errorf("paymentToken = %q, want it empty", stub.got.GetPaymentToken())
	}
}

func TestCreateFeeReceiversNormalisesTheInvoiceID(t *testing.T) {
	stub := &stubFeeReceiver{response: &pb.CreateFeeReceiversResponse{}}
	h := newFeeReceiverHandler(t, stub)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"invoiceId":"  0042  ","processor":"simple"}`))

	successBody(t, rec)
	if stub.got.GetInvoiceId() != "42" {
		t.Errorf("invoiceId = %q, want the canonical 42", stub.got.GetInvoiceId())
	}
}

func TestCreateFeeReceiversBadRequests(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		wantErr string
	}{
		{"malformed json", "not json", "invalid request body"},
		{"missing processor", `{"invoiceId":"42"}`, "processor must be simple or intermediated"},
		{"unknown processor", `{"invoiceId":"42","processor":"escrow"}`, "processor must be simple or intermediated"},
		{"unknown kind", `{"invoiceId":"42","processor":"simple","kind":"batch"}`, "kind must be single or meta"},
		{"invoice id is not a number", `{"invoiceId":"abc","processor":"simple"}`, "invoiceId must be a base-10 integer"},
		{"missing invoice id", `{"processor":"simple"}`, "invoiceId must be a base-10 integer"},
		{"unknown payment token", `{"invoiceId":"42","processor":"simple","paymentToken":"DOGE"}`, "unknown token DOGE"},
		{"meta invoice on the simple processor", `{"invoiceId":"42","processor":"simple","kind":"meta"}`, "meta invoices are intermediated only"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := &stubFeeReceiver{response: &pb.CreateFeeReceiversResponse{}}
			h := newFeeReceiverHandler(t, stub)
			rec := httptest.NewRecorder()

			h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers", tt.body))

			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
			}
			if body := decodeError(t, rec); !strings.Contains(body["error"], tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", body["error"], tt.wantErr)
			}

			if stub.got != nil {
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
			h := newFeeReceiverHandler(t, &stubFeeReceiver{err: tt.err})
			rec := httptest.NewRecorder()

			h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
				`{"invoiceId":"42","processor":"simple"}`))

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if body := decodeError(t, rec); !strings.Contains(body["error"], "error creating fee receivers") {
				t.Errorf("error = %q, want the endpoint's message", body["error"])
			}
		})
	}
}

func TestCreateFeeReceiversKeepsTheSidecarReason(t *testing.T) {
	h := newFeeReceiverHandler(t, &stubFeeReceiver{
		err: status.Error(codes.InvalidArgument, "A single invoice takes one fee receiver"),
	})
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"invoiceId":"42","processor":"simple","quantity":2}`))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if body := decodeError(t, rec); !strings.Contains(body["reason"], "A single invoice takes one fee receiver") {
		t.Errorf("reason = %q, want the sidecar's message", body["reason"])
	}
}

func TestCreateFeeReceiversWithoutTheSidecar(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"invoiceId":"42","processor":"simple"}`))

	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
	}
	if body := decodeError(t, rec); !strings.Contains(body["error"], "fee receivers are unavailable") {
		t.Errorf("error = %q, want it to say the endpoint is unavailable", body["error"])
	}
}

func TestCreateFeeReceiversChecksTheInvoiceState(t *testing.T) {
	const (
		simpleSingle        = `{"invoiceId":"%d","processor":"simple"}`
		intermediatedSingle = `{"invoiceId":"%d","processor":"intermediated"}`
		intermediatedMeta   = `{"invoiceId":"%d","processor":"intermediated","kind":"meta","quantity":2}`
	)

	tests := []struct {
		name    string
		body    string
		id      int64
		prepare func(*fakeChain)
		allowed bool
	}{
		{"simple paid", simpleSingle, 42, nil, true},
		{"simple not yet paid", simpleSingle, 42, func(c *fakeChain) { c.simple[42] = 1 }, false},
		{"simple never created", simpleSingle, 99, nil, false},
		{"intermediated created", intermediatedSingle, 42, nil, true},
		{"intermediated already paid", intermediatedSingle, 42, func(c *fakeChain) { c.intermediated[42] = 2 }, false},
		{"intermediated never created", intermediatedSingle, 99, nil, false},
		{"meta with one sub-invoice paid", intermediatedMeta, 42, nil, true},
		{"meta with no sub-invoice paid", intermediatedMeta, 42, func(c *fakeChain) { c.intermediated[43] = 1 }, false},
		{"meta never created", intermediatedMeta, 99, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stub := &stubFeeReceiver{response: &pb.CreateFeeReceiversResponse{}}
			h := newFeeReceiverHandler(t, stub)
			chain := readyChain()
			if tt.prepare != nil {
				tt.prepare(chain)
			}
			chain.install(t, h)
			rec := httptest.NewRecorder()

			h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers", fmt.Sprintf(tt.body, tt.id)))

			if tt.allowed {
				successBody(t, rec)
				if stub.got == nil {
					t.Error("an invoice in the right state was not forwarded to the sidecar")
				}
				return
			}

			if rec.Code != http.StatusConflict {
				t.Fatalf("status = %d, want %d (body %q)", rec.Code, http.StatusConflict, rec.Body.String())
			}
			body := decodeError(t, rec)
			if body["error"] != "invoice cannot take fee receivers" {
				t.Errorf("error = %q, want %q", body["error"], "invoice cannot take fee receivers")
			}
			if want := fmt.Sprintf("invoice %d is not in a state that takes fee receivers", tt.id); !strings.Contains(body["reason"], want) {
				t.Errorf("reason = %q, want it to contain %q", body["reason"], want)
			}
			if stub.got != nil {
				t.Error("an invoice in the wrong state was still forwarded to the sidecar")
			}
		})
	}
}

func TestCreateFeeReceiversWhenTheStateReadFails(t *testing.T) {
	stub := &stubFeeReceiver{response: &pb.CreateFeeReceiversResponse{}}
	h := newFeeReceiverHandler(t, stub)
	chain := readyChain()
	chain.failWith = `Post "https://rpc.example.com/v2/SECRET": dial tcp: refused`
	chain.install(t, h)
	rec := httptest.NewRecorder()

	h.CreateFeeReceivers(rec, feeReceiverRequest("/v1/fee-receivers",
		`{"invoiceId":"42","processor":"intermediated"}`))

	if rec.Code != http.StatusBadGateway {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadGateway)
	}
	if strings.Contains(rec.Body.String(), "SECRET") {
		t.Errorf("response leaks the RPC error: %s", rec.Body.String())
	}
	if stub.got != nil {
		t.Error("a request whose invoice could not be verified was forwarded to the sidecar")
	}
}
