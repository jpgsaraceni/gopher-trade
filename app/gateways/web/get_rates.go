package web

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/jpgsaraceni/gopher-trade/app/domain/currency"
	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
	"github.com/jpgsaraceni/gopher-trade/extensions"
)

func (c Client) GetRates(ctx context.Context) (vos.DefaultRates, error) {
	const operation = "Web.GetRates"

	output := make(vos.DefaultRates, 0)
	g := new(errgroup.Group)

	g.Go(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, exchangeRateURL, nil)
		if err != nil {
			return err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return extensions.ErrStack(operation, currency.ErrCurrencyAPI)
		}

		var erResponse ExchangeRateClientResponse
		err = json.NewDecoder(resp.Body).Decode(&erResponse)
		if err != nil {
			return err
		}
		output[vos.BRL] = erResponse.Rates.BRL
		output[vos.BTC] = erResponse.Rates.BTC
		output[vos.EUR] = erResponse.Rates.EUR

		return nil
	})

	g.Go(func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, cryptoCompareURL, nil)
		if err != nil {
			return err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return extensions.ErrStack(operation, currency.ErrCurrencyAPI)
		}

		var ccResponse CryptoCompareClientResponse
		err = json.NewDecoder(resp.Body).Decode(&ccResponse)
		if err != nil {
			return err
		}
		output[vos.ETH] = ccResponse.ETH

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, extensions.ErrStack(operation, err)
	}

	return output, nil
}
