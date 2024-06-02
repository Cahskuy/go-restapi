package middlewares

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/Cahskuy/go-restapi/models"
	"github.com/Cahskuy/go-restapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// InputValidation holds the validator v10 instance
type InputValidation struct {
	Validator *validator.Validate
}

// NewInputValidation returns a new instance of Validator
func NewInputValidation() *InputValidation {
	return &InputValidation{Validator: validator.New()}
}

// ValidationHandler is a middleware that validates JSON requests
func (iv *InputValidation) ValidationHandler(schema interface{}) gin.HandlerFunc {
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
		if err := validate(iv, payload); err != nil {
			log.Println("Validation error:", err.Error())
			utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		// Set the payload to the context
		ctx.Set("payload", payload)
		ctx.Next()
	}
}

// Method Validate() adalah logic dari validasi input yang akan kita buat
// Selain itu, kita dapat membuat logic validasi input secara kustom
func validate(iv *InputValidation, data interface{}) error {
	// Buat sebuah slice dari struct untuk menampung error input
	// Error dalam input bisa hanya satu atau lebih dari satu
	var errFields []models.ErrorInputResponse

	// Menambah kriteria validasi baru agar dapat memvalidasi nomor telepon
	iv.Validator.RegisterValidation("phone", phoneValidation)
	// Kita akan memanggil method Struct() dari package golang validator untuk memvalidasi input yang diterima
	err := iv.Validator.Struct(data)
	// Jika input yang diterima tidak sesuai dengan kriteria, maka akan menghasilkan error.
	// Kita dapat membuat pesan error secara custom berdasarkan kriteria yang sudah diberikan
	if err != nil {
		log.Println(err)
		for _, err := range err.(validator.ValidationErrors) {
			var errField models.ErrorInputResponse
			switch err.Tag() {
			case "email":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = "Email format is invalid"
			case "min":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = strings.ToLower(err.Field()) + " must be minimum " + err.Param() + " characters"
			case "max":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = strings.ToLower(err.Field()) + " maximum allowed is" + err.Param() + " characters"
			case "required":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = strings.ToLower(err.Field()) + " is required"
			case "phone":
				errField.FieldName = strings.ToLower(err.Field())
				errField.Message = "Phone number format is invalid"
			}
			errFields = append(errFields, errField)
		}
	}
	// Kalau tidak ada error, kita bisa mengembalikan nilainya menjadi nil
	if len(errFields) == 0 {
		return nil
	}
	// Jadikan slice dari errFields menjadi JSON Array of Objects
	// Hasilnya bisa digunakan oleh Frontend Developer untuk membuat pesan error input.
	marshaledErr, _ := json.Marshal(errFields)
	return errors.New(string(marshaledErr))
}

func phoneValidation(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if len(phoneNumber) < 10 || len(phoneNumber) > 13 {
		return false
	}
	phoneRegex, _ := regexp.Compile(`^(0|\\+62|062|62)[0-9]+$`)
	return phoneRegex.MatchString(phoneNumber)
}
