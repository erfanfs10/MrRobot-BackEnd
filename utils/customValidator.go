package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func CreateCustomValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func CustomValidationError(c echo.Context, code int, err error, message string) error {

	// Collect validation errors
	validationErrors := []ValidationError{}
	ve, _ := err.(validator.ValidationErrors)
	for _, fe := range ve {
		validationErrors = append(validationErrors, ValidationError{
			Field:   fe.Field(),
			Message: getCustomMessage(fe), // Get a custom message for the error
		})
	}
	return c.JSON(code, map[string]interface{}{
		"message": message,
		"errors":  validationErrors,
	})
}

// Function to return custom error messages based on field and tag
func getCustomMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email format"
	case "gte":
		return "between 8 to 10 characters"
	case "lte":
		return "between 8 to 10 characters"
	default:
		return fe.Field() + " is not valid"
	}
}
