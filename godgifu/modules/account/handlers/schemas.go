package handlers

type signupSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=16,lte=512"`
}

type updateEmployee struct {
	Email               string  `json:"email,omitempty" validate:"required,email"`
	Password            string  `json:"password"`
	PhoneNumber         *string `json:"phone_number,omitempty" validate:"required"`
	EmploymentTitle     *string `json:"employment_title,omitempty" validate:"required"`
	OfficeAddress       *string `json:"office_address,omitempty" validate:"required"`
	SecurityAccessLevel *string `json:"security_access_level,omitempty" validate:"required"`
	EmploymentDateStart *string `json:"employment_date_start,omitempty" validate:"required"`
	EmploymentDateEnd   *string `json:"employment_date_end,omitempty" validate:"required"`
}
type updateIdentity struct {
	FirstName   *string `json:"first_name,omitempty" validate:"required"`
	MiddleName  *string `json:"middle_name,omitempty" validate:"required"`
	LastName    *string `json:"last_name,omitempty" validate:"required"`
	Age         *string `json:"age,omitempty" validate:"required"`
	Sex         *string `json:"sex,omitempty" validate:"required"`
	Gender      *string `json:"gender,omitempty" validate:"required"`
	Height      *string `json:"height,omitempty" validate:"required"`
	HomeAddress *string `json:"home_address,omitempty" validate:"required"`
	Birthdate   *string `json:"birthdate,omitempty" validate:"required"`
	Birthplace  *string `json:"birthplace,omitempty" validate:"required"`
}
