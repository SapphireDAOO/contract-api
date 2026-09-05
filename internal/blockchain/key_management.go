package blockchain

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
)

// Auth returns transaction options signed with the configured key. Which key
// that is comes from the selected network's signerKey, so a local run signs as
// the local chain's account without touching the deployed networks' key.
func (c *Client) Auth() (*bind.TransactOpts, error) {
	if c == nil || c.signer == nil {
		return nil, errors.New("signer key is not configured")
	}
	return bind.NewKeyedTransactor(c.signer, c.ChainId), nil
}
