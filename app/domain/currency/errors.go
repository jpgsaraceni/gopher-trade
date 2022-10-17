package currency

import "errors"

var (
	ErrConflict    = errors.New("currency code already exists in db")
	ErrNotFound    = errors.New("currency not found")
	ErrCurrencyAPI = errors.New("currency API client")
)
