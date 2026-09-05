// Package config loads deployment settings from a YAML file.
//
// The file holds one section per network; Load resolves exactly one of them,
// so sections for networks that are not deployed yet cannot break startup.
package config

import (
	"fmt"
	"maps"
	"net/url"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v3"
)

const (
	// DefaultPath is used when CONFIG_PATH is unset.
	DefaultPath = "config.yaml"

	// NetworkEnv overrides the network selected by the file.
	NetworkEnv = "NETWORK"

	// defaultDialTimeout applies when rpc.dialTimeout is omitted.
	defaultDialTimeout = 10 * time.Second
)

// RPC describes how to reach the chain. A URL may carry a provider API key, so
// values can be written as ${ENV_VAR} references and resolved at load time.
type RPC struct {
	HTTP        string        `yaml:"http"`
	WS          string        `yaml:"ws"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}

type URLs struct {
	// Explorer is a block explorer root. Empty on chains without one.
	Explorer string `yaml:"explorer"`
	// Checkout is the payment page, ending in its data query parameter.
	Checkout string `yaml:"checkout"`
	// Dashboard is the multisig UI linked from Discord notifications.
	Dashboard string `yaml:"dashboard"`
	// Subgraph answers invoice queries.
	Subgraph string `yaml:"subgraph"`
	// Callback is the endpoint that payment events are posted to.
	Callback string `yaml:"callback"`
	// DiscordWebhook receives contract notifications. Empty disables them.
	DiscordWebhook string `yaml:"discordWebhook"`
}

// ContractAddresses holds the deployed address of each contract the API talks
// to, as written in the config file.
type ContractAddresses struct {
	PaymentProcessor        string `yaml:"paymentProcessor"`
	PaymentProcessorStorage string `yaml:"paymentProcessorStorage"`
	SimplePaymentProcessor  string `yaml:"simplePaymentProcessor"`
	Multisig                string `yaml:"multisig"`
	PaymentAutomation       string `yaml:"paymentAutomation"`
	Notes                   string `yaml:"notes"`
}

// Addresses is ContractAddresses parsed into the type the contracts take.
type Addresses struct {
	PaymentProcessor        common.Address
	PaymentProcessorStorage common.Address
	SimplePaymentProcessor  common.Address
	Multisig                common.Address
	PaymentAutomation       common.Address
	Notes                   common.Address
}

// Config is the resolved settings for the selected network.
type Config struct {
	Network   string
	RPC       RPC
	URLs      URLs
	Contracts ContractAddresses
}

// file mirrors the config file's layout.
type file struct {
	Network  string             `yaml:"network"`
	Networks map[string]network `yaml:"networks"`
}

type network struct {
	RPC       RPC               `yaml:"rpc"`
	URLs      URLs              `yaml:"urls"`
	Contracts ContractAddresses `yaml:"contracts"`
}

// Path returns CONFIG_PATH, falling back to DefaultPath.
func Path() string {
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		return path
	}
	return DefaultPath
}

// Load reads path and resolves the selected network: NETWORK if set, otherwise
// the file's own network key. Only that section is expanded and validated.
func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}

	var parsed file
	if err := yaml.Unmarshal(raw, &parsed); err != nil {
		return nil, fmt.Errorf("parsing config %s: %w", path, err)
	}

	name := parsed.Network
	if override := os.Getenv(NetworkEnv); override != "" {
		name = override
	}
	if name == "" {
		return nil, fmt.Errorf("config %s: no network selected: set the network key or %s", path, NetworkEnv)
	}

	selected, ok := parsed.Networks[name]
	if !ok {
		return nil, fmt.Errorf("config %s: network %q is not defined (available: %s)",
			path, name, strings.Join(slices.Sorted(maps.Keys(parsed.Networks)), ", "))
	}

	for _, field := range []struct {
		key   string
		value *string
	}{
		{"rpc.http", &selected.RPC.HTTP},
		{"rpc.ws", &selected.RPC.WS},
		{"urls.explorer", &selected.URLs.Explorer},
		{"urls.checkout", &selected.URLs.Checkout},
		{"urls.dashboard", &selected.URLs.Dashboard},
		{"urls.subgraph", &selected.URLs.Subgraph},
		{"urls.callback", &selected.URLs.Callback},
		{"urls.discordWebhook", &selected.URLs.DiscordWebhook},
	} {
		if *field.value, err = expandEnv(*field.value); err != nil {
			return nil, fmt.Errorf("config %s: networks.%s.%s: %w", path, name, field.key, err)
		}
	}

	if selected.RPC.DialTimeout == 0 {
		selected.RPC.DialTimeout = defaultDialTimeout
	}

	if err := selected.RPC.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.URLs.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.Contracts.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}

	return &Config{
		Network:   name,
		RPC:       selected.RPC,
		URLs:      selected.URLs,
		Contracts: selected.Contracts,
	}, nil
}

// Addresses parses the configured addresses. Load has already checked that
// every one of them is a valid, non-zero address.
func (c ContractAddresses) Addresses() Addresses {
	return Addresses{
		PaymentProcessor:        common.HexToAddress(c.PaymentProcessor),
		PaymentProcessorStorage: common.HexToAddress(c.PaymentProcessorStorage),
		SimplePaymentProcessor:  common.HexToAddress(c.SimplePaymentProcessor),
		Multisig:                common.HexToAddress(c.Multisig),
		PaymentAutomation:       common.HexToAddress(c.PaymentAutomation),
		Notes:                   common.HexToAddress(c.Notes),
	}
}

// expandEnv resolves ${VAR} references, reporting every variable that is unset
// rather than silently substituting an empty string.
func expandEnv(raw string) (string, error) {
	var missing []string

	expanded := os.Expand(raw, func(key string) string {
		value, ok := os.LookupEnv(key)
		if !ok || value == "" {
			if !slices.Contains(missing, key) {
				missing = append(missing, key)
			}
			return ""
		}
		return value
	})

	if len(missing) > 0 {
		return "", fmt.Errorf("environment variables not set: %s", strings.Join(missing, ", "))
	}
	return expanded, nil
}

// validate checks the URLs that must be present and the scheme of any that are
// set. Everything except checkout is optional: a chain may have no explorer,
// and an unset webhook simply disables Discord notifications.
func (u URLs) validate() error {
	for _, field := range []struct {
		key      string
		raw      string
		required bool
	}{
		{"checkout", u.Checkout, true},
		{"explorer", u.Explorer, false},
		{"dashboard", u.Dashboard, false},
		{"subgraph", u.Subgraph, false},
		{"callback", u.Callback, false},
		{"discordWebhook", u.DiscordWebhook, false},
	} {
		if field.raw == "" {
			if field.required {
				return fmt.Errorf("urls.%s is required", field.key)
			}
			continue
		}
		parsed, err := url.Parse(field.raw)
		if err != nil {
			return fmt.Errorf("urls.%s is not a valid URL: %w", field.key, err)
		}
		if parsed.Scheme != "http" && parsed.Scheme != "https" {
			return fmt.Errorf("urls.%s must use http/https, got %q", field.key, parsed.Scheme)
		}
	}

	if strings.HasSuffix(u.Explorer, "/") {
		return fmt.Errorf("urls.explorer must not end in a slash")
	}
	return nil
}

func (r RPC) validate() error {
	for _, field := range []struct {
		key     string
		raw     string
		schemes []string
	}{
		{"http", r.HTTP, []string{"http", "https"}},
		{"ws", r.WS, []string{"ws", "wss"}},
	} {
		if field.raw == "" {
			return fmt.Errorf("rpc.%s is required", field.key)
		}
		parsed, err := url.Parse(field.raw)
		if err != nil {
			return fmt.Errorf("rpc.%s is not a valid URL: %w", field.key, err)
		}
		if !slices.Contains(field.schemes, parsed.Scheme) {
			return fmt.Errorf("rpc.%s must use one of %s, got %q",
				field.key, strings.Join(field.schemes, "/"), parsed.Scheme)
		}
	}

	if r.DialTimeout < 0 {
		return fmt.Errorf("rpc.dialTimeout cannot be negative")
	}
	return nil
}

func (c ContractAddresses) validate() error {
	for _, field := range []struct {
		key     string
		address string
	}{
		{"paymentProcessor", c.PaymentProcessor},
		{"paymentProcessorStorage", c.PaymentProcessorStorage},
		{"simplePaymentProcessor", c.SimplePaymentProcessor},
		{"multisig", c.Multisig},
		{"paymentAutomation", c.PaymentAutomation},
		{"notes", c.Notes},
	} {
		if field.address == "" {
			return fmt.Errorf("contracts.%s is required", field.key)
		}
		if !common.IsHexAddress(field.address) {
			return fmt.Errorf("contracts.%s %q is not a valid address", field.key, field.address)
		}
		if common.HexToAddress(field.address) == (common.Address{}) {
			return fmt.Errorf("contracts.%s cannot be the zero address", field.key)
		}
	}
	return nil
}
