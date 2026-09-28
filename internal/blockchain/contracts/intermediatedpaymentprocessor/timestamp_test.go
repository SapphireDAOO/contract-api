package intermediatedpaymentprocessor

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestBlockTimestampMillisUsesTheLogTimestamp(t *testing.T) {
	c := &PaymentProcessor{}

	got, err := c.blockTimestampMillis(context.Background(), &types.Log{BlockTimestamp: 1_700_000_000})
	if err != nil {
		t.Fatalf("blockTimestampMillis returned an error: %v", err)
	}
	if want := int64(1_700_000_000_000); got != want {
		t.Errorf("blockTimestampMillis = %d, want %d", got, want)
	}
}

func TestBlockTimestampMillisFallsBackToTheHeader(t *testing.T) {
	c := &PaymentProcessor{}

	if _, err := c.blockTimestampMillis(context.Background(), &types.Log{}); err == nil {
		t.Error("a log without a timestamp and no client to fetch the header returned no error")
	}
}
