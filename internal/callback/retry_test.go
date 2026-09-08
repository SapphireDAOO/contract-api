package callback

import (
	"net/http"
	"testing"
	"time"
)

func TestCallbackRetryDelay(t *testing.T) {
	tests := []struct {
		attempt int
		want    time.Duration
	}{
		{-1, callbackRetryBaseDelay},
		{0, callbackRetryBaseDelay},
		{1, 1 * time.Second},
		{2, 2 * time.Second},
		{3, 4 * time.Second},
		{4, 8 * time.Second},
		{5, 16 * time.Second},
		{6, 30 * time.Second},
		{10, 30 * time.Second},
		{20, 30 * time.Second},
	}

	for _, tt := range tests {
		t.Run(time.Duration(tt.attempt).String(), func(t *testing.T) {
			if got := callbackRetryDelay(tt.attempt); got != tt.want {
				t.Errorf("callbackRetryDelay(%d) = %v, want %v", tt.attempt, got, tt.want)
			}
		})
	}
}

func TestCallbackRetryDelayIsMonotonicAndPositive(t *testing.T) {
	previous := time.Duration(0)

	for attempt := range 30 {
		delay := callbackRetryDelay(attempt)

		if delay <= 0 {
			t.Fatalf("callbackRetryDelay(%d) = %v, want a positive delay", attempt, delay)
		}
		if delay > 30*time.Second {
			t.Fatalf("callbackRetryDelay(%d) = %v, want at most 30s", attempt, delay)
		}
		if delay < previous {
			t.Fatalf("callbackRetryDelay(%d) = %v, less than the previous %v", attempt, delay, previous)
		}
		previous = delay
	}
}

func TestCallbackRetryTotalDelayIsBounded(t *testing.T) {
	var total time.Duration
	for attempt := 1; attempt < callbackRetryAttempts; attempt++ {
		total += callbackRetryDelay(attempt)
	}

	if total > time.Minute {
		t.Errorf("the %d configured attempts wait %v in total, want under a minute", callbackRetryAttempts, total)
	}
}

func TestCallbackStatusHint(t *testing.T) {
	tests := []struct {
		status int
		want   string
	}{
		{http.StatusBadRequest, "invalid request data"},
		{http.StatusForbidden, "API key rejected"},
		{http.StatusGone, "order not found"},
		{http.StatusInternalServerError, "marketplace internal error"},
		{http.StatusOK, ""},
		{http.StatusNotFound, ""},
		{http.StatusUnauthorized, ""},
		{http.StatusBadGateway, ""},
		{0, ""},
	}

	for _, tt := range tests {
		t.Run(http.StatusText(tt.status), func(t *testing.T) {
			if got := callbackStatusHint(tt.status); got != tt.want {
				t.Errorf("callbackStatusHint(%d) = %q, want %q", tt.status, got, tt.want)
			}
		})
	}
}

func TestParseCallbackResponse(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		wantError   string
		wantMessage string
	}{
		{"empty body", "", "", ""},
		{"nil body", "", "", ""},
		{
			"error and string message",
			`{"status":400,"error":"BadRequest","message":"order is closed"}`,
			"BadRequest",
			"order is closed",
		},
		{
			"error only",
			`{"error":"Forbidden"}`,
			"Forbidden",
			"",
		},
		{
			"message only",
			`{"message":"accepted"}`,
			"",
			"accepted",
		},
		{
			"non-string message is kept verbatim",
			`{"error":"Invalid","message":["a is required","b is required"]}`,
			"Invalid",
			`["a is required","b is required"]`,
		},
		{
			"object message is kept verbatim",
			`{"message":{"field":"amount"}}`,
			"",
			`{"field":"amount"}`,
		},
		{
			"numeric message is kept verbatim",
			`{"message":42}`,
			"",
			"42",
		},
		{
			"non-JSON body becomes the message",
			"  <html>502 Bad Gateway</html>  ",
			"",
			"<html>502 Bad Gateway</html>",
		},
		{
			"plain text body becomes the message",
			"upstream timeout",
			"",
			"upstream timeout",
		},
		{
			"JSON with neither field",
			`{"status":200}`,
			"",
			"",
		},
		{
			"null message",
			`{"error":"Oops","message":null}`,
			"Oops",
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotError, gotMessage := parseCallbackResponse([]byte(tt.body))

			if gotError != tt.wantError {
				t.Errorf("error = %q, want %q", gotError, tt.wantError)
			}
			if gotMessage != tt.wantMessage {
				t.Errorf("message = %q, want %q", gotMessage, tt.wantMessage)
			}
		})
	}
}

func TestParseCallbackResponseNilSlice(t *testing.T) {
	gotError, gotMessage := parseCallbackResponse(nil)

	if gotError != "" || gotMessage != "" {
		t.Errorf("parseCallbackResponse(nil) = (%q, %q), want two empty strings", gotError, gotMessage)
	}
}
