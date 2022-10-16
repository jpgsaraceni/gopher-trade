package currencypg_test

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/currencypg"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func Test_Repository_CreateCurrency(t *testing.T) {
	t.Parallel()

	tableTests := []struct {
		name      string
		runBefore func(*pgxpool.Pool)
		wantErr   error
		input     entities.Currency
	}{
		{
			name:      "should create exchange",
			runBefore: func(*pgxpool.Pool) {},
			input:     extensions.CurrencyFixtures[0],
			wantErr:   nil,
		},
		{
			name:      "should return ErrConflict when currency code already exists",
			runBefore: func(pool *pgxpool.Pool) { insertTestCur(t, pool, extensions.CurrencyFixtures[0]) },
			input: entities.Currency{
				ID:        "7c35948d-f7c3-4995-a5c0-84e2700a954f",
				Code:      extensions.CurrencyFixtures[0].Code,
				CreatedAt: extensions.CurrencyFixtures[0].CreatedAt,
				UpdatedAt: extensions.CurrencyFixtures[0].UpdatedAt,
				USDRate:   decimal.NewFromFloat(1),
			},
			wantErr: currency.ErrConflict,
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
			err := testRepo.CreateCurrency(testContext, tt.input)
			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assertInsertedCurrency(t, testPool, tt.input)
			}
		})
	}
}
