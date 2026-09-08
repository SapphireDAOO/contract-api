package query

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type capturedRequest struct {
	method      string
	contentType string
	query       string
	variables   map[string]any
}

func newSubgraph(t *testing.T, status int, body string) (*Client, *capturedRequest) {
	t.Helper()

	captured := &capturedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured.method = r.Method
		captured.contentType = r.Header.Get("Content-Type")

		raw, _ := io.ReadAll(r.Body)
		var payload struct {
			Query     string         `json:"query"`
			Variables map[string]any `json:"variables"`
		}
		json.Unmarshal(raw, &payload)
		captured.query = payload.Query
		captured.variables = payload.Variables

		w.WriteHeader(status)
		io.WriteString(w, body)
	}))
	t.Cleanup(server.Close)

	return NewClient(server.URL), captured
}

func TestGetInvoiceData(t *testing.T) {
	const body = `{"data":{"smartInvoice":{
	  "invoiceId":"42","state":"PAID","price":"1000000","amountPaid":"1000000",
	  "buyer":{"id":"0xbuyer"},"seller":{"id":"0xseller"},
	  "paymentToken":{"id":"0xtoken","name":"USDC","decimal":"6"}
	}}}`

	client, captured := newSubgraph(t, http.StatusOK, body)

	got, err := client.GetInvoiceData("42")
	if err != nil {
		t.Fatalf("GetInvoiceData returned %v", err)
	}

	if got.InvoiceId != "42" {
		t.Errorf("invoiceId = %q, want 42", got.InvoiceId)
	}
	if got.State != "PAID" {
		t.Errorf("state = %q, want PAID", got.State)
	}
	if got.PaymentToken == nil || got.PaymentToken.Name != "USDC" {
		t.Errorf("paymentToken = %+v, want USDC", got.PaymentToken)
	}

	if captured.method != http.MethodPost {
		t.Errorf("method = %s, want POST", captured.method)
	}
	if captured.contentType != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", captured.contentType)
	}
	if captured.variables["id"] != "42" {
		t.Errorf("id variable = %v, want 42", captured.variables["id"])
	}
	if !strings.Contains(captured.query, "smartInvoice") {
		t.Errorf("query = %q, want the smartInvoice query", captured.query)
	}
}

func TestGetInvoiceDataUnknownInvoice(t *testing.T) {
	client, _ := newSubgraph(t, http.StatusOK, `{"data":{"smartInvoice":null}}`)

	got, err := client.GetInvoiceData("999")
	if err != nil {
		t.Fatalf("GetInvoiceData returned %v", err)
	}
	if got == nil {
		t.Fatal("GetInvoiceData returned nil")
	}
	if got.InvoiceId != "" {
		t.Errorf("invoiceId = %q, want an empty invoice", got.InvoiceId)
	}
}

func TestGetUserInvoiceData(t *testing.T) {
	const body = `{"data":{"user":{
	  "issuedInvoices":[{"invoiceId":"1","state":"PAID"}],
	  "receivedInvoices":[{"invoiceId":"2","state":"CREATED"}],
	  "metaInvoices":[]
	}}}`

	client, captured := newSubgraph(t, http.StatusOK, body)

	got, err := client.GetUserInvoiceData("0xuser", 10, 0)
	if err != nil {
		t.Fatalf("GetUserInvoiceData returned %v", err)
	}

	if got.Address != "0xuser" {
		t.Errorf("address = %q, want 0xuser", got.Address)
	}
	if len(got.User.IssuedInvoices) != 1 || got.User.IssuedInvoices[0].InvoiceId != "1" {
		t.Errorf("issued invoices = %+v", got.User.IssuedInvoices)
	}
	if len(got.User.ReceivedInvoices) != 1 || got.User.ReceivedInvoices[0].InvoiceId != "2" {
		t.Errorf("received invoices = %+v", got.User.ReceivedInvoices)
	}

	if captured.variables["address"] != "0xuser" {
		t.Errorf("address variable = %v, want 0xuser", captured.variables["address"])
	}
	if captured.variables["first"] != float64(10) {
		t.Errorf("first variable = %v, want 10", captured.variables["first"])
	}
	if captured.variables["skip"] != float64(0) {
		t.Errorf("skip variable = %v, want 0", captured.variables["skip"])
	}
}

