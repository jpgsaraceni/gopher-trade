package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/api"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/currencypg"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/web"
	"github.com/jpgsaraceni/gopher-trade/docs"
)

const (
	defaultTimeout = time.Minute
	graceTime      = 5 * time.Second
)

// @title Gopher Trade API
// @version 0.1.0
// @description Gopher Trade is an api to get monetary exchange values.

// @contact.name João Saraceni
// @contact.url https://www.linkedin.com/in/joaosaraceni/
// @contact.email jpgome@id.uff.br

// @license.name MIT
// @license.url https://github.com/jpgsaraceni/gopher-trade/blob/main/LICENSE
func main() {
	ctx := context.Background()
	// allow graceful shutdown of goroutines.
	appShutdown := &sync.WaitGroup{}

	// connect to db
	pgPool, err := postgres.ConnectPool(
		ctx,
		"postgres://postgres:postgres@gopher_db:5432/postgres?sslmode=disable",
	) // TODO: move to config
	if err != nil {
		log.Panic(err)
	}
	defer pgPool.Close()

	// inject dependencies
	currencyClient := web.NewClient()
	currencyRepo := currencypg.NewRepository(pgPool)
	currencyUC := currency.NewUseCase(currencyRepo, currencyClient)

	// build http router
	router := api.NewRouter(currencyUC)
	startAPI(ctx, appShutdown, router)

	// wait for graceful shutdown
	appShutdown.Wait()
}

func startAPI(ctx context.Context, shutdown *sync.WaitGroup, router http.Handler) {
	address := fmt.Sprintf("0.0.0.0:%s", "3000") // TODO: move to env
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
	}

	docs.SwaggerInfo.Host = address
	// tell the wait group there is a go routine running
	shutdown.Add(1)
	go runServer(shutdown, address, srv)
	go listenForShutdown(ctx, srv)
}

func runServer(shutdown *sync.WaitGroup, address string, srv *http.Server) {
	defer shutdown.Done()
	log.Printf("starting http server on %s", address)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panicf("http server failed to listen and serve %s", err)
	}
}

// listenForShutdown creates a channel to receive signals from OS
// to trigger a graceful shutdown by cancelling context.
func listenForShutdown(ctx context.Context, srv *http.Server) {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)
	// block until there is a signal
	<-shutdownChan
	serverTimeout, cancel := context.WithTimeout(ctx, graceTime)
	defer cancel()
	log.Println("shutting down http server")
	if err := srv.Shutdown(serverTimeout); err != nil {
		log.Panicf("failed to shutdown http server: %s", err)
	}
}
