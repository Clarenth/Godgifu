package handlers

import (
	"bytes"
	"io"
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

	/*
		BindJSON is broken. Must fix to solve the comparison problem.
	*/

	accountID := ctx.Get("account")

	account, err := handler.AccountService.GetAccount(ctx, accountID.(*jwt.JWTToken).ID)
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

	err := handler.AccountService.DeleteAccount(ctx, *accountID.(*models.AccountEmployee).ID)
	if err != nil {
		return echo.NewHTTPError(500, "Account was not deleted")
	}
	return ctx.NoContent(204)
}

func (handler *handler) UpdateAccount(ctx echo.Context) error {
	accountID := ctx.Get("account")

	// We store the Request body inside 'var body' so that it can be used on one or more functions.
	// This is because each time we read the request body the body is consumed, and cannot be used again.
	// Must investigate why later.
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(400, "error: Invalid request")
	}
	ctx.Request().Body.Close()

	// Reset the body after reading into var body
	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	var reqEmp updateEmployee
	if tools.BindJSON(ctx, &reqEmp) {
		log.Print(reqEmp)
		reqData := &models.AccountEmployee{
			Email:               reqEmp.Email,
			Password:            reqEmp.Password,
			PhoneNumber:         reqEmp.PhoneNumber,
			EmploymentTitle:     reqEmp.EmploymentTitle,
			OfficeAddress:       reqEmp.OfficeAddress,
			SecurityAccessLevel: reqEmp.SecurityAccessLevel,
		}
		log.Print(reqData)
		result, err := handler.AccountService.UpdateEmployee(ctx, accountID.(*jwt.JWTToken).ID, reqData)
		if err != nil {
			log.Print(result)
			log.Print(err)
		}
		// log.Print(result)
		return ctx.JSON(200, "Made it to reqEmp with struct data populated.")
	}

	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
	var reqId updateIdentity
	if tools.BindJSON(ctx, &reqId) {
		log.Print(reqId)
		reqData := &models.AccountIdentity{
			FirstName:   reqId.FirstName,
			MiddleName:  reqId.MiddleName,
			LastName:    reqId.LastName,
			Sex:         reqId.Sex,
			Gender:      reqId.Gender,
			Height:      reqId.Height,
			HomeAddress: reqId.HomeAddress,
			Birthdate:   reqId.Birthdate,
			Birthplace:  reqId.Birthplace,
		}
		log.Print(reqData)
		result, err := handler.AccountService.UpdateIdentity(ctx, accountID.(*jwt.JWTToken).ID, reqData)
		if err != nil {
			log.Print(result)
			log.Print(err)
		}
		// log.Print(result)

		return ctx.JSON(200, "Made it to reqId with struct data populated.")
	}

	return ctx.JSON(200, "Outer scope return")
}
