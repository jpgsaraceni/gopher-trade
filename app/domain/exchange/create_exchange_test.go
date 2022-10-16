package exchange_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
)

func Test_UseCase_CreateExchange(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output exchange.CreateExchangeOutput
	}

	tableTests := []struct {
		name   string
		fields func(t *testing.T) exchange.UseCase
		input  exchange.CreateExchangeInput
		want
	}{
		{
			name: "should create exchange",
			fields: func(t *testing.T) exchange.UseCase {
				return exchange.NewUseCase(&exchange.RepositoryMock{
					CreateExchangeFunc: func(ctx context.Context, exc entities.Exchange) error {
						return nil
					},
				})
			},
			input: exchange.CreateExchangeInput{
				From: testExc01.From,
				To:   testExc01.To,
				Rate: testExc01.Rate,
			},
			want: want{
				output: exchange.CreateExchangeOutput{
					Exc: testExc01,
				},
				err: nil,
			},
		},
		{
			name: "should return error when something goes wrong in repository",
			fields: func(t *testing.T) exchange.UseCase {
				return exchange.NewUseCase(&exchange.RepositoryMock{
					CreateExchangeFunc: func(ctx context.Context, exc entities.Exchange) error {
						return testErrRepository
					},
				})
			},
			input: exchange.CreateExchangeInput{},
			want: want{
				output: exchange.CreateExchangeOutput{},
				err:    testErrRepository,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := tt.fields(t)

			got, err := uc.CreateExchange(testContext, tt.input)
			assertExchange(t, tt.want.output.Exc, got.Exc)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
