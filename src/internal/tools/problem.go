package tools

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)

type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`

	Errors []ProblemDetail `json:"errors"`
}

func (p Problem) ToHttpError() ([]byte, error) {
	return json.Marshal(&p)
}

func (p Problem) Error() string {
	return ""
}

func (p *Problem) AddValidationDetails(err error) *Problem {
	var (
		ve validator.ValidationErrors
	)

	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			validationError := ProblemDetail{
				Field:  err.Field(),
				Value:  err.Value(),
				Type:   err.Tag(),
				Detail: getErr(err.Tag(), err.Error(), err.Param()),
			}

			p.Errors = append(p.Errors, validationError)
		}
	}

	return p
}

func NewProblem(problemType, title, detail, instance string) *Problem {
	return &Problem{
		Type:     problemType,
		Title:    title,
		Detail:   detail,
		Instance: instance,
	}
}
