package service

import "context"

func (s *service) SuscribeStocks(ctx context.Context, symbol string) (<-chan float64, error) {
	return s.exchange.SuscribeStocks(ctx, symbol)
}
