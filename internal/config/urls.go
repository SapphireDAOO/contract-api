package config

import (
	"fmt"
	"net/url"
	"strings"
)

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

func (u URLs) validate() error {
	for _, field := range []struct {
		key      string
		raw      string
		required bool
	}{
		{"checkout", u.Checkout, true},
		{"explorer", u.Explorer, false},
		{"dashboard", u.Dashboard, false},
		{"subgraph", u.Subgraph, true},
		{"callback", u.Callback, true},
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
