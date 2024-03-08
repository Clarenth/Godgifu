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

func (services *authServices) CreateAccount(ctx echo.Context, account account.Account) (err error) {
	// implement argon2 hashing and salting here
	// to protect passwords from plain text exposing
	hashedPassword, err := argon2id.CreateHash(account.AccountEmployee.Password, argon2Params)
	if err != nil {
		log.Panicf("Unable to encrypt accountData password from new Signup with given email: %v\n", account.AccountEmployee.Email)
		return ctx.String(http.StatusInternalServerError, "Unable to create user account.")
	}
	account.AccountEmployee.Password = hashedPassword
	account.AccountEmployee.ID = uuid.New()
	account.AccountEmployee.CreatedAt = time.Now().UTC()
	account.AccountEmployee.UpdatedAt = time.Now().UTC()

	account.AccountIdentity.ID = account.AccountEmployee.ID
	account.AccountIdentity.CreatedAt = account.AccountEmployee.CreatedAt
	account.AccountIdentity.CreatedAt = account.AccountEmployee.UpdatedAt

	err = services.postgres.CreateAccount(ctx, &account)
	if err != nil {
		return err
	}
	// if err := services.postgres.CreateAccount(ctx, &account); err != nil {
	// 	log.Print(err)
	// 	return err
	// }
	return nil
}
