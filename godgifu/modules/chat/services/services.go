package services

import (
	"godgifu/modules/account/db/postgres"

	"github.com/labstack/echo/v4"
)

type services struct {
	repository postgres.PostgresDB
}

func NewMessageServices(repository postgres.PostgresDB) MessageService {
	return &services{
		repository: repository,
	}
}

func (services *services) CreateMessage(ctx echo.Context) (err error) {
	panic("not done")
}
