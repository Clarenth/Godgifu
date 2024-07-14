package postgres

import (
	"context"
	"log"

	"godgifu/modules/account/models"

	"github.com/jmoiron/sqlx"
	// "github.com/jackc/pgconn"
	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	DB *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) PostgresDB {
	return &postgres{
		DB: db,
	}
}

// func (pg *postgres) CreateAccount(ctx context.Context, accountData *models.Account) error {
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

func (pg *postgres) SelectFullAccountData(ctx context.Context, accountID string) (*models.Account, error) {
	accountEmployee := &models.AccountEmployee{}
	accountIdentity := &models.AccountIdentity{}

	log.Print("Hello from SelectFullAccountData")

	query := `SELECT id, email, phone_number, employment_title, office_address, employment_date_start, employment_date_end 
						FROM accounts_employee WHERE id = $1;`
	if err := pg.DB.GetContext(ctx, accountEmployee, query, accountID); err != nil {
		log.Printf("Unable to get account with id: %v. Error:%v\n", accountID, err)
	}

	query = `SELECT first_name, middle_name, last_name, age, sex, gender, height, home_address, birthdate, birthplace 
						FROM accounts_identity WHERE id = $1;`
	if err := pg.DB.GetContext(ctx, accountIdentity, query, accountID); err != nil {
		log.Printf("Unable to get account with id: %v. Error:%v\n", accountID, err)
	}

	account := &models.Account{
		AccountEmployee: &models.AccountEmployee{
			ID:                  accountEmployee.ID,
			Email:               accountEmployee.Email,
			PhoneNumber:         accountEmployee.PhoneNumber,
			EmploymentTitle:     accountEmployee.EmploymentTitle,
			OfficeAddress:       accountEmployee.OfficeAddress,
			EmploymentDateStart: accountEmployee.EmploymentDateStart,
			EmploymentDateEnd:   accountEmployee.EmploymentDateEnd,
		},
		AccountIdentity: &models.AccountIdentity{
			FirstName:   accountIdentity.FirstName,
			MiddleName:  accountIdentity.MiddleName,
			LastName:    accountIdentity.LastName,
			Age:         accountIdentity.Age,
			Sex:         accountIdentity.Sex,
			Gender:      accountIdentity.Gender,
			Height:      accountIdentity.Height,
			HomeAddress: accountIdentity.HomeAddress,
			Birthdate:   accountIdentity.Birthdate,
			Birthplace:  accountIdentity.Birthplace,
		},
	}

	return account, nil
}

func (pg *postgres) DeleteFullAccountData(ctx context.Context, accountID string) error {
	query := `DELETE FROM accounts_employee WHERE id = $1`

	err := pg.DB.QueryRowContext(ctx, query, accountID).Scan(&accountID)
	if err != nil {
		log.Printf("Error when deleting account with id %v", accountID)
		return err
	}

	query = `DELETE FROM accounts_identity WHERE id = $1`
	err = pg.DB.QueryRowContext(ctx, query, accountID).Scan(&accountID)
	if err != nil {
		log.Printf("Error when deleting account with id: %v", accountID)
		return err
	}
	return nil
}
