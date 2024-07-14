package services

import (
	"log"
	"net/http"
	"time"

	"godgifu/modules/account/db/postgres"
	"godgifu/modules/account/models"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type services struct {
	PG postgres.PostgresDB
}

func NewAccountServices(postgres postgres.PostgresDB) AccountService {
	return &services{
		PG: postgres,
	}
}

func (service *services) CreateAccount(ctx echo.Context, account *models.Account) (err error) {
	var argon2Params = &argon2id.Params{
		Memory:      62500,
		Iterations:  2,
		Parallelism: 16,
		SaltLength:  32,
		KeyLength:   64,
	}

	// implement argon2 hashing and salting here
	// to protect passwords from plain text exposing
	hashedPassword, err := argon2id.CreateHash(account.AccountEmployee.Password, argon2Params)
	if err != nil {
		log.Panicf("Unable to encrypt accountData password from new Signup with given email: %v\n", account.AccountEmployee.Email)
		return ctx.String(http.StatusInternalServerError, "Unable to create user account.")
	}
	account.AccountEmployee.Password = hashedPassword

	accountID := uuid.New()
	account.AccountEmployee.ID = &accountID
	account.AccountEmployee.ID = &accountID

	timeCreatedAt := time.Now().UTC()
	account.AccountEmployee.CreatedAt = &timeCreatedAt
	account.AccountEmployee.UpdatedAt = &timeCreatedAt

	account.AccountIdentity.ID = account.AccountEmployee.ID
	account.AccountIdentity.CreatedAt = account.AccountEmployee.CreatedAt
	account.AccountIdentity.CreatedAt = account.AccountEmployee.UpdatedAt

	ctxRequest := ctx.Request().Context()

	err = service.PG.CreateAccount(ctxRequest, account)
	if err != nil {
		return err
	}

	return nil
}

func (services *services) GetAccountData(ctx echo.Context, accountID uuid.UUID) (*models.Account, error) {
	// panic("not done")
	context := ctx.Request().Context()
	account, err := services.PG.SelectAccountDataForClientProfile(context, accountID.String())
	if err != nil {
		return nil, err
	}

	return account, nil

}

func (services *services) DeleteAccountData(ctx echo.Context, accountID uuid.UUID) error {
	ctxRequest := ctx.Request().Context()
	err := services.PG.DeleteFullAccountData(ctxRequest, accountID.String())
	if err != nil {
		return err
	}
	return nil
}
