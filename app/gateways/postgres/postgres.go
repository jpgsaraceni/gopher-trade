package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	errConfigureDB = errors.New("failed to configure db connection")
	errConnectDB   = errors.New("failed to connect to db")
)

// ConnectPool builds a config using the url passed as argument,
// then creates a new pool and connects using that config.
func ConnectPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errConfigureDB, err.Error())
	}

	log.Printf("attempting to connect to postgres on %s...\n", databaseURL)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errConnectDB, err.Error())
	}

	log.Println("successfully connected to postgres server. running migrations...")
	err = Migrate(databaseURL)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
