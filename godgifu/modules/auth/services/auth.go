package services

// Change the name of this file to auth.go. This is now a file with functions for authentication, and not an all encompassing services file.

import (
	"log"

	account "godgifu/modules/account/models"
	postgres "godgifu/modules/auth/db"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type authServices struct {
	Postgres postgres.PostgresDB
}

/*
	Look into eliminating specific names and instead be using generic names
	like repository or database
*/

func NewAuthServices(postgres postgres.PostgresDB) AuthService {
	return &authServices{
		Postgres: postgres,
	}
}

var argon2Params = &argon2id.Params{
	Memory:      62500,
	Iterations:  2,
	Parallelism: 16,
	SaltLength:  32,
	KeyLength:   64,
}

// func (service *authServices) CreateAccount(ctx echo.Context, account *account.Account) (err error) {
// 	// implement argon2 hashing and salting here
// 	// to protect passwords from plain text exposing
// 	hashedPassword, err := argon2id.CreateHash(account.AccountEmployee.Password, argon2Params)
// 	if err != nil {
// 		log.Panicf("Unable to encrypt accountData password from new Signup with given email: %v\n", account.AccountEmployee.Email)
// 		return ctx.String(http.StatusInternalServerError, "Unable to create user account.")
// 	}
// 	account.AccountEmployee.Password = hashedPassword
// 	account.AccountEmployee.ID = uuid.New()
// 	account.AccountEmployee.CreatedAt = time.Now().UTC()
// 	account.AccountEmployee.UpdatedAt = time.Now().UTC()

// 	account.AccountIdentity.ID = account.AccountEmployee.ID
// 	account.AccountIdentity.CreatedAt = account.AccountEmployee.CreatedAt
// 	account.AccountIdentity.CreatedAt = account.AccountEmployee.UpdatedAt

// 	err = service.Postgres.CreateAccount(ctx, account)
// 	if err != nil {
// 		return err
// 	}
// 	// if err := service.postgres.CreateAccount(ctx, &account); err != nil {
// 	// 	log.Print(err)
// 	// 	return err
// 	// }
// 	return nil
// }

func (service *authServices) Signin(ctx echo.Context, payload *account.AccountEmployee) (result *account.AccountEmployee, err error) {
	account, err := service.Postgres.FindAccountByEmail(ctx, payload.Email)
	if err != nil {
		log.Print("Could not verify email address in service Signin")
		return nil, err
	}

	// Verify the password against the hash
	accountMatch, err := argon2id.ComparePasswordAndHash(payload.Password, account.Password)
	if err != nil {
		log.Printf("Error: Could not ComparePasswordAndHash successfully")
		return nil, err
	}
	if !accountMatch {
		return nil, err
	}

	account.Password = ""

	return account, nil
}

func (service *jwtService) Signout(ctx echo.Context, accountID uuid.UUID) error {
	ctxRequest := ctx.Request().Context()
	return service.TokenRepository.DeleteAccountRefreshTokens(ctxRequest, accountID.String())
}
