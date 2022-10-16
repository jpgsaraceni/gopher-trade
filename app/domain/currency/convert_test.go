package currency_test

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func Test_UseCase_Convert(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output currency.ConvertOutput
	}

	tableTests := []struct {
		name   string
		fields func(t *testing.T) currency.Repository
		input  currency.ConvertInput
		want
	}{
		{
			name: "should return currencies",
			fields: func(t *testing.T) currency.Repository {
				return &currency.RepositoryMock{
					GetCurrenciesByCodeFunc: func(ctx context.Context, code ...vos.CurrencyCode) (map[vos.CurrencyCode]entities.Currency, error) { //nolint
						assert.Equal(t, "FIXT1", code[0].String())
						assert.Equal(t, "FIXT2", code[1].String())

						return map[vos.CurrencyCode]entities.Currency{
							"FIXT1": extensions.CurrencyFixtures[0], // 1.5
							"FIXT2": extensions.CurrencyFixtures[1], // 2.134
						}, nil
					},
				}
			},
			input: currency.ConvertInput{
				From:       "FIXT1",
				To:         "FIXT2",
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: currency.ConvertOutput{
					ConvertedAmount: decimal.NewFromFloat(1.58154),
				},
			},
		},
		{
			name: "should return repository error",
			fields: func(t *testing.T) currency.Repository {
				return &currency.RepositoryMock{
					GetCurrenciesByCodeFunc: func(ctx context.Context, code ...vos.CurrencyCode) (map[vos.CurrencyCode]entities.Currency, error) { //nolint
						return nil, testErrRepository
					},
				}
			},
			input: currency.ConvertInput{
				From:       "USD",
				To:         "BRL",
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: currency.ConvertOutput{},
				err:    testErrRepository,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := currency.NewUseCase(tt.fields(t))

			got, err := uc.Convert(testContext, tt.input)
			assert.ErrorIs(t, err, tt.want.err, got)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
