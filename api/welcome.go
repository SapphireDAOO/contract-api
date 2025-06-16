package handler

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	dataJson := map[string]any{
		"data": map[string]any{
			"message": "Hello World!",
		},
	}

	err := json.NewEncoder(w).Encode(dataJson)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
