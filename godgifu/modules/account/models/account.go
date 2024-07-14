package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountEmployee *AccountEmployee `json:"account_employee"`
	AccountIdentity *AccountIdentity `json:"account_identity"`
}

type AccountEmployee struct {
	ID                  *uuid.UUID `db:"id" json:"id,omitempty"`
	Email               string     `db:"email" json:"email,omitempty"`
	Password            string     `db:"password" json:"-"`
	PhoneNumber         *string    `db:"phone_number" json:"phone_number,omitempty"`
	EmploymentTitle     *string    `db:"employment_title" json:"employment_title,omitempty"`
	OfficeAddress       *string    `db:"office_address" json:"office_address,omitempty"`
	SecurityAccessLevel *string    `db:"security_access_level" json:"security_access_level,omitempty"`
	EmploymentDateStart *time.Time `db:"employment_date_start" json:"employment_date_start,omitempty"`
	EmploymentDateEnd   *time.Time `db:"employment_date_end" json:"employment_date_end,omitempty"`
	Verified            *bool      `db:"verified" json:"-,omitempty"`
	CreatedAt           *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt           *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type AccountIdentity struct {
	ID          *uuid.UUID `db:"id" json:"id,omitempty"`
	FirstName   *string    `db:"first_name" json:"first_name,omitempty"`
	MiddleName  *string    `db:"middle_name" json:"middle_name,omitempty"`
	LastName    *string    `db:"last_name" json:"last_name,omitempty"`
	Age         *string    `db:"age" json:"age,omitempty"`
	Sex         *string    `db:"sex" json:"sex,omitempty"`
	Gender      *string    `db:"gender" json:"gender,omitempty"`
	Height      *string    `db:"height" json:"height,omitempty"`
	HomeAddress *string    `db:"home_address" json:"home_address,omitempty"`
	Birthdate   *string    `db:"birthdate" json:"birthdate,omitempty"`
	Birthplace  *string    `db:"birthplace" json:"birthplace,omitempty"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
