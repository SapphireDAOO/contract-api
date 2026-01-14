package callback

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// get the paymnet token address and amount via event

func buildPaymentReceivedCallbackPayload(orderId, transactionURL, paymentToken string, amount *big.Int,
	transactionTimestamp int64) ([]byte, error) {
	data, _ := tokenCurrencyAndDecimals[paymentToken]
	// fetch default release time

	if amount == nil {
		return []byte{}, fmt.Errorf("%s", "invalid amount")
	}

	payload := paymentReceivedCallbackPayload{
		Currency:             data.Symbol,
		Amount:               formatTokenAmount(amount, data.Decimal),
		TransactionAmount:    formatTokenAmount(amount, data.Decimal),
		TransactionUrl:       transactionURL,
		TransactionTimestamp: transactionTimestamp,
		Releases:             parseTimestampMillis("0"),
	}

	return json.Marshal(payload)
}

func buildRefundCallbackPayload(paymentToken string, amount *big.Int,
	refundShare *big.Int, transactionURL string, transactionTimestamp int64) ([]byte, error) {

	data, _ := tokenCurrencyAndDecimals[paymentToken]

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

func buildReleaseCallbackPayload(orderId string, paymentToken string, releaseAmount *big.Int, transactionURL string,
	transactionTimestamp int64) ([]byte, error) {

	address := ""
	data, _ := tokenCurrencyAndDecimals[paymentToken]

	payload := releaseCallbackPayload{
		Currency:             data.Symbol,
		Amount:               formatTokenAmount(releaseAmount, data.Decimal),
		Address:              address,
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}
