package routes

import (
	"godgifu/modules/account/handlers"

	"godgifu/modules/auth"
	mw "godgifu/modules/auth/middleware"

	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"
)

func InitRoutes(server *echo.Echo, apiEndpoint string, accountHandlers handlers.AccountHandlers, auth auth.Auth) {
	apiRoutes := server.Group(apiEndpoint)
	{
		accountRoutes := apiRoutes.Group("/account")
		{
			accountRoutes.GET("/", accountHandlers.GetAccount, mw.AuthAccount(auth.JWTServices), echoMW.AddTrailingSlash())
			accountRoutes.DELETE("/", accountHandlers.DeleteAccount, mw.AuthAccount(auth.JWTServices), echoMW.AddTrailingSlash() /*mw.AuthAdmin(auth.JWTServices)*/)
			// accountRoutes.PATCH("/", accountHandlers.EditAccount, mw.AuthAccount(auth.JWTServices), echoMW.AddTrailingSlash())
			editRoutes := accountRoutes.Group("/edit")
			{
				editRoutes.PATCH("/employee", accountHandlers.UpdateEmployee, mw.AuthAdmin(auth.JWTServices))
				editRoutes.PATCH("/identity", accountHandlers.UpdateIdentity, mw.AuthAccount(auth.JWTServices))
			}

			accountRoutes.POST("/signup", accountHandlers.CreateAccount)
		}
	}
}
