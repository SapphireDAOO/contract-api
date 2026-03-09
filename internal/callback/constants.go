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
	"0x31ca89eea4dde88f8b044106117994a7fabb46b6": {Symbol: "USDC", Decimal: 6},
	"0x97e3109d772c969a9c16ecd8916a57ead5b77487": {Symbol: "wBTC", Decimal: 8},
	"0x0000000000000000000000000000000000000000": {Symbol: "ETH", Decimal: 18},
}
