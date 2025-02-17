package stockexchange

import (
	"context"
	"externalserver/internal/entities"
)

type Exchange interface {
	SuscribeStocks(ctx context.Context, symbol string) (<-chan float64, error)
}

type exchange struct {
	stocks map[string]*entities.Stock
}

func New() Exchange {
	return &exchange{
		stocks: stocks,
	}
}
