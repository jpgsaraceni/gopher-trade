package exchange

import (
	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

type UseCase struct {
	repo Repository
}

func NewUseCase(r Repository) UseCase {
	return UseCase{
		repo: r,
	}
}

type ConvertInput struct {
	From       vos.CurrencyCode
	To         vos.CurrencyCode
	FromAmount decimal.Decimal
}

type ConvertOutput struct {
	ConvertedAmount decimal.Decimal
}

type CreateExchangeInput struct {
	From vos.CurrencyCode
	To   vos.CurrencyCode
	Rate decimal.Decimal
}

type CreateExchangeOutput struct {
	Exc entities.Exchange
}
