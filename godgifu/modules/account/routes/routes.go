package routes

import (
	"net/http"

	"godgifu/modules/account/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo, handlers handlers.AccountHandlers) {
	// Setup routes
	// authRoutes := server.Group("/account")
	// {
	// 	authRoutes.POST("/signup", handlers.Signup)
	// }

	server.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}
