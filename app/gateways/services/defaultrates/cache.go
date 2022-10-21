package defaultrates

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

//go:generate moq -fmt goimports -out cache_mock.go . Cache
type Cache interface {
	SetRate(_ context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error
}
