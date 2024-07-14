package handlers

import (
	"github.com/labstack/echo/v4"
)

type AuthHandlers interface {
	Signin(ctx echo.Context) error
	Signout(ctx echo.Context) error
	// Signup(ctx echo.Context) error
}

type JWTHandler interface {
	NewTokens(ctx echo.Context) error
}

type PASETOHandler interface {
}
