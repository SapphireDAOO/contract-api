package revert

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

type revertError struct {
	code int
	data any
}

func (e revertError) Error() string          { return "execution reverted" }
func (e revertError) ErrorCode() int         { return e.code }
func (e revertError) ErrorData() interface{} { return e.data }

func revertWith(selector string) error {
	return revertError{code: 3, data: selector}
}

func TestSelector(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil error", nil, ""},
		{"plain error", errors.New("connection refused"), ""},
		{"revert with a known selector", revertWith("0xb12e2421"), "0xb12e2421"},
		{"revert with trailing argument data", revertWith("0xb12e2421deadbeef"), "0xb12e2421"},
		{"uppercase hex is normalised", revertWith("0xB12E2421"), "0xb12e2421"},
		{"wrong error code is not a revert", revertError{code: -32000, data: "0xb12e2421"}, ""},
		{"payload shorter than a selector", revertWith("0xb12e"), ""},
		{"empty payload", revertWith("0x"), ""},
		{"payload is not hex", revertError{code: 3, data: "not hex"}, ""},
		{"payload is not a string", revertError{code: 3, data: 42}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Selector(tt.err); got != tt.want {
				t.Errorf("Selector() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSelectorUnwrapsWrappedErrors(t *testing.T) {
	wrapped := fmt.Errorf("calling getUsdPerToken: %w", revertWith(UnsupportedToken))

	if got := Selector(wrapped); got != UnsupportedToken {
		t.Errorf("Selector(wrapped) = %q, want %q", got, UnsupportedToken)
	}
}

func TestIs(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		selector string
		want     bool
	}{
		{"matching selector", revertWith(UnsupportedToken), UnsupportedToken, true},
		{"different selector", revertWith("0xb12e2421"), UnsupportedToken, false},
		{"nil error", nil, UnsupportedToken, false},
		{"plain error", errors.New("boom"), UnsupportedToken, false},
		{"nil error against the empty selector", nil, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.err, tt.selector); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReason(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil error", nil, ""},
		{
			"known selector maps to its description",
			revertWith("0xb12e2421"),
			"The buyer and seller addresses cannot be the same.",
		},
		{
			"unsupported token selector",
			revertWith(UnsupportedToken),
			"The oracle has no price feed for this token.",
		},
		{
			"unknown selector falls back to the error text",
			revertWith("0xdeadbeef"),
			"execution reverted",
		},
		{
			"non-revert error falls back to the error text",
			errors.New("connection refused"),
			"connection refused",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reason(tt.err); got != tt.want {
				t.Errorf("Reason() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestUnsupportedTokenIsDescribed(t *testing.T) {
	if _, ok := Descriptions[UnsupportedToken]; !ok {
		t.Errorf("Descriptions is missing the UnsupportedToken selector %q", UnsupportedToken)
	}
}

func TestStatusCodesKeysAreDescriptions(t *testing.T) {
	descriptions := make(map[string]bool, len(Descriptions))
	for _, description := range Descriptions {
		descriptions[description] = true
	}

	for reason := range StatusCodes {
		if !descriptions[reason] {
			t.Errorf("StatusCodes has %q, which no selector in Descriptions produces", reason)
		}
	}
}

func TestStatusCodesAreClientOrConflictErrors(t *testing.T) {
	for reason, code := range StatusCodes {
		if code < 400 || code >= 500 {
			t.Errorf("StatusCodes[%q] = %d, want a 4xx status", reason, code)
		}
	}
}

func TestDescriptionKeysAreFourByteSelectors(t *testing.T) {
	for selector, description := range Descriptions {
		if len(selector) != 10 || selector[:2] != "0x" {
			t.Errorf("selector %q is not a 0x-prefixed four-byte selector", selector)
		}
		if description == "" {
			t.Errorf("selector %q has an empty description", selector)
		}
	}
}

func TestStatusCodeLookupFallsBackToServerError(t *testing.T) {
	if code := StatusCodes[Reason(revertWith("0xdeadbeef"))]; code != 0 {
		t.Errorf("unmapped reason has status %d, want 0 so callers can default", code)
	}
	if code := StatusCodes["The specified invoice does not exist."]; code != http.StatusNotFound {
		t.Errorf("missing invoice status = %d, want %d", code, http.StatusNotFound)
	}
}
