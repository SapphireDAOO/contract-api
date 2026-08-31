package notes

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/Notes"
)

type Notes struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Notescontract
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
