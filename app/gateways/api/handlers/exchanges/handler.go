package exchanges

import (
	"github.com/jpgsaraceni/gopher-trade/app/domain"
)

type Handler struct {
	uc domain.Exchange
}

func NewHandler(uc domain.Exchange) *Handler {
	return &Handler{
		uc: uc,
	}
}
