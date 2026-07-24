package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type FieldErr struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func Parse(err error) []FieldErr {
	var errs validator.ValidationErrors

	if !errors.As(err, &errs) {
		return []FieldErr{
			{
				Field: "request",
				Msg:   "Invalid request",
			},
		}
	}

	res := make([]FieldErr, 0, len(errs))

	for _, e := range errs {
		res = append(res, FieldErr{
			Field: e.Field(),
			Msg:   msg(e),
		})
	}

	return res
}

func msg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"

	case "email":
		return "Must be a valid email address"

	case "min":
		return fmt.Sprintf(
			"Must be at least %s characters",
			err.Param(),
		)

	case "max":
		return fmt.Sprintf(
			"Must be at most %s characters",
			err.Param(),
		)

	case "oneof":
		return fmt.Sprintf(
			"Must be one of: %s",
			err.Param(),
		)

	default:
		return "Invalid value"
	}
}
