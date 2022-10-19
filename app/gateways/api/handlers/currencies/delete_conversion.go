package currencies

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/responses"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

// @Summary Get a conversion for an existent rate
// @Tags Exchange
// @Accept json
// @Produce json
// @Param currency-code path string true "Currency code"
// @Success 204 {object} GetConversionResponse
// @Failure 404 {object} responses.ErrorPayload
// @Failure 422 {object} responses.ErrorPayload
// @Failure 500 {object} responses.ErrorPayload
// @Router /currencies/{currency-code} [delete]
func (h Handler) DeleteCurrencyByCode(w http.ResponseWriter, r *http.Request) {
	const operation = "Handler.Currencies.DeleteCurrencyByCode"

	codeParam := chi.URLParam(r, "currency-code")
	code := vos.CurrencyCode(strings.ToUpper(codeParam))

	input := currency.DeleteCurrencyByCodeInput{
		Code: code,
	}

	err := h.uc.DeleteCurrencyByCode(r.Context(), input)
	if err != nil {
		err = extensions.ErrStack(operation, err)

		switch {
		case errors.Is(err, currency.ErrNotFound):
			responses.NotFound(w, responses.ErrCurrencyCodeNotFound, err)
		case errors.Is(err, currency.ErrDefaultRate):
			responses.UnprocessableEntity(w, responses.ErrIsDefaultRate, err)
		default:
			responses.InternalServerError(w, err)
		}

		return
	}

	responses.NoContent(w)
}
