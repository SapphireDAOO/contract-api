package callback

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// get the paymnet token address and amount via event

func buildPaymentReceivedCallbackPayload(transactionURL, paymentToken string, amount *big.Int,
	transactionTimestamp int64) ([]byte, error) {
	if amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	data, ok := tokenData(paymentToken)
	if !ok {
		return nil, fmt.Errorf("unsupported payment token %s", paymentToken)
	}

	currentDefaultReleaseTime := 10 * time.Minute
	releaseAt := time.Now().Add(currentDefaultReleaseTime).UnixMilli()

	amountText := formatTokenAmount(amount, data.Decimal)
	payload := paymentReceivedCallbackPayload{
		Currency:             data.Symbol,
		Amount:               amountText,
		TransactionAmount:    amountText,
		TransactionUrl:       transactionURL,
		TransactionTimestamp: transactionTimestamp,
		Releases:             releaseAt,
	}

	return json.Marshal(payload)
}

func buildRefundCallbackPayload(paymentToken string, amount *big.Int,
	refundShare *big.Int, transactionURL string, transactionTimestamp int64) ([]byte, error) {

	if amount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	data, ok := tokenData(paymentToken)
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
		Currency:             data.Symbol,
		Amount:               formatTokenAmount(refundAmount, data.Decimal),
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}

func buildReleaseCallbackPayload(paymentToken, receiver string, releaseAmount *big.Int, transactionURL string,
	transactionTimestamp int64) ([]byte, error) {

	if releaseAmount == nil {
		return nil, fmt.Errorf("invalid amount")
	}

	data, ok := tokenData(paymentToken)
	if !ok {
		return nil, fmt.Errorf("unsupported payment token %s", paymentToken)
	}

	payload := releaseCallbackPayload{
		Currency:             data.Symbol,
		Amount:               formatTokenAmount(releaseAmount, data.Decimal),
		Address:              receiver,
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}
