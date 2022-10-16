package currencies

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/responses"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

type CreateCurrencyRequest struct {
	Code    string `json:"code" example:"FAKEMONEY"`
	USDRate string `json:"usd_rate" example:"200.132"`
}

type CreateCurrencyResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id" example:"2171f348-54b4-4a1e-8643-0972a3daf400"`
	Code      string    `json:"code" example:"USD"`
	USDRate   string    `json:"usd_rate" example:"2.132"`
}

var errMissingFields = errors.New("missing required fields")

// @Summary Create a new currency exchange rate
// @Description Creates an exchange rate from a specified currency to USD.
// @Tags Currency
// @Accept json
// @Produce json
// @Param currency body CreateCurrencyRequest true "Currency Info"
// @Success 201 {object} CreateCurrencyResponse
// @Failure 400 {object} responses.ErrorPayload
// @Failure 409 {object} responses.ErrorPayload
// @Failure 500 {object} responses.ErrorPayload
// @Router /currencies [post]
func (h Handler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	const operation = "Handler.Exchanges.CreateCurrency"

	var req CreateCurrencyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = extensions.ErrStack(operation, err)
		responses.BadRequest(w, responses.ErrMalformedBody, err)

		return
	}
	switch "" {
	case req.Code, req.USDRate:
		err := extensions.ErrStack(operation, errMissingFields)
		responses.BadRequest(w, responses.ErrMissingFields, err)

		return
	}
	rate, err := decimal.NewFromString(req.USDRate)
	if err != nil {
		err = extensions.ErrStack(operation, err)
		responses.BadRequest(w, responses.ErrInvalidRate, err)

		return
	}

	input := currency.CreateCurrencyInput{
		Code:    vos.CurrencyCode(req.Code),
		USDRate: rate,
	}

	output, err := h.uc.CreateCurrency(r.Context(), input)
	if err != nil {
		err = extensions.ErrStack(operation, err)

		if errors.Is(err, currency.ErrConflict) {
			responses.Conflict(w, responses.ErrConflictCode, err)

			return
		}

		responses.InternalServerError(w, err)

		return
	}
	res := CreateCurrencyResponse{
		ID:        output.Currency.ID,
		Code:      output.Currency.Code.String(),
		USDRate:   output.Currency.USDRate.String(),
		CreatedAt: output.Currency.CreatedAt,
		UpdatedAt: output.Currency.UpdatedAt,
	}

	responses.Created(w, res)
}
