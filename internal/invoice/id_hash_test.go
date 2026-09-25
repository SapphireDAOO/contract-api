package invoice

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func abiEncodeString(s string) []byte {
	encoded := make([]byte, 0, 64+len(s)+31)
	encoded = append(encoded, common.LeftPadBytes(big.NewInt(32).Bytes(), 32)...)
	encoded = append(encoded, common.LeftPadBytes(big.NewInt(int64(len(s))).Bytes(), 32)...)

	data := []byte(s)
	if len(data)%32 != 0 {
		data = common.RightPadBytes(data, ((len(data)/32)+1)*32)
	}
	return append(encoded, data...)
}

func wantUint216(s string) *big.Int {
	hash := new(big.Int).SetBytes(crypto.Keccak256(abiEncodeString(s)))
	mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 216), big.NewInt(1))
	return new(big.Int).And(hash, mask)
}

func TestInvoiceIDToUint216MatchesABIEncodedKeccak(t *testing.T) {
	invoiceIds := []string{
		"",
		"a",
		"ORDER-1",
		"order-0000-1111-2222",
		strings.Repeat("x", 32),
		strings.Repeat("y", 33),
		strings.Repeat("z", 100),
		"unicode ☃ invoice",
	}

	for _, invoiceId := range invoiceIds {
		t.Run(invoiceId, func(t *testing.T) {
			got := InvoiceIDToUint216(invoiceId)

			if want := wantUint216(invoiceId).String(); got != want {
				t.Errorf("InvoiceIDToUint216(%q) = %s, want %s", invoiceId, got, want)
			}
		})
	}
}

func TestInvoiceIDToUint216FitsIn216Bits(t *testing.T) {
	limit := new(big.Int).Lsh(big.NewInt(1), 216)

	for _, invoiceId := range []string{"", "a", "ORDER-1", strings.Repeat("q", 200)} {
		id, ok := new(big.Int).SetString(InvoiceIDToUint216(invoiceId), 10)
		if !ok {
			t.Fatalf("InvoiceIDToUint216(%q) is not a base-10 integer", invoiceId)
		}
		if id.Sign() < 0 {
			t.Errorf("InvoiceIDToUint216(%q) = %s, want a non-negative value", invoiceId, id)
		}
		if id.Cmp(limit) >= 0 {
			t.Errorf("InvoiceIDToUint216(%q) = %s, which does not fit in 216 bits", invoiceId, id)
		}
	}
}

func TestInvoiceIDToUint216IsDeterministic(t *testing.T) {
	const invoiceId = "ORDER-42"

	first := InvoiceIDToUint216(invoiceId)
	for range 5 {
		if got := InvoiceIDToUint216(invoiceId); got != first {
			t.Fatalf("InvoiceIDToUint216(%q) = %s then %s; it must be stable", invoiceId, first, got)
		}
	}
}

func TestInvoiceIDToUint216IsCaseAndWhitespaceSensitive(t *testing.T) {
	pairs := [][2]string{
		{"order-1", "ORDER-1"},
		{"order-1", " order-1"},
		{"order-1", "order-1 "},
		{"order-1", "order-2"},
	}

	for _, pair := range pairs {
		if InvoiceIDToUint216(pair[0]) == InvoiceIDToUint216(pair[1]) {
			t.Errorf("InvoiceIDToUint216(%q) collides with InvoiceIDToUint216(%q)", pair[0], pair[1])
		}
	}
}
