package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func hashKey(key string) string {
	sum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(sum[:])
}

func called(flag *bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		*flag = true
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
	}
}

func TestAccessControlMiddleWare(t *testing.T) {
	const apiKey = "the-secret-api-key"

	tests := []struct {
		name        string
		envKey      string
		header      string
		sendHeader  bool
		wantStatus  int
		wantThrough bool
	}{
		{
			name:        "correct key passes through",
			envKey:      hashKey(apiKey),
			header:      apiKey,
			sendHeader:  true,
			wantStatus:  http.StatusOK,
			wantThrough: true,
		},
		{
			name:       "wrong key is forbidden",
			envKey:     hashKey(apiKey),
			header:     "not-the-key",
			sendHeader: true,
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "empty header is unauthorized",
			envKey:     hashKey(apiKey),
			header:     "",
			sendHeader: true,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "missing header is unauthorized",
			envKey:     hashKey(apiKey),
			sendHeader: false,
			wantStatus: http.StatusUnauthorized,
		},
		{
			name:       "unset KEY is a server misconfiguration",
			envKey:     "",
			header:     apiKey,
			sendHeader: true,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name:       "sending the hash instead of the key is forbidden",
			envKey:     hashKey(apiKey),
			header:     hashKey(apiKey),
			sendHeader: true,
			wantStatus: http.StatusForbidden,
		},
		{
			name:       "key comparison is case sensitive",
			envKey:     hashKey(apiKey),
			header:     strings.ToUpper(apiKey),
			sendHeader: true,
			wantStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("KEY", tt.envKey)
			log.SetOutput(io.Discard)
			t.Cleanup(func() { log.SetOutput(os.Stderr) })

			var reached bool
			handler := AccessControlMiddleWare(called(&reached))

			req := httptest.NewRequest(http.MethodPost, "/v1/invoice", nil)
			if tt.sendHeader {
				req.Header.Set("X-API-KEY", tt.header)
			}
			rec := httptest.NewRecorder()

			handler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if reached != tt.wantThrough {
				t.Errorf("wrapped handler reached = %v, want %v", reached, tt.wantThrough)
			}
		})
	}
}

func TestAccessControlMiddleWareWithATruncatedConfiguredKey(t *testing.T) {
	t.Setenv("KEY", hashKey("secret")[:10])

	var reached bool
	handler := AccessControlMiddleWare(called(&reached))

	req := httptest.NewRequest(http.MethodPost, "/v1/invoice", nil)
	req.Header.Set("X-API-KEY", "secret")
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusForbidden)
	}
	if reached {
		t.Error("wrapped handler ran despite a mismatched key")
	}
}

func TestAccessControlMiddleWarePreservesTheRequest(t *testing.T) {
	const apiKey = "the-secret-api-key"
	t.Setenv("KEY", hashKey(apiKey))

	var body string
	handler := AccessControlMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		body = string(b)
	})

	req := httptest.NewRequest(http.MethodPost, "/v1/invoice", strings.NewReader(`{"a":1}`))
	req.Header.Set("X-API-KEY", apiKey)

	handler(httptest.NewRecorder(), req)

	if body != `{"a":1}` {
		t.Errorf("handler read %q, want the original body", body)
	}
}
