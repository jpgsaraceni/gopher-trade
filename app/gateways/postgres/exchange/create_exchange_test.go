package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jpgsaraceni/gopher-trade/app/domain/entities"
	"github.com/jpgsaraceni/gopher-trade/app/gateways/postgres/postgrestest"
)

func Test_Repository_CreateExchange(t *testing.T) {
	t.Parallel()

	testPool, tearDown := postgrestest.GetTestPool()
	testRepo := NewRepository(testPool)

	t.Cleanup(tearDown)

	tableTests := []struct {
		name    string
		wantErr error
		input   entities.Exchange
	}{
		{
			name:    "should create exchange",
			input:   testExc01,
			wantErr: nil,
		},
	}

	for _, tt := range tableTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := testRepo.CreateExchange(testContext, tt.input)
			assert.ErrorIs(t, err, tt.wantErr)
			assertInsertedExchange(t, testPool)
		})
	}
}
