package exchange

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) CreateExchange(ctx context.Context, exc entities.Exchange) error {
	const operation = "Repository.Exchange.CreateExchange"

	const query = `
		INSERT INTO exchanges (
			id,
			"from",
			"to",
			created_at,
			updated_at,
			rate
		)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		exc.ID,
		exc.From.String(),
		exc.To.String(),
		exc.CreatedAt.UTC(),
		exc.UpdatedAt.UTC(),
		exc.Rate,
	)
	if err != nil {
		return extensions.ErrStack(operation, err)
	}

	return nil
}
