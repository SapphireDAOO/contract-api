package query

import (
	"encoding/json"
	"testing"
)

func TestSmartInvoiceMarshalJSONFlattensRelations(t *testing.T) {
	invoice := SmartInvoice{
		InvoiceId:    "42",
		CreatedAt:    "1700000000",
		Price:        "1000000",
		State:        "PAID",
		AmountPaid:   "1000000",
		PaidAt:       "1700000100",
		Balance:      "0",
		ReleasedAt:   "1700000200",
		Escrow:       "3600",
		Buyer:        &Buyer{ID: "0xbuyer"},
		Seller:       &Seller{ID: "0xseller"},
		PaymentToken: &PaymentToken{ID: "0xtoken", Name: "USDC", Decimal: "6"},
		MetaInvoice:  &MetaInvoice{InvoiceID: "7"},
	}

	raw, err := json.Marshal(invoice)
	if err != nil {
		t.Fatalf("Marshal returned %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("the marshalled invoice is not valid JSON: %v", err)
	}

	want := map[string]string{
		"invoiceId": "42",
		"state":     "PAID",
		"buyer":     "0xbuyer",
		"seller":    "0xseller",

		"paymentToken": "USDC",
		"metaInvoice":  "7",
	}
	for key, wantValue := range want {
		value, ok := got[key].(string)
		if !ok {
			t.Errorf("%s is %T (%v), want a string", key, got[key], got[key])
			continue
		}
		if value != wantValue {
			t.Errorf("%s = %q, want %q", key, value, wantValue)
		}
	}
}

func TestSmartInvoiceMarshalJSONWithNilRelations(t *testing.T) {
	invoice := SmartInvoice{InvoiceId: "42", State: "CREATED"}

	raw, err := json.Marshal(invoice)
	if err != nil {
		t.Fatalf("Marshal returned %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("the marshalled invoice is not valid JSON: %v", err)
	}

	for _, key := range []string{"buyer", "seller", "paymentToken", "metaInvoice"} {
		value, ok := got[key].(string)
		if !ok {
			t.Errorf("%s is %T (%v), want an empty string", key, got[key], got[key])
			continue
		}
		if value != "" {
			t.Errorf("%s = %q, want an empty string", key, value)
		}
	}
}

func TestSmartInvoiceMarshalJSONDoesNotRecurse(t *testing.T) {
	done := make(chan []byte, 1)

	go func() {
		raw, err := json.Marshal(SmartInvoice{InvoiceId: "1"})
		if err != nil {
			done <- nil
			return
		}
		done <- raw
	}()

	raw := <-done
	if len(raw) == 0 {
		t.Fatal("Marshal produced nothing")
	}
}

func TestSmartInvoiceMarshalJSONThroughAContainer(t *testing.T) {
	data := Data{
		Address: "0xuser",
		User: User{
			IssuedInvoices: []SmartInvoice{
				{InvoiceId: "1", Buyer: &Buyer{ID: "0xbuyer"}},
			},
			ReceivedInvoices: []SmartInvoice{
				{InvoiceId: "2", Seller: &Seller{ID: "0xseller"}},
			},
		},
	}

	raw, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal returned %v", err)
	}

	var got struct {
		Address string `json:"address"`
		User    struct {
			IssuedInvoices []struct {
				InvoiceId string `json:"invoiceId"`
				Buyer     string `json:"buyer"`
			} `json:"issuedInvoices"`
			ReceivedInvoices []struct {
				InvoiceId string `json:"invoiceId"`
				Seller    string `json:"seller"`
			} `json:"receivedInvoices"`
		} `json:"user"`
	}
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("the marshalled data is not valid JSON: %v", err)
	}

	if got.Address != "0xuser" {
		t.Errorf("address = %q, want 0xuser", got.Address)
	}
	if len(got.User.IssuedInvoices) != 1 || got.User.IssuedInvoices[0].Buyer != "0xbuyer" {
		t.Errorf("issued invoices = %+v, want a flattened buyer", got.User.IssuedInvoices)
	}
	if len(got.User.ReceivedInvoices) != 1 || got.User.ReceivedInvoices[0].Seller != "0xseller" {
		t.Errorf("received invoices = %+v, want a flattened seller", got.User.ReceivedInvoices)
	}
}

func TestSmartInvoiceUnmarshalsFromASubgraphResponse(t *testing.T) {
	const body = `{
	  "data": {
	    "user": {
	      "issuedInvoices": [
	        {
	          "invoiceId": "42",
	          "amountPaid": "1000000",
	          "price": "1000000",
	          "state": "PAID",
	          "createdAt": "1700000000",
	          "paidAt": "1700000100",
	          "releasedAt": "",
	          "buyer": {"id": "0xbuyer"},
	          "metaInvoice": {"invoiceId": "7"}
	        }
	      ],
	      "receivedInvoices": [],
	      "metaInvoices": []
	    }
	  }
	}`

	var response Response
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("Unmarshal returned %v", err)
	}

	issued := response.Data.User.IssuedInvoices
	if len(issued) != 1 {
		t.Fatalf("got %d issued invoices, want 1", len(issued))
	}
	if issued[0].InvoiceId != "42" {
		t.Errorf("invoiceId = %q, want 42", issued[0].InvoiceId)
	}
	if issued[0].Buyer == nil || issued[0].Buyer.ID != "0xbuyer" {
		t.Errorf("buyer = %+v, want 0xbuyer", issued[0].Buyer)
	}
	if issued[0].MetaInvoice == nil || issued[0].MetaInvoice.InvoiceID != "7" {
		t.Errorf("metaInvoice = %+v, want 7", issued[0].MetaInvoice)
	}

	if issued[0].Seller != nil {
		t.Errorf("seller = %+v, want nil", issued[0].Seller)
	}
}
