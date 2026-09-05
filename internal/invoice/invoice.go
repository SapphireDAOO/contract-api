package invoice

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	"github.com/ethereum/go-ethereum/common"
)

type CreateInvoiceParam struct {
	OrderId          string
	Seller           string
	Price            int
	EscrowHoldPeriod uint32
	Currency         string
	PaymentTokens    []string
}

func isValidAddress(addr string) bool {
	addr = strings.TrimSpace(addr)

	if !common.IsHexAddress(addr) {
		return false
	}

	if (common.HexToAddress(addr) == common.Address{}) {
		return false
	}

	return true
}

// isValidPaymentToken accepts the zero address, which the contracts use to mean
// the native token.
func isValidPaymentToken(addr string) bool {
	return common.IsHexAddress(strings.TrimSpace(addr))
}

func toPaymentTokens(tokens []string) []common.Address {
	addresses := make([]common.Address, 0, len(tokens))
	for _, token := range tokens {
		addresses = append(addresses, common.HexToAddress(strings.TrimSpace(token)))
	}
	return addresses
}

func ValidateCreateInvoiceParams(params []CreateInvoiceParam) error {
	if len(params) == 0 {
		return fmt.Errorf("no invoice parameters provided")
	}
	for i, p := range params {
		if strings.TrimSpace(p.OrderId) == "" {
			return fmt.Errorf("invoice %d: orderId is required", i)
		}
		if !isValidAddress(p.Seller) {
			return fmt.Errorf("invoice %d: seller %q is not a valid address", i, p.Seller)
		}
		if p.Price <= 0 {
			return fmt.Errorf("invoice %d: price must be greater than zero", i)
		}

		if len(p.PaymentTokens) == 0 {
			return fmt.Errorf("invoice %d: at least one payment token is required", i)
		}
		for j, token := range p.PaymentTokens {
			if !isValidPaymentToken(token) {
				return fmt.Errorf("invoice %d: payment token %d %q is not a valid address", i, j, token)
			}
		}
	}
	return nil
}

func ConvertParam(param []CreateInvoiceParam) []intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam {
	var results []intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam

	for _, v := range param {
		var result intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam
		precision := CurrencyPrecision[v.Currency]
		multiple := precision - 2
		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(multiple)), nil)
		price := new(big.Int).Mul(big.NewInt(int64(v.Price)), multiplier)

		result = intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam{
			InvoiceId:        v.OrderId,
			Seller:           common.HexToAddress(strings.TrimSpace(v.Seller)),
			Price:            price,
			EscrowHoldPeriod: v.EscrowHoldPeriod,
			PaymentTokens:    toPaymentTokens(v.PaymentTokens),
		}
		results = append(results, result)

	}

	return results
}

func ValidateInvoices(invoices []intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam) error {
	for i, inv := range invoices {
		if strings.TrimSpace(inv.InvoiceId) == "" {
			return fmt.Errorf("invoice %d missing orderId", i)
		}
		if (inv.Seller == common.Address{}) {
			return fmt.Errorf("invoice %d missing seller", i)
		}
		if inv.Price == nil {
			return fmt.Errorf("invoice %d missing price", i)
		}
		if len(inv.PaymentTokens) == 0 {
			return fmt.Errorf("invoice %d missing payment tokens", i)
		}
	}
	return nil
}
