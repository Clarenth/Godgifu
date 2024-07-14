package db

import (
	"log"

	"godgifu/modules/account/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type postgres struct {
	DB *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) PostgresDB {
	return &postgres{
		DB: db,
	}
}

// func (pg *postgres) CreateAccount(ctx echo.Context, accountData *models.Account) error {
// 	query := `INSERT INTO accounts_employee (id, email, password, created_at, updated_at)
// 						VALUES ($1, $2, $3, $4, $5) RETURNING *;
// 						`
// 	if _, err := pg.DB.Exec(query, accountData.AccountEmployee.ID, accountData.AccountEmployee.Email, accountData.AccountEmployee.Password, accountData.AccountEmployee.CreatedAt, accountData.AccountEmployee.UpdatedAt); err != nil {
// 		log.Printf("error in Postgres accounts_employee, err: %v", err)
// 		if err, ok := err.(*pgconn.PgError); ok && err.Code == "unique_violation" {
// 			log.Printf("error1: Could not create an account with email: %v, or phone number: %v. Reason: %v\n", accountData.AccountEmployee.Email, accountData.AccountEmployee.PhoneNumber, err.Code)
// 			return err
// 		}
// 		return err
// 	}

// 	query = `INSERT INTO accounts_identity (id, first_name, last_name, created_at, updated_at)
// 						VALUES ($1, $2, $3, $4, $5) RETURNING *;
// 						`
// 	if result, err := pg.DB.Exec(query, accountData.AccountIdentity.ID, accountData.AccountIdentity.FirstName, accountData.AccountIdentity.LastName, accountData.AccountEmployee.CreatedAt, accountData.AccountEmployee.UpdatedAt); err != nil {
// 		log.Printf("error in Postgres accounts_identity, err: %v", err)
// 		log.Printf("Inserted %v", result)
// 		return err
// 	}
// 	return nil
// }

// FindAccountByEmail searches the DB for the account that matches the email parameter and returns the account.
func (pg *postgres) FindAccountByEmail(ctx echo.Context, email string) (*models.AccountEmployee, error) {
	accountData := &models.AccountEmployee{}
	query := "SELECT id, email, password FROM accounts_employee WHERE email = $1"
	if err := pg.DB.Get(accountData, query, email); err != nil {
		log.Printf("Failed to get account with email: %v. Error:%v\n", email, err)
		return nil, err
	}
	return accountData, nil
}
