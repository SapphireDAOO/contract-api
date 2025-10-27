package utils

import (
	"bytes"
	"net/http"
	"time"
)

func Cb(payload []byte) (*http.Response, error) {
	client := http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", "")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
