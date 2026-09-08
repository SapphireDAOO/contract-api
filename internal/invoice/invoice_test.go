package invoice

import (
	"maps"
	"math/big"
	"slices"
	"strings"
	"testing"

	"github.com/SapphireDAOO/contract-api/internal/blockchain/gen/intermediatedpaymentprocessor"
	"github.com/ethereum/go-ethereum/common"
)

type intermediatedInvoice = intermediatedpaymentprocessor.IIntermediatedPaymentProcessorInvoiceCreationParam

type stubTokens map[string]string

func (s stubTokens) Address(symbol string) (common.Address, bool) {
	for configured, address := range s {
		if strings.EqualFold(configured, symbol) {
			return common.HexToAddress(address), true
		}
	}
	return common.Address{}, false
}

func (s stubTokens) Symbols() []string { return slices.Sorted(maps.Keys(s)) }

const (
	usdcAddress   = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	sellerAddress = "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"
)

func testTokens() stubTokens {
	return stubTokens{
		"ETH":  "0x0000000000000000000000000000000000000000",
		"USDC": usdcAddress,
	}
}

func validParam() CreateInvoiceParam {
	return CreateInvoiceParam{
		OrderId:          "ORDER-1",
		Seller:           sellerAddress,
		Price:            1000,
		EscrowHoldPeriod: 3600,
		Currency:         "USD",
		PaymentTokens:    []string{"USDC"},
	}
}

