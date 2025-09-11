package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

func OrderIDToUint216(orderId string) string {
	stringType, _ := abi.NewType("string", "", nil)
	arguments := abi.Arguments{{Type: stringType}}
	encoded, err := arguments.Pack(orderId)
	if err != nil {
		panic(err)
	}

	hash := crypto.Keccak256(encoded)

	hashInt := new(big.Int).SetBytes(hash)

	mask := new(big.Int).Lsh(big.NewInt(1), 216)
	mask.Sub(mask, big.NewInt(1))

	result := new(big.Int).And(hashInt, mask)

	return result.String()
}
