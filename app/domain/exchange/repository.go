package exchange

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	CreateExchange(ctx context.Context, exc entities.Exchange) error
	GetExchangeByCurrencies(ctx context.Context, from, to vos.CurrencyCode) (entities.Exchange, error)
}
