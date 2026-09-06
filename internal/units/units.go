package units

import (
	"math/big"
	"strings"
)

func Format(value *big.Int, decimals int) string {
	if value == nil {
		return ""
	}
	if decimals <= 0 {
		return value.String()
	}

	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	whole, fraction := new(big.Int).QuoRem(value, scale, new(big.Int))

	// Quo truncates toward zero, so a value between -1 and 0 leaves a zero
	// whole part and the sign only on the fraction.
	sign := ""
	if fraction.Sign() < 0 {
		fraction.Neg(fraction)
		if whole.Sign() == 0 {
			sign = "-"
		}
	}

	digits := fraction.String()

	return sign + whole.String() + "." + strings.Repeat("0", decimals-len(digits)) + digits
}

//	10^(decimals+priceDecimals) / price

func Invert(price *big.Int, decimals, priceDecimals int) *big.Int {
	if price == nil || price.Sign() <= 0 {
		return nil
	}

	numerator := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals+priceDecimals)), nil)

	return new(big.Int).Quo(numerator, price)
}
