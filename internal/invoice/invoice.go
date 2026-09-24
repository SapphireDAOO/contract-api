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
	OrderId          string `json:"orderId"`
	Seller           string `json:"seller"`
	Price            int    `json:"price"`
	EscrowHoldPeriod uint32 `json:"escrowHoldPeriod"`
	Currency         string `json:"currency"`
}

type CreateInvoiceRequest struct {
	Invoices      []CreateInvoiceParam `json:"invoices"`
	PaymentTokens []string             `json:"paymentTokens"`
}

// ParseAddress reads a value as a non-zero account address. The zero address
// is rejected: it is never a real account, and the contracts read it as the
// native token.
func ParseAddress(value string) (common.Address, bool) {
	trimmed := strings.TrimSpace(value)

	if !common.IsHexAddress(trimmed) {
		return common.Address{}, false
	}

	address := common.HexToAddress(trimmed)
	if address == (common.Address{}) {
		return common.Address{}, false
	}

	return address, true
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

func ValidateCreateInvoiceParams(request CreateInvoiceRequest, tokens TokenResolver) error {
	if len(request.Invoices) == 0 {
		return fmt.Errorf("no invoice parameters provided")
	}

	if len(request.PaymentTokens) == 0 {
		return fmt.Errorf("at least one payment token is required")
	}
	for _, symbol := range request.PaymentTokens {
		if _, ok := tokens.Address(strings.TrimSpace(symbol)); !ok {
			return fmt.Errorf("unknown payment token %q (known: %s)",
				symbol, strings.Join(tokens.Symbols(), ", "))
		}
	}

	for i, p := range request.Invoices {
		if strings.TrimSpace(p.OrderId) == "" {
			return fmt.Errorf("invoice %d: orderId is required", i)
		}
		if _, ok := ParseAddress(p.Seller); !ok {
			return fmt.Errorf("invoice %d: seller %q is not a valid address", i, p.Seller)
		}
		if p.Price <= 0 {
			return fmt.Errorf("invoice %d: price must be greater than zero", i)
		}
	}
	return nil
}

func ConvertParam(request CreateInvoiceRequest, tokens TokenResolver) ([]intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam, error) {
	var results []intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam

	paymentTokens, err := toPaymentTokens(request.PaymentTokens, tokens)
	if err != nil {
		return nil, err
	}

	for _, v := range request.Invoices {
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
