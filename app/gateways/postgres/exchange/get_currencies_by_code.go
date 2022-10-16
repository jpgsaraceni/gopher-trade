package currencypg

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) GetCurrenciesByCode(
	ctx context.Context,
	codes ...vos.CurrencyCode,
) (map[vos.CurrencyCode]entities.Currency, error) {
	const operation = "Repository.Currency.GetCurrenciesByCode"

	const query = `
		SELECT
			id,
			code,
			created_at,
			updated_at,
			usd_rate
		FROM
			currencies
		WHERE code = ANY($1);
	`
	currencies := make(map[vos.CurrencyCode]entities.Currency, 0)

	rows, err := r.pool.Query(ctx, query, vos.ListOfCodesToString(codes...))
	for rows.Next() {
		var cur entities.Currency
		err = rows.Scan(
			&cur.ID,
			&cur.Code,
			&cur.CreatedAt,
			&cur.UpdatedAt,
			&cur.USDRate,
		)
		if err != nil {
			return nil, extensions.ErrStack(operation, err)
		}
		cur.CreatedAt = cur.CreatedAt.UTC()
		cur.UpdatedAt = cur.UpdatedAt.UTC()
		currencies[cur.Code] = cur
	}
	if len(currencies) != len(codes) {
		return nil, extensions.ErrStack(operation, currency.ErrNotFound)
	}
	if err != nil {
		return nil, extensions.ErrStack(operation, err)
	}

	return currencies, nil
}
