package discord

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestShortHex(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"a full address", "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed", "0x5aAe…eAed"},
		{"exactly twelve characters", "0x1234567890", "0x1234567890"},
		{"thirteen characters", "0x12345678901", "0x1234…8901"},
		{"short input", "0xabc", "0xabc"},
		{"empty", "", ""},
		{
			"a transaction hash",
			"0x0000000000000000000000000000000000000000000000000000000000000001",
			"0x0000…0001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortHex(tt.s); got != tt.want {
				t.Errorf("ShortHex(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestShortHexKeepsBothEnds(t *testing.T) {
	const address = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	got := ShortHex(address)

	if !strings.HasPrefix(got, address[:6]) {
		t.Errorf("ShortHex = %q, want it to start with %q", got, address[:6])
	}
	if !strings.HasSuffix(got, address[len(address)-4:]) {
		t.Errorf("ShortHex = %q, want it to end with %q", got, address[len(address)-4:])
	}
}

func TestAddressLink(t *testing.T) {
	address := common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")

	t.Run("with an explorer", func(t *testing.T) {
		got := AddressLink("https://etherscan.io", address)

		want := "[`" + ShortHex(address.Hex()) + "`](https://etherscan.io/address/" + address.Hex() + ")"
		if got != want {
			t.Errorf("AddressLink = %q, want %q", got, want)
		}
	})

	t.Run("without an explorer", func(t *testing.T) {
		got := AddressLink("", address)

		if !strings.Contains(got, ShortHex(address.Hex())) {
			t.Errorf("AddressLink = %q, want the shortened address", got)
		}
		if strings.Contains(got, "http") {
			t.Errorf("AddressLink = %q, want no explorer URL", got)
		}
	})
}

func TestAddressLinkIsMarkdown(t *testing.T) {
	address := common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")

	got := AddressLink("https://etherscan.io", address)

	if !strings.HasPrefix(got, "[`") || !strings.Contains(got, "`](") || !strings.HasSuffix(got, ")") {
		t.Errorf("AddressLink = %q, want a markdown link around a code span", got)
	}
}
