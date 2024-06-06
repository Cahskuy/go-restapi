package middlewares

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/Cahskuy/go-restapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var inputValidation *validator.Validate

func init() {
	inputValidation = validator.New()
}

// ValidationHandler is a middleware that validates JSON requests
func ValidationHandler(schema interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Create a pointer to the schema struct
		payload := reflect.New(reflect.TypeOf(schema)).Interface()

		// Bind JSON payload to the schema
		if err := ctx.ShouldBindJSON(payload); err != nil {
			log.Println("Failed to bind JSON:", err)
			utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Validate the bound struct
		if err := validate(payload); err != nil {
			log.Println("Validation error:", err.Error())
			utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		// Set the payload to the context
		ctx.Set("payload", payload)
		ctx.Next()
	}
}

func validate(data interface{}) error {
	// Add new validation criteria for phone numbers
	inputValidation.RegisterValidation("phone", phoneValidation)
	// Call the Struct() method from the golang validator package to validate the received input
	err := inputValidation.Struct(data)
	// If the input received does not match the criteria, it will produce an error.
	if err != nil {
		// Take the first error from the validation error
		validationErr := err.(validator.ValidationErrors)[0]
		var errMsg string
		switch validationErr.Tag() {
		case "email":
			errMsg = "Email format is invalid"
		case "min":
			errMsg = strings.ToLower(validationErr.Field()) + " must be minimum " + validationErr.Param() + " characters"
		case "max":
			errMsg = strings.ToLower(validationErr.Field()) + " maximum allowed is " + validationErr.Param() + " characters"
		case "required":
			errMsg = strings.ToLower(validationErr.Field()) + " is required"
		case "phone":
			errMsg = "Phone number format is invalid"
		default:
			errMsg = "Invalid input for " + strings.ToLower(validationErr.Field())
		}
		return errors.New(errMsg)
	}

	return nil
}

func phoneValidation(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return false
	}
	phoneRegex, _ := regexp.Compile(`^(0|\\+62|062|62)[0-9]+$`)
	return phoneRegex.MatchString(phoneNumber)
}
