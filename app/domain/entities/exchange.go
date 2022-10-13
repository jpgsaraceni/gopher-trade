package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

type Exchange struct {
	id        string
	from      vos.CurrencyCode
	to        vos.CurrencyCode
	createdAt time.Time
	updatedAt time.Time
	rate      decimal.Decimal
}

// NewExchange generates an ID (UUID) and timestamps and returns an Exchange struct.
func NewExchange(from, to vos.CurrencyCode, rate decimal.Decimal) Exchange {
	return Exchange{
		id:        uuid.NewString(),
		from:      from,
		to:        to,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		rate:      rate,
	}
}

func (e *Exchange) UpdateRate(r decimal.Decimal) {
	e.rate = r
}

// Convert method applies Exchange rate and converts from the From currency
// to the To currency.
func (e Exchange) Convert(fromAmount decimal.Decimal) decimal.Decimal {
	return e.rate.Mul(fromAmount)
}

// getters

func (e Exchange) From() vos.CurrencyCode {
	return e.from
}

func (e Exchange) To() vos.CurrencyCode {
	return e.to
}

func (e Exchange) Rate() decimal.Decimal {
	return e.rate
}
