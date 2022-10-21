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

	var (
		fromRate decimal.Decimal
		toRate   decimal.Decimal
	)

	fromRate, err := uc.getRateFromCache(ctx, input.From)
	if err != nil {
		return ConvertOutput{}, extensions.ErrStack(operation, err)
	}
	toRate, err = uc.getRateFromCache(ctx, input.To)
	if err != nil {
		return ConvertOutput{}, extensions.ErrStack(operation, err)
	}

	g := new(errgroup.Group)

	g.Go(func() error {
		if fromRate == decimal.Zero {
			rate, err := uc.getRate(ctx, input.From)
			if err != nil {
				return err
			}
			fromRate = rate
		}

		return nil
	})

	g.Go(func() error {
		if toRate == decimal.Zero {
			rate, err := uc.getRate(ctx, input.To)
			if err != nil {
				return err
			}
			toRate = rate
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return ConvertOutput{}, extensions.ErrStack(operation, err)
	}

	return ConvertOutput{
		ConvertedAmount: entities.Convert(fromRate, toRate, input.FromAmount),
	}, nil
}

func (uc UseCase) getRateFromCache(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
	const operation = "UseCase.Currency.getRateFromCache"

	rate, err := uc.Cache.GetRate(ctx, code)
	if err != nil {
		return decimal.Zero, extensions.ErrStack(operation, err)
	}

	return rate, nil
}

func (uc UseCase) getRate(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
	const operation = "UseCase.Currency.getRate"

	switch code {
	case vos.USD:
		return decimal.NewFromInt(1), nil
	case vos.BRL, vos.BTC, vos.ETH, vos.EUR:
		currencies, err := uc.Client.GetRates(ctx)
		if err != nil {
			return decimal.Decimal{}, extensions.ErrStack(operation, err)
		}
		err = uc.setRateToCache(ctx, code, currencies[code])
		if err != nil {
			return decimal.Zero, extensions.ErrStack(operation, err)
		}

		return currencies[code], nil
	default:
		cur, err := uc.Repo.GetCurrencyByCode(ctx, code)
		if err != nil {
			return decimal.Decimal{}, extensions.ErrStack(operation, err)
		}
		err = uc.setRateToCache(ctx, code, cur.USDRate)
		if err != nil {
			return decimal.Zero, extensions.ErrStack(operation, err)
		}

		return cur.USDRate, nil
	}
}

func (uc UseCase) setRateToCache(ctx context.Context, code vos.CurrencyCode, rate decimal.Decimal) error {
	const operation = "UseCase.Currency.setRateToCache"

	err := uc.Cache.SetRate(ctx, map[vos.CurrencyCode]decimal.Decimal{code: rate})
	if err != nil {
		return extensions.ErrStack(operation, err)
	}

	return nil
}
