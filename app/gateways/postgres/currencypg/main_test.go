package currencypg_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

var testContext = context.Background()

func insertTestCur(t *testing.T, pool *pgxpool.Pool, cur entities.Currency) {
	t.Helper()

	const query = `
	INSERT INTO currencies (
		id,
		code,
		created_at,
		updated_at,
		usd_rate
	)
	VALUES ($1, $2, $3, $4, $5);
`
	_, err := pool.Exec(
		testContext,
		query,
		cur.ID,
		cur.Code.String(),
		cur.CreatedAt.UTC(),
		cur.UpdatedAt.UTC(),
		cur.USDRate,
	)

	assert.NoError(t, err)
}
