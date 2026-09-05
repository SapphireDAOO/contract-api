package invoice

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	"github.com/ethereum/go-ethereum/common"
)

// TokenResolver maps a payment token symbol to the address deployed on the
// selected network. config.Tokens implements it.
type TokenResolver interface {
	Address(symbol string) (common.Address, bool)
	Symbols() []string
}

type CreateInvoiceParam struct {
	OrderId          string
	Seller           string
	Price            int
	EscrowHoldPeriod uint32
	Currency         string
	// PaymentTokens are token symbols such as "ETH" or "USDC", resolved to
	// addresses for the selected network.
	PaymentTokens []string
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

// toPaymentTokens resolves symbols to the addresses the contract is called
// with. An unresolved symbol is an error rather than a zero address, which the
// contracts would read as the native token.
func toPaymentTokens(symbols []string, tokens TokenResolver) ([]common.Address, error) {
	addresses := make([]common.Address, 0, len(symbols))
	for _, symbol := range symbols {
		address, ok := tokens.Address(strings.TrimSpace(symbol))
		if !ok {
			return nil, fmt.Errorf("unknown payment token %q (known: %s)",
				symbol, strings.Join(tokens.Symbols(), ", "))
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func ValidateCreateInvoiceParams(params []CreateInvoiceParam, tokens TokenResolver) error {
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
		for _, symbol := range p.PaymentTokens {
			if _, ok := tokens.Address(strings.TrimSpace(symbol)); !ok {
				return fmt.Errorf("invoice %d: unknown payment token %q (known: %s)",
					i, symbol, strings.Join(tokens.Symbols(), ", "))
			}
		}
	}
	return nil
}

func ConvertParam(param []CreateInvoiceParam, tokens TokenResolver) ([]intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam, error) {
	var results []intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam

	for i, v := range param {
		var result intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam
		precision := CurrencyPrecision[v.Currency]
		multiple := precision - 2
		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(multiple)), nil)
		price := new(big.Int).Mul(big.NewInt(int64(v.Price)), multiplier)

		paymentTokens, err := toPaymentTokens(v.PaymentTokens, tokens)
		if err != nil {
			return nil, fmt.Errorf("invoice %d: %w", i, err)
		}

		result = intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam{
			InvoiceId:        v.OrderId,
			Seller:           common.HexToAddress(strings.TrimSpace(v.Seller)),
			Price:            price,
			EscrowHoldPeriod: v.EscrowHoldPeriod,
			PaymentTokens:    paymentTokens,
		}
		results = append(results, result)

	}

	return results, nil
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
