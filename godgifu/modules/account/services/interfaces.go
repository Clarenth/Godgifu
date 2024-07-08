package services

import (
	"godgifu/modules/account/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AccountService interface {
	GetAccountData(ctx echo.Context, accountID uuid.UUID) (*models.Account, error)
	DeleteAccountData(ctx echo.Context, accountID uuid.UUID) error
}
