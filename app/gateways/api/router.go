package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/handlers/currencies"
)

func NewRouter(currencyUC currency.UseCase) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PATCH", "POST", "PUT"},
		AllowedHeaders: []string{"X-Requested-With", "Origin", "Content-Type", "Authorization"},
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(corsOptions))

	currencyHandler := currencies.NewHandler(currencyUC)
	r.Put("/currencies", currencyHandler.UpsertCurrency)
	r.Get("/currencies/conversion", currencyHandler.GetConversion)

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "swagger/index.html", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
