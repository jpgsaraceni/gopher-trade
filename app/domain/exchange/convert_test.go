package exchange_test

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

func Test_UseCase_Convert(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output exchange.ConvertOutput
	}

	tableTests := []struct {
		name   string
		fields func(t *testing.T) exchange.Repository
		input  exchange.ConvertInput
		want
	}{
		{
			name: "should return conversion",
			fields: func(t *testing.T) exchange.Repository {
				return &exchange.RepositoryMock{
					GetExchangeByCurrenciesFunc: func(ctx context.Context, from, to vos.CurrencyCode) (entities.Exchange, error) {
						assert.Equal(t, "USD", from.String())
						assert.Equal(t, "BRL", to.String())

						return testExc01, nil
					},
				}
			},
			input: exchange.ConvertInput{
				From:       "USD",
				To:         "BRL",
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: exchange.ConvertOutput{
					decimal.NewFromFloat(2.777625),
				},
			},
		},
		{
			name: "should return repository error",
			fields: func(t *testing.T) exchange.Repository {
				return &exchange.RepositoryMock{
					GetExchangeByCurrenciesFunc: func(ctx context.Context, from, to vos.CurrencyCode) (entities.Exchange, error) {
						assert.Equal(t, "USD", from.String())
						assert.Equal(t, "BRL", to.String())

						return entities.Exchange{}, testErrRepository
					},
				}
			},
			input: exchange.ConvertInput{
				From:       "USD",
				To:         "BRL",
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: exchange.ConvertOutput{},
				err:    testErrRepository,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := exchange.NewUseCase(tt.fields(t))

			got, err := uc.Convert(testContext, tt.input)
			assert.ErrorIs(t, err, tt.want.err, got)
			assert.Equal(t, tt.want.output, got)
		})
	}
}
