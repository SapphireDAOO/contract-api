package callback

import (
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"time"
)

func (c *Client) SendRefundCallback(orderId string, paymentToken string, amount *big.Int,
	refundShare *big.Int, transactionURL string, transactionTimestamp int64) {

	payload, err := c.buildRefundCallbackPayload(paymentToken, amount,
		refundShare, transactionURL, transactionTimestamp)
	if err != nil {
		slog.Error("refund callback payload build failed", "orderId", orderId, "error", err)
		return
	}

	c.sendCallbackWithRetry(payload, orderId, refundCallbackAction)
}

func (c *Client) SendReleaseCallback(orderId, paymentToken, receiver string, releaseAmount *big.Int,
	transactionURL string, transactionTimestamp int64) {
	payload, err := c.buildReleaseCallbackPayload(paymentToken, receiver,
		releaseAmount, transactionURL, transactionTimestamp)
	if err != nil {
		slog.Error("release callback payload build failed", "orderId", orderId, "error", err)
		return
	}

	c.sendCallbackWithRetry(payload, orderId, releaseCallbackAction)
}

func (c *Client) SendPaymentReceivedCallback(orderId, transactionURL, paymentToken string, amount *big.Int, transactionTimestamp int64) {
	payload, err := c.buildPaymentReceivedCallbackPayload(transactionURL, paymentToken,
		amount, transactionTimestamp)
	if err != nil {
		slog.Error("payment received callback payload build failed", "orderId", orderId, "error", err)
		return
	}

	c.sendCallbackWithRetry(payload, orderId, paymentReceivedCallbackAction)
}

func (c *Client) sendCallbackWithRetry(payload []byte, orderId, action string) {
	for attempt := 1; attempt <= callbackRetryAttempts; attempt++ {
		res, err := c.post(payload, orderId, action)
		if err != nil {
			if attempt < callbackRetryAttempts {
				slog.Warn("callback attempt failed, retrying",
					"attempt", attempt, "attempts", callbackRetryAttempts,
					"orderId", orderId, "action", action, "error", err)
				time.Sleep(callbackRetryDelay(attempt))
				continue
			}
			slog.Error("callback failed, giving up",
				"attempts", attempt, "orderId", orderId, "action", action, "error", err)
			return
		}

		status := res.StatusCode
		if status == http.StatusOK {
			_, _ = io.Copy(io.Discard, res.Body)
			res.Body.Close()
			slog.Info("callback delivered",
				"orderId", orderId, "action", action, "status", status)
			return
		}

		body, readErr := io.ReadAll(res.Body)
		res.Body.Close()
		if readErr != nil {
			slog.Error("callback response read failed", "orderId", orderId, "action", action, "error", readErr)
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
			slog.Error("callback returned a server error", "orderId", orderId, "action", action, "details", details)
			if attempt < callbackRetryAttempts {
				time.Sleep(callbackRetryDelay(attempt))
				continue
			}
			return
		}

		slog.Error("callback rejected", "orderId", orderId, "action", action, "details", details)
		return
	}
}
