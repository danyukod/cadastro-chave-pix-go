package handler

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
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

}

func defaultErrorHandler(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
}

func validationErrorHandler(c *gin.Context, err validator.ValidationErrors) {
	out := make([]response.ErrorResponse, len(err))
	for i, e := range err {
		out[i] = response.ErrorResponse{Field: e.Field(), Message: response.GetErrorMsg(e)}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, out)
}

func businessErrorHandler(c *gin.Context, err application.BusinessErrors) {
	out := make([]response.ErrorResponse, len(err))
	for i, e := range err {
		out[i] = response.ErrorResponse{Field: e.Field(), Message: e.Error()}
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, out)
}
