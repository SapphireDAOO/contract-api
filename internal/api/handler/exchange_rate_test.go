package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func exchangeRateRequest(query string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/v1/exchange-rate?"+query, nil)
}

func TestExchangeRateRejectsANonUSDBaseCurrency(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.ExchangeRate(rec, exchangeRateRequest("from=EUR&to=ETH"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if body := decodeError(t, rec); !strings.Contains(body["error"], "from must be USD") {
		t.Errorf("error = %q, want it to name the only supported currency", body["error"])
	}
}

func TestExchangeRateAcceptsUSDInAnyForm(t *testing.T) {
	for _, from := range []string{"", "USD", "usd", "Usd", "  USD  "} {
		t.Run(from, func(t *testing.T) {
			h := &ContractHandler{Tokens: testTokens()}
			rec := httptest.NewRecorder()

			query := url.Values{"from": {from}, "to": {"ETH"}}
			h.ExchangeRate(rec, exchangeRateRequest(query.Encode()))

			if rec.Code == http.StatusBadRequest {
				t.Errorf("from=%q was rejected as a bad currency: %s", from, rec.Body.String())
			}
		})
	}
}

func TestExchangeRateRequiresATargetToken(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.ExchangeRate(rec, exchangeRateRequest("from=USD"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	body := decodeError(t, rec)
	if !strings.Contains(body["error"], "to is required") {
		t.Errorf("error = %q, want it to say the parameter is required", body["error"])
	}

	if !strings.Contains(body["error"], "to=ETH&to=USDC") {
		t.Errorf("error = %q, want it to show an example", body["error"])
	}
}

func TestExchangeRateWithoutAnOracle(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.ExchangeRate(rec, exchangeRateRequest("to=ETH"))

	if rec.Code != http.StatusServiceUnavailable {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
	}
	if body := decodeError(t, rec); !strings.Contains(body["error"], "exchange rates are unavailable") {
		t.Errorf("error = %q, want it to say rates are unavailable", body["error"])
	}
}

func TestExchangeRateChecksTheCurrencyBeforeTheOracle(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.ExchangeRate(rec, exchangeRateRequest("from=EUR&to=ETH"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d, so a client sees the real problem", rec.Code, http.StatusBadRequest)
	}
}

func TestPriceDecimalsMatchesTheFeedConvention(t *testing.T) {
	if priceDecimals != 8 {
		t.Errorf("priceDecimals = %d, want 8", priceDecimals)
	}
	if usdCurrency != "USD" {
		t.Errorf("usdCurrency = %q, want USD", usdCurrency)
	}
}

func TestExchangeRateRejectsAnUnknownTokenBeforeCallingTheOracle(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.ExchangeRate(rec, exchangeRateRequest("to=DOGE"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	body := decodeError(t, rec)
	if !strings.Contains(body["error"], "unknown token DOGE") {
		t.Errorf("error = %q, want it to name the token", body["error"])
	}
	if !strings.Contains(body["error"], "ETH, USDC") {
		t.Errorf("error = %q, want it to list the known symbols", body["error"])
	}
}

func TestExchangeRateChecksEveryTokenBeforeTheOracle(t *testing.T) {
	h := &ContractHandler{Tokens: testTokens()}
	rec := httptest.NewRecorder()

	h.ExchangeRate(rec, exchangeRateRequest("to=ETH&to=DOGE"))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d; a bad token anywhere in the batch is a 400", rec.Code, http.StatusBadRequest)
	}
}
