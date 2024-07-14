package services

import (
	account "godgifu/modules/account/models"
	"godgifu/modules/auth/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	// CreateAccount(ctx echo.Context, accountData *account.Account) error
	Signin(ctx echo.Context, accountData *account.AccountEmployee) (result *account.AccountEmployee, err error)
}

type JWTService interface {
	NewTokenPairFromAccount(ctx echo.Context, account *account.AccountEmployee, previousTokenID string) (*models.JWTTokenPair, error)
	NewIDToken(ctx echo.Context, account *account.AccountEmployee, previousTokenID string) (*models.JWTIDToken, error)
	ValidateIDToken(tokenString string) (*models.JWTToken, error)
	ValidateRefreshToken(tokenString string) (*models.JWTRefreshToken, error)
	Signout(ctx echo.Context, accountID uuid.UUID) error
}

type PASETOService interface {
}
