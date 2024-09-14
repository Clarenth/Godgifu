package handlers

import "github.com/labstack/echo/v4"

type AccountHandlers interface {
	CreateAccount(ctx echo.Context) error
	GetAccount(ctx echo.Context) error
	DeleteAccount(ctx echo.Context) error
	UpdateAccount(ctx echo.Context) error
	UpdateEmployee(ctx echo.Context) error
	UpdateIdentity(ctx echo.Context) error
}
