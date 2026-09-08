package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"go.yaml.in/yaml/v4"
)

const validFile = `
network: sepolia
networks:
  sepolia:
    signerKey: "0x4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"
    rpc:
      http: https://sepolia.example.com/v3/key
      ws: wss://sepolia.example.com/v3/key
      dialTimeout: 5s
    urls:
      explorer: https://sepolia.etherscan.io
      checkout: https://pay.example.com/checkout?data=
      dashboard: https://dashboard.example.com
      subgraph: https://subgraph.example.com/graphql
      callback: https://market.example.com/api/callback
      discordWebhook: https://discord.com/api/webhooks/1/abc
    tokens:
      ETH:
        address: "0x0000000000000000000000000000000000000000"
        decimals: 18
      USDC:
        address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
        decimals: 6
    contracts:
      paymentProcessor: "0x1111111111111111111111111111111111111111"
      paymentProcessorStorage: "0x2222222222222222222222222222222222222222"
      simplePaymentProcessor: "0x3333333333333333333333333333333333333333"
      multisig: "0x4444444444444444444444444444444444444444"
      paymentAutomation: "0x5555555555555555555555555555555555555555"
      notes: "0x6666666666666666666666666666666666666666"
      oracleManager: "0x7777777777777777777777777777777777777777"
    services:
      feeReceiver: fee-receiver:50051
  mainnet:
    signerKey: "0x4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"
    rpc:
      http: https://mainnet.example.com/v3/key
      ws: wss://mainnet.example.com/v3/key
    urls:
      checkout: https://pay.example.com/checkout?data=
      subgraph: https://subgraph.example.com/graphql
      callback: https://market.example.com/api/callback
    tokens:
      ETH:
        address: "0x0000000000000000000000000000000000000000"
        decimals: 18
    contracts:
      paymentProcessor: "0xaaaa111111111111111111111111111111111111"
      paymentProcessorStorage: "0x2222222222222222222222222222222222222222"
      simplePaymentProcessor: "0x3333333333333333333333333333333333333333"
      multisig: "0x4444444444444444444444444444444444444444"
      paymentAutomation: "0x5555555555555555555555555555555555555555"
      notes: "0x6666666666666666666666666666666666666666"
`

func writeConfig(t *testing.T, contents string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatalf("writing the test config: %v", err)
	}
	return path
}

func TestLoad(t *testing.T) {
	t.Setenv(NetworkEnv, "")

	cfg, err := Load(writeConfig(t, validFile))
	if err != nil {
		t.Fatalf("Load returned %v", err)
	}

	if cfg.Network != "sepolia" {
		t.Errorf("Network = %q, want sepolia", cfg.Network)
	}
	if cfg.RPC.HTTP != "https://sepolia.example.com/v3/key" {
		t.Errorf("RPC.HTTP = %q", cfg.RPC.HTTP)
	}
	if cfg.RPC.DialTimeout != 5*time.Second {
		t.Errorf("RPC.DialTimeout = %v, want 5s", cfg.RPC.DialTimeout)
	}
	if cfg.URLs.Explorer != "https://sepolia.etherscan.io" {
		t.Errorf("URLs.Explorer = %q", cfg.URLs.Explorer)
	}
	if len(cfg.Tokens) != 2 {
		t.Errorf("Tokens has %d entries, want 2", len(cfg.Tokens))
	}
	if cfg.Contracts.PaymentProcessor != "0x1111111111111111111111111111111111111111" {
		t.Errorf("Contracts.PaymentProcessor = %q", cfg.Contracts.PaymentProcessor)
	}
	if cfg.Services.FeeReceiver != "fee-receiver:50051" {
		t.Errorf("Services.FeeReceiver = %q", cfg.Services.FeeReceiver)
	}
}

