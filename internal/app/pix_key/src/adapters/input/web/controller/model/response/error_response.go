package response

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "business":
		return "Value " + fe.Param() + " is invalid to field " + fe.Field() + "."
	}
	return "Unknown error"
}
