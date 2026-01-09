package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/orgs/SapphireDAOO/contract-api/internal/query"
)

const refundCallbackAction = "refundSent"
const releaseCallbackAction = "escrowReleased"
const paymentReceivedCallbackAction = "paymentReceived"
const callbackRetryAttempts = 5
const callbackRetryBaseDelay = time.Second

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

type paymentReceivedCallbackPayload struct {
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	TransactionAmount    string `json:"transactionAmount"`
	TransactionUrl       string `json:"transactionUrl"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	Releases             int64  `json:"releases"`
}

type callbackResponse struct {
	Status  int             `json:"status"`
	Error   string          `json:"error"`
	Message json.RawMessage `json:"message"`
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

func callbackRetryDelay(attempt int) time.Duration {
	if attempt < 1 {
		return callbackRetryBaseDelay
	}
	delay := callbackRetryBaseDelay * time.Duration(1<<uint(attempt-1))
	if delay > 30*time.Second {
		return 30 * time.Second
	}
	return delay
}

func callbackStatusHint(status int) string {
	switch status {
	case http.StatusBadRequest:
		return "invalid request data"
	case http.StatusForbidden:
		return "API key rejected"
	case http.StatusGone:
		return "order not found"
	case http.StatusInternalServerError:
		return "marketplace internal error"
	default:
		return ""
	}
}

func parseCallbackResponse(body []byte) (string, string) {
	if len(body) == 0 {
		return "", ""
	}
	var response callbackResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", strings.TrimSpace(string(body))
	}
	message := ""
	if len(response.Message) > 0 {
		var msgText string
		if err := json.Unmarshal(response.Message, &msgText); err == nil {
			message = msgText
		} else {
			message = strings.TrimSpace(string(response.Message))
		}
	}
	return response.Error, message
}

func parseTimestampMillis(value string) int64 {
	if value == "" {
		return 0
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	if parsed >= 1_000_000_000_000 {
		return parsed
	}
	return parsed * 1000
}

func sendCallbackWithRetry(payload []byte, orderId, action string) {
	for attempt := 1; attempt <= callbackRetryAttempts; attempt++ {
		res, err := Cb(payload, orderId, action)
		if err != nil {
			if attempt < callbackRetryAttempts {
				log.Printf("callback attempt %d/%d failed for orderId %s action %s: %v",
					attempt, callbackRetryAttempts, orderId, action, err)
				time.Sleep(callbackRetryDelay(attempt))
				continue
			}
			log.Printf("callback failed after %d attempts for orderId %s action %s: %v",
				attempt, orderId, action, err)
			return
		}

		status := res.StatusCode
		if status == http.StatusOK {
			_, _ = io.Copy(io.Discard, res.Body)
			res.Body.Close()
			return
		}

		body, readErr := io.ReadAll(res.Body)
		res.Body.Close()
		if readErr != nil {
			log.Printf("callback response read failed for orderId %s action %s: %v", orderId, action, readErr)
			if status >= http.StatusInternalServerError && attempt < callbackRetryAttempts {
				time.Sleep(callbackRetryDelay(attempt))
				continue
			}
			return
		}

		errCode, message := parseCallbackResponse(body)
		details := fmt.Sprintf("status %d", status)
		if hint := callbackStatusHint(status); hint != "" {
			details += ", hint: " + hint
		}
		if errCode != "" {
			details += ", error: " + errCode
		}
		if message != "" {
			details += ", message: " + message
		}

		if status >= http.StatusInternalServerError {
			log.Printf("callback server error for orderId %s action %s (%s)", orderId, action, details)
			if attempt < callbackRetryAttempts {
				time.Sleep(callbackRetryDelay(attempt))
				continue
			}
			return
		}

		log.Printf("callback rejected for orderId %s action %s (%s)", orderId, action, details)
		return
	}
}

func formatTokenAmount(amount *big.Int, decimals int) string {
	if amount == nil {
		return "0"
	}
	if decimals <= 0 {
		return amount.String()
	}
	amountText := amount.Text(10)

	if len(amountText) <= decimals {
		amountText = strings.Repeat("0", decimals-len(amountText)+1) + amountText
	}

	intPart := amountText[:len(amountText)-decimals]
	fracPart := strings.TrimRight(amountText[len(amountText)-decimals:], "0")

	if fracPart == "" {
		return intPart
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

	sendCallbackWithRetry(payload, orderId, refundCallbackAction)
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

	sendCallbackWithRetry(payload, orderId, releaseCallbackAction)
}

func buildPaymentReceivedCallbackPayload(orderId string, transactionURL string,
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

	payload := paymentReceivedCallbackPayload{
		Currency:             currency,
		Amount:               formatTokenAmount(amountPaid, decimals),
		TransactionAmount:    formatTokenAmount(amountPaid, decimals),
		TransactionUrl:       transactionURL,
		TransactionTimestamp: transactionTimestamp,
		Releases:             parseTimestampMillis(invoice.ReleasedAt),
	}

	return json.Marshal(payload)
}

func SendPaymentReceivedCallback(orderId string, transactionURL string, transactionTimestamp int64) {
	payload, err := buildPaymentReceivedCallbackPayload(orderId, transactionURL, transactionTimestamp)
	if err != nil {
		log.Printf("payment received callback payload error for orderId %s: %v", orderId, err)
		return
	}

	sendCallbackWithRetry(payload, orderId, paymentReceivedCallbackAction)
}
