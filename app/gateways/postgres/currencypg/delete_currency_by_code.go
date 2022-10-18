package currencypg

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (r Repository) DeleteCurrencyByCode(ctx context.Context, code vos.CurrencyCode) error {
	const operation = "Repository.Currency.DeleteCurrencyByCode"

	const query = `
		DELETE FROM currencies
		WHERE code = $1;
	`

	ct, err := r.pool.Exec(ctx, query, code)
	if err != nil {
		return extensions.ErrStack(operation, err)
	}

	if ct.RowsAffected() == 0 {
		return extensions.ErrStack(operation, currency.ErrNotFound)
	}

	return nil
}
