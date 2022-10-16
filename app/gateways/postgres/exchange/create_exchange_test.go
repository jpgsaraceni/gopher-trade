package exchange

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
)

func Test_Repository_CreateExchange(t *testing.T) {
	t.Parallel()

	tableTests := []struct {
		name      string
		runBefore func(*pgxpool.Pool)
		wantErr   error
		input     entities.Exchange
	}{
		{
			name:      "should create exchange",
			runBefore: func(*pgxpool.Pool) {},
			input:     testExc01,
			wantErr:   nil,
		},
		{
			name:      "should return ErrConflict when from-to pair already exists",
			runBefore: func(pool *pgxpool.Pool) { insertTestExc(t, pool, testExc01) },
			input: entities.Exchange{
				ID:        "9af120fa-3d86-4c88-8c3e-7d8003eed92a",
				From:      testExc01.From,
				To:        testExc01.To,
				Rate:      decimal.NewFromFloat(1.2),
				CreatedAt: time.Date(2010, time.January, 11, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2010, time.January, 11, 10, 0, 0, 0, time.UTC),
			},
			wantErr: exchange.ErrConflict,
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
			err := testRepo.CreateExchange(testContext, tt.input)
			assert.ErrorIs(t, err, tt.wantErr)
			assertInsertedExchange(t, testPool)
		})
	}
}
