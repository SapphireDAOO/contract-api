package callback

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"
)

// get the paymnet token address and amount via event

func buildPaymentReceivedCallbackPayload(transactionURL, paymentToken string, amount *big.Int,
	transactionTimestamp int64) ([]byte, error) {
	data, _ := tokenCurrencyAndDecimals[paymentToken]

	if amount == nil {
		return []byte{}, fmt.Errorf("%s", "invalid amount")
	}

	currentDefaultReleaseTime := 10 * time.Minute
	releaseAt := time.Now().Add(currentDefaultReleaseTime).UnixMilli() / 1000

	payload := paymentReceivedCallbackPayload{
		Currency:             data.Symbol,
		Amount:               formatTokenAmount(amount, data.Decimal),
		TransactionAmount:    formatTokenAmount(amount, data.Decimal),
		TransactionUrl:       transactionURL,
		TransactionTimestamp: transactionTimestamp,
		Releases:             parseTimestampMillis(fmt.Sprint(releaseAt)),
	}

	return json.Marshal(payload)
}

func buildRefundCallbackPayload(paymentToken string, amount *big.Int,
	refundShare *big.Int, transactionURL string, transactionTimestamp int64) ([]byte, error) {

	data, _ := tokenCurrencyAndDecimals[paymentToken]

	if amount == nil {
		return []byte{}, fmt.Errorf("%s", "invalid amount")
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

	data, ok := tokenCurrencyAndDecimals[strings.ToLower(paymentToken)]
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
