package service

import "context"

func (s *service) GetPrice(ctx context.Context, symbol string) (<-chan float64, error) {
	return s.fetchStockPrice(ctx, symbol)
}
