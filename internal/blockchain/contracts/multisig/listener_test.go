package multisig

import (
	"math/big"
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
			if got := shortHex(tt.s); got != tt.want {
				t.Errorf("shortHex(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestShortHexKeepsBothEnds(t *testing.T) {
	const address = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"

	got := shortHex(address)

	if !strings.HasPrefix(got, address[:6]) {
		t.Errorf("shortHex = %q, want it to start with %q", got, address[:6])
	}
	if !strings.HasSuffix(got, address[len(address)-4:]) {
		t.Errorf("shortHex = %q, want it to end with %q", got, address[len(address)-4:])
	}
}

func TestFormatEth(t *testing.T) {
	tests := []struct {
		name string
		wei  string
		want string
	}{
		{"nil", "", "0 ETH"},
		{"zero", "0", "0 ETH"},
		{"one ether", "1000000000000000000", "1 ETH"},
		{"a fraction", "1500000000000000000", "1.5 ETH"},
		{"one wei", "1", "0.000000000000000001 ETH"},
		{"sub-ether", "10000000000000000", "0.01 ETH"},
		{"many ether", "1234000000000000000000", "1234 ETH"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wei *big.Int
			if tt.wei != "" {
				wei, _ = new(big.Int).SetString(tt.wei, 10)
			}

			if got := formatEth(wei); got != tt.want {
				t.Errorf("formatEth(%v) = %q, want %q", wei, got, tt.want)
			}
		})
	}
}

func TestFormatEthAvoidsScientificNotation(t *testing.T) {
	tiny, _ := new(big.Int).SetString("1", 10)
	huge, _ := new(big.Int).SetString("1000000000000000000000000000", 10)

	for _, wei := range []*big.Int{tiny, huge} {
		got := formatEth(wei)

		if strings.ContainsAny(strings.TrimSuffix(got, " ETH"), "eE") {
			t.Errorf("formatEth(%s) = %q, want plain decimal notation", wei, got)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "A"},
		{"proposed", "Proposed"},
		{"already Capitalized", "Already Capitalized"},
		{"Proposed", "Proposed"},
		{"3 confirmations", "3 confirmations"},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := capitalize(tt.s); got != tt.want {
				t.Errorf("capitalize(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestPlural(t *testing.T) {
	tests := []struct {
		name string
		n    string
		unit string
		want string
	}{
		{"nil", "", "confirmation", "0 confirmations"},
		{"zero", "0", "confirmation", "0 confirmations"},
		{"one", "1", "confirmation", "1 confirmation"},
		{"two", "2", "confirmation", "2 confirmations"},
		{"many", "17", "owner", "17 owners"},
		{"large", "1000000", "block", "1000000 blocks"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n *big.Int
			if tt.n != "" {
				n, _ = new(big.Int).SetString(tt.n, 10)
			}

			if got := plural(n, tt.unit); got != tt.want {
				t.Errorf("plural(%v, %q) = %q, want %q", n, tt.unit, got, tt.want)
			}
		})
	}
}

func TestActionTitle(t *testing.T) {
	if got := actionTitle(nil); got != "Transaction" {
		t.Errorf("actionTitle(nil) = %q, want Transaction", got)
	}
	if got := actionTitle(&action{Name: "Set Fee Rate"}); got != "Set Fee Rate" {
		t.Errorf("actionTitle = %q, want Set Fee Rate", got)
	}
}

func TestActionName(t *testing.T) {
	if got := actionName(nil); got != "the transaction" {
		t.Errorf("actionName(nil) = %q, want %q", got, "the transaction")
	}
	if got := actionName(&action{Name: "Set Fee Rate"}); got != "**Set Fee Rate**" {
		t.Errorf("actionName = %q, want it bolded", got)
	}
}

func TestActionArgLines(t *testing.T) {
	tests := []struct {
		name string
		act  *action
		want []string
	}{
		{"nil action", nil, nil},
		{"no args", &action{Name: "Pause"}, nil},
		{
			"one arg",
			&action{Args: []argValue{{Name: "New Fee Rate", Value: "300"}}},
			[]string{"• New Fee Rate: **300**"},
		},
		{
			"several args are one block",
			&action{Args: []argValue{
				{Name: "New Fee Rate", Value: "300"},
				{Name: "Receiver", Value: "0xabc"},
			}},
			[]string{"• New Fee Rate: **300**\n• Receiver: **0xabc**"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := actionArgLines(tt.act)

			if len(got) != len(tt.want) {
				t.Fatalf("actionArgLines returned %d lines (%q), want %d", len(got), got, len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("line %d = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestProposalLine(t *testing.T) {
	var txHash [32]byte
	txHash[31] = 42

	got := proposalLine(txHash)

	if !strings.HasPrefix(got, "Transaction id: `") {
		t.Errorf("proposalLine = %q, want it to label the id", got)
	}
	if !strings.Contains(got, shortHex(common.Hash(txHash).Hex())) {
		t.Errorf("proposalLine = %q, want it to contain the shortened hash", got)
	}
}

func TestLink(t *testing.T) {
	t.Run("with an explorer", func(t *testing.T) {
		c := &Multisig{explorerURL: "https://etherscan.io"}

		if got, want := c.link("/tx/", "0xabc"), "https://etherscan.io/tx/0xabc"; got != want {
			t.Errorf("link = %q, want %q", got, want)
		}
		if got, want := c.link("/address/", "0xdef"), "https://etherscan.io/address/0xdef"; got != want {
			t.Errorf("link = %q, want %q", got, want)
		}
	})

	t.Run("without an explorer", func(t *testing.T) {
		c := &Multisig{}

		if got := c.link("/tx/", "0xabc"); got != "0xabc" {
			t.Errorf("link = %q, want the bare value", got)
		}
	})
}

func TestAddressLink(t *testing.T) {
	address := common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")

	t.Run("with an explorer", func(t *testing.T) {
		c := &Multisig{explorerURL: "https://etherscan.io"}

		got := c.addressLink(address)

		want := "[`" + shortHex(address.Hex()) + "`](https://etherscan.io/address/" + address.Hex() + ")"
		if got != want {
			t.Errorf("addressLink = %q, want %q", got, want)
		}
	})

	t.Run("without an explorer", func(t *testing.T) {
		c := &Multisig{}

		got := c.addressLink(address)

		if !strings.Contains(got, shortHex(address.Hex())) {
			t.Errorf("addressLink = %q, want the shortened address", got)
		}
	})
}

func TestTargetName(t *testing.T) {
	processor := common.HexToAddress("0x1111111111111111111111111111111111111111")
	unknown := common.HexToAddress("0x9999999999999999999999999999999999999999")

	c := &Multisig{
		explorerURL: "https://etherscan.io",
		known: buildKnownContracts(
			common.HexToAddress("0x4444444444444444444444444444444444444444"),
			Peers{PaymentProcessor: processor},
		),
	}

	t.Run("a known address is named", func(t *testing.T) {
		got := c.targetName(nil, processor)

		if !strings.Contains(got, "Payment Processor") {
			t.Errorf("targetName = %q, want it to name the contract", got)
		}
		if !strings.Contains(got, shortHex(processor.Hex())) {
			t.Errorf("targetName = %q, want it to link the address", got)
		}
	})

	t.Run("an unknown address falls back to the decoded contract", func(t *testing.T) {
		got := c.targetName(&action{Contract: "Simple Payment Processor"}, unknown)

		if !strings.Contains(got, "Simple Payment Processor") {
			t.Errorf("targetName = %q, want the decoded contract name", got)
		}
	})

	t.Run("an unknown address with no action is just a link", func(t *testing.T) {
		got := c.targetName(nil, unknown)

		if got != c.addressLink(unknown) {
			t.Errorf("targetName = %q, want a bare address link", got)
		}
	})
}
