package routes

import (
	"godgifu/modules/chat/db/postgres"
	"godgifu/modules/chat/handlers"
	"godgifu/modules/chat/services"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo, db *sqlx.DB) {
	// Initialise layers
	repository := postgres.NewPostgresDB(&sqlx.DB{})

	services := services.NewMessageServices(repository)

	handlers := handlers.NewMessageHandlers(services)

	// Setup routes
	server.Group("/chat")
	{
		messageRoutes := server.Group("/message")
		{
			messageRoutes.GET("/test", handlers.NewMessage)
		}
	}
}
