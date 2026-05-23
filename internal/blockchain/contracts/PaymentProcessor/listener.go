package paymentprocesor

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/orgs/SapphireDAOO/contract-api/internal/callback"
)

const txURL = "https://sepolia.basescan.org/tx/"

func (c *PaymentProcessor) subscribeLogs(ctx context.Context, query ethereum.FilterQuery, logs chan types.Log, label string) ethereum.Subscription {
	for {
		sub, err := c.client.WS.SubscribeFilterLogs(ctx, query, logs)
		if err == nil {
			return sub
		}
		log.Printf("Failed to subscribe to %s logs: %v", label, err)

		select {
		case <-ctx.Done():
			return nil
		case <-time.After(5 * time.Second):
		}
	}
}

func (c *PaymentProcessor) ListenToPaymentReceivedEvent(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		log.Printf("payment listener disabled: client or contract address not initialized")
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

	log.Println("Listening for InvoicePaid events...")

	for {
		select {
		case <-ctx.Done():
			log.Println("InvoicePaid listener stopping")
			return

		case err := <-sub.Err():
			log.Printf("InvoicePaid subscription error: %v", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs, "InvoicePaid")
			if sub == nil {
				return
			}

		case vLog := <-logs:
			event, err := c.contract.UnpackInvoicePaidEvent(&vLog)
			if err != nil {
				log.Printf("Failed to parse InvoicePaid event: %v", err)
				continue
			}

			log.Printf("InvoicePaid Event:\n")
			log.Printf("  OrderId: %s\n", event.InvoiceId.String())
			log.Printf("  Amount: %s\n", event.Amount.String())

			transactionTimestamp := time.Now().UTC().UnixMilli()
			if c.client.HTTP != nil {
				headerCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				header, err := c.client.HTTP.HeaderByHash(headerCtx, vLog.BlockHash)
				cancel()
				if err != nil {
					log.Printf("Failed to fetch block header for InvoicePaid event: %v", err)
				} else {
					transactionTimestamp = int64(header.Time) * 1000
				}
			}

			transactionURL := txURL + vLog.TxHash.Hex()
			go callback.
				SendPaymentReceivedCallback(event.InvoiceId.String(), transactionURL, event.PaymentToken.Hex(),
					event.Amount, transactionTimestamp)
		}
	}
}

func (c *PaymentProcessor) ListenToReleaseEvent(ctx context.Context) {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		log.Printf("release listener disabled: client or contract address not initialized")
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

	log.Println("Listening for PaymentReleased events...")

	for {
		select {
		case <-ctx.Done():
			log.Println("PaymentReleased listener stopping")
			return

		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)
			sub.Unsubscribe()
			sub = c.subscribeLogs(ctx, query, logs, "PaymentReleased")
			if sub == nil {
				return
			}

		case vLog := <-logs:
			event, err := c.contract.UnpackPaymentReleasedEvent(&vLog)
			if err != nil {
				log.Printf("Failed to parse PaymentReleased event: %v", err)
				continue
			}

			log.Printf("PaymentReleased Event:\n")
			log.Printf("  OrderId: %s\n", event.InvoiceId.String())
			log.Printf("  SellerAmount: %s\n", event.SellerAmount.String())

			transactionTimestamp := time.Now().UTC().UnixMilli()
			if c.client.HTTP != nil {
				headerCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				header, err := c.client.HTTP.HeaderByHash(headerCtx, vLog.BlockHash)
				cancel()
				if err != nil {
					log.Printf("Failed to fetch block header for release event: %v", err)
				} else {
					transactionTimestamp = int64(header.Time) * 1000
				}
			}

			transactionURL := txURL + vLog.TxHash.Hex()
			go callback.
				SendReleaseCallback(event.InvoiceId.String(), event.Currency.Hex(),
					event.Receiver.Hex(), event.SellerAmount, transactionURL, transactionTimestamp)
		}
	}
}
