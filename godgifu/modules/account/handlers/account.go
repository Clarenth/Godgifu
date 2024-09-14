package handlers

import (
	"bytes"
	"io/ioutil"
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

func (handler *handler) UpdateAccount(ctx echo.Context) error {
	accountID := ctx.Get("account")
	log.Print(accountID)

	// we store the Request body inside var body so that it can be used on more functions
	// this is because each time we read the request body, the body is exhaust. Must investigate why later.
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Reset the body after reading into var body
	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var reqEmp updateEmployee
	var reqId updateIdentity
	log.Print(ctx)
	if tools.BindJSON(ctx, &reqEmp) {
		log.Print(reqEmp)
		// accountID = accountID.(*jwt.JWTToken).ID
		updateData := &models.AccountEmployee{
			// ID:                  accountID.(*uuid.UUID),
			Email:               reqEmp.Email,
			Password:            reqEmp.Password,
			PhoneNumber:         reqEmp.PhoneNumber,
			EmploymentTitle:     reqEmp.EmploymentTitle,
			OfficeAddress:       reqEmp.OfficeAddress,
			SecurityAccessLevel: reqEmp.SecurityAccessLevel,
		}
		log.Print(updateData)
		return ctx.JSON(200, "Made it to reqEmp with struct data populated.")
	}
	log.Print(ctx)

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if tools.BindJSON(ctx, &reqId) {
		log.Print(reqId)
		updateData := &models.AccountIdentity{
			// ID:          accountID.(*uuid.UUID),
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
		log.Print(updateData)
		return ctx.JSON(200, "Made it to reqId with struct data populated.")
	}

	return ctx.JSON(200, "Outer scope return")
}

func (handler *handler) UpdateEmployee(ctx echo.Context) error {
	accountID := ctx.Get("account")
	// accountID := uuid.MustParse(string(authHeader.IDCode.String()))

	// var request models.AccountEmployee
	// var v interface{}
	// var request UpdateEmployee
	// var request1 UpdateIdentity

	// switch ctx. {
	// case UpdateEmployee:
	// 	log.Print("Hello case UpdateEmployee")
	// 	return ctx.JSON(200, "UpdateEmployee")
	// case UpdateIdentity:
	// 	log.Print("Hello case UpdateIdentity")
	// 	return ctx.JSON(200, "UpdateEmployee")
	// }

	// if ok := tools.BindJSON(ctx, &request); !ok {
	// 	log.Printf("Error in Signup, BindJSON failed request: %v", &request)
	// 	return echo.NewHTTPError(http.StatusBadRequest)
	// }

	// if reflect.DeepEqual(struct1, struct2) {
	// 	fmt.Println("The JSON data matches one of the structs.")
	// } else {
	// 	fmt.Println("The JSON data does not match either struct.")
	// }

	// if tools.BindJSON(ctx, &request) {
	// 	log.Print("reached bindJSON employee")
	// 	// return ctx.JSON(200, "This is a en edit employee")
	// } else if tools.BindJSON(ctx, &request1) {
	// 	log.Print("reached bindJSON identity")
	// 	return ctx.JSON(200, "This is a en edit identity")
	// } else {
	// 	return ctx.JSON(200, "We successfully passed the checks.")
	// }

	// switch {
	// case tools.BindJSON(ctx, &request1):
	// 	log.Print("test1")
	// 	// return ctx.JSON(200, "test1")
	// case tools.BindJSON(ctx, &request2):
	// 	log.Print("test2")
	// 	return ctx.JSON(200, "test2")
	// default:
	// 	log.Print("default")
	// 	return ctx.JSON(200, "default")
	// }
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
