package multisig

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestProgressLine(t *testing.T) {
	three := big.NewInt(3)

	tests := []struct {
		name      string
		current   stage
		approvals *big.Int
		threshold *big.Int
		want      string
	}{
		{
			name: "proposed", current: stageProposed,
			approvals: big.NewInt(1), threshold: three,
			want: "**📝 Proposed**  →  ✍️ Approvals 1/3  →  🚀 Execution",
		},
		{
			name: "collecting approvals", current: stageApprovals,
			approvals: big.NewInt(2), threshold: three,
			want: "📝 Proposed  →  **✍️ Approvals 2/3**  →  🚀 Execution",
		},
		{
			name: "threshold reached", current: stageApprovals,
			approvals: three, threshold: three,
			want: "📝 Proposed  →  **✍️ Approvals 3/3**  →  🚀 Execution",
		},
		{
			name: "executed", current: stageExecuted,
			approvals: three, threshold: three,
			want: "📝 Proposed  →  ✍️ Approvals 3/3  →  **🚀 Executed**",
		},
		{
			name: "canceled", current: stageCanceled,
			approvals: big.NewInt(1), threshold: three,
			want: "📝 Proposed  →  ✍️ Approvals 1/3  →  **🚫 Canceled**",
		},
		{
			name: "chain read failed", current: stageApprovals,
			want: "📝 Proposed  →  **✍️ Approvals**  →  🚀 Execution",
		},
		{
			name: "threshold unknown", current: stageApprovals,
			approvals: big.NewInt(2),
			want:      "📝 Proposed  →  **✍️ Approvals 2**  →  🚀 Execution",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := progressLine(tt.current, tt.approvals, tt.threshold); got != tt.want {
				t.Errorf("progressLine = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestProgressLineAlwaysShowsAllThreeStages(t *testing.T) {
	three := big.NewInt(3)

	for _, current := range []stage{stageProposed, stageApprovals, stageExecuted, stageCanceled} {
		got := progressLine(current, big.NewInt(1), three)

		if !strings.Contains(got, "Proposed") {
			t.Errorf("progressLine(%d) = %q, want the propose stage", current, got)
		}
		if !strings.Contains(got, "Approvals") {
			t.Errorf("progressLine(%d) = %q, want the approvals stage", current, got)
		}
		if strings.Count(got, "→") != 2 {
			t.Errorf("progressLine(%d) = %q, want two arrows", current, got)
		}
		if strings.Count(got, "**") != 2 {
			t.Errorf("progressLine(%d) = %q, want exactly one emphasised stage", current, got)
		}
	}
}

func TestThresholdReached(t *testing.T) {
	tests := []struct {
		name      string
		approvals *big.Int
		threshold *big.Int
		want      bool
	}{
		{name: "below", approvals: big.NewInt(1), threshold: big.NewInt(3)},
		{name: "one short", approvals: big.NewInt(2), threshold: big.NewInt(3)},
		{name: "exactly at the threshold", approvals: big.NewInt(3), threshold: big.NewInt(3), want: true},
		{name: "above", approvals: big.NewInt(4), threshold: big.NewInt(3), want: true},
		{name: "unknown approvals", threshold: big.NewInt(3)},
		{name: "unknown threshold", approvals: big.NewInt(3)},
		{name: "both unknown"},
		{name: "zero threshold is not met", approvals: big.NewInt(0), threshold: big.NewInt(0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thresholdReached(tt.approvals, tt.threshold); got != tt.want {
				t.Errorf("thresholdReached(%v, %v) = %v, want %v",
					tt.approvals, tt.threshold, got, tt.want)
			}
		})
	}
}

func TestApprovalsText(t *testing.T) {
	tests := []struct {
		name      string
		approvals *big.Int
		threshold *big.Int
		want      string
	}{
		{name: "both known", approvals: big.NewInt(2), threshold: big.NewInt(3), want: "Approvals 2/3"},
		{name: "threshold unknown", approvals: big.NewInt(2), want: "Approvals 2"},
		{name: "approvals unknown", threshold: big.NewInt(3), want: "Approvals"},
		{name: "neither known", want: "Approvals"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := approvalsText(tt.approvals, tt.threshold); got != tt.want {
				t.Errorf("approvalsText = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTransactionApprovedProducesNoSeparateMessage(t *testing.T) {
	c := &Multisig{address: &multisigAddress}

	embed, err := c.buildEmbed(context.Background(), &types.Log{
		Topics: []common.Hash{transactionApprovedTopic},
	})

	if err != nil {
		t.Fatalf("buildEmbed returned %v", err)
	}
	if embed != nil {
		t.Errorf("TransactionApproved built %q; the approval that reached the "+
			"threshold reports it instead", embed.Title)
	}
}

func TestUnknownAndEmptyEventsProduceNoMessage(t *testing.T) {
	c := &Multisig{address: &multisigAddress}

	for _, name := range []string{"empty topics", "unknown topic"} {
		vLog := &types.Log{}
		if name == "unknown topic" {
			vLog.Topics = []common.Hash{crypto.Keccak256Hash([]byte("SomethingElse()"))}
		}

		embed, err := c.buildEmbed(context.Background(), vLog)

		if err != nil {
			t.Fatalf("%s: buildEmbed returned %v", name, err)
		}
		if embed != nil {
			t.Errorf("%s: built %q, want no message", name, embed.Title)
		}
	}
}
