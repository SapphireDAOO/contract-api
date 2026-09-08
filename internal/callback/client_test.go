package callback

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBuildCallbackURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		orderId string
		action  string
		want    string
		wantErr string
	}{
		{
			name:    "path is appended",
			baseURL: "https://market.example.com/api/callback",
			orderId: "ORDER-1",
			action:  releaseCallbackAction,
			want:    "https://market.example.com/api/callback/ORDER-1/escrowReleased",
		},
		{
			name:    "a trailing slash is not doubled",
			baseURL: "https://market.example.com/api/callback/",
			orderId: "ORDER-1",
			action:  refundCallbackAction,
			want:    "https://market.example.com/api/callback/ORDER-1/refundSent",
		},
		{
			name:    "several trailing slashes are trimmed",
			baseURL: "https://market.example.com/api/callback///",
			orderId: "ORDER-1",
			action:  paymentReceivedCallbackAction,
			want:    "https://market.example.com/api/callback/ORDER-1/paymentReceived",
		},
		{
			name:    "a format template is filled in",
			baseURL: "https://market.example.com/orders/%s/events/%s",
			orderId: "ORDER-1",
			action:  releaseCallbackAction,
			want:    "https://market.example.com/orders/ORDER-1/events/escrowReleased",
		},
		{
			name:    "an indexed template is filled in",
			baseURL: "https://market.example.com/%[2]s/%[1]s",
			orderId: "ORDER-1",
			action:  releaseCallbackAction,
			want:    "https://market.example.com/escrowReleased/ORDER-1",
		},
		{
			name:    "empty order id",
			baseURL: "https://market.example.com/api/callback",
			orderId: "",
			action:  releaseCallbackAction,
			wantErr: "orderId is required",
		},
		{
			name:    "whitespace order id",
			baseURL: "https://market.example.com/api/callback",
			orderId: "   ",
			action:  releaseCallbackAction,
			wantErr: "orderId is required",
		},
		{
			name:    "empty action",
			baseURL: "https://market.example.com/api/callback",
			orderId: "ORDER-1",
			action:  "",
			wantErr: "action is required",
		},
		{
			name:    "whitespace action",
			baseURL: "https://market.example.com/api/callback",
			orderId: "ORDER-1",
			action:  "  ",
			wantErr: "action is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildCallbackURL(tt.baseURL, tt.orderId, tt.action)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("buildCallbackURL = %q, want an error containing %q", got, tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("buildCallbackURL returned %v", err)
			}
			if got != tt.want {
				t.Errorf("buildCallbackURL = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestBuildCallbackURLPerAction(t *testing.T) {
	const base = "https://market.example.com/api/callback"

	seen := make(map[string]string)
	for _, action := range []string{refundCallbackAction, releaseCallbackAction, paymentReceivedCallbackAction} {
		url, err := buildCallbackURL(base, "ORDER-1", action)
		if err != nil {
			t.Fatalf("buildCallbackURL(%s) returned %v", action, err)
		}
		if previous, ok := seen[url]; ok {
			t.Errorf("actions %s and %s both build %q", previous, action, url)
		}
		seen[url] = action

		if !strings.HasSuffix(url, "/"+action) {
			t.Errorf("buildCallbackURL(%s) = %q, want it to end in the action", action, url)
		}
	}
}

func TestPostRequiresConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		client  *Client
		wantErr string
	}{
		{"nil client", nil, "callback URL is not configured"},
		{
			"no base URL",
			NewClient("", "key", stubTokens{}),
			"callback URL is not configured",
		},
		{
			"no API key",
			NewClient("https://market.example.com/cb", "", stubTokens{}),
			"callback API key is not configured",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.client.post([]byte(`{}`), "ORDER-1", releaseCallbackAction)

			if err == nil {
				t.Fatalf("post returned %v, want an error", res)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestPostRejectsAnEmptyOrderID(t *testing.T) {
	client := NewClient("https://market.example.com/cb", "key", stubTokens{})

	if _, err := client.post([]byte(`{}`), "", releaseCallbackAction); err == nil {
		t.Error("post with an empty orderId returned no error")
	}
}

func TestPostSendsTheExpectedRequest(t *testing.T) {
	var (
		gotMethod      string
		gotPath        string
		gotAPIKey      string
		gotContentType string
		gotBody        string
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotAPIKey = r.Header.Get("APIKey")
		gotContentType = r.Header.Get("Content-Type")
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		gotBody = string(body)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL+"/api/callback", "the-api-key", stubTokens{})

	res, err := client.post([]byte(`{"amount":"1.5"}`), "ORDER-1", releaseCallbackAction)
	if err != nil {
		t.Fatalf("post returned %v", err)
	}
	defer res.Body.Close()

	if gotMethod != http.MethodPost {
		t.Errorf("method = %s, want POST", gotMethod)
	}
	if want := "/api/callback/ORDER-1/escrowReleased"; gotPath != want {
		t.Errorf("path = %q, want %q", gotPath, want)
	}
	if gotAPIKey != "the-api-key" {
		t.Errorf("APIKey header = %q, want the configured key", gotAPIKey)
	}
	if gotContentType != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", gotContentType)
	}
	if gotBody != `{"amount":"1.5"}` {
		t.Errorf("body = %q, want the payload", gotBody)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", res.StatusCode, http.StatusOK)
	}
}

func TestPostReturnsErrorStatusesToTheCaller(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.URL, "key", stubTokens{})

	res, err := client.post([]byte(`{}`), "ORDER-1", releaseCallbackAction)
	if err != nil {
		t.Fatalf("post returned %v, want the response", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", res.StatusCode, http.StatusInternalServerError)
	}
}

func TestPostTransportErrorIsReturned(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	url := server.URL
	server.Close()

	client := NewClient(url, "key", stubTokens{})

	if _, err := client.post([]byte(`{}`), "ORDER-1", releaseCallbackAction); err == nil {
		t.Error("post to a closed server returned no error")
	}
}

func TestNewClient(t *testing.T) {
	tokens := stubTokens{"ETH": {symbol: "ETH", decimals: 18}}

	client := NewClient("https://market.example.com/cb", "key", tokens)

	if client.baseURL != "https://market.example.com/cb" {
		t.Errorf("baseURL = %q", client.baseURL)
	}
	if client.apiKey != "key" {
		t.Errorf("apiKey = %q", client.apiKey)
	}
	if client.tokens == nil {
		t.Error("tokens was not stored")
	}

	if client.http == nil || client.http.Timeout <= 0 {
		t.Error("the HTTP client has no timeout")
	}
}
