package exchange

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (uc UseCase) CreateExchange(ctx context.Context, input CreateExchangeInput) (CreateExchangeOutput, error) {
	const operation = "UseCase.Exchange.CreateExchange"

	exc := entities.NewExchange(input.From, input.To, input.Rate)
	err := uc.repo.CreateExchange(ctx, exc)
	if err != nil {
		return CreateExchangeOutput{}, extensions.ErrStack(operation, err)
	}

	return CreateExchangeOutput{
		Exc: exc,
	}, nil
}
