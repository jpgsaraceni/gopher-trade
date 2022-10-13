package exchange_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
)

func Test_UseCase_Convert(t *testing.T) {
	t.Parallel()

	tableTests := []struct {
		name   string
		fields func(t *testing.T) exchange.UseCase
		input  exchange.ConvertInput
		want   exchange.ConvertOutput
	}{
		{
			name: "should return conversion",
			fields: func(t *testing.T) exchange.UseCase {
				return exchange.UseCase{}
			},
			input: exchange.ConvertInput{
				Exchange:   testExchange,
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: exchange.ConvertOutput{
				decimal.NewFromFloat(3.375),
			},
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			uc := tt.fields(t)

			got := uc.Convert(testContext, tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
