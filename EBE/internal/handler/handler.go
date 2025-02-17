package handler

import (
	"externalserver/internal/service"

	"github.com/gofiber/contrib/websocket"
)

type handler struct {
	service service.Service
}

type Handler interface {
	GetPrice(c *websocket.Conn) error
}

func New(s service.Service) Handler {
	return &handler{
		service: s,
	}
}
