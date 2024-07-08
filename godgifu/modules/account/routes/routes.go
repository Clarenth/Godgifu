package routes

import (
	"godgifu/config"
	"godgifu/modules/account/handlers"

	"godgifu/modules/auth"
	"godgifu/modules/auth/middleware"
)

func InitRoutes(server *config.DevConfiguration, accountHandlers handlers.AccountHandlers, auth auth.Auth) {
	accountRoutes := server.Echo.Group("/account")
	{
		accountRoutes.GET("/", accountHandlers.GetAccount, middleware.AuthAccount(auth.JWTServices))
		accountRoutes.DELETE("/", accountHandlers.DeleteAccount, middleware.AuthAccount(auth.JWTServices))
	}
}
