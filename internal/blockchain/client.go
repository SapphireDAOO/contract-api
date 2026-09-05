package blockchain

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	HTTP    *ethclient.Client
	WS      *ethclient.Client
	ChainId *big.Int

	// signer is parsed once here rather than on every transaction. It is never
	// logged or returned.
	signer *ecdsa.PrivateKey
}

func NewClient(rpc config.RPC, signerKey string) (*Client, error) {
	signer, err := parseSignerKey(signerKey)
	if err != nil {
		return nil, err
	}

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
		signer:  signer,
	}, nil
}

// parseSignerKey decodes the configured key. Its error messages never include
// the key material.
func parseSignerKey(key string) (*ecdsa.PrivateKey, error) {
	raw := strings.TrimPrefix(strings.TrimSpace(key), "0x")
	if raw == "" {
		return nil, fmt.Errorf("signer key is not configured")
	}

	decoded, err := hex.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("signer key is not valid hex")
	}

	signer, err := crypto.ToECDSA(decoded)
	if err != nil {
		return nil, fmt.Errorf("signer key is not a valid private key")
	}
	return signer, nil
}
