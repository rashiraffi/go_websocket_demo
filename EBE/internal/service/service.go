package service

import (
	"context"
	"externalserver/pkg/stockexchange"
)

type service struct {
	exchange stockexchange.Exchange
}

type Service interface {
	SuscribeStocks(ctx context.Context, symbol string) (<-chan float64, error)
}

func New() Service {
	return &service{
		exchange: stockexchange.New(),
	}
}
