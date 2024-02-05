package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/vinialeixo/email-markegint/internal/domain/campaign"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		fmt.Println("No erros found it")
	} else {
		validationErros := err.(validator.ValidationErrors)
		for _, v := range validationErros {
			switch v.Tag() {
			case "required":
				fmt.Println(v.StructField() + " is required")
			case "min":
				fmt.Println(v.StructField() + " is required with min: " + v.Param())
			case "max":
				fmt.Println(v.StructField() + " is required with max: " + v.Param())
			case "email":
				fmt.Println(v.StructField() + " is invalid")
			}
		}
	}
}
