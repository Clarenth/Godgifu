package services

import "github.com/labstack/echo/v4"

type AccountService interface {
	CreateAccount(ctx echo.Context) (err error)
}
