package exchange

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

//go:generate moq -fmt goimports -out repository_mock.go . Repository
type Repository interface {
	CreateExchange(ctx context.Context, exc entities.Exchange) error
}
