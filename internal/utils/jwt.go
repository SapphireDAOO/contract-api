package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const tokenTTL = 24 * time.Hour

func GenerateToken(orderId string) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", errors.New("SECRET_KEY not set")
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"orderId": orderId,
		"iat":     now.Unix(),
		"exp":     now.Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(secret))
}
