package pkg_utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateHttpReqBody(input any) []map[string]string {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return FormatValidationErrors(err)
	}
	return nil
}

func FormatValidationErrors(err error) []map[string]string {
	var errors []map[string]string
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			errors = append(errors, map[string]string{
				"field":   fieldErr.Field(),
				"message": fmt.Sprintf("Validation failed on the '%s' tag", fieldErr.Tag()),
			})
		}
	}
	return errors
}
