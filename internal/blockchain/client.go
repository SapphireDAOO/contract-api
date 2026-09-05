package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	HTTP    *ethclient.Client
	WS      *ethclient.Client
	ChainId *big.Int
}

func NewClient(rpc config.RPC) (*Client, error) {
	httpClient, err := ethclient.Dial(rpc.HTTP)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %w", err)
	}

	wsClient, err := ethclient.Dial(rpc.WS)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WSS: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), rpc.DialTimeout)
	defer cancel()

	chainId, err := httpClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chain ID: %w", err)
	}

	return &Client{
		HTTP:    httpClient,
		WS:      wsClient,
		ChainId: chainId,
	}, nil
}
