package middlewares

import (
	"net/http"
	"reflect"

	"github.com/Cahskuy/go-restapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validator holds the validator v10 instance
type Validator struct {
	validator *validator.Validate
}

// NewValidator returns a new instance of Validator
func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

// ValidateJSON validates the JSON request body using the validator
func (v *Validator) validateJSON(ctx *gin.Context, req interface{}) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		return err
	}

	if err := v.validator.Struct(req); err != nil {
		return err
	}

	return nil
}

// ValidationHandler is a middleware that validates JSON requests
func ValidationHandler(validator *Validator, schema interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Create a pointer to the request struct
		payload := reflect.New(reflect.TypeOf(schema))

		en := en.New()
		uni := ut.New(en, en)
		trans, _ := uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(validator.validator, trans)

		err := validator.validateJSON(ctx, payload.Interface())
		if err != nil {
			errMsgs := translateError(err, trans)
			if len(errMsgs) > 0 {
				// access the first error message
				firstErrMsg := errMsgs[0]
				utils.ErrorResponse(ctx, http.StatusBadRequest, firstErrMsg)
				return
			}

			utils.ErrorResponse(ctx, http.StatusBadRequest, "Error translating validation errors")
			return
		}

		ctx.Set("payload", payload.Interface())
		ctx.Next()
	}
}

func translateError(err error, trans ut.Translator) []string {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	var errMsgs []string
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errMsgs = append(errMsgs, translatedErr)
	}
	return errMsgs
}
