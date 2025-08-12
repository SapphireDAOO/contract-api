package blockchain

import (
	"context"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	eth     *ethclient.Client
	chainId *big.Int
}

func NewClient() *Client {
	rpcURL := os.Getenv("TEST_NET_RPC_URL")

	conn, err := ethclient.Dial(rpcURL)

	if err != nil {
		log.Fatalf("Failed to connect to blockchain: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chainId, err := conn.ChainID(ctx)

	if err != nil {
		log.Fatalf("failed to retrieve chain ID: %v", err)
	}

	return &Client{eth: conn, chainId: chainId}
}
