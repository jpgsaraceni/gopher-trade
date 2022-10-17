package currencies_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain"
	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/handlers/currencies"
)

func Test_Handler_UpsertCurrency(t *testing.T) {
	t.Parallel()

	const target = "/currencies"

	tests := []struct {
		name       string
		uc         domain.Currency
		args       currencies.CreateCurrencyRequest
		wantBody   json.RawMessage
		wantStatus int
	}{
		{
			name: "should create currency receiving decimal rate",
			uc: &domain.CurrencyMock{
				UpsertCurrencyFunc: func(
					ctx context.Context,
					input currency.CreateCurrencyInput,
				) (currency.CreateCurrencyOutput, error) {
					rate := decimal.NewFromFloat(1.234)
					assert.Equal(t, currency.CreateCurrencyInput{
						Code:    "BRL",
						USDRate: rate,
					}, input)

					return currency.CreateCurrencyOutput{
						Currency: entities.Currency{
							ID:        "b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
							Code:      "BRL",
							USDRate:   rate,
							CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
						},
						IsNew: true,
					}, nil
				},
			},
			args: currencies.CreateCurrencyRequest{
				Code:    "BRL",
				USDRate: "1.234",
			},
			wantBody: json.RawMessage(`{
				"id":"b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
				"code":"BRL",
				"usd_rate":"1.234",
				"created_at":"2010-01-10T10:00:00Z",
				"updated_at":"2010-01-10T10:00:00Z"
			}`),
			wantStatus: http.StatusCreated,
		},
		{
			name: "should create currency receiving integer rate",
			uc: &domain.CurrencyMock{
				UpsertCurrencyFunc: func(
					ctx context.Context,
					input currency.CreateCurrencyInput,
				) (currency.CreateCurrencyOutput, error) {
					rate := decimal.NewFromFloat(5)
					assert.Equal(t, currency.CreateCurrencyInput{
						Code:    "BRL",
						USDRate: rate,
					}, input)

					return currency.CreateCurrencyOutput{
						Currency: entities.Currency{
							ID:        "b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
							Code:      "BRL",
							USDRate:   rate,
							CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
						},
						IsNew: true,
					}, nil
				},
			},
			args: currencies.CreateCurrencyRequest{
				Code:    "BRL",
				USDRate: "5",
			},
			wantBody: json.RawMessage(`{
				"id":"b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
				"code":"BRL",
				"usd_rate":"5",
				"created_at":"2010-01-10T10:00:00Z",
				"updated_at":"2010-01-10T10:00:00Z"
			}`),
			wantStatus: http.StatusCreated,
		},
		{
			name: "should update currency receiving integer rate",
			uc: &domain.CurrencyMock{
				UpsertCurrencyFunc: func(
					ctx context.Context,
					input currency.CreateCurrencyInput,
				) (currency.CreateCurrencyOutput, error) {
					rate := decimal.NewFromFloat(5)
					assert.Equal(t, currency.CreateCurrencyInput{
						Code:    "BRL",
						USDRate: rate,
					}, input)

					return currency.CreateCurrencyOutput{
						Currency: entities.Currency{
							ID:        "b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
							Code:      "BRL",
							USDRate:   rate,
							CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2010, time.January, 11, 10, 0, 0, 0, time.UTC),
						},
						IsNew: false,
					}, nil
				},
			},
			args: currencies.CreateCurrencyRequest{
				Code:    "BRL",
				USDRate: "5",
			},
			wantBody: json.RawMessage(`{
				"id":"b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
				"code":"BRL",
				"usd_rate":"5",
				"created_at":"2010-01-10T10:00:00Z",
				"updated_at":"2010-01-11T10:00:00Z"
			}`),
			wantStatus: http.StatusOK,
		},
		{
			name: "should return 400 when body is empty",
			uc:   &domain.CurrencyMock{},
			args: currencies.CreateCurrencyRequest{},
			wantBody: json.RawMessage(`{
				"error":"Missing required fields."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 400 when rate is not a number",
			uc:   &domain.CurrencyMock{},
			args: currencies.CreateCurrencyRequest{
				Code:    "some",
				USDRate: "NaN",
			},
			wantBody: json.RawMessage(`{
				"error":"Invalid rate. Must be an integer or point separated decimal number."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 422 when rate is a default rate",
			uc: &domain.CurrencyMock{
				UpsertCurrencyFunc: func(ctx context.Context, input currency.CreateCurrencyInput) (currency.CreateCurrencyOutput, error) { //nolint
					return currency.CreateCurrencyOutput{}, currency.ErrDefaultRate
				},
			},
			args: currencies.CreateCurrencyRequest{
				Code:    "BRL",
				USDRate: "0.41",
			},
			wantBody: json.RawMessage(`{
				"error":"Code belongs to a default rate."
			}`),
			wantStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "should return 500 when something goes wrong in use case",
			uc: &domain.CurrencyMock{
				UpsertCurrencyFunc: func(
					ctx context.Context,
					input currency.CreateCurrencyInput,
				) (currency.CreateCurrencyOutput, error) {
					return currency.CreateCurrencyOutput{}, fmt.Errorf("uh oh in use case")
				},
			},
			args: currencies.CreateCurrencyRequest{
				Code:    "some",
				USDRate: "100",
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
			h := currencies.NewHandler(tt.uc)

			req := newTestPutRequest(t, target, tt.args)
			res := newTestPutResponse(h.UpsertCurrency, req, target)
			assertResponse(t, tt.wantStatus, tt.wantBody, res)
		})
	}
}
