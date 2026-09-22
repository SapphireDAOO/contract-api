package paymentprocessorstorage

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/discord"
)

const (
	pauseCheckInterval = time.Minute
	pauseCallTimeout   = 10 * time.Second
)

func (c *PaymentProcessorStorage) WatchEmergencyPause(ctx context.Context) {
	if c == nil || c.client == nil || c.instance == nil || c.address == nil {
		slog.Warn("emergency pause watcher disabled",
			"reason", "client or contract address not initialized")
		return
	}

	ticker := time.NewTicker(pauseCheckInterval)
	defer ticker.Stop()

	slog.Info("watching for emergency pause expiry",
		"contract", c.address.Hex(), "interval", pauseCheckInterval)

	var episode pauseEpisode
	episode.observe(c.pauseState(ctx))

	for {
		select {
		case <-ctx.Done():
			slog.Info("emergency pause watcher stopping")
			return

		case <-ticker.C:
			if ended, at := episode.observe(c.pauseState(ctx)); ended {
				c.reportPauseEnded(at)
			}
		}
	}
}

type pauseEpisode struct {
	paused bool
	expiry *big.Int
}

func (e *pauseEpisode) observe(paused bool, expiry *big.Int) (bool, *big.Int) {
	if paused && expiry != nil && expiry.Sign() > 0 {
		e.expiry = expiry
	}

	ended := e.paused && !paused
	at := e.expiry

	if ended {
		e.expiry = nil
	}
	e.paused = paused

	return ended, at
}

func (c *PaymentProcessorStorage) pauseState(ctx context.Context) (bool, *big.Int) {
	callCtx, cancel := context.WithTimeout(ctx, pauseCallTimeout)
	defer cancel()

	paused, expiry, err := c.PauseState(callCtx)
	if err != nil {
		slog.Error("pause state read failed", "error", err)
		return true, nil
	}

	return paused, expiry
}

func (c *PaymentProcessorStorage) reportPauseEnded(expiry *big.Int) {
	if expiry == nil || expiry.Sign() <= 0 {
		slog.Info("pause ended with no known expiry; leaving it to the Unpaused event")
		return
	}

	if time.Now().Unix() < expiry.Int64() {
		slog.Info("pause lifted before it expired, reported by the Unpaused event",
			"expiry", expiry.Int64())
		return
	}

	slog.Info("emergency pause elapsed", "expiry", expiry.Int64())

	c.notifier.SendEmbed(discord.Embed{
		Title: "⏱️ Emergency pause elapsed",
		Color: discord.ColorGreen,
		URL:   c.link("/address/", c.address.Hex()),
		Description: fmt.Sprintf(
			"The emergency pause ran its full term and ended <t:%d:R>. "+
				"Payment processing has resumed automatically — nobody unpaused it.\n\n"+
				"[View on Basescan](%s)",
			expiry.Int64(), c.link("/address/", c.address.Hex())),
		Footer: &discord.Footer{
			Text: "Payment Processor Storage " + discord.ShortHex(c.address.Hex()),
		},
	})
}
