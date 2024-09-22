package postgres

import (
	"context"
	"log"
	"time"

	"godgifu/modules/account/models"

	"github.com/jackc/pgconn"
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

func (pg *postgres) CreateAccount(ctx context.Context, accountData *models.Account) error {
	insertQuery := `INSERT INTO accounts_employee (id, email, password, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5) RETURNING *;
						`
	if _, err := pg.DB.ExecContext(ctx, insertQuery, accountData.AccountEmployee.ID, accountData.AccountEmployee.Email, accountData.AccountEmployee.Password, accountData.AccountEmployee.CreatedAt, accountData.AccountEmployee.UpdatedAt); err != nil {
		log.Printf("error in Postgres accounts_employee, err: %v", err)
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "unique_violation" {
			log.Printf("error1: Could not create an account with email: %v, or phone number: %v. Reason: %v\n", accountData.AccountEmployee.Email, accountData.AccountEmployee.PhoneNumber, err.Code)
			return err
		}
		return err
	}

	insertQuery = `INSERT INTO accounts_identity (id, first_name, last_name, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5) RETURNING *;
						`
	if result, err := pg.DB.ExecContext(ctx, insertQuery, accountData.AccountIdentity.ID, accountData.AccountIdentity.FirstName, accountData.AccountIdentity.LastName, accountData.AccountEmployee.CreatedAt, accountData.AccountEmployee.UpdatedAt); err != nil {
		log.Printf("error in Postgres accounts_identity, err: %v", err)
		log.Printf("Inserted %v", result)
		return err
	}
	return nil
}

func (pg *postgres) GetFullAccountData(ctx context.Context, accountID string) (*models.Account, error) {
	accountEmployee := &models.AccountEmployee{}
	accountIdentity := &models.AccountIdentity{}

	selectQuery := `SELECT email, phone_number, employment_title, office_address, employment_date_start, employment_date_end
						FROM accounts_employee WHERE id = $1;`
	// err := pg.DB.QueryRowxContext(ctx, query, accountID).Scan(&accountEmployee)

	if err := pg.DB.GetContext(ctx, accountEmployee, selectQuery, accountID); err != nil {
		log.Printf("Unable to get account with id: %v. Error:%v\n", accountID, err)
		return nil, err
	}

	selectQuery = `SELECT first_name, middle_name, last_name, age, sex, gender, height, home_address, birthdate, birthplace
						FROM accounts_identity WHERE id = $1;`
	if err := pg.DB.GetContext(ctx, accountIdentity, selectQuery, accountID); err != nil {
		log.Printf("Unable to get account with id: %v. Error:%v\n", accountID, err)
		return nil, err
	}

	account := &models.Account{
		AccountEmployee: &models.AccountEmployee{
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

	log.Print(account.AccountEmployee)
	log.Print(account.AccountIdentity)

	return account, nil
}

func (pg *postgres) DeleteFullAccountData(ctx context.Context, accountID string) error {
	deleteQuery := `DELETE FROM accounts_employee WHERE id = $1`

	err := pg.DB.QueryRowContext(ctx, deleteQuery, accountID).Scan(&accountID)
	if err != nil {
		log.Printf("Error when deleting account with id %v", accountID)
		return err
	}

	deleteQuery = `DELETE FROM accounts_identity WHERE id = $1`
	err = pg.DB.QueryRowContext(ctx, deleteQuery, accountID).Scan(&accountID)
	if err != nil {
		log.Printf("Error when deleting account with id: %v", accountID)
		return err
	}
	return nil
}

// https://stackoverflow.com/questions/13305878/dont-update-column-if-update-value-is-null
// https://stackoverflow.com/questions/38051311/sql-query-for-if-not-null-then-update-or-else-keep-the-same-data
func (pg *postgres) UpdateEmployeeAccount(ctx context.Context, accountID string, updateData *models.AccountEmployee) (*models.AccountEmployee, error) {
	ut := time.Now().UTC()
	updateData.UpdatedAt = &ut
	updateQuery := `UPDATE accounts_employee 
	SET email = CASE WHEN email IS DISTINCT FROM $2::varchar THEN $2::varchar ELSE email END, 
			phone_number = CASE WHEN phone_number IS DISTINCT FROM $3::varchar AND $3::varchar IS NOT '' THEN $3::varchar ELSE phone_number END,
			employment_title = CASE WHEN employment_title IS DISTINCT FROM $4::text AND $4::text IS NOT '' THEN $4::text ELSE employment_title END,
			office_address = CASE WHEN office_address IS DISTINCT FROM $5::text AND $5::text IS NOT '' THEN $5::text ELSE office_address END,
			updated_at = CASE WHEN updated_at IS DISTINCT FROM $6::timestamp AND $6::text IS NOT '' THEN $6::timestamp ELSE updated_at END
	WHERE id = $1
	;
	`
	result, err := pg.DB.ExecContext(ctx, updateQuery, accountID, updateData.Email, updateData.PhoneNumber,
		updateData.EmploymentTitle, updateData.OfficeAddress, updateData.UpdatedAt)
	if err != nil {
		log.Printf("Unable to UPDATE account with id: %v. Error:%v\n", accountID, err)
		return nil, err
	}
	log.Print(result.RowsAffected())

	return nil, nil
}

func (pg *postgres) UpdateIdentityAccount(ctx context.Context, accountID string, updateData *models.AccountIdentity) (*models.AccountIdentity, error) {
	updateQuery := `UPDATE accounts_identity 
	SET email = CASE WHEN email IS DISTINCT FROM $2::varchar THEN $2::varchar ELSE email END, 
			phone_number = CASE WHEN phone_number IS DISTINCT FROM $3::varchar AND $3::varchar IS NOT '' THEN $3::varchar ELSE phone_number END,
			employment_title = CASE WHEN employment_title IS DISTINCT FROM $4::text AND $4::text IS NOT '' THEN $4::text ELSE employment_title END,
			office_address = CASE WHEN office_address IS DISTINCT FROM $5::text AND $5::text IS NOT '' THEN $5::text ELSE office_address END,
			updated_at = CASE WHEN updated_at IS DISTINCT FROM $6::timestamp AND $6::text IS NOT '' THEN $6::timestamp ELSE updated_at END
	WHERE id = $1
	;
	`
	result, err := pg.DB.ExecContext(ctx, updateQuery, accountID, updateData.FirstName)
	if err != nil {
		log.Printf("Unable to UPDATE account with id: %v. Error:%v\n", accountID, err)
		return nil, err
	}
	log.Print(result.RowsAffected())

	return nil, nil
}
