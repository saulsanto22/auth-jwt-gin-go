package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type validationResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(ctx *gin.Context, s interface{}) bool {
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		var errors []validationResponse

		for _, fieldErr := range err.(validator.ValidationErrors) {
			var msg string

			switch fieldErr.Tag() {
			case "required":
				msg = fmt.Sprintf("field %s is required", fieldErr.Field())
			case "email":
				msg = fmt.Sprintf("field %s must be a valid email", fieldErr.Field())
			case "min":
				msg = fmt.Sprintf("field %s must be at least %s characters long", fieldErr.Field(), fieldErr.Param())
			case "max":
				msg = fmt.Sprintf("field %s must be at most %s characters long", fieldErr.Field(), fieldErr.Param())
			default:
				msg = fmt.Sprintf("field %s is invalid", fieldErr.Field())

			}

			errors = append(errors, validationResponse{
				Field: fieldErr.Field(), Message: msg})
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "errors": errors})
		return false

	}
	return true
}
