package currency

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (uc UseCase) CreateCurrency(ctx context.Context, input CreateCurrencyInput) (CreateCurrencyOutput, error) {
	const operation = "UseCase.Currency.CreateCurrency"
	if entities.IsDefaultRate(input.Code) {
		return CreateCurrencyOutput{}, extensions.ErrStack(operation, ErrDefaultRate)
	}

	cur := entities.NewCurrency(input.Code, input.USDRate)
	createdCur, err := uc.Repo.CreateCurrency(ctx, cur)
	if err != nil {
		return CreateCurrencyOutput{}, extensions.ErrStack(operation, err)
	}
	isNew := createdCur.CreatedAt.Equal(createdCur.UpdatedAt)

	return CreateCurrencyOutput{
		Currency: createdCur,
		IsNew:    isNew,
	}, nil
}
