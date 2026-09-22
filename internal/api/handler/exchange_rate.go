package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/SapphireDAOO/contract-api/internal/httpx"
	"github.com/SapphireDAOO/contract-api/internal/units"
	"github.com/ethereum/go-ethereum/common"
)

// Chainlink USD feed convention, and is reported in the response so a client
const priceDecimals = 8

// usdCurrency is the only currency the oracle prices against.
const usdCurrency = "USD"

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
			errors.New("no token requested"), "to is required, e.g. to=ETH&to=USDC")
		return
	}

	addresses := make([]common.Address, 0, len(symbols))
	for _, symbol := range symbols {
		address, ok := h.Tokens.Address(symbol)
		if !ok {
			httpx.WriteHTTPErrorWithStatus(w, http.StatusBadRequest,
				errors.New("unknown token "+symbol),
				"unknown token "+symbol+" (known: "+strings.Join(h.Tokens.Symbols(), ", ")+")")
			return
		}
		addresses = append(addresses, address)
	}

	if h.Oracle == nil {
		httpx.WriteHTTPErrorWithStatus(w, http.StatusServiceUnavailable,
			errors.New("oracle unavailable"), "exchange rates are unavailable")
		return
	}

	// One call prices every token asked for. The contract reverts the whole
	// batch if any of them has no feed, which surfaces as a 400 rather than a
	// rate the caller cannot trust.
	prices, err := h.Oracle.UsdPerTokenBatch(addresses)
	if err != nil {
		httpx.WriteMappedRevertError(w, err, "failed to read token prices")
		return
	}

	// json.Number keeps the exact digits: encoding a float64 here would
	// round an 18-decimal rate.
	rates := make(map[string]json.Number, len(symbols))
	for i, symbol := range symbols {
		price := prices[i]
		if price == nil || price.Sign() <= 0 {
			httpx.WriteHTTPErrorWithStatus(w, http.StatusBadGateway,
				errors.New("oracle returned a non-positive price"),
				"invalid price for "+symbol)
			return
		}

		_, decimals, ok := h.Tokens.ByAddress(addresses[i].Hex())
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
	if err := json.NewEncoder(w).Encode(map[string]any{
		"from": usdCurrency,
		"to":   rates,
	}); err != nil {
		slog.Error("writing the response failed", "error", err)
	}
}
