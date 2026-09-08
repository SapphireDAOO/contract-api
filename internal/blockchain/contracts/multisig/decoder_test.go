package multisig

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestHumanize(t *testing.T) {
	tests := []struct {
		identifier string
		want       string
	}{
		{"setFeeRate", "Set Fee Rate"},
		{"_newFeeRate", "New Fee Rate"},
		{"__doubleUnderscore", "Double Underscore"},
		{"transfer", "Transfer"},
		{"a", "A"},
		{"", ""},
		{"setFeeReceiver", "Set Fee Receiver"},

		{"setURL", "Set URL"},
		{"token0", "Token 0"},

		{"newURLValue", "New URLValue"},
		{"setToken2Address", "Set Token 2Address"},
		{"AlreadyCapitalized", "Already Capitalized"},
		{"_", ""},
	}

	for _, tt := range tests {
		t.Run(tt.identifier, func(t *testing.T) {
			if got := humanize(tt.identifier); got != tt.want {
				t.Errorf("humanize(%q) = %q, want %q", tt.identifier, got, tt.want)
			}
		})
	}
}

func TestHumanizeStripsLeadingUnderscores(t *testing.T) {
	for _, identifier := range []string{"_a", "__a", "___a"} {
		if got := humanize(identifier); got != "A" {
			t.Errorf("humanize(%q) = %q, want %q", identifier, got, "A")
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		max  int
		want string
	}{
		{"under the limit", "short", 10, "short"},
		{"at the limit", "exactly10!", 10, "exactly10!"},
		{"over the limit", "this is far too long", 10, "this is fa…"},
		{"empty", "", 10, ""},
		{"zero limit", "abc", 0, "…"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncate(tt.s, tt.max); got != tt.want {
				t.Errorf("truncate(%q, %d) = %q, want %q", tt.s, tt.max, got, tt.want)
			}
		})
	}
}

func TestTruncateBoundsTheLength(t *testing.T) {
	long := strings.Repeat("x", 500)

	got := truncate(long, 120)

	if len([]rune(got)) != 121 {
		t.Errorf("truncate produced %d runes, want 121", len([]rune(got)))
	}
	if !strings.HasSuffix(got, "…") {
		t.Errorf("truncate produced %q, want it to end in an ellipsis", got)
	}
}

