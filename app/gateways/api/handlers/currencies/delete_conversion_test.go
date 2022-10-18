package currencies_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain"
	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/handlers/currencies"
)

func Test_Handler_DeleteConversion(t *testing.T) {
	t.Parallel()

	const targetTemplate = "/currencies/%s"

	tests := []struct {
		name       string
		uc         func(t *testing.T) domain.Currency
		param      string
		wantBody   json.RawMessage
		wantStatus int
	}{
		{
			name: "should delete a conversion",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					DeleteCurrencyByCodeFunc: func(ctx context.Context, input currency.DeleteCurrencyByCodeInput) error {
						assert.Equal(t, "TEST", input.Code.String())

						return nil
					},
				}
			},
			param:      "TEST",
			wantBody:   nil,
			wantStatus: http.StatusNoContent,
		},
		{
			name: "should return 404 when currency code does not exist",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					DeleteCurrencyByCodeFunc: func(ctx context.Context, input currency.DeleteCurrencyByCodeInput) error {
						assert.Equal(t, "TEST", input.Code.String())

						return currency.ErrNotFound
					},
				}
			},
			param: "TEST",
			wantBody: json.RawMessage(`{
				"error":"Currency code not found."
			}`),
			wantStatus: http.StatusNotFound,
		},
		{
			name: "should return 422 when currency code is of a default currency",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					DeleteCurrencyByCodeFunc: func(ctx context.Context, input currency.DeleteCurrencyByCodeInput) error {
						assert.Equal(t, "TEST", input.Code.String())

						return currency.ErrDefaultRate
					},
				}
			},
			param: "TEST",
			wantBody: json.RawMessage(`{
				"error":"Code belongs to a default rate."
			}`),
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return 500 when use case returns unexpected error",
			uc: func(t *testing.T) domain.Currency {
				return &domain.CurrencyMock{
					DeleteCurrencyByCodeFunc: func(ctx context.Context, input currency.DeleteCurrencyByCodeInput) error {
						assert.Equal(t, "TEST", input.Code.String())

						return fmt.Errorf("uh oh in use case")
					},
				}
			},
			param: "TEST",
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

			target := fmt.Sprintf(targetTemplate, tt.param)
			req, err := http.NewRequestWithContext(testContext, http.MethodDelete, target, nil)
			assert.NoError(t, err)

			router := chi.NewRouter()
			target = fmt.Sprintf(targetTemplate, "{currency-code}")
			router.Delete(target, h.DeleteCurrencyByCode)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assertResponse(t, tt.wantStatus, tt.wantBody, res)
		})
	}
}
