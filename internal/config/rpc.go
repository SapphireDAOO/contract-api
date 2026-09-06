package config

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
	"time"
)

// RPC describes how to reach the chain. A URL may carry a provider API key, so
// values can be written as ${ENV_VAR} references and resolved at load time.
type RPC struct {
	HTTP        string        `yaml:"http"`
	WS          string        `yaml:"ws"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
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
