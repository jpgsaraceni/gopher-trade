package currency

import (
	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

type UseCase struct {
	Repo   Repository
	Client Client
}

func NewUseCase(r Repository, c Client) UseCase {
	return UseCase{
		Repo:   r,
		Client: c,
	}
}

type CreateCurrencyInput struct {
	Code    vos.CurrencyCode
	USDRate decimal.Decimal
}

type CreateCurrencyOutput struct {
	Currency entities.Currency
	IsNew    bool
}

type ConvertInput struct {
	From       vos.CurrencyCode
	To         vos.CurrencyCode
	FromAmount decimal.Decimal
}

type ConvertOutput struct {
	ConvertedAmount decimal.Decimal
}

type UpdateCurrencyByCodeInput struct {
	Code vos.CurrencyCode
	Rate decimal.Decimal
}

type UpdateCurrencyByCodeOutput struct{ entities.Currency }

type DeleteCurrencyByCodeInput struct {
	Code vos.CurrencyCode
}
