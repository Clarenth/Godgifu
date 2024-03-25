package handlers

import (
	"log"
	"net/http"

	"godgifu/modules/account/models"
	account "godgifu/modules/account/models"
	"godgifu/modules/auth/services"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	AuthService services.AuthService
}

func NewAuthHandlers(service services.AuthService) AuthHandlers {
	return &authHandler{
		AuthService: service,
	}
}

func (handler *authHandler) Signin(ctx echo.Context) error {
	var request signinSchema
	if ok := BindJSON(ctx, &request); !ok {
		log.Printf("Error in Signin, BindJSON failed request: %v", &request)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if request.Email == "" || request.Password == "" {
		log.Printf("Signin request failed with empty string fields: %v, %v", request.Email, request.Password)
		return ctx.String(http.StatusBadRequest, "Email and Password cannot be empty")
	}

	account := &account.AccountEmployee{
		Email:    request.Email,
		Password: request.Password,
	}

	result, err := handler.AuthService.Signin(ctx, account)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "Signin failed.")
		return err
	}

	log.Print(result)

	return ctx.String(http.StatusOK, "Signin success")
}

func (handler *authHandler) Signout(ctx echo.Context) (err error) {
	panic("Not done yet")
}

func (handler *authHandler) Signup(ctx echo.Context) error {
	var request signupSchema
	if ok := BindJSON(ctx, &request); !ok {
		log.Printf("Error in Signup, BindJSON failed request: %v", &request)
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if request.Email == "" || request.Password == "" {
		log.Printf("Signup request failed with empty string fields: %v, %v", request.Email, request.Password)
		return ctx.String(http.StatusBadRequest, "Email and Password cannot be empty")
	}

	account := models.Account{
		AccountEmployee: &account.AccountEmployee{
			Email:    request.Email,
			Password: request.Password,
		},
		AccountIdentity: &account.AccountIdentity{},
	}

	err := handler.AuthService.CreateAccount(ctx, &account)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user account.")
		return err
	}

	return ctx.String(http.StatusOK, "Success")
}

func GetDeviceInfo() {

}
