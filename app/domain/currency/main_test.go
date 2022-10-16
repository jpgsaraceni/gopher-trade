package currency_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
)

var (
	testContext       = context.Background()
	testErrRepository = fmt.Errorf("uh oh in repository") //nolint
)

func assertCurrency(t *testing.T, want, got entities.Currency) {
	t.Helper()

	assert.Equal(t, want.Code, got.Code)
	assert.Equal(t, want.USDRate, got.USDRate)
}
