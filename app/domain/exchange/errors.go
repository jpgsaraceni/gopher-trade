package exchange

import "errors"

var ErrConflict = errors.New("exchange rate already exists in db")
