package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func captureLog(t *testing.T, run func()) []map[string]any {
	t.Helper()

	var buf bytes.Buffer
	previous := slog.Default()
	slog.SetDefault(slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug})))
	t.Cleanup(func() { slog.SetDefault(previous) })

	run()

	var records []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
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

func TestErrorResponsesAreLogged(t *testing.T) {
	tests := []struct {
		name       string
		write      func(http.ResponseWriter)
		wantLevel  string
		wantStatus float64
		wantReason string
	}{
		{
			name: "client error logs as a warning",
			write: func(w http.ResponseWriter) {
				WriteHTTPErrorWithStatus(w, http.StatusBadRequest, errors.New("bad input"), "invalid request body")
			},
			wantLevel: "WARN", wantStatus: 400, wantReason: "bad input",
		},
		{
			name: "server error logs as an error",
			write: func(w http.ResponseWriter) {
				WriteHTTPErrorWithStatus(w, http.StatusServiceUnavailable, errors.New("oracle down"), "unavailable")
			},
			wantLevel: "ERROR", wantStatus: 503, wantReason: "oracle down",
		},
		{
			name: "mapped revert logs the description",
			write: func(w http.ResponseWriter) {
				WriteMappedRevertError(w, revertWith("0x715d9228"), "transaction failed")
			},
			wantLevel: "WARN", wantStatus: 404, wantReason: "The specified invoice does not exist.",
		},
		{
			name: "unmapped revert logs as a server error",
			write: func(w http.ResponseWriter) {
				WriteMappedRevertError(w, errors.New("connection refused"), "transaction failed")
			},
			wantLevel: "ERROR", wantStatus: 500, wantReason: "connection refused",
		},
		{
			name: "failure envelope is logged too",
			write: func(w http.ResponseWriter) {
				WriteFailure(w, http.StatusRequestEntityTooLarge, "content is too large")
			},
			wantLevel: "WARN", wantStatus: 413, wantReason: "content is too large",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			records := captureLog(t, func() { tt.write(httptest.NewRecorder()) })

			if len(records) != 1 {
				t.Fatalf("got %d log records, want 1: %v", len(records), records)
			}
			record := records[0]

			if record["level"] != tt.wantLevel {
				t.Errorf("level = %v, want %v", record["level"], tt.wantLevel)
			}
			if record["status"] != tt.wantStatus {
				t.Errorf("status = %v, want %v", record["status"], tt.wantStatus)
			}
			if record["reason"] != tt.wantReason {
				t.Errorf("reason = %v, want %q", record["reason"], tt.wantReason)
			}
		})
	}
}

// A revert's description replaces the error text in the response, so the
// original is kept in the log where it would otherwise be lost.
func TestTheOriginalErrorIsLoggedAlongsideAMappedReason(t *testing.T) {
	records := captureLog(t, func() {
		WriteMappedRevertError(httptest.NewRecorder(), revertWith("0x715d9228"), "transaction failed")
	})

	if len(records) != 1 {
		t.Fatalf("got %d log records, want 1", len(records))
	}
	if records[0]["error"] != "execution reverted" {
		t.Errorf("error = %v, want the underlying error text", records[0]["error"])
	}
}

func TestTheErrorIsNotRepeatedWhenItIsAlreadyTheReason(t *testing.T) {
	records := captureLog(t, func() {
		WriteHTTPErrorWithStatus(httptest.NewRecorder(), http.StatusBadRequest,
			errors.New("bad input"), "invalid request body")
	})

	if len(records) != 1 {
		t.Fatalf("got %d log records, want 1", len(records))
	}
	if _, ok := records[0]["error"]; ok {
		t.Errorf("error was logged as well as reason, duplicating it: %v", records[0])
	}
}

func TestSuccessfulResponsesAreNotLoggedAsFailures(t *testing.T) {
	records := captureLog(t, func() {
		WriteSuccess(httptest.NewRecorder(), map[string]any{"txHash": "0xabc"})
	})

	if len(records) != 0 {
		t.Errorf("a success wrote %d log records, want none: %v", len(records), records)
	}
}
