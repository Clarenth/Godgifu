package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountEmployeeData *AccountEmployeeData `json:"-"`
	AccountIdentityData *AccountIdentityData `json:"-"`
}

type AccountEmployeeData struct {
	ID                  uuid.UUID `db:"id" json:"id"`
	Email               string    `db:"email" json:"username"`
	Password            string    `db:"password" json:"-"`
	PhoneNumber         string    `db:"phone_number" json:"phone_number"`
	EmploymentTitle     string    `db:"employment_title" json:"employment_title"`
	OfficeAddress       string    `db:"office_address" json:"office_address"`
	SecurityAccessLevel string    `db:"security_access_level" json:"security_access_level"`
	EmploymentDateStart time.Time `db:"employment_date_start" json:"employment_date_start"`
	EmploymentDateEnd   time.Time `db:"employment_date_end" json:"employment_date_end"`
	Verified            bool      `db:"verified"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
}

type AccountIdentityData struct {
	ID          uuid.UUID `db:"id" json:"id"`
	FirstName   string    `db:"first_name" json:"first_name"`
	MiddleName  string    `db:"middle_name" json:"middle_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	Age         string    `db:"age" json:"age"`
	Sex         string    `db:"sex" json:"sex"`
	Gender      string    `db:"gender" json:"gender"`
	Height      string    `db:"height" json:"height"`
	HomeAddress string    `db:"home_address" json:"home_address"`
	Birthdate   string    `db:"birthdate" json:"birthdate"`
	Birthplace  string    `db:"birthplace" json:"birthplace"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
