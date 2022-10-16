package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jpgsaraceni/gopher-trade/app/domain/exchange"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api/handlers/exchanges"
)

func NewRouter(exchangeUC exchange.UseCase) http.Handler {
	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"X-Requested-With", "Origin", "Content-Type", "Authorization"},
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(corsOptions))

	exchangeHandler := exchanges.NewHandler(exchangeUC)
	r.Post("/exchanges", exchangeHandler.CreateExchange)
	r.Get("/exchanges/conversion", exchangeHandler.GetConversion)

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "swagger/index.html", http.StatusMovedPermanently)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
