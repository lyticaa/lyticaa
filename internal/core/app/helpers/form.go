package helpers

import (
	"reflect"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/go-playground/validator.v9"
)

type ValidateDateRange struct {
	DateRange string `validate:"required,oneof=today last_thirty_days this_month last_month last_three_months last_six_months this_year all_time"`
}

func ValidateInput(data interface{}, l *zerolog.Logger) (bool, map[string]string) {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			l.Panic().Err(err).Msg("unable to validate inputs")
		}

		errors := make(map[string]string)
		reflected := reflect.ValueOf(data)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())

			var name string
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = name + " is required"
			default:
				errors[name] = name + " is invalid"
			}
		}

		return false, errors
	}

	return true, nil
}
