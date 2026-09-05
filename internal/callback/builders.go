package callback

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// get the paymnet token address and amount via event

func (c *Client) buildPaymentReceivedCallbackPayload(transactionURL, paymentToken string, amount *big.Int,
	transactionTimestamp int64) ([]byte, error) {
	if amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	symbol, decimals, ok := c.tokens.ByAddress(paymentToken)
	if !ok {
		return nil, fmt.Errorf("unsupported payment token %s", paymentToken)
	}

	currentDefaultReleaseTime := 10 * time.Minute
	releaseAt := time.Now().Add(currentDefaultReleaseTime).UnixMilli()

	amountText := formatTokenAmount(amount, decimals)
	payload := paymentReceivedCallbackPayload{
		Currency:             symbol,
		Amount:               amountText,
		TransactionAmount:    amountText,
		TransactionUrl:       transactionURL,
		TransactionTimestamp: transactionTimestamp,
		Releases:             releaseAt,
	}

	return json.Marshal(payload)
}

func (c *Client) buildRefundCallbackPayload(paymentToken string, amount *big.Int,
	refundShare *big.Int, transactionURL string, transactionTimestamp int64) ([]byte, error) {

	if amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	symbol, decimals, ok := c.tokens.ByAddress(paymentToken)
	if !ok {
		return nil, fmt.Errorf("unsupported payment token %s", paymentToken)
	}

	refundAmount := new(big.Int)
	if refundShare != nil {
		refundAmount.Mul(amount, refundShare)
		refundAmount.Div(refundAmount, big.NewInt(10000))
	} else {
		refundAmount.Set(amount)
	}

	payload := refundCallbackPayload{
		Currency:             symbol,
		Amount:               formatTokenAmount(refundAmount, decimals),
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}

func (c *Client) buildReleaseCallbackPayload(paymentToken, receiver string, releaseAmount *big.Int, transactionURL string,
	transactionTimestamp int64) ([]byte, error) {

	if releaseAmount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	symbol, decimals, ok := c.tokens.ByAddress(paymentToken)
	if !ok {
		return nil, fmt.Errorf("unsupported payment token %s", paymentToken)
	}

	payload := releaseCallbackPayload{
		Currency:             symbol,
		Amount:               formatTokenAmount(releaseAmount, decimals),
		Address:              receiver,
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}