func TestLoadNetworkEnvOverride(t *testing.T) {
	t.Setenv(NetworkEnv, "mainnet")

	cfg, err := Load(writeConfig(t, validFile))
	if err != nil {
		t.Fatalf("Load returned %v", err)
	}

	if cfg.Network != "mainnet" {
		t.Errorf("Network = %q, want mainnet", cfg.Network)
	}
	if cfg.Contracts.PaymentProcessor != "0xaaaa111111111111111111111111111111111111" {
		t.Errorf("the mainnet section was not selected: PaymentProcessor = %q", cfg.Contracts.PaymentProcessor)
	}
}

func TestLoadDefaultsTheDialTimeout(t *testing.T) {
	t.Setenv(NetworkEnv, "mainnet")

	cfg, err := Load(writeConfig(t, validFile))
	if err != nil {
		t.Fatalf("Load returned %v", err)
	}

	if cfg.RPC.DialTimeout != defaultDialTimeout {
		t.Errorf("RPC.DialTimeout = %v, want the %v default", cfg.RPC.DialTimeout, defaultDialTimeout)
	}
}

func TestLoadExpandsEnvironmentReferences(t *testing.T) {
	t.Setenv(NetworkEnv, "")
	t.Setenv("TEST_RPC_KEY", "the-secret-key")

	contents := strings.Replace(validFile,
		"https://sepolia.example.com/v3/key",
		"https://sepolia.example.com/v3/${TEST_RPC_KEY}", 1)

	cfg, err := Load(writeConfig(t, contents))
	if err != nil {
		t.Fatalf("Load returned %v", err)
	}

	if want := "https://sepolia.example.com/v3/the-secret-key"; cfg.RPC.HTTP != want {
		t.Errorf("RPC.HTTP = %q, want %q", cfg.RPC.HTTP, want)
	}
}

