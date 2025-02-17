package router

import (
	"wsserver/internal/handler"
	"wsserver/pkg/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New(h handler.Handler) *fiber.App {
	app := fiber.New()
	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                // Allow all
		AllowMethods: "GET,POST,OPTIONS", // Allowed HTTP methods
		AllowHeaders: "Content-Type,Authorization",
	}))
	app.Post("/login", h.Login)
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Use(middleware.JWTAuthentication())
	app.Use("/ws", middleware.WebsocketUpgrade)

	app.Get("/ws/price", websocket.New(func(c *websocket.Conn) {
		h.GetPrice(c)
	}))

	return app
}
