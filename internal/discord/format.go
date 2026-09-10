package discord

import (
	"fmt"

	"github.com/SapphireDAOO/contract-api/internal/config"
	"github.com/ethereum/go-ethereum/common"
)

// ShortHex shortens a hex string to the 0x1234…abcd form, so an address or
// hash stays readable inside an embed.
func ShortHex(s string) string {
	if len(s) <= 12 {
		return s
	}
	return s[:6] + "…" + s[len(s)-4:]
}

// AddressLink renders an address as a markdown link to the explorer. Chains
// without an explorer configured fall back to the bare address.
func AddressLink(explorerURL string, addr common.Address) string {
	return fmt.Sprintf("[`%s`](%s)", ShortHex(addr.Hex()), config.Link(explorerURL, "/address/", addr.Hex()))
}
