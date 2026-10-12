package config

import (
	"fmt"
	"strings"
)

type Auth struct {
	Domains []string `yaml:"domains"`
}

func (a Auth) validate() error {
	for i, domain := range a.Domains {
		domain = strings.TrimSpace(domain)
		if domain == "" {
			return fmt.Errorf("auth.domains[%d] %w", i, ErrEmpty)
		}
		if strings.Contains(domain, "://") || strings.ContainsAny(domain, "/ ") {
			return fmt.Errorf("auth.domains[%d] %q %w", i, domain, ErrNotHost)
		}
	}
	return nil
}
