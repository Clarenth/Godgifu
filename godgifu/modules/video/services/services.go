package services

import (
	"godgifu/modules/account/db/postgres"

	"github.com/labstack/echo/v4"
)

type services struct {
	repository postgres.PostgresDB
}

func NewVideoServices(repository postgres.PostgresDB) VideoService {
	return &services{
		repository: repository,
	}
}

func (services *services) CreateVideo(ctx echo.Context) (err error) {
	panic("not done")
}
