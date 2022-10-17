package currencypg_test

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/currencypg"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func Test_Repository_GetCurrencyByCode(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output entities.Currency
	}

	tableTests := []struct {
		name      string
		runBefore func(*pgxpool.Pool)
		want
		args vos.CurrencyCode
	}{
		{
			name: "should get currency",
			runBefore: func(pool *pgxpool.Pool) {
				insertTestCur(t, pool, extensions.CurrencyFixtures[0])
			},
			args: extensions.CurrencyFixtures[0].Code,
			want: want{
				output: extensions.CurrencyFixtures[0],
			},
		},
		{
			name:      "should return ErrNotFound when currency code does not exist",
			runBefore: func(pool *pgxpool.Pool) {},
			args:      "I DON'T EXIST",
			want: want{
				output: entities.Currency{},
				err:    currency.ErrNotFound,
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
			got, err := testRepo.GetCurrencyByCode(testContext, tt.args)
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
