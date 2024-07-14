package db

import (
	"context"
	"time"

	"godgifu/modules/account/models"

	"github.com/labstack/echo/v4"
)

type PostgresDB interface {
	// CreateAccount(ctx echo.Context, accountData *models.Account) error
	FindAccountByEmail(ctx echo.Context, email string) (account *models.AccountEmployee, err error)
}

type RedisDB interface {
	SetRefreshToken(ctx context.Context, accountID string, tokenID string, tokenExpireTime time.Duration) error
	DeleteAccountRefreshTokens(ctx context.Context, accountID string) error
	DeleteRefreshToken(ctx context.Context, accountID string, tokenID string) error
}
