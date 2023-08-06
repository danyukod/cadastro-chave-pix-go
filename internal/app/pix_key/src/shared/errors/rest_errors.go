package errors

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CheckForError(c *gin.Context, err error) bool {
	if err != nil {
		out := createErrorResponse(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, out)
		return true
	}
	return false
}

func createErrorResponse(err error) []response.ErrorResponse {
	var ve validator.ValidationErrors
	var be BusinessErrors
	if errors.As(err, &ve) || errors.As(err, &be) {
		out := make([]response.ErrorResponse, len(ve)+len(be))
		getResponse(be, out)
		getResponse(ve, out)
		return out
	}
	return []response.ErrorResponse{{Message: err.Error()}}
}

func getResponse[errors validator.ValidationErrors | BusinessErrors](e errors, out []response.ErrorResponse) {
	for i, e := range e {
		out[i] = response.ErrorResponse{Field: e.Field(), Message: response.GetErrorMsg(e)}
	}
}
