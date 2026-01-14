package callback

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func Cb(payload []byte, orderId, action string) (*http.Response, error) {
	baseUrl := os.Getenv("URL")
	if baseUrl == "" {
		return nil, fmt.Errorf("URL env var not set")
	}

	key := os.Getenv("API_KEY")
	if key == "" {
		return nil, fmt.Errorf("API_KEY env var not set")
	}

	url, err := buildCallbackURL(baseUrl, orderId, action)
	if err != nil {
		return nil, err
	}

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
