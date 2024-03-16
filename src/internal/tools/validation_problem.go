package tools

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationProblem struct {
	Property   string `json:"property"`
	InputValue any    `json:"inputValue"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func GetValidationErrors(err error) *[]ValidationProblem {
	var (
		validationErrors []ValidationProblem
		ve               validator.ValidationErrors
	)

	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			validationError := ValidationProblem{
				Property:   err.Field(),
				InputValue: err.Value(),
				Type:       err.Tag(),
				Message:    getErr(err.Tag(), err.Error(), err.Param()),
			}

			validationErrors = append(validationErrors, validationError)
		}

		return &validationErrors
	}

	return nil
}

var typeErr = map[string]string{
	"gte": "O atributo deve conter no minimo %s caracteres",
	"lte": "O atributo deve ter até no máximo %s",
}

func getErr(tag string, tagErr, param string) string {
	msg := fmt.Sprintf(typeErr[tag], param)

	if msg != "" {
		return msg
	}

	return tagErr
}
