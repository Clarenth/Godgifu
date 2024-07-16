package postgres

import (
	"context"

	"godgifu/modules/account/models"
)

type PostgresDB interface {
	CreateAccount(ctx context.Context, accountData *models.Account) error

	GetFullAccountData(ctx context.Context, accountID string) (*models.Account, error)

	DeleteFullAccountData(ctx context.Context, accountID string) error

	UpdateAccountEmployee(ctx context.Context, accountID string, updateData *models.AccountEmployee) (*models.AccountEmployee, error)
	UpdateAccountIdentity(ctx context.Context, accountID string, updateData *models.AccountIdentity) (*models.AccountIdentity, error)
}
