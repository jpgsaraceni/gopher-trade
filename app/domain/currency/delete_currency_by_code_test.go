package currency_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

func Test_UseCase_DeleteCurrencyByCode(t *testing.T) {
	t.Parallel()

	type want struct {
		err error
	}

	tableTests := []struct {
		want
		name   string
		fields func(t *testing.T) currency.UseCase
		input  currency.DeleteCurrencyByCodeInput
	}{
		{
			name: "should delete currency",
			fields: func(t *testing.T) currency.UseCase {
				return currency.NewUseCase(
					&currency.RepositoryMock{
						DeleteCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) error {
							assert.Equal(t, vos.CurrencyCode("TEST"), code)

							return nil
						},
					},
					&currency.ClientMock{},
					&currency.CacheMock{},
				)
			},
			input: currency.DeleteCurrencyByCodeInput{
				Code: "TEST",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "should return error when something goes wrong in repository",
			fields: func(t *testing.T) currency.UseCase {
				return currency.NewUseCase(
					&currency.RepositoryMock{
						DeleteCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) error {
							assert.Equal(t, vos.CurrencyCode("TEST"), code)

							return testErrRepository
						},
					},
					&currency.ClientMock{},
					&currency.CacheMock{},
				)
			},
			input: currency.DeleteCurrencyByCodeInput{
				Code: "TEST",
			},
			want: want{
				err: testErrRepository,
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := tt.fields(t)

			err := uc.DeleteCurrencyByCode(testContext, tt.input)
			assert.ErrorIs(t, err, tt.want.err)
		})
	}
}
