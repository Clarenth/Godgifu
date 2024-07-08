package postgres

import (
	"context"

	"godgifu/modules/account/models"
)

type PostgresDB interface {
	SelectFullAccountData(ctx context.Context, accountID string) (*models.Account, error)
	DeleteFullAccountData(ctx context.Context, accountID string) error
}
