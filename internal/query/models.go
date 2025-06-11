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
	OwnedSmartInvoices []SmartInvoice `json:"ownedSmartInvoices"`
	PaidSmartInvoices  []SmartInvoice `json:"paidSmartInvoices"`
	MetaInvoices       []any          `json:"metaInvoices"`
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

	newSmartInvoice := struct {
		AmountPaid  *string      `json:"amountPaid"`
		CreatedAt   string       `json:"createdAt"`
		InvoiceId   string       `json:"invoiceId"`
		PaidAt      *string      `json:"paidAt"`
		Price       string       `json:"price"`
		State       string       `json:"state"`
		ReleasedAt  *string      `json:"releasedAt"`
		Seller      string       `json:"seller"`
		MetaInvoice *MetaInvoice `json:"metaInvoice"`
	}{
		AmountPaid:  s.AmountPaid,
		CreatedAt:   s.CreatedAt,
		InvoiceId:   s.InvoiceId,
		PaidAt:      s.PaidAt,
		Price:       s.Price,
		State:       s.State,
		ReleasedAt:  s.ReleasedAt,
		Seller:      s.Seller.ID,
		MetaInvoice: s.MetaInvoice,
	}

	return json.Marshal(&newSmartInvoice)

}
