package routes

import (
	"godgifu/modules/video/db/postgres"
	"godgifu/modules/video/handlers"
	"godgifu/modules/video/services"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo, db *sqlx.DB) {
	// Initialise layers
	repository := postgres.NewPostgresDB(&sqlx.DB{})

	services := services.NewVideoServices(repository)

	handlers := handlers.NewVideoHandlers(services)

	// Setup routes
	server.Group("/test1")
	{
		videoRoutes := server.Group("/video")
		{
			videoRoutes.POST("/test", handlers.NewVideo)
		}
	}
}
