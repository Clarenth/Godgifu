package services

import (
	account "godgifu/modules/account/models"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	CreateAccount(ctx echo.Context, accountData account.AccountEmployeeData) (err error)
}
