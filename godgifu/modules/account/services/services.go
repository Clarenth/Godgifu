package services

import (
	"godgifu/modules/account/db/postgres"
	"godgifu/modules/account/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type services struct {
	PG postgres.PostgresDB
}

func NewAccountServices(postgres postgres.PostgresDB) AccountService {
	return &services{
		PG: postgres,
	}
}

func (services *services) GetAccountData(ctx echo.Context, accountID uuid.UUID) (*models.Account, error) {
	// panic("not done")
	context := ctx.Request().Context()
	account, err := services.PG.SelectFullAccountData(context, accountID.String())
	if err != nil {
		return nil, err
	}

	return account, nil

}

func (services *services) DeleteAccountData(ctx echo.Context, accountID uuid.UUID) error {
	ctxRequest := ctx.Request().Context()
	err := services.PG.DeleteFullAccountData(ctxRequest, accountID.String())
	if err != nil {
		return err
	}
	return nil
}
