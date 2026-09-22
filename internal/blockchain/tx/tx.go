package tx

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	client *blockchain.Client, intermediatedOperatorAddress, contractAddress common.Address,
	data []byte) (*Response, error) {
	msg := ethereum.CallMsg{
		From: intermediatedOperatorAddress,
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.HTTP.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	auth, err := client.Auth()

	if err != nil {
		return nil, err
	}

	tx, err := bind.Transact(instance, auth, data)

	if err != nil {
		return nil, err
	}

	slog.Info("transaction broadcast",
		"txHash", tx.Hash().Hex(), "contract", contractAddress.Hex(), "from", auth.From.Hex())

	resultChan := make(chan TransactionResult, 1)

	go func() {
		receipt, err := bind.WaitMined(ctx, client.HTTP, tx.Hash())
		resultChan <- TransactionResult{Receipt: receipt, Error: err}
	}()

	select {
	case <-ctx.Done():
		// The transaction is already broadcast; only the wait is abandoned.
		// The caller still gets a success, so without this line an unconfirmed
		// transaction would leave no trace at all.
		slog.Warn("stopped waiting for the transaction, it is still in flight",
			"txHash", tx.Hash().Hex(), "reason", ctx.Err())

		result := new(big.Int).SetBytes(result).String()
		return &Response{Result: &result}, nil

	case res := <-resultChan:
		if res.Error != nil {
			slog.Error("waiting for the transaction receipt failed",
				"txHash", tx.Hash().Hex(), "error", res.Error)
			return nil, res.Error
		}
		if res.Receipt != nil && res.Receipt.Status != types.ReceiptStatusSuccessful {
			slog.Error("transaction reverted",
				"txHash", tx.Hash().Hex(), "block", res.Receipt.BlockNumber)
		}
		return &Response{Receipt: res.Receipt}, nil
	}

}
