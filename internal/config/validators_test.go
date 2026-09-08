package config

import (
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func requireErr(t *testing.T, err error, wantErr string) {
	t.Helper()

	if wantErr == "" {
		if err != nil {
			t.Fatalf("validate() returned %v, want nil", err)
		}
		return
	}
	if err == nil {
		t.Fatalf("validate() returned nil, want an error containing %q", wantErr)
	}
	if !strings.Contains(err.Error(), wantErr) {
		t.Errorf("error = %q, want it to contain %q", err, wantErr)
	}
}

func TestValidateSignerKey(t *testing.T) {
	const key = "4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"

	tests := []struct {
		name    string
		key     string
		wantErr string
	}{
		{name: "bare hex", key: key},
		{name: "0x prefixed", key: "0x" + key},
		{name: "surrounding whitespace is trimmed", key: "  " + key + "  "},
		{name: "uppercase hex", key: strings.ToUpper(key)},
		{name: "empty", key: "", wantErr: "signerKey is required"},
		{name: "whitespace only", key: "   ", wantErr: "signerKey is required"},
		{name: "just the prefix", key: "0x", wantErr: "signerKey is required"},
		{name: "too short", key: key[:63], wantErr: "must be 32 bytes of hex"},
		{name: "too long", key: key + "0", wantErr: "must be 32 bytes of hex"},
		{
			name:    "right length but not hex",
			key:     strings.Repeat("z", 64),
			wantErr: "not valid hex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requireErr(t, validateSignerKey(tt.key), tt.wantErr)
		})
	}
}

func TestValidateSignerKeyReportsTheLength(t *testing.T) {
	err := validateSignerKey("0xdeadbeef")

	if err == nil {
		t.Fatal("validateSignerKey returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "8 characters") {
		t.Errorf("error = %q, want it to state the length seen", err)
	}
}

func TestRPCValidate(t *testing.T) {
	valid := RPC{
		HTTP:        "https://mainnet.example.com/v3/key",
		WS:          "wss://mainnet.example.com/v3/key",
		DialTimeout: 10 * time.Second,
	}

	tests := []struct {
		name    string
		mutate  func(*RPC)
		wantErr string
	}{
		{name: "valid", mutate: func(*RPC) {}},
		{name: "plain http", mutate: func(r *RPC) { r.HTTP = "http://localhost:8545" }},
		{name: "plain ws", mutate: func(r *RPC) { r.WS = "ws://localhost:8546" }},
		{name: "zero dial timeout is allowed", mutate: func(r *RPC) { r.DialTimeout = 0 }},
		{name: "missing http", mutate: func(r *RPC) { r.HTTP = "" }, wantErr: "rpc.http is required"},
		{name: "missing ws", mutate: func(r *RPC) { r.WS = "" }, wantErr: "rpc.ws is required"},
		{
			name:    "http with a ws scheme",
			mutate:  func(r *RPC) { r.HTTP = "wss://mainnet.example.com" },
			wantErr: "rpc.http must use one of http/https",
		},
		{
			name:    "ws with an http scheme",
			mutate:  func(r *RPC) { r.WS = "https://mainnet.example.com" },
			wantErr: "rpc.ws must use one of ws/wss",
		},
		{
			name:    "no scheme",
			mutate:  func(r *RPC) { r.HTTP = "mainnet.example.com" },
			wantErr: "rpc.http must use one of http/https",
		},
		{
			name:    "unsupported scheme",
			mutate:  func(r *RPC) { r.HTTP = "ftp://mainnet.example.com" },
			wantErr: "rpc.http must use one of http/https",
		},
		{
			name:    "unparseable url",
			mutate:  func(r *RPC) { r.HTTP = "http://[::1" },
			wantErr: "rpc.http is not a valid URL",
		},
		{
			name:    "negative dial timeout",
			mutate:  func(r *RPC) { r.DialTimeout = -time.Second },
			wantErr: "rpc.dialTimeout cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpc := valid
			tt.mutate(&rpc)

			requireErr(t, rpc.validate(), tt.wantErr)
		})
	}
}

