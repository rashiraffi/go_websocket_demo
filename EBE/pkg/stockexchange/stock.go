package stockexchange

import (
	"externalserver/internal/entities"
	"time"
)

var (
	stocks map[string]*entities.Stock = map[string]*entities.Stock{
		"HDFCBANK": {
			Symbol:    "HDFCBANK",
			Name:      "HDFC Bank Limited",
			Price:     1717.35,
			UpdatedAt: time.Now(),
		},
		"ICICIBANK": {
			Symbol:    "ICICIBANK",
			Name:      "ICICI Bank Limited",
			Price:     1251.35,
			UpdatedAt: time.Now(),
		},
		"AXISBANK": {
			Symbol:    "AXISBANK",
			Name:      "Axis Bank Limited",
			Price:     992.35,
			UpdatedAt: time.Now(),
		},
	}
)
