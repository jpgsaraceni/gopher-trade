package entities_test

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

func Test_Domain_Currency_Convert(t *testing.T) {
	t.Parallel()

	type args struct {
		from   entities.Currency
		to     entities.Currency
		amount decimal.Decimal
	}

	tableTests := []struct {
		name string
		args
		want decimal.Decimal
	}{
		{
			name: "should convert successfully when all values are decimals",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(1.23)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(0.47)),
				amount: decimal.NewFromFloat(10.2),
			},
			want: decimal.NewFromFloat(26.69362),
		},
		{
			name: "should convert successfully when all values are integer",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(4)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(2)),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(2),
		},
		{
			name: "should convert successfully when args are integer resulting in decimal",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(4)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(3)),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(1.33333),
		},
		{
			name: "should convert successfully when only from value and result are decimal",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(1.1)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(2)),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(0.55),
		},
		{
			name: "should convert successfully when only to value and result are decimal",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(2)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(2.4)),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(0.83333),
		},
		{
			name: "should convert successfully when only amount arg and result are decimal",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(2)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(4)),
				amount: decimal.NewFromFloat(2.2),
			},
			want: decimal.NewFromFloat(1.1),
		},
		{
			name: "should convert successfully when to and from have same USD rate",
			args: args{
				from:   entities.NewCurrency("FROM", decimal.NewFromFloat(1.111111)),
				to:     entities.NewCurrency("TO", decimal.NewFromFloat(1.111111)),
				amount: decimal.NewFromFloat(2.2),
			},
			want: decimal.NewFromFloat(2.2),
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := entities.Convert(tt.from, tt.to, tt.amount)
			assert.Equal(t, tt.want.String(), got.String())
		})
	}
}
