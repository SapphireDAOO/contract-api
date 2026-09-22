package intermediatedpaymentprocessor

import (
	"context"
	"log/slog"
	"time"

	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c *PaymentProcessor) subscribeLogs(ctx context.Context, query ethereum.FilterQuery,
	logs chan types.Log, label string) ethereum.Subscription {
	return c.client.SubscribeLogs(ctx, query, logs, label)
}
func (c *PaymentProcessor) ListenToPaymentReceivedEvent(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		slog.Warn("invoice paid listener disabled", "reason", "client or contract address not initialized")
		return
	}

	invoicePaidTopic := crypto.Keccak256Hash([]byte("InvoicePaid(uint216,address,address,uint256,uint40)"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*c.address},
		Topics:    [][]common.Hash{{invoicePaidTopic}},
	}

	logs := make(chan types.Log)
	sub := c.subscribeLogs(ctx, query, logs, "InvoicePaid")
	if sub == nil {
		return
	}
	defer sub.Unsubscribe()

	slog.Info("listening for InvoicePaid events", "contract", c.address.Hex())

	for {
		select {
		case <-ctx.Done():
			slog.Info("InvoicePaid listener stopping")
			return

		case err := <-sub.Err():
			slog.Error("InvoicePaid subscription failed, resubscribing", "error", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs, "InvoicePaid")
			if sub == nil {
				return
			}

		case vLog := <-logs:
			event, err := c.contract.UnpackInvoicePaidEvent(&vLog)
			if err != nil {
				slog.Error("InvoicePaid event parse failed", "txHash", vLog.TxHash.Hex(), "error", err)
				continue
			}

			slog.Info("InvoicePaid",
				"invoiceId", event.InvoiceId.String(),
				"amount", event.Amount.String(),
				"txHash", vLog.TxHash.Hex())

			transactionTimestamp := time.Now().UTC().UnixMilli()
			if c.client.HTTP != nil {
				headerCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				header, err := c.client.HTTP.HeaderByHash(headerCtx, vLog.BlockHash)
				cancel()
				if err != nil {
					slog.Error("block header fetch failed for InvoicePaid", "block", vLog.BlockNumber, "error", err)
				} else {
					transactionTimestamp = int64(header.Time) * 1000
				}
			}

			transactionURL := c.txURL(vLog.TxHash.Hex())
			go c.callbacks.SendPaymentReceivedCallback(event.InvoiceId.String(), transactionURL, event.PaymentToken.Hex(),
				event.Amount, transactionTimestamp)
		}
	}
}

func (c *PaymentProcessor) ListenToReleaseEvent(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		slog.Warn("payment released listener disabled", "reason", "client or contract address not initialized")
		return
	}

	paymentReleasedTopic := crypto.Keccak256Hash([]byte("PaymentReleased(uint216,address,address,uint256)"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*c.address},
		Topics:    [][]common.Hash{{paymentReleasedTopic}},
	}

	logs := make(chan types.Log)
	sub := c.subscribeLogs(ctx, query, logs, "PaymentReleased")
	if sub == nil {
		return
	}
	defer sub.Unsubscribe()

	slog.Info("listening for PaymentReleased events", "contract", c.address.Hex())

	for {
		select {
		case <-ctx.Done():
			slog.Info("PaymentReleased listener stopping")
			return

		case err := <-sub.Err():
			slog.Error("PaymentReleased subscription failed, resubscribing", "error", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs, "PaymentReleased")
			if sub == nil {
				return
			}

		case vLog := <-logs:
			event, err := c.contract.UnpackPaymentReleasedEvent(&vLog)
			if err != nil {
				slog.Error("PaymentReleased event parse failed", "txHash", vLog.TxHash.Hex(), "error", err)
				continue
			}

			slog.Info("PaymentReleased",
				"invoiceId", event.InvoiceId.String(),
				"sellerAmount", event.SellerAmount.String(),
				"txHash", vLog.TxHash.Hex())

			transactionTimestamp := time.Now().UTC().UnixMilli()
			if c.client.HTTP != nil {
				headerCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				header, err := c.client.HTTP.HeaderByHash(headerCtx, vLog.BlockHash)
				cancel()
				if err != nil {
					slog.Error("block header fetch failed for PaymentReleased", "block", vLog.BlockNumber, "error", err)
				} else {
					transactionTimestamp = int64(header.Time) * 1000
				}
			}

			transactionURL := c.txURL(vLog.TxHash.Hex())
			go c.callbacks.SendReleaseCallback(event.InvoiceId.String(), event.Currency.Hex(),
				event.Receiver.Hex(), event.SellerAmount, transactionURL, transactionTimestamp)
		}
	}
}

// txURL links a transaction on the configured explorer. Chains without one
// (a local node) fall back to the bare hash.
func (c *PaymentProcessor) txURL(txHash string) string {
	return config.Link(c.explorerURL, "/tx/", txHash)
}
