package tools

// import (
// 	"log"
// 	"net/http"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/labstack/echo/v4"
// )

// func SchemaCompare(ctx echo.Context, req interface{}, schema interface{}) (interface{}, error) {
// 	// Bind incoming json to struct and check for validation errors
// 	if err := ctx.Bind(req); err != nil {
// 		log.Printf("Error binding data: %+v\n", err)

// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			// could probably extract this, it is also in middleware_auth_user
// 			var invalidArgs []invalidArgument

// 			for _, err := range errs {
// 				invalidArgs = append(invalidArgs, invalidArgument{
// 					err.Field(),
// 					err.Value().(string),
// 					err.Tag(),
// 					err.Param(),
// 				})
// 			}

// 			// err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

// 			ctx.JSON(http.StatusBadRequest, invalidArgs)
// 			return false
// 		}

// 		// later we'll add code for validating max body size here!

// 		// if we aren't able to properly extract validation errors,
// 		// we'll fallback and return an internal server error
// 		ctx.JSON(http.StatusInternalServerError, "sheet")
// 		return false
// 	}
// 	if err := ctx.Validate(req); err != nil {
// 		return false
// 	}

// 	/*
// 		// Leave below incase we need to instanciate the validator here instead of in config
// 		v := validator.New()
// 		if err := v.Struct(req); err != nil {
// 			return false
// 		}
// 	*/
// 	return true
// }

// import (
// 	"fmt"
// 	"reflect"

// 	"github.com/labstack/echo/v4"
// )

// // SchemaCompare compares the fields of the JSON request against the schema struct supplied in args it must be equal with
// func SchemaCompare(ctx echo.Context, req interface{}, schema ...interface{}) bool {
// 	panic("not done")

// 	/*
// 		Perhaps we cam compare multiple schemas here that way we know what is true and what is not
// 		If we return the correct schema struct then we can continue on with our day and not write
// 		more code in the function to handle checking the two.
// 		Perform the comparison like in BindData
// 	*/

// 	if reflect.DeepEqual(struct1, struct2) {
// 		fmt.Println("The JSON data matches one of the structs.")
// 	} else {
// 		fmt.Println("The JSON data does not match either struct.")
// 	}
// }
