package main

import (
	"fmt"
	"log"

	"gopkg.in/go-playground/validator.v9"
)

type Fruit struct {
	ID   int    `validate:"min=1,max=99"`
	Name string `validate:"max=9"`
}

func main() {
	validate := validator.New()

	orange := Fruit{
		ID:   10,
		Name: "orange",
	}

	if err := validate.Struct(orange); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", orange)
}
