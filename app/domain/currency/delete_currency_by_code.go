package currency

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (uc UseCase) DeleteCurrencyByCode(ctx context.Context, input DeleteCurrencyByCodeInput) error {
	const operation = "UseCase.Currency.DeleteCurrencyByCode"
	if entities.IsDefaultRate(input.Code) {
		return extensions.ErrStack(operation, ErrDefaultRate)
	}

	err := uc.Repo.DeleteCurrencyByCode(ctx, input.Code)
	if err != nil {
		return extensions.ErrStack(operation, err)
	}

	return nil
}
