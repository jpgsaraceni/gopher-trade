package redis

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) GetRate(_ context.Context, key vos.CurrencyCode) (decimal.Decimal, error) {
	const operation = "Redis.GetRate"

	conn := r.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GET", key.String())
	if err != nil {
		return decimal.Zero, extensions.ErrStack(operation, err)
	}

	if reply == nil { // key does not exist
		return decimal.Zero, nil
	}

	rateBytes, ok := reply.([]byte)
	rateString := string(rateBytes)

	if !ok {
		return decimal.Zero, extensions.ErrStack(operation, errType)
	}

	rate, err := decimal.NewFromString(rateString)
	if err != nil {
		return decimal.Zero, extensions.ErrStack(operation, err)
	}

	return rate, nil
}
