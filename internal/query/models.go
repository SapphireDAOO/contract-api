package query

import "encoding/json"

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
	AmountPaid  *string      `json:"amountPaid"`
	CreatedAt   string       `json:"createdAt"`
	InvoiceId   string       `json:"invoiceId"`
	PaidAt      *string      `json:"paidAt"`
	Price       string       `json:"price"`
	State       string       `json:"state"`
	ReleasedAt  *string      `json:"releasedAt"`
	Buyer       Buyer        `json:"buyer,omitzero"`
	Seller      Seller       `json:"seller,omitzero"`
	MetaInvoice *MetaInvoice `json:"metaInvoice"`
}

type Buyer struct {
	ID string `json:"id,omitempty"`
}

type Seller struct {
	ID string `json:"id,omitempty"`
}

type MetaInvoice struct {
	InvoiceID string `json:"invoiceId"`
}

func (s SmartInvoice) MarshalJSON() ([]byte, error) {
	type Alias SmartInvoice

	var metaID string

	if s.MetaInvoice != nil {
		metaID = s.MetaInvoice.InvoiceID
	}
	return json.Marshal(&struct {
		*Alias
		Buyer       string `json:"buyer"`
		Seller      string `json:"seller"`
		MetaInvoice string `json:"metaInvoice"`
	}{
		Alias:       (*Alias)(&s),
		Buyer:       s.Buyer.ID,
		Seller:      s.Seller.ID,
		MetaInvoice: metaID,
	})

}