func TestValidateCreateInvoiceParams(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*CreateInvoiceParam)
		params  []CreateInvoiceParam
		wantErr string
	}{
		{name: "valid", mutate: func(*CreateInvoiceParam) {}},
		{
			name:    "no parameters",
			params:  []CreateInvoiceParam{},
			wantErr: "no invoice parameters provided",
		},
		{
			name:    "nil parameters",
			params:  nil,
			wantErr: "no invoice parameters provided",
		},
		{
			name:    "empty order id",
			mutate:  func(p *CreateInvoiceParam) { p.OrderId = "" },
			wantErr: "orderId is required",
		},
		{
			name:    "whitespace order id",
			mutate:  func(p *CreateInvoiceParam) { p.OrderId = "   " },
			wantErr: "orderId is required",
		},
		{
			name:    "empty seller",
			mutate:  func(p *CreateInvoiceParam) { p.Seller = "" },
			wantErr: "is not a valid address",
		},
		{
			name:    "malformed seller",
			mutate:  func(p *CreateInvoiceParam) { p.Seller = "0x123" },
			wantErr: "is not a valid address",
		},
		{
			name:    "zero seller address",
			mutate:  func(p *CreateInvoiceParam) { p.Seller = "0x0000000000000000000000000000000000000000" },
			wantErr: "is not a valid address",
		},
		{
			name:    "zero price",
			mutate:  func(p *CreateInvoiceParam) { p.Price = 0 },
			wantErr: "price must be greater than zero",
		},
		{
			name:    "negative price",
			mutate:  func(p *CreateInvoiceParam) { p.Price = -1 },
			wantErr: "price must be greater than zero",
		},
		{
			name:    "no payment tokens",
			mutate:  func(p *CreateInvoiceParam) { p.PaymentTokens = nil },
			wantErr: "at least one payment token is required",
		},
		{
			name:    "unknown payment token",
			mutate:  func(p *CreateInvoiceParam) { p.PaymentTokens = []string{"DOGE"} },
			wantErr: `unknown payment token "DOGE"`,
		},
		{
			name:    "one unknown token among known ones",
			mutate:  func(p *CreateInvoiceParam) { p.PaymentTokens = []string{"USDC", "DOGE"} },
			wantErr: `unknown payment token "DOGE"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := tt.params
			if tt.mutate != nil {
				p := validParam()
				tt.mutate(&p)
				params = []CreateInvoiceParam{p}
			}

			err := ValidateCreateInvoiceParams(params, testTokens())

			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("ValidateCreateInvoiceParams returned %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("ValidateCreateInvoiceParams returned nil, want an error containing %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCreateInvoiceParamsTrimsAndFolds(t *testing.T) {
	p := validParam()
	p.Seller = "  " + sellerAddress + "  "
	p.PaymentTokens = []string{" usdc ", "eth"}

	if err := ValidateCreateInvoiceParams([]CreateInvoiceParam{p}, testTokens()); err != nil {
		t.Errorf("ValidateCreateInvoiceParams returned %v, want nil", err)
	}
}

func TestValidateCreateInvoiceParamsReportsTheIndex(t *testing.T) {
	good := validParam()
	bad := validParam()
	bad.Price = 0

	err := ValidateCreateInvoiceParams([]CreateInvoiceParam{good, good, bad}, testTokens())

	if err == nil {
		t.Fatal("ValidateCreateInvoiceParams returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "invoice 2") {
		t.Errorf("error = %q, want it to name invoice 2", err)
	}
}

func TestValidateCreateInvoiceParamsListsKnownTokens(t *testing.T) {
	p := validParam()
	p.PaymentTokens = []string{"DOGE"}

	err := ValidateCreateInvoiceParams([]CreateInvoiceParam{p}, testTokens())

	if err == nil {
		t.Fatal("ValidateCreateInvoiceParams returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "ETH, USDC") {
		t.Errorf("error = %q, want it to list the known symbols", err)
	}
}

func TestConvertParam(t *testing.T) {
	p := validParam()
	p.Price = 1000

	got, err := ConvertParam([]CreateInvoiceParam{p}, testTokens())
	if err != nil {
		t.Fatalf("ConvertParam returned %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("ConvertParam returned %d invoices, want 1", len(got))
	}

	invoice := got[0]
	if invoice.InvoiceId != "ORDER-1" {
		t.Errorf("InvoiceId = %q, want %q", invoice.InvoiceId, "ORDER-1")
	}
	if invoice.Seller != common.HexToAddress(sellerAddress) {
		t.Errorf("Seller = %s, want %s", invoice.Seller, sellerAddress)
	}
	if invoice.EscrowHoldPeriod != 3600 {
		t.Errorf("EscrowHoldPeriod = %d, want 3600", invoice.EscrowHoldPeriod)
	}

	if want := big.NewInt(1000 * 1_000_000); invoice.Price.Cmp(want) != 0 {
		t.Errorf("Price = %s, want %s", invoice.Price, want)
	}
	if len(invoice.PaymentTokens) != 1 || invoice.PaymentTokens[0] != common.HexToAddress(usdcAddress) {
		t.Errorf("PaymentTokens = %v, want [%s]", invoice.PaymentTokens, usdcAddress)
	}
}

func TestConvertParamScalesByCurrencyPrecision(t *testing.T) {
	tests := []struct {
		name     string
		currency string
		price    int
		want     string
	}{
		{"USD scales cents to 8 decimals", "USD", 1, "1000000"},
		{"USD larger amount", "USD", 250000, "250000000000"},

		{"unknown currency is not scaled", "EUR", 1000, "1000"},
		{"empty currency is not scaled", "", 1000, "1000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := validParam()
			p.Currency = tt.currency
			p.Price = tt.price

			got, err := ConvertParam([]CreateInvoiceParam{p}, testTokens())
			if err != nil {
				t.Fatalf("ConvertParam returned %v", err)
			}
			if got[0].Price.String() != tt.want {
				t.Errorf("Price = %s, want %s", got[0].Price, tt.want)
			}
		})
	}
}

func TestConvertParamRejectsUnknownTokens(t *testing.T) {
	p := validParam()
	p.PaymentTokens = []string{"DOGE"}

	got, err := ConvertParam([]CreateInvoiceParam{p}, testTokens())

	if err == nil {
		t.Fatalf("ConvertParam returned %v, want an error", got)
	}
	if !strings.Contains(err.Error(), `unknown payment token "DOGE"`) {
		t.Errorf("error = %q, want it to name the unknown token", err)
	}
	if !strings.Contains(err.Error(), "invoice 0") {
		t.Errorf("error = %q, want it to name the invoice index", err)
	}
	if got != nil {
		t.Errorf("ConvertParam returned %v alongside an error, want nil", got)
	}
}

func TestConvertParamResolvesTheNativeToken(t *testing.T) {
	p := validParam()
	p.PaymentTokens = []string{"ETH"}

	got, err := ConvertParam([]CreateInvoiceParam{p}, testTokens())
	if err != nil {
		t.Fatalf("ConvertParam returned %v", err)
	}
	if got[0].PaymentTokens[0] != (common.Address{}) {
		t.Errorf("ETH resolved to %s, want the zero address", got[0].PaymentTokens[0])
	}
}

func TestConvertParamPreservesOrderAndCount(t *testing.T) {
	first := validParam()
	first.OrderId = "ORDER-1"
	second := validParam()
	second.OrderId = "ORDER-2"
	second.PaymentTokens = []string{"ETH", "USDC"}

	got, err := ConvertParam([]CreateInvoiceParam{first, second}, testTokens())
	if err != nil {
		t.Fatalf("ConvertParam returned %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("ConvertParam returned %d invoices, want 2", len(got))
	}
	if got[0].InvoiceId != "ORDER-1" || got[1].InvoiceId != "ORDER-2" {
		t.Errorf("order not preserved: %q, %q", got[0].InvoiceId, got[1].InvoiceId)
	}
	if len(got[1].PaymentTokens) != 2 {
		t.Errorf("second invoice has %d payment tokens, want 2", len(got[1].PaymentTokens))
	}
}

func TestConvertParamEmpty(t *testing.T) {
	got, err := ConvertParam(nil, testTokens())

	if err != nil {
		t.Fatalf("ConvertParam(nil) returned %v", err)
	}
	if len(got) != 0 {
		t.Errorf("ConvertParam(nil) = %v, want no invoices", got)
	}
}

func TestValidateInvoices(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(inv []intermediatedInvoice)
		wantErr string
	}{
		{name: "valid"},
		{
			name:    "missing order id",
			mutate:  func(inv []intermediatedInvoice) { inv[0].InvoiceId = "" },
			wantErr: "missing orderId",
		},
		{
			name:    "whitespace order id",
			mutate:  func(inv []intermediatedInvoice) { inv[0].InvoiceId = "  " },
			wantErr: "missing orderId",
		},
		{
			name:    "zero seller",
			mutate:  func(inv []intermediatedInvoice) { inv[0].Seller = common.Address{} },
			wantErr: "missing seller",
		},
		{
			name:    "nil price",
			mutate:  func(inv []intermediatedInvoice) { inv[0].Price = nil },
			wantErr: "missing price",
		},
		{
			name:    "no payment tokens",
			mutate:  func(inv []intermediatedInvoice) { inv[0].PaymentTokens = nil },
			wantErr: "missing payment tokens",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invoices, err := ConvertParam([]CreateInvoiceParam{validParam()}, testTokens())
			if err != nil {
				t.Fatalf("ConvertParam returned %v", err)
			}
			if tt.mutate != nil {
				tt.mutate(invoices)
			}

			err = ValidateInvoices(invoices)

			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("ValidateInvoices returned %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("ValidateInvoices returned nil, want an error containing %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestValidateInvoicesEmpty(t *testing.T) {
	if err := ValidateInvoices(nil); err != nil {
		t.Errorf("ValidateInvoices(nil) returned %v, want nil", err)
	}
}

func TestValidateInvoicesReportsTheIndex(t *testing.T) {
	invoices, err := ConvertParam([]CreateInvoiceParam{validParam(), validParam()}, testTokens())
	if err != nil {
		t.Fatalf("ConvertParam returned %v", err)
	}
	invoices[1].Price = nil

	err = ValidateInvoices(invoices)

	if err == nil {
		t.Fatal("ValidateInvoices returned nil, want an error")
	}
	if !strings.Contains(err.Error(), "invoice 1") {
		t.Errorf("error = %q, want it to name invoice 1", err)
	}
}

func TestConvertParamOutputPassesValidateInvoices(t *testing.T) {
	params := []CreateInvoiceParam{validParam(), validParam()}

	if err := ValidateCreateInvoiceParams(params, testTokens()); err != nil {
		t.Fatalf("ValidateCreateInvoiceParams returned %v", err)
	}
	invoices, err := ConvertParam(params, testTokens())
	if err != nil {
		t.Fatalf("ConvertParam returned %v", err)
	}
	if err := ValidateInvoices(invoices); err != nil {
		t.Errorf("ValidateInvoices returned %v for converted valid params", err)
	}
}

func TestCurrencyPrecision(t *testing.T) {
	precision, ok := CurrencyPrecision["USD"]
	if !ok {
		t.Fatal("CurrencyPrecision is missing USD")
	}
	if precision <= 2 {
		t.Errorf("CurrencyPrecision[USD] = %d, want more than the 2 decimals a price arrives with", precision)
	}
}
