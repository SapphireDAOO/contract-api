package contracts

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
)

type TransactionResult struct {
	Receipt *types.Receipt
	Error   error
}

type Response struct {
	Receipt *types.Receipt
	Result  *string
}

func SimulateAndBroadcast(ctx context.Context, instance *bind.BoundContract,
	client *blockchain.Client, marketplaceAddress, contractAddress common.Address,
	data []byte) (*Response, error) {
	msg := ethereum.CallMsg{
		From: marketplaceAddress,
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.HTTP.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	auth, err := blockchain.Auth(client.ChainId)

	if err != nil {
		return nil, err
	}

	tx, err := bind.Transact(instance, auth, data)

	if err != nil {
		return nil, err
	}

	resultChan := make(chan TransactionResult, 1)

	go func() {
		receipt, err := bind.WaitMined(ctx, client.HTTP, tx.Hash())
		resultChan <- TransactionResult{Receipt: receipt, Error: err}
	}()

	select {
	case <-ctx.Done():
		result := new(big.Int).SetBytes(result).String()
		return &Response{Result: &result}, nil

	case res := <-resultChan:
		if res.Error != nil {
			return nil, res.Error
		}
		return &Response{Receipt: res.Receipt}, nil
	}

}
