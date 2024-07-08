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
		accountAuth := authRoutes.Group("/account")
		{
			accountAuth.POST("/signin", authHandlers.Signin)
			accountAuth.POST("/signup", authHandlers.Signup)
			accountAuth.POST("/signout", authHandlers.Signout, middleware.AuthAccount(jwtServices))

			accountAuth.POST("/token", jwtHandlers.RefreshJWT, middleware.AuthAccount(jwtServices))
		}
	}
}
