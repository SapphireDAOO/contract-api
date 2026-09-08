package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SapphireDAOO/contract-api/internal/note"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func noteResponse(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v (body %q)", err, rec.Body.String())
	}
	return body
}

func noteRequestFor(body string) *http.Request {
	return httptest.NewRequest(http.MethodPost, "/v1/note", strings.NewReader(body))
}

func TestParseAddress(t *testing.T) {
	const valid = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	tests := []struct {
		name   string
		value  string
		want   string
		wantOK bool
	}{
		{name: "checksummed", value: valid, want: valid, wantOK: true},
		{name: "lowercase", value: strings.ToLower(valid), want: valid, wantOK: true},
		{name: "surrounding whitespace is trimmed", value: "  " + valid + "  ", want: valid, wantOK: true},
		{name: "empty", value: ""},
		{name: "not hex", value: "not-an-address"},
		{name: "too short", value: "0x123"},
		{name: "missing prefix", value: "5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed", wantOK: true, want: valid},

		{name: "zero address", value: zeroAddress},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := parseAddress(tt.value)

			if ok != tt.wantOK {
				t.Fatalf("parseAddress(%q) ok = %v, want %v", tt.value, ok, tt.wantOK)
			}
			if !ok {
				if got != (common.Address{}) {
					t.Errorf("parseAddress(%q) = %s on a miss, want the zero address", tt.value, got)
				}
				return
			}
			if got != common.HexToAddress(tt.want) {
				t.Errorf("parseAddress(%q) = %s, want %s", tt.value, got, tt.want)
			}
		})
	}
}

func TestWriteNoteError(t *testing.T) {
	rec := httptest.NewRecorder()

	writeNoteError(rec, http.StatusBadRequest, "Invalid request body")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if got, want := rec.Header().Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}

	body := noteResponse(t, rec)
	if body["success"] != false {
		t.Errorf("success = %v, want false", body["success"])
	}
	if body["error"] != "Invalid request body" {
		t.Errorf("error = %v, want the message", body["error"])
	}
}

func TestWriteNoteSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	writeNoteSuccess(rec, map[string]any{"payload": "0xdeadbeef"})

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	body := noteResponse(t, rec)
	if body["success"] != true {
		t.Errorf("success = %v, want true", body["success"])
	}
	if body["payload"] != "0xdeadbeef" {
		t.Errorf("payload = %v, want the body passed in", body["payload"])
	}
}

func TestWriteNoteSuccessAlwaysReportsSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	writeNoteSuccess(rec, map[string]any{"success": false})

	if body := noteResponse(t, rec); body["success"] != true {
		t.Errorf("success = %v, want true", body["success"])
	}
}

func TestHandleNoteRejectsAMalformedBody(t *testing.T) {
	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	h.HandleNote(rec, noteRequestFor("not json"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if body := noteResponse(t, rec); body["error"] != "Invalid request body" {
		t.Errorf("error = %v, want the malformed-body message", body["error"])
	}
}

func TestHandleNoteRejectsAnUnknownAction(t *testing.T) {
	h := &ContractHandler{}

	for _, action := range []string{"", "delete", "Encrypt", "unknown"} {
		t.Run(action, func(t *testing.T) {
			rec := httptest.NewRecorder()

			h.HandleNote(rec, noteRequestFor(`{"action":"`+action+`"}`))

			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
			}
			if body := noteResponse(t, rec); body["error"] != "Unknown action" {
				t.Errorf("error = %v, want %q", body["error"], "Unknown action")
			}
		})
	}
}

func TestHandleNoteEncrypt(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", "the-test-secret")

	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	h.HandleNote(rec, noteRequestFor(`{"action":"encrypt","content":"  a note  "}`))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d (body %q)", rec.Code, http.StatusOK, rec.Body.String())
	}

	body := noteResponse(t, rec)
	if body["success"] != true {
		t.Errorf("success = %v, want true", body["success"])
	}

	payload, ok := body["payload"].(string)
	if !ok {
		t.Fatalf("payload is %T, want a hex string", body["payload"])
	}

	raw, err := hexutil.Decode(payload)
	if err != nil {
		t.Fatalf("payload %q is not hex: %v", payload, err)
	}

	got, err := note.DecryptNote(string(raw), "the-test-secret")
	if err != nil {
		t.Fatalf("DecryptNote returned %v", err)
	}
	if got != "a note" {
		t.Errorf("decrypted %q, want the trimmed content", got)
	}
}

func TestHandleNoteEncryptRejectsEmptyContent(t *testing.T) {
	h := &ContractHandler{}

	for _, content := range []string{"", "   ", "\\t"} {
		t.Run(content, func(t *testing.T) {
			rec := httptest.NewRecorder()

			h.HandleNote(rec, noteRequestFor(`{"action":"encrypt","content":"`+content+`"}`))

			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
			}
			if body := noteResponse(t, rec); body["error"] != note.ErrNoteContentRequired.Error() {
				t.Errorf("error = %v, want %q", body["error"], note.ErrNoteContentRequired)
			}
		})
	}
}

func TestHandleNoteEncryptRejectsAnOverlongNote(t *testing.T) {
	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	long := strings.Repeat("a", note.MaxNoteLength+1)
	h.HandleNote(rec, noteRequestFor(`{"action":"encrypt","content":"`+long+`"}`))

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusRequestEntityTooLarge)
	}
	if body := noteResponse(t, rec); body["error"] != note.ErrNoteTooLong.Error() {
		t.Errorf("error = %v, want %q", body["error"], note.ErrNoteTooLong)
	}
}

func TestHandleNoteEncryptAcceptsANoteAtTheLimit(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", "the-test-secret")

	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	atLimit := strings.Repeat("a", note.MaxNoteLength)
	h.HandleNote(rec, noteRequestFor(`{"action":"encrypt","content":"`+atLimit+`"}`))

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d (body %q)", rec.Code, http.StatusOK, rec.Body.String())
	}
}

func TestHandleNoteEncryptWithoutASecret(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", "")

	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	h.HandleNote(rec, noteRequestFor(`{"action":"encrypt","content":"a note"}`))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d (body %q)", rec.Code, http.StatusOK, rec.Body.String())
	}

	payload := noteResponse(t, rec)["payload"].(string)
	raw, err := hexutil.Decode(payload)
	if err != nil {
		t.Fatalf("payload %q is not hex: %v", payload, err)
	}
	if string(raw) != "a note" {
		t.Errorf("payload decoded to %q, want the plaintext", raw)
	}
}

func TestHandleNoteEncryptIsRandomised(t *testing.T) {
	t.Setenv("NOTES_SECRET_KEY", "the-test-secret")

	h := &ContractHandler{}

	payloads := make([]string, 2)
	for i := range payloads {
		rec := httptest.NewRecorder()
		h.HandleNote(rec, noteRequestFor(`{"action":"encrypt","content":"a note"}`))
		payloads[i] = noteResponse(t, rec)["payload"].(string)
	}

	if payloads[0] == payloads[1] {
		t.Error("encrypting the same note twice produced identical payloads")
	}
}
