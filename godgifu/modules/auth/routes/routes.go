package routes

import (
	"godgifu/modules/auth/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo, handlers handlers.AuthHandlers) {
	// Setup routes
	authRoutes := server.Group("/auth")
	{
		authRoutes.POST("/signup", handlers.Signup)
	}
}
