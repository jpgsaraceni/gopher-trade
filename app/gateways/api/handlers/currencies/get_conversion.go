package currencies

import (
	"errors"
	"net/http"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/responses"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

type GetConversionResponse struct {
	ConvertedAmount string `json:"converted_amount" example:"23.431"`
}

var errMissingParams = errors.New("missing required query params")

// @Summary Get a conversion for an existent rate
// @Tags Exchange
// @Accept json
// @Produce json
// @Param from query string true "From currency code"
// @Param to query string true "To currency code"
// @Param amount query string true "Amount to be converted"
// @Success 200 {object} GetConversionResponse
// @Failure 400 {object} responses.ErrorPayload
// @Failure 404 {object} responses.ErrorPayload
// @Failure 500 {object} responses.ErrorPayload
// @Router /exchanges [get]
func (h Handler) GetConversion(w http.ResponseWriter, r *http.Request) {
	const operation = "Handler.Exchanges.GetExchange"

	fromParam := r.URL.Query().Get("from")
	toParam := r.URL.Query().Get("to")
	amountParam := r.URL.Query().Get("amount")
	switch "" {
	case fromParam, toParam, amountParam:
		responses.BadRequest(w, responses.ErrMissingParams, errMissingParams)

		return
	}
	amount, err := decimal.NewFromString(amountParam)
	if err != nil {
		responses.BadRequest(w, responses.ErrInvalidAmount, err)

		return
	}
	from := vos.CurrencyCode(strings.ToUpper(fromParam))
	to := vos.CurrencyCode(strings.ToUpper(toParam))

	input := currency.ConvertInput{
		From:       from,
		To:         to,
		FromAmount: amount,
	}

	output, err := h.uc.Convert(r.Context(), input)
	if err != nil {
		err = extensions.ErrStack(operation, err)

		if errors.Is(err, currency.ErrNotFound) {
			responses.NotFound(w, responses.ErrCurrenciesNotFound, err)

			return
		}

		responses.InternalServerError(w, err)

		return
	}
	res := GetConversionResponse{
		ConvertedAmount: output.ConvertedAmount.String(),
	}

	responses.OK(w, res)
}
