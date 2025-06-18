package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256(data string) (*common.Hash, error) {
	encoded, err := abi.Arguments{{Type: abi.Type{T: abi.StringTy}}}.Pack(data)

	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256Hash(encoded)

	fmt.Println(hash, crypto.Keccak256Hash(encoded))

	return &hash, nil
}
