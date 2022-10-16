package currencypg

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) CreateCurrency(ctx context.Context, cur entities.Currency) error {
	const operation = "Repository.Currency.CreateCurrency"

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

	_, err := r.pool.Exec(
		ctx,
		query,
		cur.ID,
		cur.Code.String(),
		cur.CreatedAt.UTC(),
		cur.UpdatedAt.UTC(),
		cur.USDRate,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.SQLState() == uniqueKeyViolationCode && pgErr.ConstraintName == currenciesCodeConstraint {
				return extensions.ErrStack(operation, currency.ErrConflict)
			}
		}

		return extensions.ErrStack(operation, err)
	}

	return nil
}
