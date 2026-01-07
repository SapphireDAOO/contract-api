package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
)

const refundCallbackAction = "refund"
const releaseCallbackAction = "escrowReleased"

type refundCallbackPayload struct {
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	TransactionUrl       string `json:"transactionUrl"`
}

type releaseCallbackPayload struct {
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	Address              string `json:"address"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	TransactionUrl       string `json:"transactionUrl"`
}

func Cb(payload []byte, orderId, action string) (*http.Response, error) {
	baseUrl := os.Getenv("URL")
	if baseUrl == "" {
		return nil, fmt.Errorf("URL env var not set")
	}

	key := os.Getenv("API_KEY")
	if key == "" {
		return nil, fmt.Errorf("API_KEY env var not set")
	}

	url := fmt.Sprintf(baseUrl, orderId, "/", action)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("APIKey", key)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func formatTokenAmount(amount *big.Int, decimals int) string {
	if amount == nil {
		return "0"
	}
	if decimals <= 0 {
		return amount.String()
	}
	amountText := amount.Text(10)
	negative := strings.HasPrefix(amountText, "-")
	if negative {
		amountText = strings.TrimPrefix(amountText, "-")
	}
	if len(amountText) <= decimals {
		amountText = strings.Repeat("0", decimals-len(amountText)+1) + amountText
	}
	intPart := amountText[:len(amountText)-decimals]
	fracPart := strings.TrimRight(amountText[len(amountText)-decimals:], "0")
	if fracPart == "" {
		if negative {
			return "-" + intPart
		}
		return intPart
	}
	if negative {
		return "-" + intPart + "." + fracPart
	}
	return intPart + "." + fracPart
}

func buildRefundCallbackPayload(orderId string, refundShare *big.Int, transactionURL string,
	transactionTimestamp int64) ([]byte, error) {
	invoice, err := query.GetInvoiceData(orderId)
	if err != nil {
		return nil, err
	}

	var (
		currency string
		decimals int
	)

	if invoice.PaymentToken != nil {
		currency = invoice.PaymentToken.Name
		if invoice.PaymentToken.Decimal != "" {
			if parsed, err := strconv.Atoi(invoice.PaymentToken.Decimal); err == nil {
				decimals = parsed
			}
		}
		if currency == "" {
			currency = invoice.PaymentToken.ID
		}
	}

	amountPaid := new(big.Int)
	if invoice.AmountPaid != "" {
		if _, ok := amountPaid.SetString(invoice.AmountPaid, 10); !ok {
			amountPaid.SetInt64(0)
		}
	}

	refundAmount := new(big.Int)
	if refundShare != nil {
		refundAmount.Mul(amountPaid, refundShare)
		refundAmount.Div(refundAmount, big.NewInt(10000))
	} else {
		refundAmount.Set(amountPaid)
	}

	payload := refundCallbackPayload{
		Currency:             currency,
		Amount:               formatTokenAmount(refundAmount, decimals),
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}

func SendRefundCallback(orderId string, refundShare *big.Int, transactionURL string, transactionTimestamp int64) {
	payload, err := buildRefundCallbackPayload(orderId, refundShare, transactionURL, transactionTimestamp)
	if err != nil {
		log.Printf("refund callback payload error for orderId %s: %v", orderId, err)
		return
	}

	res, err := Cb(payload, orderId, refundCallbackAction)
	if err != nil {
		log.Printf("refund callback error for orderId %s: %v", orderId, err)
		return
	}
	if res != nil && res.Body != nil {
		res.Body.Close()
	}
}

func buildReleaseCallbackPayload(orderId string, transactionURL string, transactionTimestamp int64) ([]byte, error) {
	invoice, err := query.GetInvoiceData(orderId)
	if err != nil {
		return nil, err
	}

	var (
		currency string
		decimals int
		address  string
	)

	if invoice.PaymentToken != nil {
		currency = invoice.PaymentToken.Name
		if invoice.PaymentToken.Decimal != "" {
			if parsed, err := strconv.Atoi(invoice.PaymentToken.Decimal); err == nil {
				decimals = parsed
			}
		}
		if currency == "" {
			currency = invoice.PaymentToken.ID
		}
	}

	if invoice.Seller != nil {
		address = invoice.Seller.ID
	}

	amountPaid := new(big.Int)
	if invoice.AmountPaid != "" {
		if _, ok := amountPaid.SetString(invoice.AmountPaid, 10); !ok {
			amountPaid.SetInt64(0)
		}
	}

	payload := releaseCallbackPayload{
		Currency:             currency,
		Amount:               formatTokenAmount(amountPaid, decimals),
		Address:              address,
		TransactionTimestamp: transactionTimestamp,
		TransactionUrl:       transactionURL,
	}

	return json.Marshal(payload)
}

func SendReleaseCallback(orderId string, transactionURL string, transactionTimestamp int64) {
	payload, err := buildReleaseCallbackPayload(orderId, transactionURL, transactionTimestamp)
	if err != nil {
		log.Printf("release callback payload error for orderId %s: %v", orderId, err)
		return
	}

	res, err := Cb(payload, orderId, releaseCallbackAction)
	if err != nil {
		log.Printf("release callback error for orderId %s: %v", orderId, err)
		return
	}
	if res != nil && res.Body != nil {
		res.Body.Close()
	}
}