func validURLs() URLs {
	return URLs{
		Explorer:       "https://etherscan.io",
		Checkout:       "https://pay.example.com/checkout?data=",
		Dashboard:      "https://dashboard.example.com",
		Subgraph:       "https://subgraph.example.com/graphql",
		Callback:       "https://market.example.com/api/callback",
		DiscordWebhook: "https://discord.com/api/webhooks/1/abc",
	}
}

func TestURLsValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*URLs)
		wantErr string
	}{
		{name: "valid", mutate: func(*URLs) {}},
		{name: "optional explorer may be empty", mutate: func(u *URLs) { u.Explorer = "" }},
		{name: "optional dashboard may be empty", mutate: func(u *URLs) { u.Dashboard = "" }},
		{name: "optional discord webhook may be empty", mutate: func(u *URLs) { u.DiscordWebhook = "" }},
		{name: "plain http is allowed", mutate: func(u *URLs) { u.Checkout = "http://localhost:3000/checkout?data=" }},
		{
			name:    "checkout is required",
			mutate:  func(u *URLs) { u.Checkout = "" },
			wantErr: "urls.checkout is required",
		},
		{
			name:    "subgraph is required",
			mutate:  func(u *URLs) { u.Subgraph = "" },
			wantErr: "urls.subgraph is required",
		},
		{
			name:    "callback is required",
			mutate:  func(u *URLs) { u.Callback = "" },
			wantErr: "urls.callback is required",
		},
		{
			name:    "non-http scheme",
			mutate:  func(u *URLs) { u.Callback = "ftp://market.example.com" },
			wantErr: "urls.callback must use http/https",
		},
		{
			name:    "websocket scheme is not a URL here",
			mutate:  func(u *URLs) { u.Subgraph = "wss://subgraph.example.com" },
			wantErr: "urls.subgraph must use http/https",
		},
		{
			name:    "no scheme",
			mutate:  func(u *URLs) { u.Dashboard = "dashboard.example.com" },
			wantErr: "urls.dashboard must use http/https",
		},
		{
			name:    "unparseable url",
			mutate:  func(u *URLs) { u.Explorer = "http://[::1" },
			wantErr: "urls.explorer is not a valid URL",
		},
		{
			name:    "explorer with a trailing slash",
			mutate:  func(u *URLs) { u.Explorer = "https://etherscan.io/" },
			wantErr: "urls.explorer must not end in a slash",
		},
		{
			name:    "an optional field that is set is still checked",
			mutate:  func(u *URLs) { u.DiscordWebhook = "not a url at all" },
			wantErr: "urls.discordWebhook must use http/https",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urls := validURLs()
			tt.mutate(&urls)

			requireErr(t, urls.validate(), tt.wantErr)
		})
	}
}

func TestServicesValidate(t *testing.T) {
	tests := []struct {
		name    string
		target  string
		wantErr string
	}{
		{name: "empty disables the fee receiver", target: ""},
		{name: "host and port", target: "fee-receiver:50051"},
		{name: "localhost", target: "localhost:50051"},
		{name: "ipv4", target: "127.0.0.1:50051"},
		{name: "ipv6", target: "[::1]:50051"},
		{
			name:    "no port",
			target:  "fee-receiver",
			wantErr: "must be host:port",
		},
		{
			name:    "a URL is not a gRPC target",
			target:  "http://fee-receiver:50051",
			wantErr: "must be host:port",
		},
		{
			name:    "empty host",
			target:  ":50051",
			wantErr: "must be host:port",
		},
		{
			name:    "empty port",
			target:  "fee-receiver:",
			wantErr: "must be host:port",
		},
		{
			name:    "too many colons",
			target:  "a:b:c",
			wantErr: "must be host:port",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			services := Services{FeeReceiver: tt.target}

			requireErr(t, services.validate(), tt.wantErr)
		})
	}
}

