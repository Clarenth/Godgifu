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

func (pgDB *postgres) CreateAccount(ctx echo.Context, accountData *models.AccountEmployeeData) error {
	// return nil
	log.Print("reached the PG DB")

	query := `INSERT INTO account_employee (id, email, password, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5) RETURNING *;
						`
	log.Print("Reached db.Exec")
	if result, err := pgDB.db.Exec(query, accountData.ID, accountData.Email, accountData.Password, accountData.CreatedAt, accountData.UpdatedAt); err != nil {
		log.Print("Hello from err: ", err)
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "unique_violation" {
			log.Printf("error1: Could not create an account with email: %v, or phone number: %v. Reason: %v\n", accountData.Email, accountData.PhoneNumber, err.Code)
			return err
		}
		log.Print(result)
		return nil
	}
	// result, err := pgDB.db.Exec(query, accountData.ID, accountData.Email, accountData.Password, accountData.CreatedAt, accountData.UpdatedAt)
	// if err != nil {
	// 	log.Printf("error in postgres CreateAccount: %v", err)
	// 	return fmt.Errorf("error: could not create account")
	// }
	log.Print("We are now at return nil")
	return nil
}
