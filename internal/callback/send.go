package callback

import (
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"time"
)

func SendRefundCallback(orderId string, paymentToken string, amount *big.Int,
	refundShare *big.Int, transactionURL string, transactionTimestamp int64) {
	payload, err := buildRefundCallbackPayload(paymentToken, amount,
		refundShare, transactionURL, transactionTimestamp)
	if err != nil {
		log.Printf("refund callback payload error for orderId %s: %v", orderId, err)
		return
	}

	sendCallbackWithRetry(payload, orderId, refundCallbackAction)
}

func SendReleaseCallback(orderId string, releaseAmount *big.Int, transactionURL string,
	transactionTimestamp int64) {
	paymentToken := ""
	payload, err := buildReleaseCallbackPayload(orderId, paymentToken,
		releaseAmount, transactionURL, transactionTimestamp)
	if err != nil {
		log.Printf("release callback payload error for orderId %s: %v", orderId, err)
		return
	}

	sendCallbackWithRetry(payload, orderId, releaseCallbackAction)
}

func SendPaymentReceivedCallback(orderId, transactionURL, paymentToken string, amount *big.Int, transactionTimestamp int64) {
	payload, err := buildPaymentReceivedCallbackPayload(orderId, transactionURL, paymentToken,
		amount, transactionTimestamp)
	if err != nil {
		log.Printf("payment received callback payload error for orderId %s: %v", orderId, err)
		return
	}

	sendCallbackWithRetry(payload, orderId, paymentReceivedCallbackAction)
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
