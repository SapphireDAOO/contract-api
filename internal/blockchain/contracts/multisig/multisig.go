package multisig

import (
	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/multisig"
	"github.com/ethereum/go-ethereum/common"
)

// Peers are the other contracts whose admin calls the multisig proposes, and
// whose calldata it therefore has to decode.
type Peers struct {
	PaymentProcessor        common.Address
	SimplePaymentProcessor  common.Address
	PaymentProcessorStorage common.Address
}

func NewMultisig(client *blockchain.Client, address common.Address, peers Peers) *Multisig {
	contract := gen.NewMultisig()
	instance := contract.Instance(client.HTTP, address)

	return &Multisig{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
		known:    buildKnownContracts(address, peers),
	}
}
