package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func Error(w http.ResponseWriter, statusCode int, err error, msg string) {
	errorMessage := fmt.Sprintf(`{"error": "%s", "reason": "%s"}`, msg, strings.ReplaceAll(err.Error(), `"`, `'`))
	http.Error(w, errorMessage, statusCode)
}
