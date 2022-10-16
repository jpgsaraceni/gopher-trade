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

func Test_Repository_GetCurrenciesByCode(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output map[vos.CurrencyCode]entities.Currency
	}

	tableTests := []struct {
		name      string
		runBefore func(*pgxpool.Pool)
		want
		args []vos.CurrencyCode
	}{
		{
			name: "should get currencies",
			runBefore: func(pool *pgxpool.Pool) {
				insertTestCur(t, pool, extensions.CurrencyFixtures[0])
				insertTestCur(t, pool, extensions.CurrencyFixtures[1])
			},
			args: []vos.CurrencyCode{
				extensions.CurrencyFixtures[0].Code,
				extensions.CurrencyFixtures[1].Code,
			},
			want: want{
				output: map[vos.CurrencyCode]entities.Currency{
					extensions.CurrencyFixtures[0].Code: extensions.CurrencyFixtures[0],
					extensions.CurrencyFixtures[1].Code: extensions.CurrencyFixtures[1],
				},
			},
		},
		{
			name:      "should return ErrNotFound when currency code does not exist",
			runBefore: func(pool *pgxpool.Pool) {},
			args: []vos.CurrencyCode{
				"I DON'T EXIST",
			},
			want: want{
				output: nil,
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
			got, err := testRepo.GetCurrenciesByCode(testContext, tt.args...)
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
