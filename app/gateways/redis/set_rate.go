package redis

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

const ttl = 60 * 5 // TODO: move to env

func (r Repository) SetRate(_ context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
	const operation = "Redis.SetRate"

	conn := r.pool.Get()
	defer conn.Close()

	err := conn.Send("MULTI")
	if err != nil {
		return extensions.ErrStack(operation, err)
	}

	for code, rate := range currencyRates {
		err = conn.Send("SETEX", code.String(), ttl, rate.String())
		if err != nil {
			return extensions.ErrStack(operation, err)
		}
	}
	_, err = conn.Do("EXEC")

	if err != nil {
		return extensions.ErrStack(operation, err)
	}

	return nil
}
