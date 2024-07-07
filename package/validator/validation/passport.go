package validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func Passport(field validator.FieldLevel) bool {
	passport := strings.Split(field.Field().String(), " ")

	if len(passport) != 2 {
		return false
	}

	series := passport[0]
	number := passport[1]

	if len(series) != 4 || len(number) != 6 {
		return false
	}

	return true
}
