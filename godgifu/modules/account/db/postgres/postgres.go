package postgres

import (
	"context"
	"godgifu/modules/account/models"
	"log"

	"github.com/jmoiron/sqlx"
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

func (pg postgres) SelectFullAccountData(ctx context.Context, accountID string) (*models.Account, error) {
	accountEmployee := models.AccountEmployee{}
	accountIdentity := models.AccountIdentity{}

	query := `SELECT id, email, phone_number, employement_title, office_address, employment_date_start, employment_date_end 
						FROM accounts_employee;`
	if err := pg.DB.GetContext(ctx, accountEmployee, query, accountID); err != nil {
		log.Printf("Unable to get account with email: %v. Error:%v\n", accountID, err)
	}

	query = `SELECT first_name, middle_name, last_name, age, ex, gender, height, home_address, birthdate, brithplace 
						FROM accounts_identity;`
	if err := pg.DB.GetContext(ctx, accountIdentity, query, accountID); err != nil {
		log.Printf("Unable to get account with email: %v. Error:%v\n", accountID, err)
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
