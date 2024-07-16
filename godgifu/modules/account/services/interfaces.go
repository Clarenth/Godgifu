package services

import (
	"godgifu/modules/account/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AccountService interface {
	CreateAccount(ctx echo.Context, account *models.Account) error
	GetAccountData(ctx echo.Context, accountID uuid.UUID) (*models.Account, error)
	DeleteAccountData(ctx echo.Context, accountID uuid.UUID) error
	UpdateEmploymeeData(ctx echo.Context, accountID uuid.UUID, updateData *models.AccountEmployee) (*models.AccountEmployee, error)
	UpdateIdentityData(ctx echo.Context, accountID uuid.UUID, updateData *models.AccountIdentity) (*models.AccountIdentity, error)
}
