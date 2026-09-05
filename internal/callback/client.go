package callback

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Client posts payment events to the marketplace callback endpoint. It is
// built once at startup from the configured URL and API key.
type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		http:    &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) post(payload []byte, orderId, action string) (*http.Response, error) {
	if c == nil || c.baseURL == "" {
		return nil, fmt.Errorf("callback URL is not configured")
	}
	if c.apiKey == "" {
		return nil, fmt.Errorf("callback API key is not configured")
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

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildCallbackURL(baseURL, orderId, action string) (string, error) {
	if strings.TrimSpace(orderId) == "" {
		return "", fmt.Errorf("orderId is required for callback URL")
	}
	if strings.TrimSpace(action) == "" {
		return "", fmt.Errorf("action is required for callback URL")
	}

	if strings.Contains(baseURL, "%s") || strings.Contains(baseURL, "%[") {
		return fmt.Sprintf(baseURL, orderId, action), nil
	}

	trimmed := strings.TrimRight(baseURL, "/")
	return trimmed + "/" + orderId + "/" + action, nil
}
