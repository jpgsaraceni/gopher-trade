package exchange

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

var (
	testContext = context.Background()

	testExc01 = entities.Exchange{
		ID:        "346e0d72-990f-4fc9-ae5d-7c90282d8e93",
		From:      "USD",
		To:        "BRL",
		CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
		Rate:      decimal.NewFromFloat(1.2345),
	}
)

func assertInsertedExchange(t *testing.T, pool *pgxpool.Pool) {
	const query = `
		SELECT
			id,
			"from",
			"to",
			created_at,
			updated_at,
			rate
		FROM
			exchanges
		WHERE id = $1;
	`
	var got entities.Exchange

	err := pool.QueryRow(testContext, query, testExc01.ID).Scan(
		&got.ID,
		&got.From,
		&got.To,
		&got.CreatedAt,
		&got.UpdatedAt,
		&got.Rate,
	)

	assert.NoError(t, err)
	assert.Equal(t, testExc01.Rate.String(), got.Rate.String())
	testExc01.Rate = got.Rate // cant compare pointers
	// convert times to utc since they come from db in local time zone
	got.CreatedAt = got.CreatedAt.UTC()
	got.UpdatedAt = got.UpdatedAt.UTC()
	assert.Equal(t, testExc01, got)
}
