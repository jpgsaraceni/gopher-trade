package currencypg

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) UpsertCurrency(ctx context.Context, cur entities.Currency) (entities.Currency, error) {
	const operation = "Repository.Currency.UpsertCurrency"

	const query = `
	INSERT INTO currencies (
		id,
		code,
		created_at,
		updated_at,
		usd_rate
	)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (code) 
	DO UPDATE SET 
		updated_at = $4,
		usd_rate = $5
	RETURNING
		id,
		code,
		created_at,
		updated_at,
		usd_rate
	;`
	var returned entities.Currency

	err := r.pool.QueryRow(
		ctx,
		query,
		cur.ID,
		cur.Code.String(),
		cur.CreatedAt.UTC(),
		cur.UpdatedAt.UTC(),
		cur.USDRate,
	).Scan(
		&returned.ID,
		&returned.Code,
		&returned.CreatedAt,
		&returned.UpdatedAt,
		&returned.USDRate,
	)
	if err != nil {
		return entities.Currency{}, extensions.ErrStack(operation, err)
	}
	returned.CreatedAt = returned.CreatedAt.UTC()
	returned.UpdatedAt = returned.UpdatedAt.UTC()

	return returned, nil
}
