package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// ContractAddresses holds the deployed address of each contract the API talks
// to, as written in the config file.
type ContractAddresses struct {
	PaymentProcessor        string `yaml:"paymentProcessor"`
	PaymentProcessorStorage string `yaml:"paymentProcessorStorage"`
	SimplePaymentProcessor  string `yaml:"simplePaymentProcessor"`
	Multisig                string `yaml:"multisig"`
	PaymentAutomation       string `yaml:"paymentAutomation"`
	Notes                   string `yaml:"notes"`
	OracleManager           string `yaml:"oracleManager"`
}

// Addresses is ContractAddresses parsed into the type the contracts take.
type Addresses struct {
	PaymentProcessor        common.Address
	PaymentProcessorStorage common.Address
	SimplePaymentProcessor  common.Address
	Multisig                common.Address
	PaymentAutomation       common.Address
	Notes                   common.Address
	OracleManager           common.Address
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
		OracleManager:           common.HexToAddress(c.OracleManager),
	}
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
