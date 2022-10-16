package currency

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (uc UseCase) Convert(ctx context.Context, input ConvertInput) (ConvertOutput, error) {
	const operation = "UseCase.Currency.Convert"

	curs, err := uc.repo.GetCurrenciesByCode(ctx, input.From, input.To)
	if err != nil {
		return ConvertOutput{}, extensions.ErrStack(operation, err)
	}

	convertedAmount := entities.Convert(curs[input.From], curs[input.To], input.FromAmount)

	return ConvertOutput{
		ConvertedAmount: convertedAmount,
	}, nil
}
