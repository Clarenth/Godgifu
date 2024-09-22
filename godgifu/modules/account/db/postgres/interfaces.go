package postgres

import (
	"context"

	"godgifu/modules/account/models"
)

type PostgresDB interface {
	CreateAccount(ctx context.Context, accountData *models.Account) error

	GetFullAccountData(ctx context.Context, accountID string) (*models.Account, error)

	DeleteFullAccountData(ctx context.Context, accountID string) error

	UpdateEmployeeAccount(ctx context.Context, accountID string, updateData *models.AccountEmployee) (*models.AccountEmployee, error)
	UpdateIdentityAccount(ctx context.Context, accountID string, updateData *models.AccountIdentity) (*models.AccountIdentity, error)
}
