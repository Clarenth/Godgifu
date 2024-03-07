package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// bindData is a helper function that returns true or false if the
// incoming request data is not bound, i.e. is not JSON.
func BindJSON(ctx echo.Context, req interface{}) bool {
	if ctx.Request().Header.Get("Content-Type") != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", ctx.Path())

		// err := apperrors.NewUnsupportedMediaType(msg)
		ctx.JSON(http.StatusUnsupportedMediaType, msg)
		return false
	}
	// Bind incoming json to struct and check for validation errors
	if err := ctx.Bind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			// could probably extract this, it is also in middleware_auth_user
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			// err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

			ctx.JSON(http.StatusBadRequest, invalidArgs)
			return false
		}

		// later we'll add code for validating max body size here!

		// if we aren't able to properly extract validation errors,
		// we'll fallback and return an internal server error
		ctx.JSON(http.StatusInternalServerError, "test2")
		return false
	}
	if err := ctx.Validate(req); err != nil {
		return false
	}

	/*
		// Leave below incase we need to instanciate the validator here instead of in config
		v := validator.New()
		if err := v.Struct(req); err != nil {
			return false
		}
	*/
	return true
}
