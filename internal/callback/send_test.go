package callback

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
	"testing"
)

func countingServer(t *testing.T) (*httptest.Server, *atomic.Int32) {
	t.Helper()
	var hits atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)
	return server, &hits
}

func TestCallbacksAreNotSentOutsideProduction(t *testing.T) {
	server, hits := countingServer(t)
	client := NewClient(server.URL, "key", stubTokens{})
	client.production = false

	client.sendCallbackWithRetry([]byte(`{}`), "42", releaseCallbackAction)

	if got := hits.Load(); got != 0 {
		t.Errorf("intermediated platform received %d callbacks outside production, want 0", got)
	}
}

func TestCallbacksAreSentInProduction(t *testing.T) {
	server, hits := countingServer(t)
	client := NewClient(server.URL, "key", stubTokens{})
	client.production = true

	client.sendCallbackWithRetry([]byte(`{}`), "42", releaseCallbackAction)

	if got := hits.Load(); got != 1 {
		t.Errorf("intermediated platform received %d callbacks in production, want 1", got)
	}
}

func TestNewClientReadsProductionFromTheEnvironment(t *testing.T) {
	tests := []struct {
		value string
		want  bool
	}{
		// docker-compose passes PRODUCTION=${PRODUCTION}, so an unset host
		// variable arrives empty rather than missing.
		{"", false},
		{"0", false},
		{"false", false},
		{"no", false},
		{"1", true},
		{"true", true},
		{" TRUE ", true},
	}
	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			t.Setenv("PRODUCTION", tt.value)
			if got := NewClient("https://platform.example.com/cb", "key", stubTokens{}).production; got != tt.want {
				t.Errorf("PRODUCTION=%q: production = %v, want %v", tt.value, got, tt.want)
			}
		})
	}

	t.Run("unset", func(t *testing.T) {
		t.Setenv("PRODUCTION", "1") // restored after the test
		os.Unsetenv("PRODUCTION")
		if NewClient("https://platform.example.com/cb", "key", stubTokens{}).production {
			t.Error("PRODUCTION unset, but the client is in production")
		}
	})
}

func TestPostBuildsButDoesNotSendOutsideProduction(t *testing.T) {
	server, hits := countingServer(t)
	client := NewClient(server.URL, "key", stubTokens{})
	client.production = false

	res, err := client.post([]byte(`{}`), "42", releaseCallbackAction)
	if err != nil {
		t.Fatalf("post returned an error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", res.StatusCode)
	}
	if got := hits.Load(); got != 0 {
		t.Errorf("intermediated platform received %d requests outside production, want 0", got)
	}

	// Configuration is still checked, so a broken setup shows up before
	// production does.
	if _, err := NewClient("", "key", stubTokens{}).post([]byte(`{}`), "42", releaseCallbackAction); err == nil {
		t.Error("post with no callback URL returned no error outside production")
	}
}
