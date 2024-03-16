package model

import (
	"problem/src/internal/tools"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name        string `json:"name" validate:"required,gte=2,lte=100"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
}

func NewPerson(name, birth string) *Person {
	return &Person{name, birth}
}

func (p *Person) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return tools.NewProblem(
			"validation",
			"Estrutura com dados inválidos",
			"A entidade Pessoa que está sendo cadastrada não possui todos os dados corretos",
			"https://github.com/silasenrique/api-heper",
		).AddValidationDetails(err)
	}

	return nil
}
