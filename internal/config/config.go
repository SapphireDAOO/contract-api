// Package config loads deployment settings from a YAML file.
//
// The file holds one section per network; Load resolves exactly one of them,
// so sections for networks that are not deployed yet cannot break startup.
//
// Load is the only entry point: every section is parsed, expanded and
// validated through it.
package config

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"time"

	"go.yaml.in/yaml/v4"
)

const (
	// DefaultPath is used when CONFIG_PATH is unset.
	DefaultPath = "config.yaml"

	// NetworkEnv overrides the network selected by the file.
	NetworkEnv = "NETWORK"

	// defaultDialTimeout applies when rpc.dialTimeout is omitted.
	defaultDialTimeout = 10 * time.Second
)

// Config is the resolved settings for the selected network.
type Config struct {
	Network   string
	SignerKey string
	RPC       RPC
	URLs      URLs
	Tokens    Tokens
	Contracts ContractAddresses
	Services  Services
}

// file mirrors the config file's layout.
type file struct {
	Network  string             `yaml:"network"`
	Networks map[string]network `yaml:"networks"`
}

type network struct {
	SignerKey string            `yaml:"signerKey"`
	RPC       RPC               `yaml:"rpc"`
	URLs      URLs              `yaml:"urls"`
	Tokens    Tokens            `yaml:"tokens"`
	Contracts ContractAddresses `yaml:"contracts"`
	Services  Services          `yaml:"services"`
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
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config %s: %w", path, err)
	}

	var parsed file
	if err := yaml.Unmarshal([]byte(os.ExpandEnv(string(data))), &parsed); err != nil {
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

	if selected.RPC.DialTimeout == 0 {
		selected.RPC.DialTimeout = defaultDialTimeout
	}

	if err := validateSignerKey(selected.SignerKey); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.RPC.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.URLs.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.Tokens.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.Contracts.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}
	if err := selected.Services.validate(); err != nil {
		return nil, fmt.Errorf("config %s: networks.%s.%w", path, name, err)
	}

	return &Config{
		Network:   name,
		SignerKey: selected.SignerKey,
		RPC:       selected.RPC,
		URLs:      selected.URLs,
		Tokens:    selected.Tokens,
		Contracts: selected.Contracts,
		Services:  selected.Services,
	}, nil
}
