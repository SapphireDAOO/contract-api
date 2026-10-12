package middleware

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"

	"github.com/SapphireDAOO/contract-api/internal/auth"
)

const (
	sessionSecret = "0123456789abcdef0123456789abcdef"
	sessionDomain = "app.example.com"
	sessionChain  = 84532
)

func signIn(t *testing.T, sessions *auth.Service) (common.Address, string) {
	t.Helper()

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generating key: %v", err)
	}
	nonce, _, err := sessions.Nonce()
	if err != nil {
		t.Fatalf("Nonce: %v", err)
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	msg, err := siwe.InitMessage(sessionDomain, address.Hex(), "https://"+sessionDomain, nonce,
		map[string]any{"chainId": sessionChain})
	if err != nil {
		t.Fatalf("building message: %v", err)
	}
	session, err := sessions.SignIn(msg.String(), sign(t, key, msg.String()))
	if err != nil {
		t.Fatalf("SignIn: %v", err)
	}
	return address, session.Token
}

func sign(t *testing.T, key *ecdsa.PrivateKey, message string) string {
	t.Helper()
	sig, err := crypto.Sign(accounts.TextHash([]byte(message)), key)
	if err != nil {
		t.Fatalf("signing: %v", err)
	}
	return hexutil.Encode(sig)
}

func newSessions(t *testing.T) *auth.Service {
	t.Helper()
	sessions, err := auth.New(sessionSecret, []string{sessionDomain}, sessionChain)
	if err != nil {
		t.Fatalf("auth.New: %v", err)
	}
	return sessions
}

func TestRequireSessionPassesTheSignedInAddress(t *testing.T) {
	sessions := newSessions(t)
	address, token := signIn(t, sessions)

	var got common.Address
	var found bool
	handler := RequireSession(sessions, func(w http.ResponseWriter, r *http.Request) {
		got, found = auth.AddressFrom(r.Context())
	})

	req := httptest.NewRequest(http.MethodPost, "/v1/notes", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	handler(httptest.NewRecorder(), req)

	if !found || got != address {
		t.Errorf("handler saw address %s (found %v), want %s", got, found, address)
	}
}

func TestRequireSessionRejects(t *testing.T) {
	sessions := newSessions(t)
	_, token := signIn(t, sessions)

	tests := []struct {
		name          string
		authorization string
		wantError     string
	}{
		{"no header", "", "Sign-in required"},
		{"not a bearer token", "Basic " + token, "Sign-in required"},
		{"empty token", "Bearer ", "Sign-in required"},
		{"forged token", "Bearer " + token + "x", "Session is invalid or expired"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			handler := RequireSession(sessions, func(http.ResponseWriter, *http.Request) { called = true })

			req := httptest.NewRequest(http.MethodPost, "/v1/notes", nil)
			if tt.authorization != "" {
				req.Header.Set("Authorization", tt.authorization)
			}
			rec := httptest.NewRecorder()
			handler(rec, req)

			if called {
				t.Error("the handler ran without a valid session")
			}
			if rec.Code != http.StatusUnauthorized {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
			}
			if !strings.HasPrefix(rec.Header().Get("WWW-Authenticate"), "Bearer") {
				t.Errorf("WWW-Authenticate = %q, want a Bearer challenge", rec.Header().Get("WWW-Authenticate"))
			}
			var body map[string]any
			if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
				t.Fatalf("response is not JSON: %v", err)
			}
			if body["success"] != false || body["error"] != tt.wantError {
				t.Errorf("body = %v, want success false and error %q", body, tt.wantError)
			}
		})
	}
}

func TestRequireSessionWithoutSignIn(t *testing.T) {
	called := false
	handler := RequireSession(nil, func(http.ResponseWriter, *http.Request) { called = true })
	rec := httptest.NewRecorder()

	handler(rec, httptest.NewRequest(http.MethodPost, "/v1/notes", nil))

	if called {
		t.Error("the handler ran with sign-in unconfigured")
	}
	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
	}
}

func TestPreflightAllowsTheAuthorizationHeader(t *testing.T) {
	rec := httptest.NewRecorder()
	Preflight(rec, httptest.NewRequest(http.MethodOptions, "/v1/notes", nil))

	if got := rec.Header().Get("Access-Control-Allow-Headers"); !strings.Contains(got, "Authorization") {
		t.Errorf("Access-Control-Allow-Headers = %q, want it to include Authorization", got)
	}
}
