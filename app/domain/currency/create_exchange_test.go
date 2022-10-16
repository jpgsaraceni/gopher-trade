package currency_test

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

func Test_UseCase_CreateCurrency(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output currency.CreateCurrencyOutput
	}

	tableTests := []struct {
		name   string
		fields func(t *testing.T) currency.UseCase
		input  currency.CreateCurrencyInput
		want
	}{
		{
			name: "should create currency",
			fields: func(t *testing.T) currency.UseCase {
				return currency.NewUseCase(&currency.RepositoryMock{
					CreateCurrencyFunc: func(ctx context.Context, cur entities.Currency) error {
						return nil
					},
				})
			},
			input: currency.CreateCurrencyInput{
				Code:    "TEST",
				USDRate: decimal.NewFromFloat(1.23),
			},
			want: want{
				output: currency.CreateCurrencyOutput{
					Currency: entities.Currency{
						Code:    "TEST",
						USDRate: decimal.NewFromFloat(1.23),
					},
				},
				err: nil,
			},
		},
		{
			name: "should return error when something goes wrong in repository",
			fields: func(t *testing.T) currency.UseCase {
				return currency.NewUseCase(&currency.RepositoryMock{
					CreateCurrencyFunc: func(ctx context.Context, cur entities.Currency) error {
						return testErrRepository
					},
				})
			},
			input: currency.CreateCurrencyInput{},
			want: want{
				output: currency.CreateCurrencyOutput{},
				err:    testErrRepository,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := tt.fields(t)

			got, err := uc.CreateCurrency(testContext, tt.input)
			assertCurrency(t, tt.want.output.Currency, got.Currency)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
