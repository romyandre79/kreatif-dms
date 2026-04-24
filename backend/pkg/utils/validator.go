package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func FormatValidationErrors(errs []*ErrorResponse) string {
	if len(errs) == 0 {
		return ""
	}
	var msg string
	for _, e := range errs {
		msg += fmt.Sprintf("Field '%s' failed on the '%s' tag. ", e.Field, e.Tag)
	}
	return msg
}
