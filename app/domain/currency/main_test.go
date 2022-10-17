package currency_test

import (
	"context"
	"fmt"
)

var (
	testContext       = context.Background()
	testErrRepository = fmt.Errorf("uh oh in repository") //nolint
)
