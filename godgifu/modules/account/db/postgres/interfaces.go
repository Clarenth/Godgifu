package postgres

import (
	"context"

	"godgifu/modules/account/models"
)

type PostgresDB interface {
	CreateAccount(ctx context.Context, accountData *models.Account) error
	SelectAccountDataForClientProfile(ctx context.Context, accountID string) (*models.Account, error)
	DeleteFullAccountData(ctx context.Context, accountID string) error
}
