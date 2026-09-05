package callback

const refundCallbackAction = "refundSent"
const releaseCallbackAction = "escrowReleased"
const paymentReceivedCallbackAction = "paymentReceived"

// TokenLookup resolves a token address seen in a chain event back to the
// symbol and decimals used to render it. config.Tokens implements it, so the
// token table lives in the config file rather than being duplicated here.
type TokenLookup interface {
	ByAddress(address string) (symbol string, decimals int, ok bool)
}
