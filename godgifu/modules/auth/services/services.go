package services

import (
	account "godgifu/modules/account/models"
	postgres "godgifu/modules/auth/db"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type authServices struct {
	postgres postgres.PostgresDB
}

/*
	Look into eliminating specific names and instead be using generic names
	like repository or database
*/

func NewAuthServices(postgres postgres.PostgresDB) AuthService {
	return &authServices{
		postgres: postgres,
	}
}

var argon2Params = &argon2id.Params{
	Memory:      62500,
	Iterations:  2,
	Parallelism: 16,
	SaltLength:  32,
	KeyLength:   64,
}

func (services *authServices) CreateAccount(ctx echo.Context, accountData account.AccountEmployeeData) (err error) {
	// implement argon2 hashing and salting here
	// to protect passwords from plain text exposing
	hashedPassword, err := argon2id.CreateHash(accountData.Password, argon2Params)
	if err != nil {
		log.Panicf("Unable to encrypt accountData password from new Signup with given email: %v\n", accountData.Email)
		return ctx.String(http.StatusInternalServerError, "Unable to create user account.")
	}
	accountData.Password = hashedPassword
	accountData.ID = uuid.New()
	accountData.CreatedAt = time.Now().UTC()
	accountData.UpdatedAt = time.Now().UTC()

	if err := services.postgres.CreateAccount(ctx, &accountData); err != nil {
		log.Print(err)
		return err
	}

	return nil
}
