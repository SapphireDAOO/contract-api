package callback

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const (
	callbackRetryAttempts  = 5
	callbackRetryBaseDelay = time.Second
)

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
