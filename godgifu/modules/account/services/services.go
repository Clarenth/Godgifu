package services

import (
	"godgifu/modules/account/db/postgres"

	"github.com/labstack/echo/v4"
)

type services struct {
	repository postgres.PostgresDB
}

func NewAccountServices(repository postgres.PostgresDB) AccountService {
	return &services{
		repository: repository,
	}
}

func (services *services) CreateAccount(ctx echo.Context) (err error) {
	panic("not done")
}
