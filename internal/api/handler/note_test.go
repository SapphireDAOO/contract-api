package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SapphireDAOO/contract-api/internal/note"
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

func noteRequestFor(path, body string) *http.Request {
	return httptest.NewRequest(http.MethodPost, path, strings.NewReader(body))
}

func TestWriteNoteRejectsAMalformedBody(t *testing.T) {
	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	h.WriteNote(rec, noteRequestFor("/v1/notes", "not json"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if body := noteResponse(t, rec); body["error"] != "Invalid request body" {
		t.Errorf("error = %v, want the malformed-body message", body["error"])
	}
}

func TestWriteNoteBadRequests(t *testing.T) {
	const author = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantErr    string
	}{
		{
			name:       "invoice id is not a number",
			body:       `{"invoiceId":"abc","author":"` + author + `","content":"0xdeadbeef"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    "invoiceId must be a base-10 integer",
		},
		{
			name:       "missing invoice id",
			body:       `{"author":"` + author + `","content":"0xdeadbeef"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    "invoiceId must be a base-10 integer",
		},
		{
			name:       "invalid author",
			body:       `{"invoiceId":"42","author":"nope","content":"0xdeadbeef"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    "Invalid author address",
		},
		{
			name:       "zero author",
			body:       `{"invoiceId":"42","author":"` + zeroAddress + `","content":"0xdeadbeef"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    "Invalid author address",
		},
		{
			name:       "missing content",
			body:       `{"invoiceId":"42","author":"` + author + `"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    "content is required",
		},
		{
			name:       "content is not hex",
			body:       `{"invoiceId":"42","author":"` + author + `","content":"Left at the door"}`,
			wantStatus: http.StatusBadRequest,
			wantErr:    "content must be 0x-prefixed hex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ContractHandler{}
			rec := httptest.NewRecorder()

			h.WriteNote(rec, noteRequestFor("/v1/notes", tt.body))

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			body := noteResponse(t, rec)
			if body["success"] != false {
				t.Errorf("success = %v, want false", body["success"])
			}
			if body["error"] != tt.wantErr {
				t.Errorf("error = %v, want %q", body["error"], tt.wantErr)
			}
		})
	}
}

func TestWriteNoteRejectsOversizedContent(t *testing.T) {
	const author = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	oversized := hexutil.Encode(make([]byte, note.MaxContentBytes+1))
	h.WriteNote(rec, noteRequestFor("/v1/notes",
		`{"invoiceId":"42","author":"`+author+`","content":"`+oversized+`"}`))

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusRequestEntityTooLarge)
	}
	if body := noteResponse(t, rec); body["error"] != "content is too large" {
		t.Errorf("error = %v, want the size message", body["error"])
	}
}

func TestOpenNoteRejectsAMalformedBody(t *testing.T) {
	h := &ContractHandler{}
	rec := httptest.NewRecorder()

	h.OpenNote(rec, noteRequestFor("/v1/notes/open", "not json"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if body := noteResponse(t, rec); body["error"] != "Invalid request body" {
		t.Errorf("error = %v, want the malformed-body message", body["error"])
	}
}

func TestOpenNoteBadRequests(t *testing.T) {
	const author = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	tests := []struct {
		name    string
		body    string
		wantErr string
	}{
		{
			name:    "invoice id is not a number",
			body:    `{"invoiceId":"abc","noteId":"1","author":"` + author + `"}`,
			wantErr: "invoiceId must be a base-10 integer",
		},
		{
			name:    "note id is not a number",
			body:    `{"invoiceId":"42","noteId":"abc","author":"` + author + `"}`,
			wantErr: "noteId must be a base-10 integer",
		},
		{
			name:    "missing note id",
			body:    `{"invoiceId":"42","author":"` + author + `"}`,
			wantErr: "noteId must be a base-10 integer",
		},
		{
			name:    "invalid author",
			body:    `{"invoiceId":"42","noteId":"1","author":"nope"}`,
			wantErr: "Invalid author address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ContractHandler{}
			rec := httptest.NewRecorder()

			h.OpenNote(rec, noteRequestFor("/v1/notes/open", tt.body))

			if rec.Code != http.StatusBadRequest {
				t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
			}
			body := noteResponse(t, rec)
			if body["success"] != false {
				t.Errorf("success = %v, want false", body["success"])
			}
			if body["error"] != tt.wantErr {
				t.Errorf("error = %v, want %q", body["error"], tt.wantErr)
			}
		})
	}
}
