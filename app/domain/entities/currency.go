package entities

import (
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/jpgsaraceni/gopher-trade/app/domain/vos"
)

const defaultDecimalPlaces = 5

type Currency struct {
	ID        string
	Code      vos.CurrencyCode
	CreatedAt time.Time
	UpdatedAt time.Time
	USDRate   decimal.Decimal
}

// NewCurrency generates an ID (UUID) and timestamps and returns an instance of Currency.
func NewCurrency(code vos.CurrencyCode, usdRate decimal.Decimal) Currency {
	now := time.Now().UTC()

	return Currency{
		ID:        uuid.NewString(),
		Code:      code,
		CreatedAt: now,
		UpdatedAt: now,
		USDRate:   usdRate,
	}
}

// Convert converts amount in original currency to dollars then to target currency
// using their USD rates.
func Convert(originalRate, targetRate, amount decimal.Decimal) decimal.Decimal {
	var decimalPlaces int
	dp, err := strconv.Atoi(os.Getenv("DECIMAL_PLACES"))
	if err == nil {
		decimalPlaces = dp
	}
	if decimalPlaces == 0 {
		decimalPlaces = defaultDecimalPlaces
	}

	originalAmountInUSD := amount.Div(originalRate)
	originalAmountInTarget := originalAmountInUSD.Mul(targetRate).Round(int32(decimalPlaces))

	return originalAmountInTarget
}

func IsDefaultRate(code vos.CurrencyCode) bool {
	defaultRates := map[vos.CurrencyCode]struct{}{
		vos.BRL: {},
		vos.BTC: {},
		vos.ETH: {},
		vos.EUR: {},
		vos.USD: {},
	}

	_, ok := defaultRates[code]

	return ok
}
