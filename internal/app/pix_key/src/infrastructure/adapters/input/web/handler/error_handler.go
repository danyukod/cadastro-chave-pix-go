package handler

import (
	"errors"
	application "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validationErrors validator.ValidationErrors
var businessErrors application.BusinessErrors

func ErrorHandler(c *gin.Context, err error) {
	switch {
	case errors.As(err, &validationErrors):
		validationErrorHandler(c, validationErrors)
	case errors.As(err, &businessErrors):
		businessErrorHandler(c, businessErrors)
	default:
		defaultErrorHandler(c, err)
	}
	return
}

func defaultErrorHandler(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
}

func validationErrorHandler(c *gin.Context, err validator.ValidationErrors) {
	out := make([]ErrorResponse, len(err))
	for i, e := range err {
		out[i] = ErrorResponse{Field: e.Field(), Message: GetErrorMsg(e)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, out)
}

func businessErrorHandler(c *gin.Context, err application.BusinessErrors) {
	out := make([]ErrorResponse, len(err))
	for i, e := range err {
		out[i] = ErrorResponse{Field: e.Field(), Message: e.Error()}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, out)
}

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
	}
	return "Unknown handler"
}
