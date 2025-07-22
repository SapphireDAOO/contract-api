package query

import (
	"encoding/json"
)

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	Address string `json:"address"`
	User    User   `json:"user"`
}

type User struct {
	IssuedInvoices   []SmartInvoice `json:"issuedInvoices"`
	ReceivedInvoices []SmartInvoice `json:"receivedInvoices"`
	MetaInvoices     []any          `json:"metaInvoices"`
}

type SmartInvoice struct {
	InvoiceId    string        `json:"invoiceId"`
	CreatedAt    string        `json:"createdAt"`
	Price        string        `json:"price"`
	State        string        `json:"state"`
	AmountPaid   string        `json:"amountPaid"`
	PaidAt       string        `json:"paidAt"`
	Balance      string        `json:"balance"`
	PaymentToken *PaymentToken `json:"paymentToken"`
	ReleasedAt   string        `json:"releasedAt"`
	Buyer        *Buyer        `json:"buyer"`
	Seller       *Seller       `json:"seller"`
	MetaInvoice  *MetaInvoice  `json:"metaInvoice"`
	Escrow       string        `json:"escrow"`
}

type PaymentToken struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Decimal string `json:"decimal"`
}

type Buyer struct {
	ID string `json:"id"`
}

type Seller struct {
	ID string `json:"id"`
}

type MetaInvoice struct {
	InvoiceID string `json:"invoiceId"`
}

func (s SmartInvoice) MarshalJSON() ([]byte, error) {
	type Alias SmartInvoice

	var (
		metaId       string
		paymentToken string
		buyerId      string
		sellerId     string
	)

	if s.MetaInvoice != nil {
		metaId = s.MetaInvoice.InvoiceID
	}

	if s.PaymentToken != nil {
		paymentToken = s.PaymentToken.Name
	}

	if s.Seller != nil {
		sellerId = s.Seller.ID
	}

	if s.Buyer != nil {
		buyerId = s.Buyer.ID
	}

	return json.Marshal(&struct {
		*Alias
		Buyer        string `json:"buyer"`
		Seller       string `json:"seller"`
		MetaInvoice  string `json:"metaInvoice"`
		PaymentToken string `json:"paymentToken"`
	}{
		Alias:        (*Alias)(&s),
		Buyer:        buyerId,
		Seller:       sellerId,
		PaymentToken: paymentToken,
		MetaInvoice:  metaId,
	})
}
