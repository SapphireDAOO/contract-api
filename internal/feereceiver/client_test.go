package feereceiver

import (
	"context"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/SapphireDAOO/contract-api/proto/feereceiverpb"
)

type stubServer struct {
	pb.UnimplementedFeeReceiverServer

	gotPrepare *pb.PrepareAddressRequest
	gotVerify  *pb.VerifyAddressesRequest

	prepareResponse *pb.PrepareAddressResponse
	prepareErr      error
	verifyResponse  *pb.VerifyAddressesResponse
	verifyErr       error

	delay time.Duration
}

func (s *stubServer) Generateaddresses(ctx context.Context, in *pb.PrepareAddressRequest) (*pb.PrepareAddressResponse, error) {
	s.gotPrepare = in
	if s.delay > 0 {
		select {
		case <-time.After(s.delay):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if s.prepareErr != nil {
		return nil, s.prepareErr
	}
	return s.prepareResponse, nil
}

func (s *stubServer) Send(ctx context.Context, in *pb.VerifyAddressesRequest) (*pb.VerifyAddressesResponse, error) {
	s.gotVerify = in
	if s.delay > 0 {
		select {
		case <-time.After(s.delay):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if s.verifyErr != nil {
		return nil, s.verifyErr
	}
	return s.verifyResponse, nil
}

func newStubClient(t *testing.T, stub *stubServer) *Client {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listening for the stub sidecar: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterFeeReceiverServer(server, stub)
	go server.Serve(listener)
	t.Cleanup(server.Stop)

	client, err := NewClient(listener.Addr().String())
	if err != nil {
		t.Fatalf("NewClient returned %v", err)
	}
	t.Cleanup(func() { client.Close() })

	return client
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("fee-receiver:50051")
	if err != nil {
		t.Fatalf("NewClient returned %v", err)
	}
	defer client.Close()

	if got := client.Target(); got != "fee-receiver:50051" {
		t.Errorf("Target() = %q, want the configured target", got)
	}
	if client.conn == nil {
		t.Error("NewClient produced no connection")
	}
	if client.client == nil {
		t.Error("NewClient produced no gRPC stub")
	}
}

func TestNewClientIsLazy(t *testing.T) {
	for _, target := range []string{"127.0.0.1:1", "nothing-here:50051", "not a target"} {
		t.Run(target, func(t *testing.T) {
			client, err := NewClient(target)
			if err != nil {
				t.Fatalf("NewClient(%q) returned %v", target, err)
			}
			defer client.Close()

			if client.Target() != target {
				t.Errorf("Target() = %q, want %q", client.Target(), target)
			}
		})
	}
}

func TestCallToAnUnreachableSidecar(t *testing.T) {
	client, err := NewClient("127.0.0.1:1")
	if err != nil {
		t.Fatalf("NewClient returned %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = client.Generate(ctx, &pb.PrepareAddressRequest{})

	if err == nil {
		t.Fatal("Generate against an unreachable sidecar returned no error")
	}
	if got := status.Code(err); got != codes.Unavailable && got != codes.DeadlineExceeded {
		t.Errorf("error code = %v, want Unavailable or DeadlineExceeded", got)
	}
}

func TestGenerate(t *testing.T) {
	stub := &stubServer{
		prepareResponse: &pb.PrepareAddressResponse{
			EphemeralPublicKey: []string{"0xkey1", "0xkey2"},
		},
	}
	client := newStubClient(t, stub)

	res, err := client.Generate(context.Background(), &pb.PrepareAddressRequest{
		Quantity:      2,
		ProcessorKind: pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED,
		ChainId:       8453,
		PaymentToken:  "0xtoken",
	})
	if err != nil {
		t.Fatalf("Generate returned %v", err)
	}

	if len(res.GetEphemeralPublicKey()) != 2 {
		t.Errorf("got %d keys, want 2", len(res.GetEphemeralPublicKey()))
	}

	if stub.gotPrepare.GetQuantity() != 2 {
		t.Errorf("quantity = %d, want 2", stub.gotPrepare.GetQuantity())
	}
	if stub.gotPrepare.GetChainId() != 8453 {
		t.Errorf("chainId = %d, want 8453", stub.gotPrepare.GetChainId())
	}
	if stub.gotPrepare.GetPaymentToken() != "0xtoken" {
		t.Errorf("paymentToken = %q, want 0xtoken", stub.gotPrepare.GetPaymentToken())
	}
}

func TestAuthorize(t *testing.T) {
	stub := &stubServer{
		verifyResponse: &pb.VerifyAddressesResponse{
			Addresses: []string{"0xreceiver"},
			Signature: "0xsignature",
		},
	}
	client := newStubClient(t, stub)

	res, err := client.Authorize(context.Background(), &pb.VerifyAddressesRequest{
		EphemeralPublicKey: []string{"0xkey1"},
		ProcessorKind:      pb.ProcessorKind_PROCESSOR_KIND_SIMPLE,
		InvoiceKind:        pb.InvoiceKind_INVOICE_KIND_SINGLE,
		ChainId:            8453,
		InvoiceId:          "42",
	})
	if err != nil {
		t.Fatalf("Authorize returned %v", err)
	}

	if res.GetSignature() != "0xsignature" {
		t.Errorf("signature = %q, want 0xsignature", res.GetSignature())
	}
	if len(res.GetAddresses()) != 1 {
		t.Errorf("got %d addresses, want 1", len(res.GetAddresses()))
	}
	if stub.gotVerify.GetInvoiceId() != "42" {
		t.Errorf("invoiceId = %q, want 42", stub.gotVerify.GetInvoiceId())
	}
}

func TestClientPropagatesStatusErrors(t *testing.T) {
	client := newStubClient(t, &stubServer{
		prepareErr: status.Error(codes.InvalidArgument, "quantity must be positive"),
		verifyErr:  status.Error(codes.NotFound, "unknown ephemeral key"),
	})

	_, err := client.Generate(context.Background(), &pb.PrepareAddressRequest{})
	if got := status.Code(err); got != codes.InvalidArgument {
		t.Errorf("Generate error code = %v, want InvalidArgument", got)
	}

	_, err = client.Authorize(context.Background(), &pb.VerifyAddressesRequest{})
	if got := status.Code(err); got != codes.NotFound {
		t.Errorf("Authorize error code = %v, want NotFound", got)
	}
}

func TestClientHonoursACancelledContext(t *testing.T) {
	client := newStubClient(t, &stubServer{
		prepareResponse: &pb.PrepareAddressResponse{},
		delay:           5 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.Generate(ctx, &pb.PrepareAddressRequest{})

	if err == nil {
		t.Fatal("Generate returned no error for a cancelled context")
	}
	if got := status.Code(err); got != codes.Canceled {
		t.Errorf("error code = %v, want Canceled", got)
	}
}

func TestClientHonoursAShorterCallerDeadline(t *testing.T) {
	client := newStubClient(t, &stubServer{
		prepareResponse: &pb.PrepareAddressResponse{},
		delay:           5 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := client.Generate(ctx, &pb.PrepareAddressRequest{})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("Generate returned no error for an expired deadline")
	}
	if got := status.Code(err); got != codes.DeadlineExceeded {
		t.Errorf("error code = %v, want DeadlineExceeded", got)
	}

	if elapsed > callTimeout {
		t.Errorf("Generate waited %v, want it bounded by the caller's deadline", elapsed)
	}
}

func TestCallTimeoutIsBounded(t *testing.T) {
	if callTimeout <= 0 {
		t.Fatalf("callTimeout = %v, want a positive bound", callTimeout)
	}
	if callTimeout > time.Minute {
		t.Errorf("callTimeout = %v, want it well under a minute", callTimeout)
	}
}

func TestClose(t *testing.T) {
	client, err := NewClient("127.0.0.1:1")
	if err != nil {
		t.Fatalf("NewClient returned %v", err)
	}

	if err := client.Close(); err != nil {
		t.Errorf("Close returned %v", err)
	}
}

func TestCallAfterClose(t *testing.T) {
	client := newStubClient(t, &stubServer{prepareResponse: &pb.PrepareAddressResponse{}})

	if err := client.Close(); err != nil {
		t.Fatalf("Close returned %v", err)
	}

	if _, err := client.Generate(context.Background(), &pb.PrepareAddressRequest{}); err == nil {
		t.Error("Generate after Close returned no error")
	}
}