func TestFormatArg(t *testing.T) {
	c := &Multisig{}

	address := common.HexToAddress("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
	var hash [32]byte
	hash[31] = 1

	tests := []struct {
		name  string
		value any
		want  string
	}{
		{"big int", big.NewInt(300), "300"},
		{"nil-safe large big int", new(big.Int).Lsh(big.NewInt(1), 200), new(big.Int).Lsh(big.NewInt(1), 200).String()},
		{"true", true, "yes"},
		{"false", false, "no"},
		{"string", "a value", "a value"},
		{"empty string", "", ""},
		{
			"bytes32 is rendered as a hash",
			hash,
			"`" + common.Hash(hash).Hex() + "`",
		},
		{"bytes", []byte{0xde, 0xad, 0xbe, 0xef}, "`0xdeadbeef`"},
		{"empty bytes", []byte{}, "`0x`"},
		{"unknown type falls back to a formatted value", uint32(7), "7"},
		{"unknown struct type", struct{ A int }{A: 1}, "{A:1}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.formatArg(tt.value); got != tt.want {
				t.Errorf("formatArg(%v) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}

	t.Run("address is rendered as a link", func(t *testing.T) {
		got := c.formatArg(address)

		if !strings.Contains(got, shortHex(address.Hex())) {
			t.Errorf("formatArg(address) = %q, want it to contain the shortened address", got)
		}
	})
}

func TestFormatArgTruncatesLongBytes(t *testing.T) {
	c := &Multisig{}

	got := c.formatArg(make([]byte, 500))

	if !strings.HasSuffix(got, "…`") {
		t.Errorf("formatArg(long bytes) = %q, want it truncated", got)
	}
	if len([]rune(got)) > 130 {
		t.Errorf("formatArg(long bytes) is %d runes, want it capped near 120", len([]rune(got)))
	}
}

func TestBuildKnownContracts(t *testing.T) {
	multisigAddress := common.HexToAddress("0x4444444444444444444444444444444444444444")
	peers := Peers{
		PaymentProcessor:        common.HexToAddress("0x1111111111111111111111111111111111111111"),
		SimplePaymentProcessor:  common.HexToAddress("0x3333333333333333333333333333333333333333"),
		PaymentProcessorStorage: common.HexToAddress("0x2222222222222222222222222222222222222222"),
	}

	known := buildKnownContracts(multisigAddress, peers)

	wantNames := map[string]common.Address{
		"Payment Processor":         peers.PaymentProcessor,
		"Simple Payment Processor":  peers.SimplePaymentProcessor,
		"Payment Processor Storage": peers.PaymentProcessorStorage,
		"Multisig":                  multisigAddress,
	}
	if len(known) != len(wantNames) {
		t.Fatalf("buildKnownContracts returned %d contracts, want %d", len(known), len(wantNames))
	}

	for _, kc := range known {
		wantAddress, ok := wantNames[kc.name]
		if !ok {
			t.Errorf("unexpected contract %q", kc.name)
			continue
		}
		if kc.address != wantAddress {
			t.Errorf("%s address = %s, want %s", kc.name, kc.address, wantAddress)
		}
		if kc.abi == nil {
			t.Errorf("%s has no parsed ABI", kc.name)
			continue
		}
		if len(kc.abi.Methods) == 0 {
			t.Errorf("%s has an ABI with no methods", kc.name)
		}
	}
}

func TestDecodeAction(t *testing.T) {
	multisigAddress := common.HexToAddress("0x4444444444444444444444444444444444444444")
	processor := common.HexToAddress("0x1111111111111111111111111111111111111111")
	peers := Peers{
		PaymentProcessor:        processor,
		SimplePaymentProcessor:  common.HexToAddress("0x3333333333333333333333333333333333333333"),
		PaymentProcessorStorage: common.HexToAddress("0x2222222222222222222222222222222222222222"),
	}
	c := &Multisig{known: buildKnownContracts(multisigAddress, peers)}

	var (
		methodName string
		calldata   []byte
	)
	for _, kc := range c.known {
		if kc.address != processor {
			continue
		}
		for name, method := range kc.abi.Methods {
			if len(method.Inputs) != 1 || method.Inputs[0].Type.String() != "uint256" {
				continue
			}
			packed, err := kc.abi.Pack(name, big.NewInt(300))
			if err != nil {
				continue
			}
			methodName, calldata = name, packed
			break
		}
		break
	}
	if calldata == nil {
		t.Skip("the payment processor ABI has no single-uint256 method to exercise")
	}

	got := c.decodeAction(processor, calldata)

	if got == nil {
		t.Fatalf("decodeAction returned nil for a call to %s", methodName)
	}
	if got.Name != humanize(methodName) {
		t.Errorf("Name = %q, want %q", got.Name, humanize(methodName))
	}
	if got.Contract != "Payment Processor" {
		t.Errorf("Contract = %q, want Payment Processor", got.Contract)
	}
	if len(got.Args) != 1 {
		t.Fatalf("got %d args, want 1", len(got.Args))
	}
	if got.Args[0].Value != "300" {
		t.Errorf("arg value = %q, want 300", got.Args[0].Value)
	}
}

func TestDecodeActionRejectsShortCalldata(t *testing.T) {
	c := &Multisig{known: buildKnownContracts(common.Address{}, Peers{})}

	for _, data := range [][]byte{nil, {}, {0x01}, {0x01, 0x02, 0x03}} {
		if got := c.decodeAction(common.Address{}, data); got != nil {
			t.Errorf("decodeAction(%v) = %+v, want nil", data, got)
		}
	}
}

func TestDecodeActionUnknownSelector(t *testing.T) {
	c := &Multisig{known: buildKnownContracts(common.Address{}, Peers{})}

	got := c.decodeAction(common.Address{}, []byte{0xff, 0xff, 0xff, 0xff, 0x00})

	if got != nil {
		t.Errorf("decodeAction with an unknown selector = %+v, want nil", got)
	}
}

func TestBuildKnownContractsExcludesERC20(t *testing.T) {
	known := buildKnownContracts(common.Address{}, Peers{})

	for _, kc := range known {
		if strings.Contains(strings.ToLower(kc.name), "erc20") {
			t.Errorf("ERC20 is listed as a known contract (%q)", kc.name)
		}
	}
}
