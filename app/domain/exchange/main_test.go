package exchange_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

var (
	testContext = context.Background()
	// base exchange instance to use in tests
	testExchange  = entities.NewExchange("USD", "BRL", decimal.NewFromFloat(1.5))
	errRepository = fmt.Errorf("uh oh in repository")
)

func assertExchange(t *testing.T, want, got entities.Exchange) {
	assert.Equal(t, want.From(), got.From())
	assert.Equal(t, want.To(), got.To())
	assert.Equal(t, want.Rate(), got.Rate())
}
