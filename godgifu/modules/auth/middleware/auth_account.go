package middleware

import (
	"godgifu/modules/auth/services"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// func AuthAccount(service services.JWTService, next echo.HandlerFunc) echo.HandlerFunc {
func AuthAccount(service services.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			header := authHeader{}
			// alternative method of retriving header data without DefaultBinder
			// header.IDToken = ctx.Request().Header.Get("Authorization")
			binder := &echo.DefaultBinder{}

			// bind the incoming authorization header and check for errors
			if err := binder.BindHeaders(ctx, &header); err != nil {
				if errs, ok := err.(validator.ValidationErrors); ok {
					var invalidArgs []invalidArgument

					for _, err := range errs {
						invalidArgs = append(invalidArgs, invalidArgument{
							err.Field(),
							err.Value().(string),
							err.Tag(),
							err.Param(),
						})
					}

					err := echo.NewHTTPError(echo.ErrBadRequest.Code, invalidArgs)
					ctx.Error(err)
					return err
				}

				//
				err := echo.NewHTTPError(echo.ErrInternalServerError.Code)
				ctx.Error(err)
				return err
			}

			log.Print(header)
			log.Print(header.IDToken)
			authHeader := strings.Split(header.IDToken, "Bearer ")
			log.Print("Hello Auth:Bearer token from middleware success") //authHeader
			if len(authHeader) < 2 {
				err := echo.NewHTTPError(echo.ErrUnauthorized.Code)
				ctx.Error(err)
				return err
			}

			// Perform IDToken validation here
			idToken, err := service.ValidateIDToken(authHeader[1])
			if err != nil {
				err := echo.NewHTTPError(echo.ErrUnauthorized.Code)
				ctx.Error(err)
				return err
			}

			ctx.Set("account", idToken)
			next(ctx)
			return nil
		}
	}
}
