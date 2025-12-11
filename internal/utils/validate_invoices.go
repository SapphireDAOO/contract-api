// internal/utils/validation.go
package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
)

type CreateInvoiceParam struct {
	OrderId          string
	Seller           string
	Price            int
	EscrowHoldPeriod int64
	Currency         string
}

//

func ConvertParam(param []CreateInvoiceParam) []advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam {
	var results []advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam

	for _, v := range param {
		var result advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam
		precision := CurrencyPrecision[v.Currency]
		multiple := precision - 4
		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(multiple)), nil)
		price := new(big.Int).Mul(big.NewInt(int64(v.Price)), multiplier)

		result = advancedprocessor.IAdvancedPaymentProcessorInvoiceCreationParam{
			OrderId:          v.OrderId,
			Seller:           common.HexToAddress(v.Seller),
			Price:            price,
			EscrowHoldPeriod: big.NewInt(v.EscrowHoldPeriod),
		}
		results = append(results, result)

	}

	return results
}

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
