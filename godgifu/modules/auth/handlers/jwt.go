package handlers

import (
	"godgifu/modules/auth/services"

	"github.com/labstack/echo/v4"
)

type jwtHandler struct {
	JWTService services.JWTService
}

func NewJWTHandlers(authServices services.AuthService, jwtServices services.JWTService) JWTHandler {
	return &jwtHandler{
		JWTService: jwtServices,
	}
}

func (handler *jwtHandler) RefreshJWT(ctx echo.Context) error {
	panic("Not done yet")
}
