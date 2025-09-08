package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(orderId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"orderId": orderId,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
