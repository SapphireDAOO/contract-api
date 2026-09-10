package revert

import (
	"errors"
	"fmt"
	"testing"
)

type revertError struct {
	code int
	data any
}

func (e revertError) Error() string  { return "execution reverted" }
func (e revertError) ErrorCode() int { return e.code }
func (e revertError) ErrorData() any { return e.data }

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
