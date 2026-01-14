package callback

import (
	"math/big"
	"strconv"
	"strings"
)

func parseTimestampMillis(value string) int64 {
	if value == "" {
		return 0
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	if parsed >= 1_000_000_000_000 {
		return parsed
	}
	return parsed * 1000
}

func formatTokenAmount(amount *big.Int, decimals int) string {
	if amount == nil {
		return "0"
	}
	if decimals <= 0 {
		return amount.String()
	}
	amountText := amount.Text(10)

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
