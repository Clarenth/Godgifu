package handlers

import "github.com/labstack/echo/v4"

type AccountHandlers interface {
	GetAccount(ctx echo.Context) error
	DeleteAccount(ctx echo.Context) error
}