func TestGetUserInvoiceDataPagination(t *testing.T) {
	client, captured := newSubgraph(t, http.StatusOK, `{"data":{"user":{}}}`)

	if _, err := client.GetUserInvoiceData("0xuser", 50, 100); err != nil {
		t.Fatalf("GetUserInvoiceData returned %v", err)
	}

	if captured.variables["first"] != float64(50) {
		t.Errorf("first variable = %v, want 50", captured.variables["first"])
	}
	if captured.variables["skip"] != float64(100) {
		t.Errorf("skip variable = %v, want 100", captured.variables["skip"])
	}
}

func TestGetUserInvoiceDataRejectsAnEmptyAddress(t *testing.T) {
	client, _ := newSubgraph(t, http.StatusOK, `{"data":{"user":{}}}`)

	got, err := client.GetUserInvoiceData("", 10, 0)

	if err == nil {
		t.Fatalf("GetUserInvoiceData returned %+v, want an error", got)
	}
	if !strings.Contains(err.Error(), "address cannot be empty") {
		t.Errorf("error = %q, want it to mention the empty address", err)
	}
}

func TestGetInvoiceDataRejectsAnEmptyID(t *testing.T) {
	client, _ := newSubgraph(t, http.StatusOK, `{"data":{"smartInvoice":null}}`)

	got, err := client.GetInvoiceData("")

	if err == nil {
		t.Fatalf("GetInvoiceData returned %+v, want an error", got)
	}
	if !strings.Contains(err.Error(), "id cannot be empty") {
		t.Errorf("error = %q, want it to mention the empty id", err)
	}
}

func TestSubgraphNon200IsAnError(t *testing.T) {
	client, _ := newSubgraph(t, http.StatusInternalServerError, "upstream exploded")

	_, err := client.GetInvoiceData("42")

	if err == nil {
		t.Fatal("GetInvoiceData returned no error for a 500")
	}
	if !strings.Contains(err.Error(), "non-200 status code: 500") {
		t.Errorf("error = %q, want it to state the status", err)
	}

	if !strings.Contains(err.Error(), "upstream exploded") {
		t.Errorf("error = %q, want it to include the response body", err)
	}
}

func TestSubgraphMalformedJSONIsAnError(t *testing.T) {
	client, _ := newSubgraph(t, http.StatusOK, "not json at all")

	if _, err := client.GetInvoiceData("42"); err == nil {
		t.Error("GetInvoiceData returned no error for a malformed body")
	}
	if _, err := client.GetUserInvoiceData("0xuser", 10, 0); err == nil {
		t.Error("GetUserInvoiceData returned no error for a malformed body")
	}
}

func TestSubgraphRequiresAnEndpoint(t *testing.T) {
	tests := []struct {
		name   string
		client *Client
	}{
		{"nil client", nil},
		{"unconfigured endpoint", NewClient("")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.client.GetInvoiceData("42")

			if err == nil {
				t.Fatal("GetInvoiceData returned no error")
			}
			if !strings.Contains(err.Error(), "subgraph URL is not configured") {
				t.Errorf("error = %q, want it to mention the missing URL", err)
			}
		})
	}
}

func TestSubgraphUnreachableEndpointIsAnError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	url := server.URL
	server.Close()

	client := NewClient(url)

	if _, err := client.GetInvoiceData("42"); err == nil {
		t.Error("GetInvoiceData against a closed server returned no error")
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient("https://subgraph.example.com/graphql")

	if client == nil {
		t.Fatal("NewClient returned nil")
	}
	if client.endpoint != "https://subgraph.example.com/graphql" {
		t.Errorf("endpoint = %q", client.endpoint)
	}
}
