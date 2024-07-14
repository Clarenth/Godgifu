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
		accountRoutes := authRoutes.Group("/account")
		{
			accountRoutes.POST("/signin", authHandlers.Signin)
			// accountRoutes.POST("/signup", authHandlers.Signup)
			accountRoutes.POST("/signout", authHandlers.Signout, middleware.AuthAccount(jwtServices))
		}

		authRoutes.GET("/refresh", jwtHandlers.NewTokens, middleware.AuthAccount(jwtServices))
	}
}
