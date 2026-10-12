package callback

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Client posts payment events to the marketplace callback endpoint. It is
// built once at startup from the configured URL and API key.
type Client struct {
	baseURL string
	apiKey  string
	tokens  TokenLookup
	http    *http.Client

	production bool
}

func NewClient(baseURL, apiKey string, tokens TokenLookup) *Client {
	production, _ := strconv.ParseBool(strings.TrimSpace(os.Getenv("PRODUCTION")))
	return &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		tokens:     tokens,
		http:       &http.Client{Timeout: 30 * time.Second},
		production: production,
	}
}

func (c *Client) post(payload []byte, orderId, action string) (*http.Response, error) {
	if c == nil || c.baseURL == "" {
		return nil, ErrURLNotConfigured
	}
	if c.apiKey == "" {
		return nil, ErrAPIKeyNotConfigured
	}

	url, err := buildCallbackURL(c.baseURL, orderId, action)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("APIKey", c.apiKey)

	if !c.production {
		slog.Info("callback not posted outside production", "action", action)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildCallbackURL(baseURL, orderId, action string) (string, error) {
	if strings.TrimSpace(orderId) == "" {
		return "", ErrOrderIDRequired
	}
	if strings.TrimSpace(action) == "" {
		return "", ErrActionRequired
	}

	if strings.Contains(baseURL, "%s") || strings.Contains(baseURL, "%[") {
		return fmt.Sprintf(baseURL, orderId, action), nil
	}

	trimmed := strings.TrimRight(baseURL, "/")
	return trimmed + "/" + orderId + "/" + action, nil
}
