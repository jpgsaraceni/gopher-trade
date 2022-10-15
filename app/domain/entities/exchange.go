package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

type Exchange struct {
	ID        string
	From      vos.CurrencyCode
	To        vos.CurrencyCode
	CreatedAt time.Time
	UpdatedAt time.Time
	Rate      decimal.Decimal
}

// NewExchange generates an ID (UUID) and timestamps and returns an Exchange struct.
func NewExchange(from, to vos.CurrencyCode, rate decimal.Decimal) Exchange {
	return Exchange{
		ID:        uuid.NewString(),
		From:      from,
		To:        to,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Rate:      rate,
	}
}

func (e *Exchange) UpdateRate(r decimal.Decimal) {
	e.UpdatedAt = time.Now().UTC()
	e.Rate = r
}

// Convert method applies Exchange rate and converts from the From currency
// to the To currency.
func (e Exchange) Convert(fromAmount decimal.Decimal) decimal.Decimal {
	return e.Rate.Mul(fromAmount)
}
