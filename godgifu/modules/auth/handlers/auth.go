package handlers

import (
	"log"
	"net/http"

	account "godgifu/modules/account/models"
	"godgifu/modules/auth/models"
	"godgifu/modules/auth/services"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	AuthService services.AuthService
	JWTService  services.JWTService
}

func NewAuthHandlers(accountServices services.AuthService, jwtServices services.JWTService) AuthHandlers {
	return &authHandler{
		AuthService: accountServices,
		JWTService:  jwtServices,
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

	payload := &account.AccountEmployee{
		Email:    request.Email,
		Password: request.Password,
	}

	accountData, err := handler.AuthService.Signin(ctx, payload)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "Signin failed.")
		return err
	}

	token, err := handler.JWTService.NewTokenPairFromAccount(ctx, accountData, "")
	log.Print(token)
	if err != nil {
		log.Print(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Signin failed.")
	}

	return ctx.JSON(http.StatusOK, token)
}

func (handler *authHandler) Signout(ctx echo.Context) (err error) {
	accountID := ctx.Get("account")

	if err := handler.JWTService.Signout(ctx, accountID.(*models.JWTToken).ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There was a problem on our end.")
	}
	return
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

	account := account.Account{
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
