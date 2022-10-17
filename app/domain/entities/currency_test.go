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
		from   decimal.Decimal
		to     decimal.Decimal
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
				from:   decimal.NewFromFloat(1.23),
				to:     decimal.NewFromFloat(0.47),
				amount: decimal.NewFromFloat(10.2),
			},
			want: decimal.NewFromFloat(3.89756),
		},
		{
			name: "should convert successfully when all values are integer",
			args: args{
				from:   decimal.NewFromFloat(2),
				to:     decimal.NewFromFloat(4),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(2),
		},
		{
			name: "should convert successfully when args are integer resulting in decimal",
			args: args{
				from:   decimal.NewFromFloat(4),
				to:     decimal.NewFromFloat(3),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(0.75),
		},
		{
			name: "should convert successfully when only from value and result are decimal",
			args: args{
				from:   decimal.NewFromFloat(1.1),
				to:     decimal.NewFromFloat(2),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(1.81818),
		},
		{
			name: "should convert successfully when only to value and result are decimal",
			args: args{
				from:   decimal.NewFromFloat(2),
				to:     decimal.NewFromFloat(2.4),
				amount: decimal.NewFromFloat(1),
			},
			want: decimal.NewFromFloat(1.2),
		},
		{
			name: "should convert successfully when only amount arg and result are decimal",
			args: args{
				from:   decimal.NewFromFloat(2),
				to:     decimal.NewFromFloat(4),
				amount: decimal.NewFromFloat(2.2),
			},
			want: decimal.NewFromFloat(4.4),
		},
		{
			name: "should convert successfully when to and from have same USD rate",
			args: args{
				from:   decimal.NewFromFloat(1.111111),
				to:     decimal.NewFromFloat(1.111111),
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
