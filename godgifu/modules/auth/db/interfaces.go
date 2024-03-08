package db

import (
	"godgifu/modules/account/models"

	"github.com/labstack/echo/v4"
)

type PostgresDB interface {
	CreateAccount(ctx echo.Context, accountData *models.Account) error
}
