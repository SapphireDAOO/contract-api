package notes

import (
	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/notes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type Notes struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Notes
	client   *blockchain.Client
}

// Note is a single note as stored by the Notes contract.
type Note struct {
	Author  common.Address
	Share   bool
	Content []byte
	Opened  bool
	Version uint8
}
