package httpx

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type revertError struct {
	data any
}

func (e revertError) Error() string          { return "execution reverted" }
func (e revertError) ErrorCode() int         { return 3 }
func (e revertError) ErrorData() interface{} { return e.data }

func revertWith(selector string) error { return revertError{data: selector} }

func decodeBody(t *testing.T, rec *httptest.ResponseRecorder) map[string]string {
	t.Helper()
	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body is not JSON: %v (body %q)", err, rec.Body.String())
	}
	return body
}

func TestWriteHTTPErrorWithStatus(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteHTTPErrorWithStatus(rec, http.StatusBadRequest, errors.New("bad input"), "invalid request body")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if got, want := rec.Header().Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}

	body := decodeBody(t, rec)
	if body["error"] != "invalid request body" {
		t.Errorf("error = %q, want %q", body["error"], "invalid request body")
	}
	if body["reason"] != "bad input" {
		t.Errorf("reason = %q, want %q", body["reason"], "bad input")
	}
}

func TestWriteHTTPErrorWithStatusUsesTheRevertDescription(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteHTTPErrorWithStatus(rec, http.StatusBadGateway, revertWith("0x2c669f0a"), "failed to read the price")

	body := decodeBody(t, rec)
	if want := "The price cannot be zero."; body["reason"] != want {
		t.Errorf("reason = %q, want %q", body["reason"], want)
	}

	if rec.Code != http.StatusBadGateway {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadGateway)
	}
}

func TestWriteHTTPErrorWithStatusNilError(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteHTTPErrorWithStatus(rec, http.StatusServiceUnavailable, nil, "unavailable")

	body := decodeBody(t, rec)
	if body["reason"] != "" {
		t.Errorf("reason = %q, want an empty string for a nil error", body["reason"])
	}
	if body["error"] != "unavailable" {
		t.Errorf("error = %q, want %q", body["error"], "unavailable")
	}
}

func TestWriteMappedRevertError(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
		wantReason string
	}{
		{
			"conflict revert",
			revertWith("0x487e4409"),
			http.StatusConflict,
			"The invoice is not in a valid state for this action.",
		},
		{
			"not found revert",
			revertWith("0x715d9228"),
			http.StatusNotFound,
			"The specified invoice does not exist.",
		},
		{
			"forbidden revert",
			revertWith("0xea8e4eb5"),
			http.StatusForbidden,
			"The caller is not authorized to perform this action.",
		},
		{
			"bad request revert",
			revertWith("0xf4d678b8"),
			http.StatusBadRequest,
			"The account balance is insufficient to perform this action.",
		},
		{
			"described revert with no mapped status falls back to 500",
			revertWith("0x00bfc921"),
			http.StatusInternalServerError,
			"The oracle reported an invalid price.",
		},
		{
			"unknown revert falls back to 500 and the error text",
			revertWith("0xdeadbeef"),
			http.StatusInternalServerError,
			"execution reverted",
		},
		{
			"plain error falls back to 500 and the error text",
			errors.New("connection refused"),
			http.StatusInternalServerError,
			"connection refused",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()

			WriteMappedRevertError(rec, tt.err, "transaction failed")

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			body := decodeBody(t, rec)
			if body["reason"] != tt.wantReason {
				t.Errorf("reason = %q, want %q", body["reason"], tt.wantReason)
			}
			if body["error"] != "transaction failed" {
				t.Errorf("error = %q, want %q", body["error"], "transaction failed")
			}
		})
	}
}

func TestErrorBodyHasOnlyErrorAndReason(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteMappedRevertError(rec, errors.New("boom"), "failed")

	body := decodeBody(t, rec)
	if len(body) != 2 {
		t.Errorf("body has %d keys (%v), want exactly error and reason", len(body), body)
	}
}
