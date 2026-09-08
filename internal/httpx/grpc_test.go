package httpx

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestWriteGRPCError(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
		wantReason string
	}{
		{
			"invalid argument becomes a bad request",
			status.Error(codes.InvalidArgument, "quantity must be positive"),
			http.StatusBadRequest,
			"quantity must be positive",
		},
		{
			"unavailable becomes a service unavailable",
			status.Error(codes.Unavailable, "connection refused"),
			http.StatusServiceUnavailable,
			"connection refused",
		},
		{
			"deadline exceeded becomes a gateway timeout",
			status.Error(codes.DeadlineExceeded, "context deadline exceeded"),
			http.StatusGatewayTimeout,
			"context deadline exceeded",
		},
		{
			"internal becomes a bad gateway",
			status.Error(codes.Internal, "signer failed"),
			http.StatusBadGateway,
			"signer failed",
		},
		{
			"not found becomes a bad gateway",
			status.Error(codes.NotFound, "no such key"),
			http.StatusBadGateway,
			"no such key",
		},
		{
			"permission denied becomes a bad gateway",
			status.Error(codes.PermissionDenied, "wrong chain"),
			http.StatusBadGateway,
			"wrong chain",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			WriteGRPCError(rec, tt.err, "error creating fee receivers")

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			body := decodeBody(t, rec)
			if body["reason"] != tt.wantReason {
				t.Errorf("reason = %q, want %q", body["reason"], tt.wantReason)
			}
			if body["error"] != "error creating fee receivers" {
				t.Errorf("error = %q, want %q", body["error"], "error creating fee receivers")
			}
		})
	}
}

func TestWriteGRPCErrorStripsTheRPCPrefix(t *testing.T) {
	err := status.Error(codes.InvalidArgument, "chain id mismatch")
	rec := httptest.NewRecorder()

	WriteGRPCError(rec, err, "failed")

	body := decodeBody(t, rec)
	if body["reason"] != "chain id mismatch" {
		t.Errorf("reason = %q, want the bare status message", body["reason"])
	}
	if body["reason"] == err.Error() {
		t.Errorf("reason still carries the rpc error prefix: %q", body["reason"])
	}
}

func TestWriteGRPCErrorNil(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteGRPCError(rec, nil, "failed")

	if rec.Code != http.StatusBadGateway {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadGateway)
	}
}

func TestWriteGRPCErrorPlainError(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteGRPCError(rec, errors.New("dial tcp: no route to host"), "failed")

	if rec.Code != http.StatusBadGateway {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadGateway)
	}
	body := decodeBody(t, rec)
	if body["reason"] != "dial tcp: no route to host" {
		t.Errorf("reason = %q, want the error text", body["reason"])
	}
}
