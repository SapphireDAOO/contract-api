package config

import (
	"fmt"
	"net"
	"strings"
)

// Services holds the sidecar services the API calls out to, as gRPC targets
// rather than URLs.
type Services struct {
	// FeeReceiver is the fee-receiver sidecar (services/fee-receiver), written
	// as host:port. Empty disables the fee receiver endpoints, the same way an
	// unset oracleManager disables exchange rates.
	FeeReceiver string `yaml:"feeReceiver"`
}

func (s Services) validate() error {
	if s.FeeReceiver == "" {
		return nil
	}

	host, port, err := net.SplitHostPort(s.FeeReceiver)
	if err != nil {
		return fmt.Errorf("services.feeReceiver %q must be host:port: %w", s.FeeReceiver, err)
	}
	if strings.TrimSpace(host) == "" || strings.TrimSpace(port) == "" {
		return fmt.Errorf("services.feeReceiver %q must be host:port", s.FeeReceiver)
	}
	return nil
}
