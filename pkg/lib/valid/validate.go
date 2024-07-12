// Package valid are valid struct fields
package valid

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// ErrorResponse holds the error response
type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

// _Validate validates the struct fields
func _Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

// Validate validates the struct fields
func Validate(data interface{}) error {
	if errs := _Validate(data); len(errs) > 0 && errs[0].Error {
		errMsg := make([]string, 0)

		for _, err := range errs {
			errMsg = append(errMsg, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return errors.New(strings.Join(errMsg, " and "))
	}
	return nil
}
