package notes

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/orgs/SapphireDAOO/contract-api/internal/blockchain"
	gen "github.com/orgs/SapphireDAOO/contract-api/internal/blockchain/gen/Notes"
	"github.com/orgs/SapphireDAOO/contract-api/internal/utils"
)

func NewNotes(client *blockchain.Client) *Notes {
	address := common.HexToAddress(utils.NOTES_ADDRESS)
	contract := gen.NewNotescontract()
	instance := contract.Instance(client.HTTP, address)

	return &Notes{
		address:  &address,
		instance: instance,
		contract: contract,
		client:   client,
	}
}

func (c *Notes) CreateNote(
	ctx context.Context,
	invoiceId *big.Int,
	author common.Address,
	encryptedContent []byte,
	share bool,
) (*common.Hash, error) {
	if c == nil || c.instance == nil || c.address == nil {
		return nil, errors.New("notes contract is not initialized")
	}

	return c.send(ctx, c.contract.PackCreateNote(invoiceId, author, encryptedContent, share))
}

func (c *Notes) SetOpened(
	ctx context.Context,
	invoiceId *big.Int,
	account common.Address,
	noteId *big.Int,
) (*common.Hash, error) {
	if c == nil || c.instance == nil || c.address == nil {
		return nil, errors.New("notes contract is not initialized")
	}

	return c.send(ctx, c.contract.PackSetOpened(invoiceId, account, noteId))
}

func (c *Notes) GetNote(ctx context.Context, invoiceId, noteId *big.Int) (*Note, error) {
	if c == nil || c.instance == nil {
		return nil, errors.New("notes contract is not initialized")
	}

	data := c.contract.PackGetNote(invoiceId, noteId)

	note, err := bind.Call(c.instance, &bind.CallOpts{Context: ctx}, data, c.contract.UnpackGetNote)
	if err != nil {
		return nil, err
	}

	return &Note{
		Author:  note.Author,
		Share:   note.Share,
		Content: note.Content,
		Opened:  note.OpenedStatus,
		Version: note.Version,
	}, nil
}

func (c *Notes) GetNoteCount(ctx context.Context, invoiceId *big.Int) (*big.Int, error) {
	if c == nil || c.instance == nil {
		return nil, errors.New("notes contract is not initialized")
	}

	data := c.contract.PackGetNoteCount(invoiceId)
	return bind.Call(c.instance, &bind.CallOpts{Context: ctx}, data, c.contract.UnpackGetNoteCount)
}

func (c *Notes) send(ctx context.Context, data []byte) (*common.Hash, error) {
	auth, err := blockchain.Auth(c.client.ChainId)
	if err != nil {
		return nil, err
	}

	if _, err := c.client.HTTP.CallContract(ctx, ethereum.CallMsg{
		From: auth.From,
		To:   c.address,
		Data: data,
	}, nil); err != nil {
		return nil, err
	}

	tx, err := bind.Transact(c.instance, auth, data)
	if err != nil {
		return nil, err
	}

	hash := tx.Hash()
	return &hash, nil
}
