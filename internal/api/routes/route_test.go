package routes

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

	"github.com/SapphireDAOO/contract-api/internal/api/handler"
	"github.com/SapphireDAOO/contract-api/internal/auth"
)

func call(t *testing.T, router http.Handler, method, path, body, token string) (*httptest.ResponseRecorder, map[string]any) {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	var decoded map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &decoded)
	return rec, decoded
}

func TestNotesSignInFlow(t *testing.T) {
	service, err := auth.New("0123456789abcdef0123456789abcdef", []string{"app.example.com"}, 84532)
	if err != nil {
		t.Fatalf("auth.New: %v", err)
	}
	router := Route(&handler.ContractHandler{Auth: service})
	note := `{"invoiceId":"42","content":"0xdeadbeef"}`

	if rec, _ := call(t, router, http.MethodPost, "/v1/notes", note, ""); rec.Code != http.StatusUnauthorized {
		t.Fatalf("note without a session: status = %d, want 401", rec.Code)
	}

	rec, body := call(t, router, http.MethodGet, "/v1/auth/nonce", "", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("nonce: status = %d, want 200", rec.Code)
	}
	if rec.Header().Get("Access-Control-Allow-Origin") == "" {
		t.Error("nonce response has no CORS header")
	}
	nonce, _ := body["nonce"].(string)

	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generating key: %v", err)
	}
	msg, err := siwe.InitMessage("app.example.com", crypto.PubkeyToAddress(key.PublicKey).Hex(),
		"https://app.example.com", nonce, map[string]any{"chainId": 84532})
	if err != nil {
		t.Fatalf("building message: %v", err)
	}
	sig, err := crypto.Sign(accounts.TextHash([]byte(msg.String())), key)
	if err != nil {
		t.Fatalf("signing: %v", err)
	}
	verify, err := json.Marshal(map[string]string{"message": msg.String(), "signature": hexutil.Encode(sig)})
	if err != nil {
		t.Fatalf("encoding: %v", err)
	}

	rec, body = call(t, router, http.MethodPost, "/v1/auth/verify", string(verify), "")
	if rec.Code != http.StatusOK {
		t.Fatalf("verify: status = %d, want 200 (body %v)", rec.Code, body)
	}
	token, _ := body["token"].(string)

	// No notes contract is configured, so a request that got past the
	// session and validation fails at the contract call.
	if rec, _ := call(t, router, http.MethodPost, "/v1/notes", note, token); rec.Code != http.StatusInternalServerError {
		t.Errorf("signed-in note: status = %d, want the contract-call failure 500", rec.Code)
	}
	if rec, _ := call(t, router, http.MethodPost, "/v1/notes", note, token+"x"); rec.Code != http.StatusUnauthorized {
		t.Errorf("note with a forged token: status = %d, want 401", rec.Code)
	}
}

func TestNotesWhenSignInIsNotConfigured(t *testing.T) {
	router := Route(&handler.ContractHandler{})

	if rec, _ := call(t, router, http.MethodPost, "/v1/notes", `{}`, "anything"); rec.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want 503", rec.Code)
	}
}
