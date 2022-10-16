package exchange

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
)

func Test_Repository_GetExchange(t *testing.T) {
	t.Parallel()

	type args struct {
		from vos.CurrencyCode
		to   vos.CurrencyCode
	}

	type want struct {
		err    error
		output entities.Exchange
	}

	tableTests := []struct {
		name      string
		runBefore func(*pgxpool.Pool)
		want
		args
	}{
		{
			name:      "should get exchange",
			runBefore: func(pool *pgxpool.Pool) { insertTestExc(t, pool, testExc01) },
			args: args{
				from: testExc01.From,
				to:   testExc01.To,
			},
			want: want{
				output: testExc01,
			},
		},
		{
			name:      "should return ErrNotFound when from-to pair does not exist",
			runBefore: func(pool *pgxpool.Pool) {},
			args: args{
				from: testExc01.From,
				to:   testExc01.To,
			},
			want: want{
				output: entities.Exchange{},
				err:    exchange.ErrNotFound,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testPool, tearDown := postgrestest.GetTestPool()
			testRepo := NewRepository(testPool)
			t.Cleanup(tearDown)

			tt.runBefore(testPool)
			got, err := testRepo.GetExchangeByCurrencies(testContext, tt.args.from, tt.args.to)
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
