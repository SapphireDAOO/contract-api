package note

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestParseContent(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    []byte
		wantErr error
	}{
		{name: "hex payload", value: "0xdeadbeef", want: []byte{0xde, 0xad, 0xbe, 0xef}},
		{name: "single byte", value: "0x01", want: []byte{0x01}},
		{name: "surrounding whitespace is trimmed", value: "  0xdeadbeef  ", want: []byte{0xde, 0xad, 0xbe, 0xef}},
		{name: "uppercase hex", value: "0xDEADBEEF", want: []byte{0xde, 0xad, 0xbe, 0xef}},
		{name: "empty", value: "", wantErr: ErrContentRequired},
		{name: "whitespace only", value: "   ", wantErr: ErrContentRequired},
		{name: "0x with no payload", value: "0x", wantErr: ErrContentRequired},
		{name: "missing 0x prefix", value: "deadbeef", wantErr: ErrContentHex},
		{name: "not hex", value: "0xnothex", wantErr: ErrContentHex},
		{name: "odd number of digits", value: "0xabc", wantErr: ErrContentHex},
		{name: "plaintext is rejected", value: "Left at the door", wantErr: ErrContentHex},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseContent(tt.value)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ParseContent(%q) error = %v, want %v", tt.value, err, tt.wantErr)
				}
				if got != nil {
					t.Errorf("ParseContent(%q) = %x alongside an error, want nil", tt.value, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseContent(%q) returned %v", tt.value, err)
			}
			if string(got) != string(tt.want) {
				t.Errorf("ParseContent(%q) = %x, want %x", tt.value, got, tt.want)
			}
		})
	}
}

func TestParseContentSizeLimit(t *testing.T) {
	atLimit := hexutil.Encode(make([]byte, MaxContentBytes))
	overLimit := hexutil.Encode(make([]byte, MaxContentBytes+1))

	if _, err := ParseContent(atLimit); err != nil {
		t.Errorf("ParseContent at the limit returned %v, want nil", err)
	}
	if _, err := ParseContent(overLimit); !errors.Is(err, ErrContentTooLong) {
		t.Errorf("ParseContent over the limit = %v, want ErrContentTooLong", err)
	}
}

func TestParseContentReturnsTheDecodedBytes(t *testing.T) {
	payload := []byte{0x00, 0xff, 0x10, 0x7f}

	got, err := ParseContent(hexutil.Encode(payload))
	if err != nil {
		t.Fatalf("ParseContent returned %v", err)
	}
	if string(got) != string(payload) {
		t.Errorf("ParseContent = %x, want %x", got, payload)
	}
}

func TestContentErrorsAreDistinct(t *testing.T) {
	all := []error{ErrContentRequired, ErrContentHex, ErrContentTooLong}

	for i, a := range all {
		for j, b := range all {
			if i != j && errors.Is(a, b) {
				t.Errorf("%v and %v are indistinguishable", a, b)
			}
		}
	}
}
