package exchanges

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
	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
)

func Test_Handler_CreateExchange(t *testing.T) {
	t.Parallel()

	const target = "/exchanges"

	tests := []struct {
		name       string
		uc         domain.Exchange
		args       CreateExchangeRequest
		wantBody   json.RawMessage
		wantStatus int
	}{
		{
			name: "should create exchange receiving decimal rate",
			uc: &domain.ExchangeMock{
				CreateExchangeFunc: func(
					ctx context.Context,
					input exchange.CreateExchangeInput,
				) (exchange.CreateExchangeOutput, error) {
					rate, err := decimal.NewFromString("1.234")
					assert.NoError(t, err)
					assert.Equal(t, exchange.CreateExchangeInput{
						From: "USD",
						To:   "BRL",
						Rate: rate,
					}, input)

					return exchange.CreateExchangeOutput{
						Exc: entities.Exchange{
							ID:        "b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
							From:      "USD",
							To:        "BRL",
							Rate:      rate,
							CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
						},
					}, nil
				},
			},
			args: CreateExchangeRequest{
				From: "USD",
				To:   "BRL",
				Rate: "1.234",
			},
			wantBody: json.RawMessage(`{
				"id":"b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
				"from":"USD",
				"to":"BRL",
				"rate":"1.234",
				"created_at":"2010-01-10T10:00:00Z",
				"updated_at":"2010-01-10T10:00:00Z"
			}`),
			wantStatus: http.StatusCreated,
		},
		{
			name: "should create exchange receiving integer rate",
			uc: &domain.ExchangeMock{
				CreateExchangeFunc: func(
					ctx context.Context,
					input exchange.CreateExchangeInput,
				) (exchange.CreateExchangeOutput, error) {
					rate, err := decimal.NewFromString("5")
					assert.NoError(t, err)
					assert.Equal(t, exchange.CreateExchangeInput{
						From: "USD",
						To:   "BRL",
						Rate: rate,
					}, input)

					return exchange.CreateExchangeOutput{
						Exc: entities.Exchange{
							ID:        "b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
							From:      "USD",
							To:        "BRL",
							Rate:      rate,
							CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
						},
					}, nil
				},
			},
			args: CreateExchangeRequest{
				From: "USD",
				To:   "BRL",
				Rate: "5",
			},
			wantBody: json.RawMessage(`{
				"id":"b94d6cbb-f5b2-4c27-8375-df5dfca13f0b",
				"from":"USD",
				"to":"BRL",
				"rate":"5",
				"created_at":"2010-01-10T10:00:00Z",
				"updated_at":"2010-01-10T10:00:00Z"
			}`),
			wantStatus: http.StatusCreated,
		},
		{
			name: "should return 400 when body is empty",
			uc:   &domain.ExchangeMock{},
			args: CreateExchangeRequest{},
			wantBody: json.RawMessage(`{
				"error":"Missing required fields."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 400 when rate is not a number",
			uc:   &domain.ExchangeMock{},
			args: CreateExchangeRequest{
				From: "some",
				To:   "other",
				Rate: "NaN",
			},
			wantBody: json.RawMessage(`{
				"error":"Invalid rate. Must be an integer or point separated decimal number."
			}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "should return 409 when rate already exists for from-to pair",
			uc: &domain.ExchangeMock{
				CreateExchangeFunc: func(
					ctx context.Context,
					input exchange.CreateExchangeInput,
				) (exchange.CreateExchangeOutput, error) {
					return exchange.CreateExchangeOutput{}, fmt.Errorf("repo error: %w", exchange.ErrConflict)
				},
			},
			args: CreateExchangeRequest{
				From: "some",
				To:   "other",
				Rate: "100",
			},
			wantBody: json.RawMessage(`{
				"error":"Rate for from-to currency pair already exists."
			}`),
			wantStatus: http.StatusConflict,
		},
		{
			name: "should return 500 when something goes wrong in use case",
			uc: &domain.ExchangeMock{
				CreateExchangeFunc: func(
					ctx context.Context,
					input exchange.CreateExchangeInput,
				) (exchange.CreateExchangeOutput, error) {
					return exchange.CreateExchangeOutput{}, fmt.Errorf("uh oh in use case")
				},
			},
			args: CreateExchangeRequest{
				From: "some",
				To:   "other",
				Rate: "100",
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
			h := NewHandler(tt.uc)

			req := newTestPostRequest(t, target, tt.args)
			res := newTestPostResponse(h.CreateExchange, req, target)
			assertResponse(t, tt.wantStatus, tt.wantBody, res)
		})
	}
}
