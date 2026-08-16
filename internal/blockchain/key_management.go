package blockchain

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/crypto"
)

func Auth(chainId *big.Int) (*bind.TransactOpts, error) {
	raw := strings.TrimPrefix(strings.TrimSpace(os.Getenv("PASS")), "0x")
	if raw == "" {
		return nil, errors.New("PASS env var not set")
	}

	k, err := hex.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("PASS is not valid hex: %w", err)
	}

	value, err := crypto.ToECDSA(k)
	if err != nil {
		return nil, fmt.Errorf("PASS is not a valid private key: %w", err)
	}

	return bind.NewKeyedTransactor(value, chainId), nil
}
