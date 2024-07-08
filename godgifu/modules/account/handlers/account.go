package handlers

import (
	"godgifu/modules/account/models"
	"godgifu/modules/account/services"

	"github.com/labstack/echo/v4"
)

type handler struct {
	AccountService services.AccountService
}

func NewAccountHandlers(service services.AccountService) AccountHandlers {
	return &handler{
		AccountService: service,
	}
}

func (handler *handler) GetAccount(ctx echo.Context) error {
	// panic("not done")
	accountID := ctx.Get("account")

	account, err := handler.AccountService.GetAccountData(ctx, accountID.(*models.AccountEmployee).ID)
	if err != nil {
		return echo.NewHTTPError(400, "Debug: An error with your ID code")
	}
	return ctx.JSON(200, account)
}

func (handler *handler) DeleteAccount(ctx echo.Context) error {
	accountID := ctx.Get("account")

	err := handler.AccountService.DeleteAccountData(ctx, accountID.(*models.AccountEmployee).ID)
	if err != nil {
		return echo.NewHTTPError(500, "Account was not deleted")
	}
	return ctx.NoContent(204)
}