func validContracts() ContractAddresses {
	return ContractAddresses{
		PaymentProcessor:        "0x1111111111111111111111111111111111111111",
		PaymentProcessorStorage: "0x2222222222222222222222222222222222222222",
		SimplePaymentProcessor:  "0x3333333333333333333333333333333333333333",
		Multisig:                "0x4444444444444444444444444444444444444444",
		PaymentAutomation:       "0x5555555555555555555555555555555555555555",
		Notes:                   "0x6666666666666666666666666666666666666666",
		OracleManager:           "0x7777777777777777777777777777777777777777",
	}
}

func TestContractAddressesValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*ContractAddresses)
		wantErr string
	}{
		{name: "valid", mutate: func(*ContractAddresses) {}},
		{
			name:   "oracle manager may be empty",
			mutate: func(c *ContractAddresses) { c.OracleManager = "" },
		},
		{
			name:    "missing payment processor",
			mutate:  func(c *ContractAddresses) { c.PaymentProcessor = "" },
			wantErr: "contracts.paymentProcessor is required",
		},
		{
			name:    "missing storage",
			mutate:  func(c *ContractAddresses) { c.PaymentProcessorStorage = "" },
			wantErr: "contracts.paymentProcessorStorage is required",
		},
		{
			name:    "missing simple processor",
			mutate:  func(c *ContractAddresses) { c.SimplePaymentProcessor = "" },
			wantErr: "contracts.simplePaymentProcessor is required",
		},
		{
			name:    "missing multisig",
			mutate:  func(c *ContractAddresses) { c.Multisig = "" },
			wantErr: "contracts.multisig is required",
		},
		{
			name:    "missing payment automation",
			mutate:  func(c *ContractAddresses) { c.PaymentAutomation = "" },
			wantErr: "contracts.paymentAutomation is required",
		},
		{
			name:    "missing notes",
			mutate:  func(c *ContractAddresses) { c.Notes = "" },
			wantErr: "contracts.notes is required",
		},
		{
			name:    "malformed address",
			mutate:  func(c *ContractAddresses) { c.Multisig = "0x123" },
			wantErr: "is not a valid address",
		},
		{
			name:    "zero address",
			mutate:  func(c *ContractAddresses) { c.Notes = zeroAddress },
			wantErr: "cannot be the zero address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contracts := validContracts()
			tt.mutate(&contracts)

			requireErr(t, contracts.validate(), tt.wantErr)
		})
	}
}

func TestContractAddressesAddresses(t *testing.T) {
	contracts := validContracts()

	got := contracts.Addresses()

	want := map[string]struct {
		got  common.Address
		want string
	}{
		"paymentProcessor":        {got.PaymentProcessor, contracts.PaymentProcessor},
		"paymentProcessorStorage": {got.PaymentProcessorStorage, contracts.PaymentProcessorStorage},
		"simplePaymentProcessor":  {got.SimplePaymentProcessor, contracts.SimplePaymentProcessor},
		"multisig":                {got.Multisig, contracts.Multisig},
		"paymentAutomation":       {got.PaymentAutomation, contracts.PaymentAutomation},
		"notes":                   {got.Notes, contracts.Notes},
		"oracleManager":           {got.OracleManager, contracts.OracleManager},
	}
	for field, pair := range want {
		if pair.got != common.HexToAddress(pair.want) {
			t.Errorf("%s = %s, want %s", field, pair.got, pair.want)
		}
	}
}

func TestContractAddressesAddressesAreDistinct(t *testing.T) {
	got := validContracts().Addresses()

	seen := make(map[common.Address]bool)
	for _, address := range []common.Address{
		got.PaymentProcessor,
		got.PaymentProcessorStorage,
		got.SimplePaymentProcessor,
		got.Multisig,
		got.PaymentAutomation,
		got.Notes,
		got.OracleManager,
	} {
		if seen[address] {
			t.Errorf("address %s is mapped to more than one contract", address)
		}
		seen[address] = true
	}
}

func TestContractAddressesAddressesWithAnUnsetOracle(t *testing.T) {
	contracts := validContracts()
	contracts.OracleManager = ""

	if got := contracts.Addresses().OracleManager; got != (common.Address{}) {
		t.Errorf("OracleManager = %s, want the zero address", got)
	}
}
