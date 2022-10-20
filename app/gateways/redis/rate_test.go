package redis

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/redis/redistest"
)

func Test_Redis_Rate(t *testing.T) {
	t.Parallel()

	testConn, tearDown := redistest.GetTestPool()
	testRepo := NewRepository(testConn)

	t.Cleanup(tearDown)

	type want struct {
		setErr error
		getErr error
		rate   decimal.Decimal
	}

	type setArgs struct {
		currencyRates map[vos.CurrencyCode]decimal.Decimal
	}

	type getArgs struct {
		code vos.CurrencyCode
	}

	tableTests := []struct {
		name string
		setArgs
		getArgs
		want
	}{
		{
			name: "should set and get keys",
			setArgs: setArgs{
				currencyRates: map[vos.CurrencyCode]decimal.Decimal{
					vos.BRL: decimal.NewFromFloat(1.234),
					vos.EUR: decimal.NewFromFloat(2.4321),
				},
			},
			getArgs: getArgs{
				code: vos.BRL,
			},
			want: want{
				rate: decimal.NewFromFloat(1.234),
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := testRepo.SetRate(context.Background(), tt.setArgs.currencyRates)
			assert.Equal(t, tt.want.setErr, err)

			got, err := testRepo.GetRate(context.Background(), tt.getArgs.code)
			assert.Equal(t, tt.want.rate, got)
			assert.ErrorIs(t, err, tt.want.getErr)
		})
	}
}
