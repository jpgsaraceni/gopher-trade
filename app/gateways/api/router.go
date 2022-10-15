package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/handlers/exchanges"
)

func NewRouter(exchangeUC exchange.UseCase) http.Handler {
	exchangeHandler := exchanges.NewHandler(exchangeUC)

	r := chi.NewRouter()

	r.Post("/exchanges", exchangeHandler.CreateExchange)

	return r
}
