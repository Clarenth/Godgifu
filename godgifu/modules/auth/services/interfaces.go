package services

import (
	account "godgifu/modules/account/models"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	CreateAccount(ctx echo.Context, accountData *account.Account) error
	Signin(ctx echo.Context, accountData *account.AccountEmployee) (match bool, err error)
}
