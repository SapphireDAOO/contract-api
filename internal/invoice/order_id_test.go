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

func TestOrderIDToUint216MatchesABIEncodedKeccak(t *testing.T) {
	orderIds := []string{
		"",
		"a",
		"ORDER-1",
		"order-0000-1111-2222",
		strings.Repeat("x", 32),
		strings.Repeat("y", 33),
		strings.Repeat("z", 100),
		"unicode ☃ order",
	}

	for _, orderId := range orderIds {
		t.Run(orderId, func(t *testing.T) {
			got := OrderIDToUint216(orderId)

			if want := wantUint216(orderId).String(); got != want {
				t.Errorf("OrderIDToUint216(%q) = %s, want %s", orderId, got, want)
			}
		})
	}
}

func TestOrderIDToUint216FitsIn216Bits(t *testing.T) {
	limit := new(big.Int).Lsh(big.NewInt(1), 216)

	for _, orderId := range []string{"", "a", "ORDER-1", strings.Repeat("q", 200)} {
		id, ok := new(big.Int).SetString(OrderIDToUint216(orderId), 10)
		if !ok {
			t.Fatalf("OrderIDToUint216(%q) is not a base-10 integer", orderId)
		}
		if id.Sign() < 0 {
			t.Errorf("OrderIDToUint216(%q) = %s, want a non-negative value", orderId, id)
		}
		if id.Cmp(limit) >= 0 {
			t.Errorf("OrderIDToUint216(%q) = %s, which does not fit in 216 bits", orderId, id)
		}
	}
}

func TestOrderIDToUint216IsDeterministic(t *testing.T) {
	const orderId = "ORDER-42"

	first := OrderIDToUint216(orderId)
	for range 5 {
		if got := OrderIDToUint216(orderId); got != first {
			t.Fatalf("OrderIDToUint216(%q) = %s then %s; it must be stable", orderId, first, got)
		}
	}
}

func TestOrderIDToUint216IsCaseAndWhitespaceSensitive(t *testing.T) {
	pairs := [][2]string{
		{"order-1", "ORDER-1"},
		{"order-1", " order-1"},
		{"order-1", "order-1 "},
		{"order-1", "order-2"},
	}

	for _, pair := range pairs {
		if OrderIDToUint216(pair[0]) == OrderIDToUint216(pair[1]) {
			t.Errorf("OrderIDToUint216(%q) collides with OrderIDToUint216(%q)", pair[0], pair[1])
		}
	}
}
