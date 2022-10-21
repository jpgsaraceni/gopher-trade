package defaultrates

import (
	"context"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

//go:generate moq -fmt goimports -out client_mock.go . Client
type Client interface {
	GetRates(ctx context.Context) (vos.DefaultRates, error)
}
