package currencypg_test

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/currencypg"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func Test_Repository_UpsertCurrency(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output entities.Currency
	}

	tableTests := []struct {
		name      string
		runBefore func(*pgxpool.Pool)
		want
		input entities.Currency
	}{
		{
			name:      "should create currency",
			runBefore: func(*pgxpool.Pool) {},
			input:     extensions.CurrencyFixtures[0],
			want: want{
				output: extensions.CurrencyFixtures[0],
			},
		},
		{
			name: "should update currency",
			runBefore: func(pool *pgxpool.Pool) {
				insertTestCur(t, pool, extensions.CurrencyFixtures[0])
			},
			input: entities.Currency{
				ID:        extensions.CurrencyFixtures[0].ID,
				Code:      extensions.CurrencyFixtures[0].Code,
				USDRate:   decimal.NewFromFloat(5.432),
				CreatedAt: extensions.CurrencyFixtures[0].CreatedAt,
				UpdatedAt: time.Date(2022, 10, 10, 10, 10, 10, 0, time.UTC),
			},
			want: want{
				output: entities.Currency{
					ID:        extensions.CurrencyFixtures[0].ID,
					Code:      extensions.CurrencyFixtures[0].Code,
					USDRate:   decimal.NewFromFloat(5.432),
					CreatedAt: extensions.CurrencyFixtures[0].CreatedAt,
					UpdatedAt: time.Date(2022, 10, 10, 10, 10, 10, 0, time.UTC),
				},
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
			got, err := testRepo.UpsertCurrency(testContext, tt.input)
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
