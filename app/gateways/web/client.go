package web

import (
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type Client struct {
	client http.Client
}

func NewClient() Client {
	return Client{
		client: http.Client{
			Timeout: time.Minute, // TODO: move to env
		},
	}
}

const (
	cryptoCompareURL = "https://min-api.cryptocompare.com/data/price?fsym=USD&tsyms=ETH"
	exchangeRateURL  = "https://api.exchangerate.host/latest?base=USD&symbols=EUR,BRL,BTC"
)

// ExchangeRateClientResponse is the payload for [get] https://api.exchangerate.host/latest?base=USD&symbols=EUR,BRL,BTC
type ExchangeRateClientResponse struct {
	Rates ExchangeRateClientRates `json:"rates"`
}

type ExchangeRateClientRates struct {
	BRL decimal.Decimal `json:"brl"`
	BTC decimal.Decimal `json:"btc"`
	EUR decimal.Decimal `json:"eur"`
}

// CryptoCompareClientResponse is the payload for [get] https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD
type CryptoCompareClientResponse struct {
	ETH decimal.Decimal `json:"eth"`
}
