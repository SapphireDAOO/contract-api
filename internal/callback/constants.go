package callback

const refundCallbackAction = "refundSent"
const releaseCallbackAction = "escrowReleased"
const paymentReceivedCallbackAction = "paymentReceived"

// to-do: make query direct from contract/subgraph
type TokenData struct {
	Symbol  string
	Decimal int
}

var tokenCurrencyAndDecimals = map[string]TokenData{
	"0x41A196b1fF165419A1320F029E689A41F30c70b0": {Symbol: "USDC", Decimal: 6},
	"0x8Cdaf12598d71cad44e91FB1c05d565a383e3dba": {Symbol: "wBTC", Decimal: 8},
	"0x0000000000000000000000000000000000000000": {Symbol: "ETH", Decimal: 18},
}
