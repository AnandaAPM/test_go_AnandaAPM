package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Validator struct{
	Validate *validator.Validate
}

func Valid() *Validator{
	v := validator.New()

	v.RegisterValidation("birthdate", func(fl validator.FieldLevel) bool {
		date, ok := fl.Field().Interface().(time.Time)
		return ok && !date.IsZero()
	})

	return &Validator{Validate:v}
}