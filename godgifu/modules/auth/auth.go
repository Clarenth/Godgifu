package auth

import (
	"godgifu/modules/auth/db"
	"godgifu/modules/auth/handlers"
	"godgifu/modules/auth/routes"
	"godgifu/modules/auth/services"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Auth struct {
}

func InitAuth(server *echo.Echo, database *sqlx.DB) {
	repository := db.NewPostgresDB(database)
	services := services.NewAuthServices(repository)
	handlers := handlers.NewAuthHandlers(services)
	routes.InitRoutes(server, handlers)
}
