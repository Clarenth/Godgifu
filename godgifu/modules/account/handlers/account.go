package handlers

import (
	"log"
	"net/http"

	"godgifu/modules/account/models"
	"godgifu/modules/account/services"
	jwt "godgifu/modules/auth/models"
	"godgifu/modules/tools"

	"github.com/labstack/echo/v4"
)

type handler struct {
	AccountService services.AccountService
}

func NewAccountHandlers(service services.AccountService) AccountHandlers {
	return &handler{
		AccountService: service,
	}
}

func (handler *handler) CreateAccount(ctx echo.Context) error {
	var request signupSchema
	if ok := tools.BindJSON(ctx, &request); !ok {
		log.Printf("Error in Signup, BindJSON failed request: %v", &request)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if request.Email == "" || request.Password == "" {
		log.Printf("Signup request failed with empty string fields: %v, %v", request.Email, request.Password)
		return ctx.String(http.StatusBadRequest, "Email and Password cannot be empty")
	}

	account := models.Account{
		AccountEmployee: &models.AccountEmployee{
			Email:    request.Email,
			Password: request.Password,
		},
		AccountIdentity: &models.AccountIdentity{},
	}

	err := handler.AccountService.CreateAccount(ctx, &account)
	if err != nil {
		// return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user account")
		return ctx.String(http.StatusInternalServerError, "Failed to create user account")
	}

	return ctx.String(http.StatusOK, "Success")
}

func (handler *handler) GetAccount(ctx echo.Context) error {
	accountID := ctx.Get("account")

	account, err := handler.AccountService.GetAccountData(ctx, accountID.(*jwt.JWTToken).ID)
	if err != nil {
		return ctx.JSON(400, "Debug: An error with your ID code")
	}

	// res := &models.Account{
	// 	AccountEmployee: account.AccountEmployee,
	// 	AccountIdentity: account.AccountIdentity,
	// }

	// res := &models.Account{
	// 	AccountEmployee: &models.AccountEmployee{
	// 		Email:               account.AccountEmployee.Email,
	// 		PhoneNumber:         account.AccountEmployee.PhoneNumber,
	// 		EmploymentTitle:     account.AccountEmployee.EmploymentTitle,
	// 		OfficeAddress:       account.AccountEmployee.OfficeAddress,
	// 		EmploymentDateStart: account.AccountEmployee.EmploymentDateStart,
	// 		EmploymentDateEnd:   account.AccountEmployee.EmploymentDateEnd,
	// 	},
	// 	AccountIdentity: &models.AccountIdentity{
	// 		FirstName:   account.AccountIdentity.FirstName,
	// 		MiddleName:  account.AccountIdentity.MiddleName,
	// 		LastName:    account.AccountIdentity.LastName,
	// 		Age:         account.AccountIdentity.Age,
	// 		Sex:         account.AccountIdentity.Sex,
	// 		Gender:      account.AccountIdentity.Gender,
	// 		Height:      account.AccountIdentity.Height,
	// 		HomeAddress: account.AccountIdentity.HomeAddress,
	// 		Birthdate:   account.AccountIdentity.Birthdate,
	// 		Birthplace:  account.AccountIdentity.Birthplace,
	// 	},
	// }
	return ctx.JSON(200, account)
}

func (handler *handler) DeleteAccount(ctx echo.Context) error {
	accountID := ctx.Get("account")

	err := handler.AccountService.DeleteAccountData(ctx, *accountID.(*models.AccountEmployee).ID)
	if err != nil {
		return echo.NewHTTPError(500, "Account was not deleted")
	}
	return ctx.NoContent(204)
}

func (handler *handler) UpdateEmployee(ctx echo.Context) error {
	accountID := ctx.Get("account")
	// accountID := uuid.MustParse(string(authHeader.IDCode.String()))

	var request models.AccountEmployee

	if ok := tools.BindJSON(ctx, &request); !ok {
		log.Printf("Error in Signup, BindJSON failed request: %v", &request)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	updateData := models.AccountEmployee{}

	result, err := handler.AccountService.UpdateEmploymeeData(ctx, accountID.(*jwt.JWTToken).ID, &updateData)
	log.Print(result)
	if err != nil {
		return ctx.JSON(400, updateData)
	}

	return nil
}

func (handler *handler) UpdateIdentity(ctx echo.Context) error {
	accountID := ctx.Get("account")

	var request models.AccountIdentity

	if ok := tools.BindJSON(ctx, &request); !ok {
		log.Printf("Error in Signup, BindJSON failed request: %v", &request)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	updateData := models.AccountIdentity{
		FirstName: request.FirstName,
	}

	result, err := handler.AccountService.UpdateIdentityData(ctx, accountID.(*jwt.JWTToken).ID, &updateData)
	log.Print(result)
	if err != nil {
		ctx.JSON(400, request)
	}

	return nil
}
