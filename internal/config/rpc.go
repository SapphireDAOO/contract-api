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
			return fmt.Errorf("rpc.%s %w", field.key, ErrRequired)
		}
		parsed, err := url.Parse(field.raw)
		if err != nil {
			return fmt.Errorf("rpc.%s %w: %w", field.key, ErrInvalidURL, err)
		}
		if !slices.Contains(field.schemes, parsed.Scheme) {
			return fmt.Errorf("rpc.%s %w %q, want %s",
				field.key, ErrUnsupportedScheme, parsed.Scheme, strings.Join(field.schemes, "/"))
		}
	}

	if r.DialTimeout < 0 {
		return fmt.Errorf("rpc.dialTimeout %w", ErrNegative)
	}
	return nil
}
