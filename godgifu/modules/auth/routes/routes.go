package routes

import (
	"godgifu/modules/auth/handlers"
	"godgifu/modules/auth/middleware"
	"godgifu/modules/auth/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo, authHandlers handlers.AuthHandlers, jwtHandlers handlers.JWTHandler, jwtServices services.JWTService) {
	// Setup routes
	authRoutes := server.Group("/auth")
	{
		authRoutes.POST("/signin", authHandlers.Signin)
		authRoutes.POST("/signup", authHandlers.Signup)
		authRoutes.POST("/signout", authHandlers.Signout, middleware.AuthAccount(jwtServices))

		authRoutes.POST("/token", jwtHandlers.RefreshJWT, middleware.AuthAccount(jwtServices))
	}
}
