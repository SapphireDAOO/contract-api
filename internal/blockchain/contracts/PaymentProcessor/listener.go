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

const RELEASE_TOPIC_HASH = "0x8ea2131e86229753e4a36a9ffc579af1b38fdada1aefe3e09a44cf2eab25befe"
const txURL = "https://sepolia.etherscan.io/tx/"

func (c *PaymentProcessor) subscribeLogs(query ethereum.FilterQuery, logs chan types.Log, label string) ethereum.Subscription {
	for {
		sub, err := c.client.WS.SubscribeFilterLogs(context.Background(), query, logs)
		if err == nil {
			return sub
		}
		log.Printf("Failed to subscribe to %s logs: %v", label, err)
		time.Sleep(5 * time.Second)
	}
}

func (c *PaymentProcessor) ListenToPaymentReceivedEvent() {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		log.Printf("payment listener disabled: client or contract address not initialized")
		return
	}

	invoicePaidTopic := crypto.Keccak256Hash([]byte("InvoicePaid(uint216,address,address,uint256)"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*c.address},
		Topics:    [][]common.Hash{{invoicePaidTopic}},
	}

	logs := make(chan types.Log)
	sub := c.subscribeLogs(query, logs, "InvoicePaid")

	log.Println("Listening for InvoicePaid events...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("InvoicePaid subscription error: %v", err)

			sub.Unsubscribe()
			sub = c.subscribeLogs(query, logs, "InvoicePaid")

		case vLog := <-logs:
			event, err := c.contract.UnpackInvoicePaidEvent(&vLog)
			if err != nil {
				log.Printf("Failed to parse InvoicePaid event: %v", err)
				continue
			}

			log.Printf("InvoicePaid Event:\n")
			log.Printf("  OrderId: %s\n", event.OrderId.String())
			log.Printf("  Amount: %s\n", event.Amount.String())

			transactionTimestamp := time.Now().UTC().UnixMilli()
			if c.client.HTTP != nil {
				header, err := c.client.HTTP.HeaderByHash(context.Background(), vLog.BlockHash)
				if err != nil {
					log.Printf("Failed to fetch block header for InvoicePaid event: %v", err)
				} else {
					transactionTimestamp = int64(header.Time) * 1000
				}
			}

			transactionURL := txURL + vLog.TxHash.Hex()
			go callback.
				SendPaymentReceivedCallback(event.OrderId.String(), transactionURL, event.PaymentToken.Hex(),
					event.Amount, transactionTimestamp)
		}
	}

}

func (c *PaymentProcessor) ListenToReleaseEvent() {
	if c == nil || c.client == nil || c.client.WS == nil || c.address == nil {
		log.Printf("release listener disabled: client or contract address not initialized")
		return
	}

	paymentReleasedTopic := common.HexToHash(RELEASE_TOPIC_HASH)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{*c.address},
		Topics:    [][]common.Hash{{paymentReleasedTopic}},
	}

	logs := make(chan types.Log)
	sub := c.subscribeLogs(query, logs, "PaymentReleased")

	log.Println("Listening for PaymentReleased events...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)

			sub.Unsubscribe()
			sub = c.subscribeLogs(query, logs, "PaymentReleased")
		case vLog := <-logs:
			event, err := c.contract.UnpackPaymentReleasedEvent(&vLog)
			if err != nil {
				log.Printf("Failed to parse PaymentReleased event: %v", err)
				continue
			}

			log.Printf("PaymentReleased Event:\n")
			log.Printf("  OrderId: %s\n", event.OrderId.String())
			log.Printf("  SellerAmount: %s\n", event.SellerAmount.String())

			transactionTimestamp := time.Now().UTC().UnixMilli()
			if c.client.HTTP != nil {
				header, err := c.client.HTTP.HeaderByHash(context.Background(), vLog.BlockHash)
				if err != nil {
					log.Printf("Failed to fetch block header for release event: %v", err)
				} else {
					transactionTimestamp = int64(header.Time) * 1000
				}
			}

			transactionURL := txURL + vLog.TxHash.Hex()
			go callback.
				SendReleaseCallback(event.OrderId.String(),
					event.SellerAmount, transactionURL, transactionTimestamp)
		}
	}
}
