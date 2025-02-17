package handler

import (
	"wsserver/internal/service"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.Service
}

type Handler interface {
	Login(c *fiber.Ctx) error
	GetPrice(c *websocket.Conn) error
}

func New(s service.Service) Handler {
	return &handler{
		service: s,
	}
}
