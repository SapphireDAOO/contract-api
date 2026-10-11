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

	got *pb.CreateFeeReceiversRequest

	response *pb.CreateFeeReceiversResponse
	err      error

	delay time.Duration
}

func (s *stubServer) CreateFeeReceivers(ctx context.Context, in *pb.CreateFeeReceiversRequest) (*pb.CreateFeeReceiversResponse, error) {
	s.got = in
	if s.delay > 0 {
		select {
		case <-time.After(s.delay):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if s.err != nil {
		return nil, s.err
	}
	return s.response, nil
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

	_, err = client.Create(ctx, &pb.CreateFeeReceiversRequest{})

	if err == nil {
		t.Fatal("Create against an unreachable sidecar returned no error")
	}
	if got := status.Code(err); got != codes.Unavailable && got != codes.DeadlineExceeded {
		t.Errorf("error code = %v, want Unavailable or DeadlineExceeded", got)
	}
}

func TestCreate(t *testing.T) {
	stub := &stubServer{
		response: &pb.CreateFeeReceiversResponse{
			Addresses:          []string{"0xreceiver1", "0xreceiver2"},
			Signature:          "0xsignature",
			EphemeralPublicKey: []string{"0xkey1", "0xkey2"},
		},
	}
	client := newStubClient(t, stub)

	res, err := client.Create(context.Background(), &pb.CreateFeeReceiversRequest{
		ProcessorKind: pb.ProcessorKind_PROCESSOR_KIND_INTERMEDIATED,
		InvoiceKind:   pb.InvoiceKind_INVOICE_KIND_META,
		ChainId:       8453,
		InvoiceId:     "42",
		PaymentToken:  "0xtoken",
		Quantity:      2,
	})
	if err != nil {
		t.Fatalf("Create returned %v", err)
	}

	if len(res.GetAddresses()) != 2 {
		t.Errorf("got %d addresses, want 2", len(res.GetAddresses()))
	}
	if res.GetSignature() != "0xsignature" {
		t.Errorf("signature = %q, want 0xsignature", res.GetSignature())
	}
	if len(res.GetEphemeralPublicKey()) != 2 {
		t.Errorf("got %d keys, want 2", len(res.GetEphemeralPublicKey()))
	}

	if stub.got.GetQuantity() != 2 {
		t.Errorf("quantity = %d, want 2", stub.got.GetQuantity())
	}
	if stub.got.GetChainId() != 8453 {
		t.Errorf("chainId = %d, want 8453", stub.got.GetChainId())
	}
	if stub.got.GetInvoiceId() != "42" {
		t.Errorf("invoiceId = %q, want 42", stub.got.GetInvoiceId())
	}
	if stub.got.GetPaymentToken() != "0xtoken" {
		t.Errorf("paymentToken = %q, want 0xtoken", stub.got.GetPaymentToken())
	}
}

func TestClientPropagatesStatusErrors(t *testing.T) {
	client := newStubClient(t, &stubServer{
		err: status.Error(codes.InvalidArgument, "quantity must be positive"),
	})

	_, err := client.Create(context.Background(), &pb.CreateFeeReceiversRequest{})
	if got := status.Code(err); got != codes.InvalidArgument {
		t.Errorf("Create error code = %v, want InvalidArgument", got)
	}
}

func TestClientHonoursACancelledContext(t *testing.T) {
	client := newStubClient(t, &stubServer{
		response: &pb.CreateFeeReceiversResponse{},
		delay:    5 * time.Second,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := client.Create(ctx, &pb.CreateFeeReceiversRequest{})

	if err == nil {
		t.Fatal("Create returned no error for a cancelled context")
	}
	if got := status.Code(err); got != codes.Canceled {
		t.Errorf("error code = %v, want Canceled", got)
	}
}

func TestClientHonoursAShorterCallerDeadline(t *testing.T) {
	client := newStubClient(t, &stubServer{
		response: &pb.CreateFeeReceiversResponse{},
		delay:    5 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := client.Create(ctx, &pb.CreateFeeReceiversRequest{})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("Create returned no error for an expired deadline")
	}
	if got := status.Code(err); got != codes.DeadlineExceeded {
		t.Errorf("error code = %v, want DeadlineExceeded", got)
	}

	if elapsed > callTimeout {
		t.Errorf("Create waited %v, want it bounded by the caller's deadline", elapsed)
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
	client := newStubClient(t, &stubServer{response: &pb.CreateFeeReceiversResponse{}})

	if err := client.Close(); err != nil {
		t.Fatalf("Close returned %v", err)
	}

	if _, err := client.Create(context.Background(), &pb.CreateFeeReceiversRequest{}); err == nil {
		t.Error("Create after Close returned no error")
	}
}
