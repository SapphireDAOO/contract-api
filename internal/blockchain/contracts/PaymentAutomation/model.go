package paymentautomation

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/PaymentAutomation"
)

type PaymentAutomation struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Paymentautomation
	client   *blockchain.Client
}
