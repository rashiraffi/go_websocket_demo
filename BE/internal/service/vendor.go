package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func (s *service) fetchStockPrice(ctx context.Context, symbol string) (<-chan float64, error) {

	// PriceMessage represents the expected WebSocket response structure
	type PriceMessage struct {
		Price float64 `json:"price"`
	}

	url := fmt.Sprintf("%s/ws/price?stockID=%s", s.stockExchangeURL, symbol)

	zap.S().Debugw("URL Formed", "URL", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		zap.S().Warnw("Error dialing websocket", "error", err.Error())
		return nil, err
	}

	priceChan := make(chan float64)

	go func() {
		defer close(priceChan)
		defer conn.Close()

		for {
			select {
			case <-ctx.Done():
				zap.S().Infow("Context done, closing connection")
				return
			default:
				_, message, err := conn.ReadMessage()
				if err != nil {
					zap.S().Warnw("Error reading message", err.Error())
					return
				}

				var priceMsg PriceMessage
				if err := json.Unmarshal(message, &priceMsg); err != nil {
					zap.S().Warnw("Error unmarshalling message", err.Error())
					continue
				}

				priceChan <- priceMsg.Price
			}
		}
	}()

	return priceChan, nil

}
