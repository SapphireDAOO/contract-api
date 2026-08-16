package callback

import "strings"

const refundCallbackAction = "refundSent"
const releaseCallbackAction = "escrowReleased"
const paymentReceivedCallbackAction = "paymentReceived"

// to-do: make query direct from contract/subgraph
type TokenData struct {
	Symbol  string
	Decimal int
}

// keys are lowercase hex addresses; use tokenData for lookups.
var tokenCurrencyAndDecimals = map[string]TokenData{
	"0x41a196b1ff165419a1320f029e689a41f30c70b0": {Symbol: "USDC", Decimal: 6},
	"0x8cdaf12598d71cad44e91fb1c05d565a383e3dba": {Symbol: "wBTC", Decimal: 8},
	"0x0000000000000000000000000000000000000000": {Symbol: "ETH", Decimal: 18},
}

func tokenData(address string) (TokenData, bool) {
	data, ok := tokenCurrencyAndDecimals[strings.ToLower(strings.TrimSpace(address))]
	return data, ok
}
