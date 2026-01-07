package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	HTTP    *ethclient.Client
	WS      *ethclient.Client
	ChainId *big.Int
}

func NewClient() (*Client, error) {
	httpURL := os.Getenv("TEST_NET_RPC_URL")
	wssURL := os.Getenv("TEST_NET_WSS")

	if httpURL == "" || wssURL == "" {
		return nil, errors.New("TEST_NET_RPC_URL and TEST_NET_WSS are not set")
	}

	httpClient, err := ethclient.Dial(httpURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to RPC: %v", err)
	}

	wsClient, err := ethclient.Dial(wssURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to WSS: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chainId, err := httpClient.ChainID(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chain ID: %v", err)
	}

	return &Client{
		HTTP: httpClient,
		WS:      wsClient,
		ChainId: chainId,
	}, nil
}
