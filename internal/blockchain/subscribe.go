package blockchain

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

// resubscribeDelay is how long to wait before retrying a failed subscription.
const resubscribeDelay = 5 * time.Second

func (c *Client) SubscribeLogs(ctx context.Context, query ethereum.FilterQuery,
	logs chan types.Log, label string) ethereum.Subscription {
	for {
		sub, err := c.WS.SubscribeFilterLogs(ctx, query, logs)
		if err == nil {
			return sub
		}
		log.Printf("Failed to subscribe to %s logs: %v", label, err)

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(resubscribeDelay):
		}
	}
}
