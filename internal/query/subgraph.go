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

func GetUserInvoiceData(address string, first, skip int) (*Data, error) {
	if address == "" {
		return nil, fmt.Errorf("address cannot be empty")
	}

	payload := map[string]any{
		"query": userQuery,
		"variables": map[string]any{
			"address": address,
			"first":   first,
			"skip":    skip,
		},
	}

	body, err := handleRequest(payload)

	if err != nil {
		return nil, err
	}

	var gqlResp Response
	if err := json.Unmarshal(body, &gqlResp); err != nil {
		return nil, err
	}

	userData := gqlResp.Data
	userData.Address = address

	return &userData, nil
}

func GetInvoiceData(id string) (*SmartInvoice, error) {
	if id == "" {
		return nil, fmt.Errorf("address cannot be empty")
	}

	payload := map[string]any{
		"query": invoiceQuery,
		"variables": map[string]any{
			"id": id,
		},
	}

	body, err := handleRequest(payload)

	if err != nil {
		return nil, err
	}

	var resp struct {
		Data struct {
			SmartInvoice SmartInvoice `json:"smartInvoice"`
		}
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return &resp.Data.SmartInvoice, err
}

func handleRequest(payload map[string]any) ([]byte, error) {
	jsonData, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", os.Getenv("END_POINT"), bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("non-200 status code: %d, body: %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}
