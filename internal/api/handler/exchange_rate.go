package handler

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/httpx"
	"github.com/SapphireDAOO/contract-api/internal/revert"
	"github.com/SapphireDAOO/contract-api/internal/units"
)

// Chainlink USD feed convention, and is reported in the response so a client
const priceDecimals = 8

// usdCurrency is the only currency the oracle prices against.
const usdCurrency = "USD"

var defaultUsdPerToken = new(big.Int).Exp(big.NewInt(10), big.NewInt(priceDecimals), nil)

func (h *ContractHandler) ExchangeRate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	from := strings.TrimSpace(query.Get("from"))

	if from == "" {
		from = usdCurrency
	}
	if !strings.EqualFold(from, usdCurrency) {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest,
			errors.New("only USD is priced"), "from must be "+usdCurrency)
		return
	}

	symbols := query["to"]
	if len(symbols) == 0 {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest,
			errors.New("no token requested"), "to is required, e.g. to=ETH,USDC")
		return
	}

	if h.Oracle == nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusServiceUnavailable,
			errors.New("oracle unavailable"), "exchange rates are unavailable")
		return
	}

	// json.Number keeps the exact digits: encoding a float64 here would
	// round an 18-decimal rate.
	rates := make(map[string]json.Number, len(symbols))
	for _, symbol := range symbols {
		address, ok := h.Tokens.Address(symbol)
		if !ok {
			httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest,
				errors.New("unknown token "+symbol),
				"unknown token "+symbol+" (known: "+strings.Join(h.Tokens.Symbols(), ", ")+")")
			return
		}

		price, err := h.Oracle.UsdPerToken(address)
		if err != nil {
			if !revert.Is(err, revert.UnsupportedToken) {
				httpx.WriteHTTPErrorWithStatus(w, http.StatusBadGateway, err,
					"failed to read the price of "+symbol)
				return
			}
			price = defaultUsdPerToken
		}

		if price.Sign() <= 0 {
			httpx.WriteHTTPErrorWithStatus(w, http.StatusBadGateway,
				errors.New("oracle returned a non-positive price"),
				"invalid price for "+symbol)
			return
		}

		_, decimals, ok := h.Tokens.ByAddress(address.Hex())
		if !ok {
			decimals = priceDecimals
		}

		rate := units.Invert(price, decimals, priceDecimals)
		if rate == nil {
			httpx.WriteHTTPErrorWithStatus(w, http.StatusBadGateway,
				errors.New("price cannot be inverted"),
				"invalid price for "+symbol)
			return
		}

		rates[symbol] = json.Number(units.Format(rate, decimals))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"from": usdCurrency,
		"to":   rates,
	})
}
