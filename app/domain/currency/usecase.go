package currency

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

type CreateCurrencyInput struct {
	Code    vos.CurrencyCode
	USDRate decimal.Decimal
}

type CreateCurrencyOutput struct {
	Currency entities.Currency
}

type ConvertInput struct {
	From       vos.CurrencyCode
	To         vos.CurrencyCode
	FromAmount decimal.Decimal
}

type ConvertOutput struct {
	ConvertedAmount decimal.Decimal
}
