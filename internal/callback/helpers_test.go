package callback

import (
	"math/big"
	"testing"
)

func TestFormatTokenAmount(t *testing.T) {
	tests := []struct {
		name     string
		amount   string
		decimals int
		want     string
	}{
		{"nil", "", 18, "0"},
		{"zero decimals returns the raw integer", "12345", 0, "12345"},
		{"negative decimals returns the raw integer", "12345", -1, "12345"},
		{"one whole token", "1000000000000000000", 18, "1"},
		{"trailing zeros are trimmed", "1500000000000000000", 18, "1.5"},
		{"smallest unit", "1", 18, "0.000000000000000001"},
		{"zero", "0", 18, "0"},
		{"below one token", "500000000000000000", 18, "0.5"},
		{"six decimal token", "1500000", 6, "1.5"},
		{"six decimal token, whole", "1000000", 6, "1"},
		{"six decimal token, sub-unit", "1", 6, "0.000001"},
		{"value shorter than the decimal count", "123", 18, "0.000000000000000123"},
		{"many significant digits", "1234567890123456789", 18, "1.234567890123456789"},
		{"large value", "123456789000000000000000", 18, "123456.789"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var amount *big.Int
			if tt.amount != "" {
				amount, _ = new(big.Int).SetString(tt.amount, 10)
			}

			if got := formatTokenAmount(amount, tt.decimals); got != tt.want {
				t.Errorf("formatTokenAmount(%v, %d) = %q, want %q", amount, tt.decimals, got, tt.want)
			}
		})
	}
}

func TestFormatTokenAmountUsesTheAbsoluteValue(t *testing.T) {
	amount, _ := new(big.Int).SetString("-1500000000000000000", 10)

	if got, want := formatTokenAmount(amount, 18), "1.5"; got != want {
		t.Errorf("formatTokenAmount(negative) = %q, want %q", got, want)
	}
}

func TestFormatTokenAmountDoesNotMutateInput(t *testing.T) {
	amount, _ := new(big.Int).SetString("-1500000000000000000", 10)
	before := new(big.Int).Set(amount)

	formatTokenAmount(amount, 18)

	if amount.Cmp(before) != 0 {
		t.Errorf("formatTokenAmount mutated its input: got %s, want %s", amount, before)
	}
}

func TestFormatTokenAmountIsAlwaysParseable(t *testing.T) {
	amounts := []string{"0", "1", "999999999999999999", "1000000000000000000", "10000000000000000000"}

	for _, raw := range amounts {
		amount, _ := new(big.Int).SetString(raw, 10)
		got := formatTokenAmount(amount, 18)

		if _, ok := new(big.Float).SetString(got); !ok {
			t.Errorf("formatTokenAmount(%s, 18) = %q, which is not a number", raw, got)
		}
		if got != "" && got[len(got)-1] == '.' {
			t.Errorf("formatTokenAmount(%s, 18) = %q, which ends in a bare separator", raw, got)
		}
	}
}
