package currencies

import (
	"github.com/jpgsaraceni/gopher-trade/app/domain"
)

type Handler struct {
	uc domain.Currency
}

func NewHandler(uc domain.Currency) *Handler {
	return &Handler{
		uc: uc,
	}
}
