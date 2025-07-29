package blockchain

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Contract) getDisputeResolution(action MarketplaceAction) uint8 {
	if action == DismissDispute {
		value, _ := bind.Call(
			c.instance, nil, c.paymentProcessor.PackDISPUTEDISMISSED(), c.paymentProcessor.UnpackDISPUTEDISMISSED)
		return value
	}

	if action == SettleDispute {
		value, _ := bind.Call(
			c.instance, nil, c.paymentProcessor.PackDISPUTESETTLED(), c.paymentProcessor.UnpackDISPUTESETTLED)
		return value
	}

	return 0
}

type TransactionResult struct {
	Receipt *types.Receipt
	Error   error
}

type Response struct {
	receipt *types.Receipt
	result  *string
}

func (c *Contract) simulateAndBroadcast(ctx context.Context, data []byte) (*Response, error) {
	result, err := c.simulateTransaction(ctx, data)
	if err != nil {
		return nil, err
	}

	// broadcast

	auth, err := auth(c.client.chainId)

	if err != nil {
		return nil, err
	}

	tx, err := bind.Transact(c.instance, auth, data)

	if err != nil {
		return nil, err
	}

	resultChan := make(chan TransactionResult, 1)

	go func() {
		receipt, err := bind.WaitMined(ctx, c.client.eth, tx.Hash())
		resultChan <- TransactionResult{Receipt: receipt, Error: err}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout occured while sending transaction to be mined", ctx.Err())
		result := "0x" + common.Bytes2Hex(result)
		return &Response{result: &result}, nil

	case res := <-resultChan:
		if res.Error != nil {
			return nil, res.Error
		}
		return &Response{receipt: res.Receipt}, nil
	}

}

func (c *Contract) simulateTransaction(ctx context.Context, data []byte) ([]byte, error) {
	msg := ethereum.CallMsg{
		From: common.HexToAddress("0x0f447989b14A3f0bbf08808020Ec1a6DE0b8cbC4"),
		To:   c.address,
		Data: data,
	}

	return c.client.eth.CallContract(ctx, msg, nil)
}
