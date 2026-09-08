package invoice

import (
	"math/big"
	"strings"
	"testing"
)

func TestEncodeID(t *testing.T) {
	tests := []struct {
		name string
		id   *big.Int
		want string
	}{
		{"nil", nil, ""},
		{"negative", big.NewInt(-1), ""},
		{"zero encodes a single zero byte", big.NewInt(0), "AA"},
		{"one", big.NewInt(1), "AQ"},
		{"255", big.NewInt(255), "_w"},
		{"256", big.NewInt(256), "AQA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeID(tt.id); got != tt.want {
				t.Errorf("EncodeID(%v) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}

func TestEncodeIDIsURLSafeAndUnpadded(t *testing.T) {
	id, _ := new(big.Int).SetString("18446744073709551615", 10)

	encoded := EncodeID(id)

	if strings.ContainsAny(encoded, "+/=") {
		t.Errorf("EncodeID produced %q, which is not unpadded base64url", encoded)
	}
}

func TestEncodeIDString(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{"empty", "", ""},
		{"whitespace only", "   ", ""},
		{"not a number", "abc", ""},
		{"hex is not accepted", "0x10", ""},
		{"decimal", "1", "AQ"},
		{"surrounding whitespace is trimmed", "  256  ", "AQA"},
		{"zero", "0", "AA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeIDString(tt.id); got != tt.want {
				t.Errorf("EncodeIDString(%q) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}

func TestEncodeIDStringMatchesEncodeID(t *testing.T) {
	id, _ := new(big.Int).SetString("123456789012345678901234567890", 10)

	if got, want := EncodeIDString(id.String()), EncodeID(id); got != want {
		t.Errorf("EncodeIDString = %q, EncodeID = %q; they must agree", got, want)
	}
}

func TestEncodeMetaID(t *testing.T) {
	tests := []struct {
		name string
		id   *big.Int
		want string
	}{
		{"nil stays empty rather than a bare prefix", nil, ""},
		{"negative stays empty", big.NewInt(-5), ""},
		{"one", big.NewInt(1), MetaPrefix + "AQ"},
		{"zero", big.NewInt(0), MetaPrefix + "AA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeMetaID(tt.id); got != tt.want {
				t.Errorf("EncodeMetaID(%v) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}

func TestEncodeMetaIDString(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{"empty stays empty rather than a bare prefix", "", ""},
		{"not a number stays empty", "nope", ""},
		{"decimal", "256", MetaPrefix + "AQA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeMetaIDString(tt.id); got != tt.want {
				t.Errorf("EncodeMetaIDString(%q) = %q, want %q", tt.id, got, tt.want)
			}
		})
	}
}

func TestIsMetaID(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		want    bool
	}{
		{"meta id", MetaPrefix + "AQ", true},
		{"leading whitespace is trimmed first", "  " + MetaPrefix + "AQ", true},
		{"single invoice id", "AQ", false},
		{"empty", "", false},
		{"a plain id that happens to start with m", "mtAQ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMetaID(tt.encoded); got != tt.want {
				t.Errorf("IsMetaID(%q) = %v, want %v", tt.encoded, got, tt.want)
			}
		})
	}
}

func TestDecodeID(t *testing.T) {
	tests := []struct {
		name    string
		encoded string
		want    string
		wantErr bool
	}{
		{"single invoice id", "AQ", "1", false},
		{"meta id decodes without stripping the prefix first", MetaPrefix + "AQ", "1", false},
		{"padded input is accepted", "AQ==", "1", false},
		{"surrounding whitespace is trimmed", "  AQA  ", "256", false},
		{"zero", "AA", "0", false},
		{"empty", "", "", true},
		{"whitespace only", "   ", "", true},
		{"bare prefix", MetaPrefix, "", true},
		{"not base64url", "!!!", "", true},
		{"standard base64 characters are rejected", "+/A", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeID(tt.encoded)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("DecodeID(%q) = %v, want an error", tt.encoded, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("DecodeID(%q) returned %v", tt.encoded, err)
			}
			if got.String() != tt.want {
				t.Errorf("DecodeID(%q) = %s, want %s", tt.encoded, got, tt.want)
			}
		})
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	ids := []string{
		"0",
		"1",
		"255",
		"256",
		"1461501637330902918203684832716283019655932542975",
		"105312291668557186697918027683670432318895095400549111254310977535",
	}

	for _, raw := range ids {
		t.Run(raw, func(t *testing.T) {
			id, _ := new(big.Int).SetString(raw, 10)

			decoded, err := DecodeID(EncodeID(id))
			if err != nil {
				t.Fatalf("DecodeID(EncodeID(%s)) returned %v", raw, err)
			}
			if decoded.Cmp(id) != 0 {
				t.Errorf("round trip of %s produced %s", raw, decoded)
			}

			metaDecoded, err := DecodeID(EncodeMetaID(id))
			if err != nil {
				t.Fatalf("DecodeID(EncodeMetaID(%s)) returned %v", raw, err)
			}
			if metaDecoded.Cmp(id) != 0 {
				t.Errorf("meta round trip of %s produced %s", raw, metaDecoded)
			}
		})
	}
}

func TestMetaAndSingleIDsAreDistinguishable(t *testing.T) {
	id := big.NewInt(4242)

	single := EncodeID(id)
	meta := EncodeMetaID(id)

	if single == meta {
		t.Fatal("meta and single encodings are identical")
	}
	if IsMetaID(single) {
		t.Errorf("IsMetaID(%q) = true for a single invoice id", single)
	}
	if !IsMetaID(meta) {
		t.Errorf("IsMetaID(%q) = false for a meta invoice id", meta)
	}
}
