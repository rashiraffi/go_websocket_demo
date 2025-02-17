package handler

import (
	"context"
	"math"

	"github.com/gofiber/contrib/websocket"
	"go.uber.org/zap"
)

func (h *handler) GetPrice(c *websocket.Conn) error {
	zap.S().Infow("GetPrice function called", "allowed", c.Locals("allowed"))

	stockID := c.Query("stockID")
	if stockID == "" {
		return c.WriteJSON(
			map[string]any{
				"error": "stockID is required",
			},
		)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Subscribe to the stock price changes
	priceChannel, err := h.service.SuscribeStocks(ctx, stockID)
	if err != nil {
		return c.WriteJSON(map[string]any{
			"error": err.Error(),
		})
	}

	zap.S().Infow("Subscription started", "stockID", stockID)

	for {
		select {
		case <-ctx.Done():
			zap.S().Infow("Subscription stopped", "stockID", stockID)
			return nil
		case price, ok := <-priceChannel:
			if !ok {
				zap.S().Infow("Subscription stopped", "stockID", stockID)
				return nil
			}
			if err := c.WriteJSON(map[string]any{"price": math.Round(price*100) / 100}); err != nil {
				zap.S().Warnw("Error writing to WebSocket", "error", err)
				return err
			}
		}
	}
}
