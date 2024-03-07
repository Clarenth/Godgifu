package handlers

import (
	"log"
	"net/http"

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

func (handler *authHandler) Signin(ctx echo.Context) (err error) {
	panic("Not done yet")
}

func (handler *authHandler) Signout(ctx echo.Context) (err error) {
	panic("Not done yet")
}

func (handler *authHandler) Signup(ctx echo.Context) error {
	var request signupSchema
	if ok := BindJSON(ctx, &request); !ok {
		log.Printf("Error in func Signup, BindJSON, request: \n %v", &request)
		return ctx.String(http.StatusBadRequest, "Failed to create user account.")
	}

	if request.Email == "" || request.Password == "" {
		return ctx.String(http.StatusBadRequest, "Email and Password cannot be empty")
	}

	account := &account.AccountEmployeeData{
		Email:    request.Email,
		Password: request.Password,
	}

	err := handler.AuthService.CreateAccount(ctx, *account)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create user account.")
	}

	return ctx.String(http.StatusOK, "Success")
}

func (auth *authHandler) Test(ctx echo.Context) (err error) {
	response := struct {
		Msg string `json:"msg"`
	}{
		Msg: "Hello, world",
	}
	return ctx.String(http.StatusOK, response.Msg)
}

func GetDeviceInfo() {

}
