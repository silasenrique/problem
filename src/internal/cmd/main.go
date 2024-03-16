package main

import (
	"problem/src/internal/model"
	"problem/src/internal/tools"
)

func main() {
	problem := model.NewPerson("a", "")
	err := problem.Validate()
	if err != nil {
		println(string(err.(*tools.Problem).ToHttpError()))
	}
}
