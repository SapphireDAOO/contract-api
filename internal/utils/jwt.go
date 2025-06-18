package utils

import (
	"encoding/hex"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(invoiceKey *common.Hash) (string, error) {
	val := "0x" + hex.EncodeToString(invoiceKey[:])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"invoiceKey": val,
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
