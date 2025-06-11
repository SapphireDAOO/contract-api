package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetInvoice(address string) (*Data, error) {
	if address == "" {
		return nil, fmt.Errorf("address cannot be empty")
	}

	payload := map[string]any{
		"query": invoiceQuery,
		"variables": map[string]any{
			"address": address,
		},
	}

	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", os.Getenv("END_POINT"), bytes.NewBuffer(jsonData))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Println("Raw response:", string(body))

	var gqlResp Response
	if err := json.Unmarshal(body, &gqlResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Println(gqlResp)

	userData := gqlResp.Data
	userData.Address = address

	return &userData, nil
}
