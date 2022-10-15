package vos

type CurrencyCode string

func (c CurrencyCode) String() string {
	return string(c)
}
