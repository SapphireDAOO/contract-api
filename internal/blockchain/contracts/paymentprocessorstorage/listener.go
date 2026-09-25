package paymentprocessorstorage

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/paymentprocessorstorage"
	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/SapphireDAOO/contract-api/internal/discord"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	pausedTopic          = blockchain.EventTopic(&gen.PaymentprocessorstorageMetaData, gen.PaymentprocessorstoragePausedEventName)
	unpausedTopic        = blockchain.EventTopic(&gen.PaymentprocessorstorageMetaData, gen.PaymentprocessorstorageUnpausedEventName)
	emergencyPausedTopic = blockchain.EventTopic(&gen.PaymentprocessorstorageMetaData, gen.PaymentprocessorstorageEmergencyPausedEventName)
)

func (c *PaymentProcessorStorage) subscribeLogs(ctx context.Context, query ethereum.FilterQuery,
	logs chan types.Log) ethereum.Subscription {
	return c.client.SubscribeLogs(ctx, query, logs, "Payment Processor Storage")
}
func (c *PaymentProcessorStorage) ListenToPauseEvents(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		slog.Warn("payment processor storage listener disabled", "reason", "client or contract address not initialized")
		return
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*c.address},
		Topics:    [][]common.Hash{{pausedTopic, unpausedTopic, emergencyPausedTopic}},
	}

	logs := make(chan types.Log)
	sub := c.subscribeLogs(ctx, query, logs)
	if sub == nil {
		return
	}
	defer sub.Unsubscribe()

	slog.Info("listening for payment processor storage pause events", "contract", c.address.Hex())

	for {
		select {
		case <-ctx.Done():
			slog.Info("payment processor storage listener stopping")
			return

		case err := <-sub.Err():
			slog.Error("payment processor storage subscription failed, resubscribing", "error", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs)
			if sub == nil {
				return
			}

		case vLog := <-logs:
			embed, err := c.buildEmbed(&vLog)
			if err != nil {
				slog.Error("payment processor storage event parse failed", "txHash", vLog.TxHash.Hex(), "error", err)
				continue
			}
			if embed == nil {
				continue
			}

			slog.Info("payment processor storage event", "event", embed.Title, "txHash", vLog.TxHash.Hex())
			go c.notifier.SendEmbed(*embed)
		}
	}
}

func (c *PaymentProcessorStorage) buildEmbed(vLog *types.Log) (*discord.Embed, error) {
	if len(vLog.Topics) == 0 {
		return nil, nil
	}

	base := discord.Embed{
		URL:    c.link("/tx/", vLog.TxHash.Hex()),
		Footer: &discord.Footer{Text: "Payment Processor Storage " + discord.ShortHex(c.address.Hex()) + " • Base Sepolia"},
	}
	var lines []string

	switch vLog.Topics[0] {
	case pausedTopic:
		event, err := c.contract.UnpackPausedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "⏸️ Contract paused"
		base.Color = discord.ColorRed
		lines = append(lines,
			fmt.Sprintf("%s paused the Payment Processor Storage. Payment processing is halted until it is unpaused.",
				c.addressLink(event.Account)))

	case unpausedTopic:
		event, err := c.contract.UnpackUnpausedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "▶️ Contract unpaused"
		base.Color = discord.ColorGreen
		lines = append(lines,
			fmt.Sprintf("%s unpaused the Payment Processor Storage. Payment processing has resumed.",
				c.addressLink(event.Account)))

	case emergencyPausedTopic:
		event, err := c.contract.UnpackEmergencyPausedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "🚨 Emergency pause activated"
		base.Color = discord.ColorRed
		lines = append(lines,
			fmt.Sprintf("%s triggered an **emergency pause** on the Payment Processor Storage.",
				c.addressLink(event.Account)))
		if event.Expiry != nil {
			lines = append(lines,
				fmt.Sprintf("It elapses <t:%d:F> (<t:%d:R>).", event.Expiry.Int64(), event.Expiry.Int64()))
		}

	default:
		return nil, nil
	}

	lines = append(lines, fmt.Sprintf("[View on Basescan](%s)", base.URL))
	// Blank lines between sections keep the message easy to scan.
	base.Description = strings.Join(lines, "\n\n")
	return &base, nil
}

func (c *PaymentProcessorStorage) addressLink(addr common.Address) string {
	return discord.AddressLink(c.explorerURL, addr)
}
func (c *PaymentProcessorStorage) link(path, value string) string {
	return config.Link(c.explorerURL, path, value)
}
