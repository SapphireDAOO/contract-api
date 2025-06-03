package blockchain

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	rpc *ethclient.Client
}

func NewClient() *Client {
	rpcURL := os.Getenv("TEST_NET_RPC_URL")

	conn, err := ethclient.Dial(rpcURL)

	if err != nil {
		log.Fatalf("Failed to connect to blockchain: %v", err)
	}

	return &Client{rpc: conn}
}
