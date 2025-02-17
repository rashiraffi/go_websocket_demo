package stockexchange

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func (e *exchange) SuscribeStocks(ctx context.Context, symbol string) (<-chan float64, error) {
	stock, ok := e.stocks[symbol]
	if !ok {
		return nil, fmt.Errorf("stock with symbol %s not found", symbol)
	}

	priceChannel := make(chan float64)

	go func() {
		defer close(priceChannel)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// Simulate stock price changes
				stock.Price += (rand.Float64() - 0.5) * 2 // Random small variation
				priceChannel <- stock.Price
			}
		}
	}()

	return priceChannel, nil
}
