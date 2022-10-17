package currency_test

import (
	"context"
	"testing"
	"time"

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
				return currency.NewUseCase(
					&currency.RepositoryMock{
						CreateCurrencyFunc: func(ctx context.Context, cur entities.Currency) (entities.Currency, error) {
							return cur, nil
						},
					},
					&currency.ClientMock{},
				)
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
					IsNew: true,
				},
				err: nil,
			},
		},
		{
			name: "should update currency",
			fields: func(t *testing.T) currency.UseCase {
				return currency.NewUseCase(
					&currency.RepositoryMock{
						CreateCurrencyFunc: func(ctx context.Context, cur entities.Currency) (entities.Currency, error) {
							return entities.Currency{
								ID:        cur.ID,
								Code:      cur.Code,
								CreatedAt: cur.CreatedAt,
								UpdatedAt: cur.CreatedAt.Add(time.Hour),
								USDRate:   decimal.NewFromFloat(1.23),
							}, nil
						},
					},
					&currency.ClientMock{},
				)
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
					IsNew: false,
				},
				err: nil,
			},
		},
		{
			name: "should return error when something goes wrong in repository",
			fields: func(t *testing.T) currency.UseCase {
				return currency.NewUseCase(
					&currency.RepositoryMock{
						CreateCurrencyFunc: func(ctx context.Context, cur entities.Currency) (entities.Currency, error) {
							return entities.Currency{}, testErrRepository
						},
					},
					&currency.ClientMock{},
				)
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
			// match randomic values
			tt.want.output.Currency.ID = got.Currency.ID
			tt.want.output.Currency.CreatedAt = got.Currency.CreatedAt
			tt.want.output.Currency.UpdatedAt = got.Currency.UpdatedAt
			assert.Equal(t, tt.output, got)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
