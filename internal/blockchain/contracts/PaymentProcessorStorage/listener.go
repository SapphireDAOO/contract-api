package paymentprocessorstorage

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/orgs/SapphireDAOO/contract-api/internal/discord"
)

const explorerURL = "https://sepolia.basescan.org"

var (
	pausedTopic          = crypto.Keccak256Hash([]byte("Paused(address)"))
	unpausedTopic        = crypto.Keccak256Hash([]byte("Unpaused(address)"))
	emergencyPausedTopic = crypto.Keccak256Hash([]byte("EmergencyPaused(address,uint256)"))
)

func (c *PaymentProcessorStorage) subscribeLogs(ctx context.Context, query ethereum.FilterQuery, logs chan types.Log) ethereum.Subscription {
	for {
		sub, err := c.client.WS.SubscribeFilterLogs(ctx, query, logs)
		if err == nil {
			return sub
		}
		log.Printf("Failed to subscribe to Payment Processor Storage logs: %v", err)

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(5 * time.Second):
		}
	}
}

// ListenToPauseEvents subscribes to the storage contract's Paused, Unpaused
// and EmergencyPaused events and posts a Discord notification for each one.
func (c *PaymentProcessorStorage) ListenToPauseEvents(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		log.Printf("payment processor storage listener disabled: client or contract address not initialized")
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

	log.Println("Listening for Payment Processor Storage pause events...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Payment Processor Storage listener stopping")
			return

		case err := <-sub.Err():
			log.Printf("Payment Processor Storage subscription error: %v", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs)
			if sub == nil {
				return
			}

		case vLog := <-logs:
			embed, err := c.buildEmbed(&vLog)
			if err != nil {
				log.Printf("Failed to parse Payment Processor Storage event: %v", err)
				continue
			}
			if embed == nil {
				continue
			}

			log.Printf("Payment Processor Storage event: %s (%s)", embed.Title, vLog.TxHash.Hex())
			go discord.SendEmbed(*embed)
		}
	}
}

func (c *PaymentProcessorStorage) buildEmbed(vLog *types.Log) (*discord.Embed, error) {
	if len(vLog.Topics) == 0 {
		return nil, nil
	}

	base := discord.Embed{
		URL:    explorerURL + "/tx/" + vLog.TxHash.Hex(),
		Footer: &discord.Footer{Text: "Payment Processor Storage " + shortHex(c.address.Hex()) + " • Base Sepolia"},
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
				addressLink(event.Account)))

	case unpausedTopic:
		event, err := c.contract.UnpackUnpausedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "▶️ Contract unpaused"
		base.Color = discord.ColorGreen
		lines = append(lines,
			fmt.Sprintf("%s unpaused the Payment Processor Storage. Payment processing has resumed.",
				addressLink(event.Account)))

	case emergencyPausedTopic:
		event, err := c.contract.UnpackEmergencyPausedEvent(vLog)
		if err != nil {
			return nil, err
		}
		base.Title = "🚨 Emergency pause activated"
		base.Color = discord.ColorRed
		lines = append(lines,
			fmt.Sprintf("%s triggered an **emergency pause** on the Payment Processor Storage.",
				addressLink(event.Account)))
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

func addressLink(addr common.Address) string {
	return fmt.Sprintf("[`%s`](%s/address/%s)", shortHex(addr.Hex()), explorerURL, addr.Hex())
}

// shortHex shortens a hex string to the 0x1234…abcd form.
func shortHex(s string) string {
	if len(s) <= 12 {
		return s
	}
	return s[:6] + "…" + s[len(s)-4:]
}
