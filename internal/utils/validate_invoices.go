package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	intermediatedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/IntermediatedPaymentProcessor"
)

type CreateInvoiceParam struct {
	OrderId          string
	Seller           string
	Price            int
	EscrowHoldPeriod uint32
	Currency         string
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
	}
	return nil
}

func ConvertParam(param []CreateInvoiceParam) []intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam {
	var results []intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam

	for _, v := range param {
		var result intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam
		precision := CurrencyPrecision[v.Currency]
		multiple := precision - 2
		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(multiple)), nil)
		price := new(big.Int).Mul(big.NewInt(int64(v.Price)), multiplier)

		result = intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam{
			InvoiceId:        v.OrderId,
			Seller:           common.HexToAddress(strings.TrimSpace(v.Seller)),
			Price:            price,
			EscrowHoldPeriod: v.EscrowHoldPeriod,
		}
		results = append(results, result)

	}

	return results
}

func ValidateInvoices(invoices []intermediatedprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam) error {
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
	}
	return nil
}
