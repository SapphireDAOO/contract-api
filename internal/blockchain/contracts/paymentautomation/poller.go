package paymentautomation

import (
	"context"
	"log"
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
		log.Println("payment automation poller disabled: client or contract address not initialized")
		return
	}

	interval := pollInterval()
	log.Printf("Polling Payment Automation for due tasks every %s...", interval)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		c.runCycle(ctx)

		select {
		case <-ctx.Done():
			log.Println("Payment Automation poller stopping")
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
			log.Printf("Failed to check for due tasks: %v", err)
		}
		return
	}
	if !due {
		return
	}

	log.Println("Due tasks found; sending processDueTasks transaction")

	processCtx, cancelProcess := context.WithTimeout(ctx, processTimeout)
	defer cancelProcess()

	receipt, err := c.ProcessDueTasks(processCtx)
	if err != nil {
		if ctx.Err() == nil {
			log.Printf("Failed to process due tasks: %v", err)
		}
		return
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Printf("processDueTasks reverted in tx %s (block %d)",
			receipt.TxHash.Hex(), receipt.BlockNumber)
		return
	}

	log.Printf("Processed due tasks in tx %s (block %d, gas used %d)",
		receipt.TxHash.Hex(), receipt.BlockNumber, receipt.GasUsed)
}

func pollInterval() time.Duration {
	raw := os.Getenv("AUTOMATION_POLL_INTERVAL")
	if raw == "" {
		return defaultPollInterval
	}

	interval, err := time.ParseDuration(raw)
	if err != nil {
		log.Printf("Invalid AUTOMATION_POLL_INTERVAL %q, using %s: %v", raw, defaultPollInterval, err)
		return defaultPollInterval
	}
	if interval <= 0 {
		log.Printf("AUTOMATION_POLL_INTERVAL must be positive, using %s", defaultPollInterval)
		return defaultPollInterval
	}
	return interval
}
