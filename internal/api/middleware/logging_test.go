package middleware

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func captureLog(t *testing.T, level slog.Level, run func()) []map[string]any {
	t.Helper()

	var buf bytes.Buffer
	previous := slog.Default()
	slog.SetDefault(slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: level})))
	t.Cleanup(func() { slog.SetDefault(previous) })

	run()

	var records []map[string]any
	for line := range strings.SplitSeq(strings.TrimSpace(buf.String()), "\n") {
		if line == "" {
			continue
		}
		var record map[string]any
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			t.Fatalf("log line is not JSON: %v (%q)", err, line)
		}
		records = append(records, record)
	}
	return records
}

func TestLoggingRecordsTheRequest(t *testing.T) {
	handler := Logging(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"ok":true}`))
	})

	records := captureLog(t, slog.LevelDebug, func() {
		req := httptest.NewRequest(http.MethodPost, "/v1/notes?a=1", nil)
		handler(httptest.NewRecorder(), req)
	})

	if len(records) != 1 {
		t.Fatalf("got %d log records, want 1", len(records))
	}
	record := records[0]

	if record["msg"] != "request" {
		t.Errorf("msg = %v, want request", record["msg"])
	}
	if record["method"] != http.MethodPost {
		t.Errorf("method = %v, want POST", record["method"])
	}
	if record["path"] != "/v1/notes" {
		t.Errorf("path = %v, want /v1/notes", record["path"])
	}
	if record["query"] != "a=1" {
		t.Errorf("query = %v, want a=1", record["query"])
	}
	if record["status"] != float64(http.StatusCreated) {
		t.Errorf("status = %v, want 201", record["status"])
	}
	if record["bytes"] != float64(len(`{"ok":true}`)) {
		t.Errorf("bytes = %v, want %d", record["bytes"], len(`{"ok":true}`))
	}
	if _, ok := record["duration"]; !ok {
		t.Error("no duration recorded")
	}
}

func TestLoggingDefaultsToOKWhenTheHandlerNeverSetsAStatus(t *testing.T) {
	handler := Logging(func(w http.ResponseWriter, r *http.Request) {})

	records := captureLog(t, slog.LevelDebug, func() {
		handler(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))
	})

	if len(records) != 1 {
		t.Fatalf("got %d log records, want 1", len(records))
	}
	if records[0]["status"] != float64(http.StatusOK) {
		t.Errorf("status = %v, want 200", records[0]["status"])
	}
}

func TestLoggingLevelFollowsTheStatus(t *testing.T) {
	tests := []struct {
		status int
		want   string
	}{
		{http.StatusOK, "INFO"},
		{http.StatusNoContent, "INFO"},
		{http.StatusBadRequest, "WARN"},
		{http.StatusForbidden, "WARN"},
		{http.StatusNotFound, "WARN"},
		{http.StatusInternalServerError, "ERROR"},
		{http.StatusBadGateway, "ERROR"},
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.status), func(t *testing.T) {
			handler := Logging(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
			})

			records := captureLog(t, slog.LevelDebug, func() {
				handler(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil))
			})

			if len(records) != 1 {
				t.Fatalf("got %d log records, want 1", len(records))
			}
			if records[0]["level"] != tt.want {
				t.Errorf("level = %v, want %v", records[0]["level"], tt.want)
			}
		})
	}
}

// The API key arrives in a header and note content in the body; neither may
// reach the log.
func TestLoggingRecordsNoHeadersOrBody(t *testing.T) {
	handler := Logging(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	records := captureLog(t, slog.LevelDebug, func() {
		req := httptest.NewRequest(http.MethodPost, "/v1/notes",
			strings.NewReader(`{"content":"0xsecretciphertext"}`))
		req.Header.Set("X-API-KEY", "super-secret-key")
		handler(httptest.NewRecorder(), req)
	})

	line, err := json.Marshal(records)
	if err != nil {
		t.Fatalf("re-marshalling the records: %v", err)
	}
	for _, secret := range []string{"super-secret-key", "0xsecretciphertext", "X-API-KEY"} {
		if strings.Contains(string(line), secret) {
			t.Errorf("the log contains %q: %s", secret, line)
		}
	}
}

func TestLoggingPassesTheResponseThrough(t *testing.T) {
	handler := Logging(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("brewing"))
	})

	rec := httptest.NewRecorder()
	captureLog(t, slog.LevelDebug, func() {
		handler(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	})

	if rec.Code != http.StatusTeapot {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusTeapot)
	}
	if rec.Body.String() != "brewing" {
		t.Errorf("body = %q, want %q", rec.Body.String(), "brewing")
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", got)
	}
}
