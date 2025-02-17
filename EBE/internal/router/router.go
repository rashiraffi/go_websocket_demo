package router

import (
	"externalserver/internal/handler"
	"externalserver/pkg/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func New(h handler.Handler) *fiber.App {
	app := fiber.New()

	app.Use("/ws", middleware.WebsocketUpgrade)

	app.Get("/ws/price", websocket.New(func(c *websocket.Conn) {
		h.GetPrice(c)
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	return app
}
