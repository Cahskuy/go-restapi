package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
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
	inputValidation.RegisterValidation("phone", phoneValidation)
}

// ValidationHandler is a middleware that validates JSON requests
func ValidationHandler(schema interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Create a pointer to the schema struct
		payload := reflect.New(reflect.TypeOf(schema)).Interface()

		// Bind JSON payload to the schema
		if err := ctx.ShouldBindJSON(payload); err != nil {
			log.Println("Failed to bind JSON:", err)
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				// Extract the field name and type from the error
				errorMessage := fmt.Sprintf("Field '%s' should be of type '%s'", err.(*json.UnmarshalTypeError).Field, err.(*json.UnmarshalTypeError).Type)
				utils.ErrorResponse(ctx, http.StatusBadRequest, errorMessage)
				return
			}

			utils.ErrorResponse(ctx, http.StatusBadRequest, "Error decoding JSON")
			return
		}

		// Validate the bound struct
		if err := validate(payload); err != nil {
			log.Println("Validation error:", err.Error())
			utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		log.Println(payload)
		// Set the payload to the context
		ctx.Set("payload", payload)
		ctx.Next()
	}
}

func validate(data interface{}) error {
	err := inputValidation.Struct(data)
	if err != nil {
		validationErr := err.(validator.ValidationErrors)[0]
		var errMsg string
		switch validationErr.Tag() {
		case "required":
			errMsg = strings.ToLower(validationErr.Field()) + " is required"
		case "min":
			switch validationErr.Type().String() {
			case "string":
				errMsg = strings.ToLower(validationErr.Field()) + " must be minimum " + validationErr.Param() + " characters"
			case "*int":
			case "int":
				errMsg = strings.ToLower(validationErr.Field()) + " must be >= " + validationErr.Param()
			default:
				errMsg = strings.ToLower(validationErr.Field()) + " must be minimum " + validationErr.Param()
			}
		case "max":
			errMsg = strings.ToLower(validationErr.Field()) + " maximum allowed is " + validationErr.Param() + " characters"
		case "email":
			errMsg = "Email format is invalid"
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
	phoneRegex := regexp.MustCompile(`^(0|\\+62|062|62)[0-9]+$`)
	return phoneRegex.MatchString(phoneNumber)
}

func BindJSONCaseSensitive(c *gin.Context, obj interface{}) error {
	// Ensure obj is a pointer to a struct
	objType := reflect.TypeOf(obj)

	// Get JSON data from request
	raw := make(map[string]json.RawMessage)
	if err := c.ShouldBindJSON(&raw); err != nil {
		return err
	}

	// Get the type of the struct
	typ := objType.Elem()

	// Create a new instance of the struct
	value := reflect.New(typ).Elem()

	// Iterate over JSON keys and set struct fields
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldType := strings.ReplaceAll(field.Type.String(), "*", "")
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		jsonName := jsonTag
		if rawValue, ok := raw[jsonName]; ok {
			if err := json.Unmarshal(rawValue, value.Field(i).Addr().Interface()); err != nil {
				return errors.New(jsonTag + " must be " + fieldType)
			}
		}
	}

	// Set the struct instance to the object
	reflect.ValueOf(obj).Elem().Set(value)

	return nil
}
