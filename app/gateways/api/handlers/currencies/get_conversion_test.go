package currencies_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain"
	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/handlers/currencies"
)

func Test_Handler_GetConversion(t *testing.T) {
	t.Parallel()

	const target = "/currencys/conversion"

	tests := []struct {
		name       string
		uc         func(t *testing.T) domain.Currency
		params     map[string]string
		wantBody   json.RawMessage
		wantStatus int
	}{
		{
			name: "should get a conversion",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					ConvertFunc: func(ctx context.Context, input currency.ConvertInput) (currency.ConvertOutput, error) {
						inputAmount, err := decimal.NewFromString("1.23")
						assert.NoError(t, err)
						convertedAmount, err := decimal.NewFromString("10.12")
						assert.NoError(t, err)
						assert.Equal(t, "USD", input.From.String())
						assert.Equal(t, "BRL", input.To.String())
						assert.Equal(t, inputAmount, input.FromAmount)

						return currency.ConvertOutput{
							ConvertedAmount: convertedAmount,
						}, nil
					},
				}
			},
			params: map[string]string{
				"from":   "USD",
				"to":     "BRL",
				"amount": "1.23",
			},
			wantBody: json.RawMessage(`{
				"converted_amount":"10.12"
			}`),
			wantStatus: http.StatusOK,
		},
		{
			name: "should return 400 when missing params",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{}
			},
			params: map[string]string{},
			wantBody: json.RawMessage(`{
				"error":"Missing required params."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 400 when amount params is not a number",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{}
			},
			params: map[string]string{
				"from":   "USD",
				"to":     "BRL",
				"amount": "NaN",
			},
			wantBody: json.RawMessage(`{
				"error":"Invalid amount. Must be an integer or point separated decimal number."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 404 when currencies pair does not exist",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					ConvertFunc: func(ctx context.Context, input currency.ConvertInput) (currency.ConvertOutput, error) {
						return currency.ConvertOutput{}, currency.ErrNotFound
					},
				}
			},
			params: map[string]string{
				"from":   "USD",
				"to":     "BRL",
				"amount": "1.23",
			},
			wantBody: json.RawMessage(`{
				"error":"Currency pair not found."
			}`),
			wantStatus: http.StatusNotFound,
		},
		{
			name: "should return 500 when use case returns unexpected error",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					ConvertFunc: func(ctx context.Context, input currency.ConvertInput) (currency.ConvertOutput, error) {
						return currency.ConvertOutput{}, fmt.Errorf("uh oh in use case")
					},
				}
			},
			params: map[string]string{
				"from":   "USD",
				"to":     "BRL",
				"amount": "1.23",
			},
			wantBody: json.RawMessage(`{
				"error":"Internal server error."
			}`),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := currencies.NewHandler(tt.uc(t))

			req := newTestGetRequest(t, target, tt.params)
			res := newTestGetResponse(h.GetConversion, req, target)
			assertResponse(t, tt.wantStatus, tt.wantBody, res)
		})
	}
}
