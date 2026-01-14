package erc20

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	erc20 "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/ERC20"
)

type Erc20 struct {
	address  *common.Address
	contract *erc20.Erc20
	client   *blockchain.Client
}

func NewErc20(client *blockchain.Client, contractAddress string) *Erc20 {
	address := common.HexToAddress(contractAddress)
	contract := erc20.NewErc20()

	return &Erc20{
		address:  &address,
		contract: contract,
		client:   client,
	}
}

type TokenData struct {
	Symbol  string
	Decimal int
}

// pass payment token, use other info

func (c *Erc20) SymbolAndDecimal(ctx context.Context) (*string, error) {
	data := c.contract.PackSymbol()

	out, err := c.client.HTTP.CallContract(ctx, ethereum.CallMsg{
		From: *c.address,
		Data: data,
	}, nil)

	if err != nil {
		return nil, err
	}

	symbol, err := c.contract.UnpackSymbol(out)

	if err != nil {
		return nil, err
	}

	return &symbol, nil
}
