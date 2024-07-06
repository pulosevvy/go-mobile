package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func IsUuid(uuid validator.FieldLevel) bool {
	regex := regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")
	return regex.MatchString(uuid.Field().String())
}
