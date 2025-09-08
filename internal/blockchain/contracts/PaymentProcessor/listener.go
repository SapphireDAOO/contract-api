package paymentprocesor

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	advancedprocessor "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/AdvancedPaymentProcessor"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

const RELEASE_TOPIC_HASH = ""

func (c *PaymentProcessor) ListenToReleaseEvent() {
	contractAddress := common.HexToAddress(utils.PAYMENT_PROCESSOR_ADDRESS)
	if contractAddress == (common.Address{}) {
		log.Fatalf("ADVANCED_PROCESSOR_CONTRACT_ADDRESS not set in environment")
	}

	contract := advancedprocessor.NewAdvancedprocessor()

	paymentReleasedTopic := common.HexToHash(RELEASE_TOPIC_HASH)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{paymentReleasedTopic}},
	}

	logs := make(chan types.Log)
	sub, err := c.client.WS.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	defer sub.Unsubscribe()

	log.Println("Listening for PaymentReleased events...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)

			time.Sleep(5 * time.Second)
			sub, err = c.client.WS.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				log.Printf("Failed to resubscribe: %v", err)
				continue
			}
		case vLog := <-logs:
			event, err := contract.UnpackPaymentReleasedEvent(&vLog)
			if err != nil {
				log.Printf("Failed to parse PaymentReleased event: %v", err)
				continue
			}

			log.Printf("PaymentReleased Event:\n")
			log.Printf("  OrderId: %s\n", event.OrderId.String())
			log.Printf("  SellerAmount: %s\n", event.SellerAmount.String())
		}
	}
}
