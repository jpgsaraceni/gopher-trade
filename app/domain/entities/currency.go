package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

type Currency struct {
	ID        string
	Code      vos.CurrencyCode
	CreatedAt time.Time
	UpdatedAt time.Time
	USDRate   decimal.Decimal
}

// NewCurrency generates an ID (UUID) and timestamps and returns an instance of Currency.
func NewCurrency(code vos.CurrencyCode, usdRate decimal.Decimal) Currency {
	return Currency{
		ID:        uuid.NewString(),
		Code:      code,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		USDRate:   usdRate,
	}
}

func (c *Currency) UpdateCurrency(r decimal.Decimal) {
	c.UpdatedAt = time.Now().UTC()
	c.USDRate = r
}

// Convert converts amount in original currency to dollars then to target currency
// using their USD rates.
func Convert(original, target Currency, amount decimal.Decimal) decimal.Decimal {
	decimal.DivisionPrecision = 5 // TODO: move to env
	fromInUSD := amount.Mul(original.USDRate)
	result := fromInUSD.Div(target.USDRate)

	return result
}
