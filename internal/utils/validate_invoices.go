// internal/utils/validation.go
package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
)

func ValidateInvoices(invoices []advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam) error {
	for i, inv := range invoices {
		if strings.TrimSpace(inv.OrderId) == "" {
			return fmt.Errorf("invoice %d missing orderId", i)
		}
		if (inv.Seller == common.Address{}) {
			return fmt.Errorf("invoice %d missing seller", i)
		}
		if inv.Price == nil {
			return fmt.Errorf("invoice %d missing price", i)
		}
		if inv.EscrowHoldPeriod == nil {
			return fmt.Errorf("invoice %d missing escrowHoldPeriod", i)
		}
	}
	return nil
}
