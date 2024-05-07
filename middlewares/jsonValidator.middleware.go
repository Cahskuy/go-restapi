package middlewares

import (
	"fmt"
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
func (cv *Validator) ValidateJSON(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}

	if err := cv.validator.Struct(req); err != nil {
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

		err := validator.ValidateJSON(ctx, payload.Interface())
		if err != nil {
			errMsg := translateError(err, trans)
			fmt.Println(errMsg)
			utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		ctx.Set("payload", payload.Interface())
		ctx.Next()
	}
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
