package currencypg

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) GetCurrencyByCode(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
	const operation = "Repository.Currency.GetCurrencyByCode"

	const query = `
		SELECT
			id,
			code,
			created_at,
			updated_at,
			usd_rate
		FROM
			currencies
		WHERE code = $1;
	`

	var cur entities.Currency
	err := r.pool.QueryRow(ctx, query, code).Scan(
		&cur.ID,
		&cur.Code,
		&cur.CreatedAt,
		&cur.UpdatedAt,
		&cur.USDRate,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Currency{}, extensions.ErrStack(operation, currency.ErrNotFound)
		}

		return entities.Currency{}, extensions.ErrStack(operation, err)
	}
	cur.CreatedAt = cur.CreatedAt.UTC()
	cur.UpdatedAt = cur.UpdatedAt.UTC()

	return cur, nil
}
