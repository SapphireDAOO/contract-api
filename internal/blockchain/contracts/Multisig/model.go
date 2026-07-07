package multisig

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/Multisig"
)

type Multisig struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Multisig
	client   *blockchain.Client
}
