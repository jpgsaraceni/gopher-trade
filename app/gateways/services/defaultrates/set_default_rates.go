package defaultrates

import (
	"context"
	"log"

	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (s Service) SetDefaultRates(ctx context.Context) error {
	const operation = "Services.Currency.SetDefaultRates"

	rates, err := s.Client.GetRates(ctx)
	if err != nil {
		return extensions.ErrStack(operation, err)
	}
	err = s.Cache.SetRate(ctx, rates)
	if err != nil {
		return extensions.ErrStack(operation, err)
	}
	log.Printf("set rates: %v\n", rates)

	return nil
}
