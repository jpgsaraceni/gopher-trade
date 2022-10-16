package exchanges

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/responses"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

type CreateExchangeRequest struct {
	From string `json:"from" example:"USD"`
	To   string `json:"to" example:"COOLCOIN"`
	Rate string `json:"rate" example:"2.132"`
}

type CreateExchangeResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id" example:"2171f348-54b4-4a1e-8643-0972a3daf400"`
	From      string    `json:"from" example:"USD"`
	To        string    `json:"to" example:"COOLCOIN"`
	Rate      string    `json:"rate" example:"2.132"`
}

var errMissingFields = errors.New("missing required fields")

// @Summary Create a new exchange rate
// @Description Creates an exchange rate from and to specified currencies.
// @Description Note that from-to currency pairs must be unique.
// @Tags Exchange
// @Accept json
// @Produce json
// @Param account body CreateExchangeRequest true "Exchange Info"
// @Success 201 {object} CreateExchangeResponse
// @Failure 400 {object} responses.ErrorPayload
// @Failure 500 {object} responses.ErrorPayload
// @Router /exchanges [post]
func (h Handler) CreateExchange(w http.ResponseWriter, r *http.Request) {
	const operation = "Handler.Exchanges.CreateExchange"

	var req CreateExchangeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = extensions.ErrStack(operation, err)
		responses.BadRequest(w, responses.ErrMalformedBody, err)

		return
	}
	switch "" {
	case req.From, req.To, req.Rate:
		err := extensions.ErrStack(operation, errMissingFields)
		responses.BadRequest(w, responses.ErrMissingFields, err)

		return
	}
	rate, err := decimal.NewFromString(req.Rate)
	if err != nil {
		err = extensions.ErrStack(operation, err)
		responses.BadRequest(w, responses.ErrInvalidRate, err)

		return
	}

	input := exchange.CreateExchangeInput{
		From: vos.CurrencyCode(req.From),
		To:   vos.CurrencyCode(req.To),
		Rate: rate,
	}

	output, err := h.uc.CreateExchange(r.Context(), input)
	if err != nil {
		responses.InternalServerError(w, extensions.ErrStack(operation, err))

		return
	}
	res := CreateExchangeResponse{
		ID:        output.Exc.ID,
		From:      output.Exc.From.String(),
		To:        output.Exc.To.String(),
		Rate:      output.Exc.Rate.String(),
		CreatedAt: output.Exc.CreatedAt,
		UpdatedAt: output.Exc.UpdatedAt,
	}

	responses.Created(w, res)
}
