package units

import (
	"math/big"
	"testing"
)

func bi(t *testing.T, s string) *big.Int {
	t.Helper()
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		t.Fatalf("bad big.Int literal %q", s)
	}
	return n
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		decimals int
		want     string
	}{
		{"nil is empty", "", 18, ""},
		{"zero decimals returns the raw integer", "12345", 0, "12345"},
		{"negative decimals returns the raw integer", "12345", -3, "12345"},
		{"whole number", "1000000000000000000", 18, "1.000000000000000000"},
		{"fraction is left padded", "1", 18, "0.000000000000000001"},
		{"mixed value", "1500000000000000000", 18, "1.500000000000000000"},
		{"zero", "0", 18, "0.000000000000000000"},
		{"six decimals", "1234567", 6, "1.234567"},
		{"value smaller than one unit", "999999", 6, "0.999999"},
		{"negative with a whole part", "-1500000000000000000", 18, "-1.500000000000000000"},
		{"negative below one keeps its sign", "-500000000000000000", 18, "-0.500000000000000000"},
		{"negative smallest unit keeps its sign", "-1", 18, "-0.000000000000000001"},
		{"large value", "123456789012345678901234567890", 8, "1234567890123456789012.34567890"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value *big.Int
			if tt.value != "" {
				value = bi(t, tt.value)
			}
			if got := Format(value, tt.decimals); got != tt.want {
				t.Errorf("Format(%v, %d) = %q, want %q", value, tt.decimals, got, tt.want)
			}
		})
	}
}

func TestFormatDoesNotMutateInput(t *testing.T) {
	value := bi(t, "-1500000000000000000")
	before := new(big.Int).Set(value)

	Format(value, 18)

	if value.Cmp(before) != 0 {
		t.Errorf("Format mutated its input: got %s, want %s", value, before)
	}
}

func TestInvert(t *testing.T) {
	tests := []struct {
		name          string
		price         string
		decimals      int
		priceDecimals int
		want          string
	}{
		{"nil price", "", 18, 8, ""},
		{"zero price", "0", 18, 8, ""},
		{"negative price", "-100000000", 18, 8, ""},
		{"one dollar per token", "100000000", 18, 8, "1000000000000000000"},
		{"two dollars per token", "200000000", 18, 8, "500000000000000000"},
		{"six decimal token", "100000000", 6, 8, "1000000"},
		{"truncates rather than rounds", "300000000", 18, 8, "333333333333333333"},
		{"zero decimals", "100000000", 0, 8, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var price *big.Int
			if tt.price != "" {
				price = bi(t, tt.price)
			}

			got := Invert(price, tt.decimals, tt.priceDecimals)

			if tt.want == "" {
				if got != nil {
					t.Errorf("Invert(%v) = %v, want nil", price, got)
				}
				return
			}
			if got == nil {
				t.Fatalf("Invert(%v) = nil, want %s", price, tt.want)
			}
			if got.String() != tt.want {
				t.Errorf("Invert(%v, %d, %d) = %s, want %s", price, tt.decimals, tt.priceDecimals, got, tt.want)
			}
		})
	}
}

func TestInvertThenFormat(t *testing.T) {
	price := bi(t, "250000000")

	rate := Invert(price, 18, 8)

	if got, want := Format(rate, 18), "0.400000000000000000"; got != want {
		t.Errorf("Format(Invert(2.50)) = %q, want %q", got, want)
	}
}

func TestInvertDoesNotMutateInput(t *testing.T) {
	price := bi(t, "300000000")
	before := new(big.Int).Set(price)

	Invert(price, 18, 8)

	if price.Cmp(before) != 0 {
		t.Errorf("Invert mutated its input: got %s, want %s", price, before)
	}
}
