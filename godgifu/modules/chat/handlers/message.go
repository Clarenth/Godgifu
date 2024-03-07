package handlers

import (
	"godgifu/modules/chat/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	MessageService services.MessageService
}

func NewMessageHandlers(service services.MessageService) MessageHandlers {
	return &handler{
		MessageService: service,
	}
}

type Message struct {
	Field1 string `json:"field1"`
}

func (handler *handler) NewMessage(ctx echo.Context) error {
	// message := &Message{
	// 	Field1: "Hello, World!",
	// }
	return ctx.String(http.StatusOK, "Hello, World!")
}
