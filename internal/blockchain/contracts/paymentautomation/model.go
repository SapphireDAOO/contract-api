package paymentautomation

import (
	"github.com/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/SapphireDAOO/contract-api/internal/blockchain/gen/paymentautomation"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

type PaymentAutomation struct {
	address  *common.Address
	instance *bind.BoundContract
	contract *gen.Paymentautomation
	client   *blockchain.Client
}
