package paymentautomation

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	defaultPollInterval = time.Minute

	checkTimeout = 15 * time.Second

	processTimeout = 5 * time.Minute
)

func (c *PaymentAutomation) PollDueTasks(ctx context.Context) {
	if c == nil || c.client == nil || c.client.HTTP == nil || c.address == nil {
		slog.Warn("payment automation poller disabled", "reason", "client or contract address not initialized")
		return
	}

	interval := pollInterval()
	slog.Info("polling payment automation for due tasks", "interval", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		c.runCycle(ctx)

		select {
		case <-ctx.Done():
			slog.Info("payment automation poller stopping")
			return
		case <-ticker.C:
		}
	}
}

func (c *PaymentAutomation) runCycle(ctx context.Context) {
	checkCtx, cancel := context.WithTimeout(ctx, checkTimeout)
	defer cancel()

	due, err := c.HasDueTasks(checkCtx)
	if err != nil {
		// A cancelled context means shutdown, not a contract problem.
		if ctx.Err() == nil {
			slog.Error("due task check failed", "error", err)
		}
		return
	}
	if !due {
		return
	}

	slog.Info("due tasks found, sending processDueTasks")

	processCtx, cancelProcess := context.WithTimeout(ctx, processTimeout)
	defer cancelProcess()

	receipt, err := c.ProcessDueTasks(processCtx)
	if err != nil {
		if ctx.Err() == nil {
			slog.Error("processDueTasks failed", "error", err)
		}
		return
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		slog.Error("processDueTasks reverted",
			"txHash", receipt.TxHash.Hex(), "block", receipt.BlockNumber)
		return
	}

	slog.Info("processed due tasks",
		"txHash", receipt.TxHash.Hex(), "block", receipt.BlockNumber, "gasUsed", receipt.GasUsed)
}

func pollInterval() time.Duration {
	raw := os.Getenv("AUTOMATION_POLL_INTERVAL")
	if raw == "" {
		return defaultPollInterval
	}

	interval, err := time.ParseDuration(raw)
	if err != nil {
		slog.Warn("invalid AUTOMATION_POLL_INTERVAL, using default", "value", raw, "default", defaultPollInterval, "error", err)
		return defaultPollInterval
	}
	if interval <= 0 {
		slog.Warn("AUTOMATION_POLL_INTERVAL must be positive, using default", "value", raw, "default", defaultPollInterval)
		return defaultPollInterval
	}
	return interval
}
