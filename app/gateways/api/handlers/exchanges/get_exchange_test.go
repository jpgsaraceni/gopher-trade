package exchanges

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
)

func Test_Handler_GetConversion(t *testing.T) {
	t.Parallel()

	const target = "/exchanges/conversion"

	tests := []struct {
		name       string
		uc         func(t *testing.T) domain.Exchange
		params     map[string]string
		wantBody   json.RawMessage
		wantStatus int
	}{
		{
			name: "should get a conversion",
			uc: func(t *testing.T) domain.Exchange {
				return &domain.ExchangeMock{
					ConvertFunc: func(ctx context.Context, input exchange.ConvertInput) (exchange.ConvertOutput, error) {
						inputAmount, err := decimal.NewFromString("1.23")
						assert.NoError(t, err)
						convertedAmount, err := decimal.NewFromString("10.12")
						assert.NoError(t, err)
						assert.Equal(t, "USD", input.From.String())
						assert.Equal(t, "BRL", input.To.String())
						assert.Equal(t, inputAmount, input.FromAmount)

						return exchange.ConvertOutput{
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
			uc: func(t *testing.T) domain.Exchange {
				return &domain.ExchangeMock{}
			},
			params: map[string]string{},
			wantBody: json.RawMessage(`{
				"error":"Missing required params."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 400 when amount params is not a number",
			uc: func(t *testing.T) domain.Exchange {
				return &domain.ExchangeMock{}
			},
			params: map[string]string{
				"from":   "USD",
				"to":     "BRL",
				"amount": "NaN",
			},
			wantBody: json.RawMessage(`{
				"error":"Invalid rate. Must be an integer or point separated decimal number."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 404 when from-to currencie pair does not exist",
			uc: func(t *testing.T) domain.Exchange {
				return &domain.ExchangeMock{
					ConvertFunc: func(ctx context.Context, input exchange.ConvertInput) (exchange.ConvertOutput, error) {
						return exchange.ConvertOutput{}, exchange.ErrNotFound
					},
				}
			},
			params: map[string]string{
				"from":   "USD",
				"to":     "BRL",
				"amount": "1.23",
			},
			wantBody: json.RawMessage(`{
				"error":"No conversion found for from-to currencies pair."
			}`),
			wantStatus: http.StatusNotFound,
		},
		{
			name: "should return 500 when use case returns unexpected error",
			uc: func(t *testing.T) domain.Exchange {
				return &domain.ExchangeMock{
					ConvertFunc: func(ctx context.Context, input exchange.ConvertInput) (exchange.ConvertOutput, error) {
						return exchange.ConvertOutput{}, fmt.Errorf("uh oh in use case")
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
			h := NewHandler(tt.uc(t))

			req := newTestGetRequest(t, target, tt.params)
			res := newTestGetResponse(h.GetConversion, req, target)
			assertResponse(t, tt.wantStatus, tt.wantBody, res)
		})
	}
}
