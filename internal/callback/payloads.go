package callback

import "encoding/json"

type refundCallbackPayload struct {
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	TransactionUrl       string `json:"transactionUrl"`
}

type releaseCallbackPayload struct {
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	Address              string `json:"address"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	TransactionUrl       string `json:"transactionUrl"`
}

type paymentReceivedCallbackPayload struct {
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	TransactionAmount    string `json:"transactionAmount"`
	TransactionUrl       string `json:"transactionUrl"`
	TransactionTimestamp int64  `json:"transactionTimestamp"`
	Releases             int64  `json:"releases"`
}

type callbackResponse struct {
	Status  int             `json:"status"`
	Error   string          `json:"error"`
	Message json.RawMessage `json:"message"`
}
