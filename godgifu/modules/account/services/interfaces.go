package services

import (
	"godgifu/modules/account/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AccountService interface {
	CreateAccount(ctx echo.Context, account *models.Account) error
	GetAccount(ctx echo.Context, accountID uuid.UUID) (*models.Account, error)
	DeleteAccount(ctx echo.Context, accountID uuid.UUID) error
	UpdateEmployee(ctx echo.Context, accountID uuid.UUID, updateData *models.AccountEmployee) (*models.AccountEmployee, error)
	UpdateIdentity(ctx echo.Context, accountID uuid.UUID, updateData *models.AccountIdentity) (*models.AccountIdentity, error)
}
