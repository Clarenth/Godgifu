package handlers

import "github.com/labstack/echo/v4"

type MessageHandlers interface {
	NewMessage(ctx echo.Context) (err error)
}
