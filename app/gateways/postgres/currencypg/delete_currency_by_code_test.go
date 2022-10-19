package currencypg_test

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/currencypg"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func Test_Repository_DeleteCurrencyByCode(t *testing.T) {
	t.Parallel()

	type want struct {
		err error
	}

	tableTests := []struct {
		want
		name      string
		runBefore func(*pgxpool.Pool)
		args      vos.CurrencyCode
	}{
		{
			name: "should delete currency",
			runBefore: func(pool *pgxpool.Pool) {
				insertTestCur(t, pool, extensions.CurrencyFixtures[0])
			},
			args: extensions.CurrencyFixtures[0].Code,
			want: want{
				err: nil,
			},
		},
		{
			name:      "should return ErrNotFound when currency code does not exist",
			runBefore: func(pool *pgxpool.Pool) {},
			args:      "I DON'T EXIST",
			want: want{
				err: currency.ErrNotFound,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testPool, tearDown := postgrestest.GetTestPool()
			testRepo := currencypg.NewRepository(testPool)
			t.Cleanup(tearDown)

			tt.runBefore(testPool)
			err := testRepo.DeleteCurrencyByCode(testContext, tt.args)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
