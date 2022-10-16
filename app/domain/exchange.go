package domain

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
)

//go:generate moq -fmt goimports -out exchange_mock.go . Exchange
type Exchange interface {
	CreateExchange(ctx context.Context, input exchange.CreateExchangeInput) (exchange.CreateExchangeOutput, error)
	Convert(ctx context.Context, input exchange.ConvertInput) (exchange.ConvertOutput, error)
}
