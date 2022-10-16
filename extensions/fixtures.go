//nolint
package extensions

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

var CurrencyFixtures = []entities.Currency{
	{
		ID:        "04eb1dfe-7e07-490a-bb9c-e8f930851426",
		Code:      "FIXT1",
		CreatedAt: time.Date(2015, time.July, 6, 16, 15, 0, 0, time.UTC),
		UpdatedAt: time.Date(2015, time.July, 6, 16, 15, 0, 0, time.UTC),
		USDRate:   decimal.NewFromFloat(1.5),
	},
	{
		ID:        "a28d4f77-942f-4f0f-8365-957a0da4e594",
		Code:      "FIXT2",
		CreatedAt: time.Date(2015, time.July, 6, 16, 15, 0, 0, time.UTC),
		UpdatedAt: time.Date(2015, time.July, 6, 16, 15, 0, 0, time.UTC),
		USDRate:   decimal.NewFromFloat(2.134),
	},
}
