package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCORSSetsTheOpenOriginHeader(t *testing.T) {
	var reached bool
	handler := CORS(called(&reached))

	req := httptest.NewRequest(http.MethodPost, "/v1/fee-receivers", nil)
	req.Header.Set("Origin", "https://example.com")
	rec := httptest.NewRecorder()

	handler(rec, req)

	if !reached {
		t.Error("CORS did not call through to the wrapped handler")
	}
	if got, want := rec.Header().Get("Access-Control-Allow-Origin"), "*"; got != want {
		t.Errorf("Access-Control-Allow-Origin = %q, want %q", got, want)
	}

	if got, want := rec.Header().Get("Vary"), "Origin"; got != want {
		t.Errorf("Vary = %q, want %q", got, want)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestCORSWithoutAnOriginHeader(t *testing.T) {
	var reached bool
	handler := CORS(called(&reached))

	rec := httptest.NewRecorder()
	handler(rec, httptest.NewRequest(http.MethodPost, "/v1/fee-receivers", nil))

	if !reached {
		t.Error("CORS did not call through to the wrapped handler")
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("Access-Control-Allow-Origin = %q, want %q", got, "*")
	}
}

func TestCORSHeadersSurviveAHandlerThatWrites(t *testing.T) {
	handler := CORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"success"}`))
	})

	rec := httptest.NewRecorder()
	handler(rec, httptest.NewRequest(http.MethodPost, "/v1/fee-receivers", nil))

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusCreated)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("Access-Control-Allow-Origin = %q, want %q", got, "*")
	}
}

func TestCORSHeadersAreSetOnErrorResponses(t *testing.T) {
	handler := CORS(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusBadRequest)
	})

	rec := httptest.NewRecorder()
	handler(rec, httptest.NewRequest(http.MethodPost, "/v1/fee-receivers", nil))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("Access-Control-Allow-Origin = %q, want %q", got, "*")
	}
}

func TestPreflight(t *testing.T) {
	req := httptest.NewRequest(http.MethodOptions, "/v1/fee-receivers", nil)
	req.Header.Set("Origin", "https://example.com")
	req.Header.Set("Access-Control-Request-Method", "POST")
	rec := httptest.NewRecorder()

	Preflight(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}

	want := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": allowedMethods,
		"Access-Control-Allow-Headers": allowedHeaders,
		"Access-Control-Max-Age":       preflightMaxAge,
		"Vary":                         "Origin",
	}
	for header, wantValue := range want {
		if got := rec.Header().Get(header); got != wantValue {
			t.Errorf("%s = %q, want %q", header, got, wantValue)
		}
	}
}

func TestPreflightHasNoBody(t *testing.T) {
	rec := httptest.NewRecorder()

	Preflight(rec, httptest.NewRequest(http.MethodOptions, "/v1/fee-receivers", nil))

	if rec.Body.Len() != 0 {
		t.Errorf("preflight wrote %q, want an empty body", rec.Body.String())
	}
}

func TestPreflightAdvertisesTheMethodsTheAPIServes(t *testing.T) {
	rec := httptest.NewRecorder()

	Preflight(rec, httptest.NewRequest(http.MethodOptions, "/v1/fee-receivers", nil))

	methods := rec.Header().Get("Access-Control-Allow-Methods")
	for _, method := range []string{http.MethodPost, http.MethodOptions} {
		if !contains(methods, method) {
			t.Errorf("Access-Control-Allow-Methods = %q, want it to include %s", methods, method)
		}
	}
	if got := rec.Header().Get("Access-Control-Allow-Headers"); !contains(got, "Content-Type") {
		t.Errorf("Access-Control-Allow-Headers = %q, want it to include Content-Type", got)
	}
}

func contains(list, value string) bool {
	for _, part := range splitAndTrim(list) {
		if part == value {
			return true
		}
	}
	return false
}

func splitAndTrim(list string) []string {
	var parts []string
	for _, part := range strings.Split(list, ",") {
		parts = append(parts, strings.TrimSpace(part))
	}
	return parts
}
