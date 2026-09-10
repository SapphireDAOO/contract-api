package httpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func decodeEnvelope(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not JSON: %v (body %q)", err, rec.Body.String())
	}
	return body
}

func TestWriteFailure(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteFailure(rec, http.StatusBadRequest, "Invalid request body")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if got, want := rec.Header().Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}

	body := decodeEnvelope(t, rec)
	if body["success"] != false {
		t.Errorf("success = %v, want false", body["success"])
	}
	if body["error"] != "Invalid request body" {
		t.Errorf("error = %v, want the message", body["error"])
	}
}

func TestWriteFailureStatuses(t *testing.T) {
	for _, status := range []int{
		http.StatusBadRequest,
		http.StatusRequestEntityTooLarge,
		http.StatusInternalServerError,
	} {
		rec := httptest.NewRecorder()

		WriteFailure(rec, status, "nope")

		if rec.Code != status {
			t.Errorf("status = %d, want %d", rec.Code, status)
		}
		if body := decodeEnvelope(t, rec); body["success"] != false {
			t.Errorf("success = %v, want false", body["success"])
		}
	}
}

func TestWriteSuccess(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteSuccess(rec, map[string]any{"txHash": "0xdeadbeef"})

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got, want := rec.Header().Get("Content-Type"), "application/json"; got != want {
		t.Errorf("Content-Type = %q, want %q", got, want)
	}

	body := decodeEnvelope(t, rec)
	if body["success"] != true {
		t.Errorf("success = %v, want true", body["success"])
	}
	if body["txHash"] != "0xdeadbeef" {
		t.Errorf("txHash = %v, want the body passed in", body["txHash"])
	}
}

func TestWriteSuccessWithAnEmptyBody(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteSuccess(rec, map[string]any{})

	body := decodeEnvelope(t, rec)
	if len(body) != 1 || body["success"] != true {
		t.Errorf("body = %v, want just the success flag", body)
	}
}

// The flag reflects the writer that was called, not whatever the caller put
// in the map.
func TestWriteSuccessOverridesAFalseFlag(t *testing.T) {
	rec := httptest.NewRecorder()

	WriteSuccess(rec, map[string]any{"success": false})

	if body := decodeEnvelope(t, rec); body["success"] != true {
		t.Errorf("success = %v, want true", body["success"])
	}
}

func TestEnvelopesAreDistinctFromTheRevertShape(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteFailure(rec, http.StatusBadRequest, "nope")
	envelope := decodeEnvelope(t, rec)

	if _, ok := envelope["reason"]; ok {
		t.Error("the success envelope carries a reason key")
	}
	if _, ok := envelope["success"]; !ok {
		t.Error("the success envelope has no success key")
	}
}
