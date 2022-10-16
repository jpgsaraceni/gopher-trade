package exchange

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (uc UseCase) Convert(ctx context.Context, input ConvertInput) (ConvertOutput, error) {
	const operation = "UseCase.Exchange.Convert"

	exc, err := uc.repo.GetExchangeByCurrencies(ctx, input.From, input.To)
	if err != nil {
		return ConvertOutput{}, extensions.ErrStack(operation, err)
	}

	convertedAmount := exc.Convert(input.FromAmount)

	return ConvertOutput{
		ConvertedAmount: convertedAmount,
	}, nil
}
