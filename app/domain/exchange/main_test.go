package exchange_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

var (
	testContext = context.Background()
	// base exchange instance to use in tests
	testExc01 = entities.Exchange{
		ID:        "346e0d72-990f-4fc9-ae5d-7c90282d8e93",
		From:      "USD",
		To:        "BRL",
		CreatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2010, time.January, 10, 10, 0, 0, 0, time.UTC),
		Rate:      decimal.NewFromFloat(1.2345),
	}
	errRepository = fmt.Errorf("uh oh in repository")
)

func assertExchange(t *testing.T, want, got entities.Exchange) {
	assert.Equal(t, want.From, got.From)
	assert.Equal(t, want.To, got.To)
	assert.Equal(t, want.Rate, got.Rate)
}
