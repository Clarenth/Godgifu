package db

import (
	"log"

	"godgifu/modules/account/models"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type postgres struct {
	db *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) PostgresDB {
	return &postgres{
		db: db,
	}
}

func (pgDB *postgres) CreateAccount(ctx echo.Context, accountData *models.Account) error {
	query := `INSERT INTO accounts_employee (id, email, password, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5) RETURNING *;
						`
	log.Print("Reached db.Exec")
	if _, err := pgDB.db.Exec(query, accountData.AccountEmployee.ID, accountData.AccountEmployee.Email, accountData.AccountEmployee.Password, accountData.AccountEmployee.CreatedAt, accountData.AccountEmployee.UpdatedAt); err != nil {
		log.Printf("error in Postgres accounts_employee, err: %v", err)
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "unique_violation" {
			log.Printf("error1: Could not create an account with email: %v, or phone number: %v. Reason: %v\n", accountData.AccountEmployee.Email, accountData.AccountEmployee.PhoneNumber, err.Code)
			return err
		}
		return err
	}

	query = `INSERT INTO accounts_identity (first_name, last_name, created_at, updated_at)
						VALUES ($1, $2, $3, $4) RETURNING *;
						`
	if result, err := pgDB.db.Exec(query, accountData.AccountIdentity.FirstName, accountData.AccountIdentity.LastName, accountData.AccountEmployee.CreatedAt, accountData.AccountEmployee.UpdatedAt); err != nil {
		log.Printf("error in Postgres accounts_identity, err: %v", err)
		log.Printf("Inserted %v", result)
		return err
	}
	return nil
}
