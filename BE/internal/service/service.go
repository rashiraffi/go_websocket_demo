package service

import (
	"context"
	"os"
	"wsserver/internal/model"
)

type service struct {
	model            model.Model
	stockExchangeURL string
}

type Service interface {
	Login(email, pass string) (string, error)
	GetPrice(ctx context.Context, symbol string) (<-chan float64, error)
}

func New(m model.Model) Service {
	return &service{
		model:            m,
		stockExchangeURL: os.Getenv("STOCK_EXCHANGE_URL"),
	}
}
