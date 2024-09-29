package requests

import "github.com/go-playground/validator/v10"

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			errors[e.Field()] = getErrorMessage(e)
		}
	}
	return errors
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters long"
	case "eqfield":
		return e.Field() + " must be equal to " + e.Param()
	default:
		return "Invalid " + e.Field()
	}
}
