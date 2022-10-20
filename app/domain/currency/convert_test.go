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

// TODO: cover cache

func Test_UseCase_Convert(t *testing.T) {
	t.Parallel()

	type want struct {
		err    error
		output currency.ConvertOutput
	}

	tableTests := []struct {
		name   string
		fields func(t *testing.T) currency.UseCase
		input  currency.ConvertInput
		want
	}{
		{
			name: "should return custom currencies conversion",
			fields: func(t *testing.T) currency.UseCase {
				return currency.UseCase{
					Repo: &currency.RepositoryMock{
						GetCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
							switch code { //nolint
							case extensions.CurrencyFixtures[0].Code:
								return extensions.CurrencyFixtures[0], nil // 1.5
							case extensions.CurrencyFixtures[1].Code:
								return extensions.CurrencyFixtures[1], nil // 2.134
							default:
								t.FailNow()

								return entities.Currency{}, nil
							}
						},
					},
					Cache: &currency.CacheMock{
						GetRateFunc: func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
							return decimal.Zero, nil
						},
						SetRateFunc: func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
							return nil
						},
					},
				}
			},
			input: currency.ConvertInput{
				From:       extensions.CurrencyFixtures[0].Code,
				To:         extensions.CurrencyFixtures[1].Code,
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: currency.ConvertOutput{
					ConvertedAmount: decimal.NewFromFloat(3.20100),
				},
			},
		},
		{
			name: "should return default currencies conversion",
			fields: func(t *testing.T) currency.UseCase {
				return currency.UseCase{
					Client: &currency.ClientMock{
						GetRatesFunc: func(ctx context.Context) (vos.DefaultRates, error) {
							return vos.DefaultRates{
								vos.BRL: extensions.CurrencyFixtures[0].USDRate, // 1.5
								vos.EUR: extensions.CurrencyFixtures[1].USDRate, // 2.134
							}, nil
						},
					},
					Cache: &currency.CacheMock{
						GetRateFunc: func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
							return decimal.Zero, nil
						},
						SetRateFunc: func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
							return nil
						},
					},
				}
			},
			input: currency.ConvertInput{
				From:       vos.BRL,
				To:         vos.EUR,
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: currency.ConvertOutput{
					ConvertedAmount: decimal.NewFromFloat(3.20100),
				},
			},
		},
		{
			name: "should return conversion from default to custom currency",
			fields: func(t *testing.T) currency.UseCase {
				return currency.UseCase{
					Repo: &currency.RepositoryMock{
						GetCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
							switch code { //nolint
							case extensions.CurrencyFixtures[1].Code:
								return extensions.CurrencyFixtures[1], nil // 2.134
							default:
								t.FailNow()

								return entities.Currency{}, nil
							}
						},
					},
					Client: &currency.ClientMock{
						GetRatesFunc: func(ctx context.Context) (vos.DefaultRates, error) {
							return vos.DefaultRates{
								vos.BRL: extensions.CurrencyFixtures[0].USDRate, // 1.5
								vos.EUR: extensions.CurrencyFixtures[1].USDRate, // 2.134
							}, nil
						},
					},
					Cache: &currency.CacheMock{
						GetRateFunc: func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
							return decimal.Zero, nil
						},
						SetRateFunc: func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
							return nil
						},
					},
				}
			},
			input: currency.ConvertInput{
				From:       vos.BRL,
				To:         extensions.CurrencyFixtures[1].Code,
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: currency.ConvertOutput{
					ConvertedAmount: decimal.NewFromFloat(3.20100),
				},
			},
		},
		{
			name: "should return conversion from custom to default currency",
			fields: func(t *testing.T) currency.UseCase {
				return currency.UseCase{
					Repo: &currency.RepositoryMock{
						GetCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
							switch code { //nolint
							case extensions.CurrencyFixtures[0].Code:
								return extensions.CurrencyFixtures[0], nil // 1.5
							default:
								t.FailNow()

								return entities.Currency{}, nil
							}
						},
					},
					Client: &currency.ClientMock{
						GetRatesFunc: func(ctx context.Context) (vos.DefaultRates, error) {
							return vos.DefaultRates{
								vos.BRL: extensions.CurrencyFixtures[1].USDRate, // 1.5
								vos.EUR: extensions.CurrencyFixtures[0].USDRate, // 2.134
							}, nil
						},
					},
					Cache: &currency.CacheMock{
						GetRateFunc: func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
							return decimal.Zero, nil
						},
						SetRateFunc: func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
							return nil
						},
					},
				}
			},
			input: currency.ConvertInput{
				From:       extensions.CurrencyFixtures[0].Code,
				To:         vos.BRL,
				FromAmount: decimal.NewFromFloat(2.25),
			},
			want: want{
				output: currency.ConvertOutput{
					ConvertedAmount: decimal.NewFromFloat(3.20100),
				},
			},
		},
		{
			name: "should return repository error",
			fields: func(t *testing.T) currency.UseCase {
				return currency.UseCase{
					Repo: &currency.RepositoryMock{
						GetCurrencyByCodeFunc: func(ctx context.Context, code vos.CurrencyCode) (entities.Currency, error) {
							return entities.Currency{}, testErrRepository
						},
					},
					Cache: &currency.CacheMock{
						GetRateFunc: func(ctx context.Context, code vos.CurrencyCode) (decimal.Decimal, error) {
							return decimal.Zero, nil
						},
						SetRateFunc: func(contextMoqParam context.Context, currencyRates map[vos.CurrencyCode]decimal.Decimal) error {
							return nil
						},
					},
				}
			},
			input: currency.ConvertInput{
				From:       "ABC",
				To:         "DEF",
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
			uc := tt.fields(t)

			got, err := uc.Convert(testContext, tt.input)
			assert.ErrorIs(t, err, tt.want.err, got)
			assert.Equal(t, tt.want.output.ConvertedAmount.String(), got.ConvertedAmount.String())
		})
	}
}
