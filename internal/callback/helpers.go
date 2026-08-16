package callback

import (
	"math/big"
	"strings"
)

func formatTokenAmount(amount *big.Int, decimals int) string {
	if amount == nil {
		return "0"
	}
	if decimals <= 0 {
		return amount.String()
	}

	absAmount := new(big.Int).Abs(amount)
	amountText := absAmount.Text(10)

	if len(amountText) <= decimals {
		amountText = strings.Repeat("0", decimals-len(amountText)+1) + amountText
	}

	intPart := amountText[:len(amountText)-decimals]
	fracPart := strings.TrimRight(amountText[len(amountText)-decimals:], "0")

	if fracPart == "" {
		return intPart
	}

	return intPart + "." + fracPart
}
