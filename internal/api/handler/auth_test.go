package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spruceid/siwe-go"

	"github.com/SapphireDAOO/contract-api/internal/auth"
)

const (
	signInDomain = "app.example.com"
	signInChain  = 84532
)

func newAuthHandler(t *testing.T) *ContractHandler {
	t.Helper()
	service, err := auth.New("0123456789abcdef0123456789abcdef", []string{signInDomain}, signInChain)
	if err != nil {
		t.Fatalf("auth.New: %v", err)
	}
	return &ContractHandler{Auth: service}
}

func authBody(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v (body %q)", err, rec.Body.String())
	}
	return body
}

func requestNonce(t *testing.T, h *ContractHandler) string {
	t.Helper()
	rec := httptest.NewRecorder()
	h.AuthNonce(rec, httptest.NewRequest(http.MethodGet, "/v1/auth/nonce", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("nonce status = %d, want 200 (body %q)", rec.Code, rec.Body.String())
	}
	nonce, _ := authBody(t, rec)["nonce"].(string)
	return nonce
}

func signInRequest(t *testing.T, message, signature string) *http.Request {
	t.Helper()
	body, err := json.Marshal(map[string]string{"message": message, "signature": signature})
	if err != nil {
		t.Fatalf("encoding request: %v", err)
	}
	return httptest.NewRequest(http.MethodPost, "/v1/auth/verify", strings.NewReader(string(body)))
}

func signedSignIn(t *testing.T, nonce string) (address, message, signature string) {
	t.Helper()
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generating key: %v", err)
	}
	address = crypto.PubkeyToAddress(key.PublicKey).Hex()

	msg, err := siwe.InitMessage(signInDomain, address, "https://"+signInDomain, nonce,
		map[string]any{"chainId": signInChain, "statement": "Sign in to write and read invoice notes."})
	if err != nil {
		t.Fatalf("building message: %v", err)
	}
	sig, err := crypto.Sign(accounts.TextHash([]byte(msg.String())), key)
	if err != nil {
		t.Fatalf("signing: %v", err)
	}
	return address, msg.String(), hexutil.Encode(sig)
}

func TestAuthNonce(t *testing.T) {
	h := newAuthHandler(t)
	rec := httptest.NewRecorder()

	h.AuthNonce(rec, httptest.NewRequest(http.MethodGet, "/v1/auth/nonce", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if got := rec.Header().Get("Cache-Control"); got != "no-store" {
		t.Errorf("Cache-Control = %q, want no-store", got)
	}
	body := authBody(t, rec)
	if body["success"] != true || body["nonce"] == "" || body["expiresAt"] == "" {
		t.Errorf("body = %v, want success with a nonce and expiresAt", body)
	}
}

func TestSignInIssuesAWorkingSession(t *testing.T) {
	h := newAuthHandler(t)
	address, message, signature := signedSignIn(t, requestNonce(t, h))
	rec := httptest.NewRecorder()

	h.SignIn(rec, signInRequest(t, message, signature))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200 (body %q)", rec.Code, rec.Body.String())
	}
	body := authBody(t, rec)
	if body["address"] != address {
		t.Errorf("address = %v, want %s", body["address"], address)
	}
	token, _ := body["token"].(string)
	got, err := h.Auth.SessionAddress(token)
	if err != nil {
		t.Fatalf("the issued token is not a valid session: %v", err)
	}
	if got.Hex() != address {
		t.Errorf("session address = %s, want %s", got.Hex(), address)
	}
}

func TestSignInRejects(t *testing.T) {
	tests := []struct {
		name       string
		request    func(t *testing.T, h *ContractHandler) *http.Request
		wantStatus int
	}{
		{"malformed body", func(t *testing.T, h *ContractHandler) *http.Request {
			return httptest.NewRequest(http.MethodPost, "/v1/auth/verify", strings.NewReader("not json"))
		}, http.StatusBadRequest},
		{"not a SIWE message", func(t *testing.T, h *ContractHandler) *http.Request {
			return signInRequest(t, "hello", "0x00")
		}, http.StatusBadRequest},
		{"signature from another message", func(t *testing.T, h *ContractHandler) *http.Request {
			_, message, _ := signedSignIn(t, requestNonce(t, h))
			_, _, otherSignature := signedSignIn(t, requestNonce(t, h))
			return signInRequest(t, message, otherSignature)
		}, http.StatusUnauthorized},
		{"nonce not issued here", func(t *testing.T, h *ContractHandler) *http.Request {
			_, message, signature := signedSignIn(t, "abcdef0123456789")
			return signInRequest(t, message, signature)
		}, http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newAuthHandler(t)
			rec := httptest.NewRecorder()

			h.SignIn(rec, tt.request(t, h))

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d (body %q)", rec.Code, tt.wantStatus, rec.Body.String())
			}
			if body := authBody(t, rec); body["success"] != false {
				t.Errorf("success = %v, want false", body["success"])
			}
		})
	}
}

func TestSignInWhenNotConfigured(t *testing.T) {
	h := &ContractHandler{}

	nonce := httptest.NewRecorder()
	h.AuthNonce(nonce, httptest.NewRequest(http.MethodGet, "/v1/auth/nonce", nil))
	verify := httptest.NewRecorder()
	h.SignIn(verify, signInRequest(t, "hello", "0x00"))

	for name, rec := range map[string]*httptest.ResponseRecorder{"nonce": nonce, "verify": verify} {
		if rec.Code != http.StatusServiceUnavailable {
			t.Errorf("%s status = %d, want %d", name, rec.Code, http.StatusServiceUnavailable)
		}
	}
}

func TestNotesRequireASignedInAccount(t *testing.T) {
	h := &ContractHandler{}
	body := `{"invoiceId":"42","noteId":"1","content":"0xdeadbeef"}`

	for name, call := range map[string]http.HandlerFunc{"write": h.WriteNote, "open": h.OpenNote} {
		rec := httptest.NewRecorder()
		call(rec, httptest.NewRequest(http.MethodPost, "/v1/notes", strings.NewReader(body)))

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("%s: status = %d, want %d", name, rec.Code, http.StatusUnauthorized)
		}
	}
}

func TestNotesRefuseAnotherAccountsAuthor(t *testing.T) {
	const someoneElse = "0x329C3E1bEa46Abc22F307eE30Cbb522B82Fe7082"
	h := &ContractHandler{}

	for name, tc := range map[string]struct {
		call http.HandlerFunc
		body string
	}{
		"write": {h.WriteNote, `{"invoiceId":"42","author":"` + someoneElse + `","content":"0xdeadbeef"}`},
		"open":  {h.OpenNote, `{"invoiceId":"42","noteId":"1","author":"` + someoneElse + `"}`},
	} {
		rec := httptest.NewRecorder()
		tc.call(rec, noteRequestFor("/v1/notes", tc.body))

		if rec.Code != http.StatusForbidden {
			t.Errorf("%s: status = %d, want %d", name, rec.Code, http.StatusForbidden)
		}
		if body := noteResponse(t, rec); body["error"] != "author is not the signed-in account" {
			t.Errorf("%s: error = %v, want the mismatch message", name, body["error"])
		}
	}
}

func TestWriteNoteWithoutAuthorUsesTheSignedInAccount(t *testing.T) {
	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	h.WriteNote(rec, noteRequestFor("/v1/notes", `{"invoiceId":"42","content":"0xdeadbeef"}`))

	// No contract is configured, so a request that passed validation fails
	// at the contract call rather than with a 4xx.
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want the contract-call failure %d (body %q)",
			rec.Code, http.StatusInternalServerError, rec.Body.String())
	}
}
