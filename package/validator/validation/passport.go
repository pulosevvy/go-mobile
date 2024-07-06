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

	number := passport[0]
	serial := passport[1]

	if len(number) != 4 || len(serial) != 6 {
		return false
	}

	return true
}
