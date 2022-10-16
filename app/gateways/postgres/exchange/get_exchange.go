package exchange

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) GetExchangeByCurrencies(ctx context.Context, from, to vos.CurrencyCode) (entities.Exchange, error) {
	const operation = "Repository.Exchange.GetExchangeByCurrencies"

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
		WHERE "from" = $1 AND "to" = $2;
	`
	var exc entities.Exchange

	err := r.pool.QueryRow(ctx, query, from.String(), to.String()).Scan(
		&exc.ID,
		&exc.From,
		&exc.To,
		&exc.CreatedAt,
		&exc.UpdatedAt,
		&exc.Rate,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entities.Exchange{}, extensions.ErrStack(operation, exchange.ErrNotFound)
		}

		return entities.Exchange{}, extensions.ErrStack(operation, err)
	}
	exc.CreatedAt = exc.CreatedAt.UTC()
	exc.UpdatedAt = exc.UpdatedAt.UTC()

	return exc, nil
}
