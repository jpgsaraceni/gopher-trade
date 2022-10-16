package currency

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	CreateCurrency(ctx context.Context, cur entities.Currency) error
	GetCurrenciesByCode(ctx context.Context, code ...vos.CurrencyCode) (map[vos.CurrencyCode]entities.Currency, error)
}
