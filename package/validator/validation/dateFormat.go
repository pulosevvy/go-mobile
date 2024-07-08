package validation

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func ValidateDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return false
	}

	return dateStr == parsedTime.Format("2006-01-02")
}
