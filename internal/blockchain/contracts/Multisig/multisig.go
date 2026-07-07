package multisig

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/Multisig"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func NewMultisig(client *blockchain.Client) *Multisig {
	address := common.HexToAddress(utils.MULTISIG_ADDRESS)
	contract := gen.NewMultisig()
	instance := contract.Instance(client.HTTP, address)

	return &Multisig{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}
