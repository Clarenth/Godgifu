package services

import "github.com/labstack/echo/v4"

type MessageService interface {
	CreateMessage(ctx echo.Context) (err error)
}
