package currency

import (
	"context"

	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (uc UseCase) Convert(ctx context.Context, input ConvertInput) (ConvertOutput, error) {
	const operation = "UseCase.Currency.Convert"

	g := new(errgroup.Group)
	var (
		fromRate decimal.Decimal
		toRate   decimal.Decimal
	)

	g.Go(func() error {
		rate, err := uc.getRate(ctx, input.From)
		if err != nil {
			return err
		}
		fromRate = rate

		return nil
	})

	g.Go(func() error {
		rate, err := uc.getRate(ctx, input.To)
		if err != nil {
			return err
		}
		toRate = rate

		return nil
	})

	if err := g.Wait(); err != nil {
		return ConvertOutput{}, extensions.ErrStack(operation, err)
	}

	return ConvertOutput{
		ConvertedAmount: entities.Convert(fromRate, toRate, input.FromAmount),
	}, nil
}

func (uc UseCase) getRate(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
	const operation = "UseCase.Currency.getRate"

	switch code {
	case vos.USD:
		return decimal.NewFromInt(1), nil
	case vos.BRL, vos.BTC, vos.ETH, vos.EUR:
		currencies, err := uc.Client.GetRates(ctx) // TODO: cache
		if err != nil {
			return decimal.Decimal{}, extensions.ErrStack(operation, err)
		}

		return currencies[code], nil
	default:
		cur, err := uc.Repo.GetCurrencyByCode(ctx, code) // TODO: cache
		if err != nil {
			return decimal.Decimal{}, extensions.ErrStack(operation, err)
		}

		return cur.USDRate, nil
	}
}
