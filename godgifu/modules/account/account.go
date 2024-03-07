package account

import (
	"godgifu/modules/account/db/postgres"
	"godgifu/modules/account/handlers"
	"godgifu/modules/account/routes"
	"godgifu/modules/account/services"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Account is a struct returned on the initialisation of the Account module.
// type Account struct {
// 	Handlers handlers.AccountHandlers
// 	Postgres postgres.PostgresDB
// 	// Routes
// 	Services services.AccountService
// }

// InitAccount initialises the Account module, loading the handler routes, services configuration,
// and Database configuration.
func InitAccount(server *echo.Echo, db *sqlx.DB) {
	//
	repository := postgres.NewPostgresDB(&sqlx.DB{})

	//
	services := services.NewAccountServices(repository)

	//
	handlers := handlers.NewAccountHandlers(services)

	routes.InitRoutes(server, handlers)

	// account := Account{
	// 	Handlers: handlers,
	// 	Postgres: repository,
	// }
}