func TestLoadErrors(t *testing.T) {
	tests := []struct {
		name    string
		network string
		mutate  func(string) string
		wantErr string
	}{
		{
			name:    "no network selected",
			mutate:  func(s string) string { return strings.Replace(s, "network: sepolia\n", "", 1) },
			wantErr: "no network selected",
		},
		{
			name:    "network is not defined",
			network: "goerli",
			mutate:  func(s string) string { return s },
			wantErr: `network "goerli" is not defined`,
		},
		{
			name:    "malformed yaml",
			mutate:  func(string) string { return "network: [unclosed\n" },
			wantErr: "parsing config",
		},
		{
			name: "invalid signer key",
			mutate: func(s string) string {
				return strings.Replace(s, `"0x4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"`, `"0xdead"`, 1)
			},
			wantErr: "signerKey must be 32 bytes of hex",
		},
		{
			name: "invalid rpc",
			mutate: func(s string) string {
				return strings.Replace(s, "http: https://sepolia.example.com/v3/key", "http: ftp://sepolia.example.com", 1)
			},
			wantErr: "rpc.http must use one of http/https",
		},
		{
			name: "invalid urls",
			mutate: func(s string) string {
				return strings.Replace(s, "explorer: https://sepolia.etherscan.io", "explorer: https://sepolia.etherscan.io/", 1)
			},
			wantErr: "urls.explorer must not end in a slash",
		},
		{
			name: "invalid token",
			mutate: func(s string) string {
				return strings.Replace(s, `address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"`, `address: "nope"`, 1)
			},
			wantErr: "is not a valid address",
		},
		{
			name: "invalid contract",
			mutate: func(s string) string {
				return strings.Replace(s, `multisig: "0x4444444444444444444444444444444444444444"`, `multisig: "0x44"`, 1)
			},
			wantErr: "contracts.multisig",
		},
		{
			name: "invalid service target",
			mutate: func(s string) string {
				return strings.Replace(s, "feeReceiver: fee-receiver:50051", "feeReceiver: fee-receiver", 1)
			},
			wantErr: "services.feeReceiver",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(NetworkEnv, tt.network)

			cfg, err := Load(writeConfig(t, tt.mutate(validFile)))

			if err == nil {
				t.Fatalf("Load returned %+v, want an error containing %q", cfg, tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestLoadMissingFile(t *testing.T) {
	t.Setenv(NetworkEnv, "")

	_, err := Load(filepath.Join(t.TempDir(), "does-not-exist.yaml"))

	if err == nil {
		t.Fatal("Load of a missing file returned no error")
	}
	if !strings.Contains(err.Error(), "reading config") {
		t.Errorf("error = %q, want it to mention reading the config", err)
	}
}

func TestLoadErrorsNameTheFileAndNetwork(t *testing.T) {
	t.Setenv(NetworkEnv, "")

	contents := strings.Replace(validFile,
		`"0x4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"`, `"0xdead"`, 1)
	path := writeConfig(t, contents)

	_, err := Load(path)

	if err == nil {
		t.Fatal("Load returned no error")
	}
	if !strings.Contains(err.Error(), path) {
		t.Errorf("error = %q, want it to name the config path", err)
	}
	if !strings.Contains(err.Error(), "networks.sepolia") {
		t.Errorf("error = %q, want it to name the network section", err)
	}
}

func TestLoadUnknownNetworkListsAvailable(t *testing.T) {
	t.Setenv(NetworkEnv, "goerli")

	_, err := Load(writeConfig(t, validFile))

	if err == nil {
		t.Fatal("Load returned no error")
	}
	if !strings.Contains(err.Error(), "mainnet, sepolia") {
		t.Errorf("error = %q, want it to list the defined networks in order", err)
	}
}

func TestLoadIgnoresOtherNetworks(t *testing.T) {
	t.Setenv(NetworkEnv, "sepolia")

	contents := strings.Replace(validFile,
		"http: https://mainnet.example.com/v3/key", "http: not-a-url-at-all", 1)

	cfg, err := Load(writeConfig(t, contents))
	if err != nil {
		t.Fatalf("Load returned %v; a broken unselected network must not fail startup", err)
	}
	if cfg.Network != "sepolia" {
		t.Errorf("Network = %q, want sepolia", cfg.Network)
	}
}

func TestPath(t *testing.T) {
	t.Run("defaults when unset", func(t *testing.T) {
		t.Setenv("CONFIG_PATH", "")

		if got := Path(); got != DefaultPath {
			t.Errorf("Path() = %q, want %q", got, DefaultPath)
		}
	})

	t.Run("uses CONFIG_PATH when set", func(t *testing.T) {
		t.Setenv("CONFIG_PATH", "/etc/contract-api/config.yaml")

		if got := Path(); got != "/etc/contract-api/config.yaml" {
			t.Errorf("Path() = %q, want the CONFIG_PATH value", got)
		}
	})
}

func TestRepositoryConfigStructure(t *testing.T) {
	path := filepath.Join("..", "..", "config.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Skipf("no config.yaml at the repository root: %v", err)
	}

	var parsed file
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("config.yaml does not parse: %v", err)
	}

	if len(parsed.Networks) == 0 {
		t.Fatal("config.yaml defines no networks")
	}
	for name, network := range parsed.Networks {
		t.Run(name, func(t *testing.T) {
			if network.SignerKey == "" {
				t.Error("signerKey is missing")
			}
			if network.RPC.HTTP == "" || network.RPC.WS == "" {
				t.Error("rpc.http and rpc.ws are both required")
			}
			if len(network.Tokens) == 0 {
				t.Error("no payment tokens are configured")
			}
			if network.URLs.Checkout == "" {
				t.Error("urls.checkout is missing")
			}

			if network.Contracts.PaymentProcessor == "" {
				return
			}
			for field, address := range map[string]string{
				"paymentProcessorStorage": network.Contracts.PaymentProcessorStorage,
				"simplePaymentProcessor":  network.Contracts.SimplePaymentProcessor,
				"multisig":                network.Contracts.Multisig,
				"paymentAutomation":       network.Contracts.PaymentAutomation,
				"notes":                   network.Contracts.Notes,
			} {
				if address == "" {
					t.Errorf("contracts.%s is blank in a deployed section", field)
				}
			}
		})
	}
}
