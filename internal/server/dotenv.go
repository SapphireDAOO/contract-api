package server

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	if _, ok := os.LookupEnv("PRODUCTION"); ok {
		return nil
	}
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}
