package handlers

import (
	"github.com/labstack/echo/v4"
)

type AuthHandlers interface {
	Signin(ctx echo.Context) (err error)
	Signout(ctx echo.Context) (err error)
	Signup(ctx echo.Context) (err error)
}
