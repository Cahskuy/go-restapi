package middlewares

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CustomValidator holds the validator instance
type CustomValidator struct {
	validator *validator.Validate
}

// NewCustomValidator returns a new instance of CustomValidator
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

// ValidateJSON validates the JSON request body using the validator
func (cv *CustomValidator) ValidateJSON(c *gin.Context, req interface{}) error {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}

	if err := cv.validator.Struct(req); err != nil {
		return err
	}

	return nil
}

// ValidationHandler is a middleware that validates JSON requests
func ValidationHandler(validator *CustomValidator, req interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a pointer to the request struct
		reqPointer := reflect.New(reflect.TypeOf(req))
		err := validator.ValidateJSON(c, reqPointer.Interface())

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
