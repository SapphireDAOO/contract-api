package blockchain

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

func EventTopic(metadata *bind.MetaData, name string) common.Hash {
	parsed, err := metadata.ParseABI()
	if err != nil {
		panic(fmt.Sprintf("parse contract ABI: %v", err))
	}
	event, ok := parsed.Events[name]
	if !ok {
		panic(fmt.Sprintf("event %q not in contract ABI", name))
	}
	return event.ID
}
