package domain

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
)

//go:generate moq -fmt goimports -out currency_mock.go . Currency
type Currency interface {
	UpsertCurrency(ctx context.Context, input currency.CreateCurrencyInput) (currency.CreateCurrencyOutput, error)
	Convert(ctx context.Context, input currency.ConvertInput) (currency.ConvertOutput, error)
	DeleteCurrencyByCode(ctx context.Context, input currency.DeleteCurrencyByCodeInput) error
}
