package account

import (
	"godgifu/config"
	"godgifu/modules/account/db/postgres"
	"godgifu/modules/account/handlers"
	"godgifu/modules/account/routes"
	"godgifu/modules/account/services"
	"godgifu/modules/auth"
)

// InitAccount initialises the Account module, loading the handler routes, services configuration,
// and Database configuration.
func InitAccount(server *config.DevConfiguration, auth auth.Auth) {
	postgresDB := postgres.NewPostgresDB(server.Postgres)
	// redisDB := tokenDB.NewRedisTokenRepository(server.Redis)
	services := services.NewAccountServices(postgresDB)
	accountHandlers := handlers.NewAccountHandlers(services)

	// jwtHandlers := authHandlers.NewJWTHandlers(auth, jwt)

	routes.InitRoutes(server.Router, accountHandlers, auth)
}
