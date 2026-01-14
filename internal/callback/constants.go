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
	"0x01C524f6ef6F58627499542F9f1C1d461a246f4e": TokenData{Symbol: "USDC", Decimal: 6},
	"0x3da0578e48842b6bc04Ec1CDCc99B473270253fD": TokenData{Symbol: "wBTC", Decimal: 8},
	"0x0000000000000000000000000000000000000000": TokenData{Symbol: "ETH", Decimal: 18},
}
