package entities

import "time"

type Stock struct {
	Symbol    string    `json:"symbol"`     // Stock ticker symbol (e.g., AAPL, TSLA)
	Name      string    `json:"name"`       // Full name of the company
	Price     float64   `json:"price"`      // Current stock price
	UpdatedAt time.Time `json:"updated_at"` // Last update timestamp
}
