package blockchain

import (
	"encoding/hex"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/crypto"
)

func Auth(chainId *big.Int) (*bind.TransactOpts, error) {

	k, _ := hex.DecodeString(os.Getenv("PASS"))

	value, err := crypto.ToECDSA(k)

	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(value, chainId), nil

}
