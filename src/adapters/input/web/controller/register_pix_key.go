package controller

import (
	"errors"
	modelrequest "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/response"
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/src/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (p *pixKeyController) RegisterPixKey(c *gin.Context) {
	var request modelrequest.RegisterPixKeyRequest

	if validateRequest(c, request) {
		return
	}

	pixKeyResponse, err := p.usecase.Execute(request)

	if validateResponse(c, err) {
		return
	}

	c.JSON(http.StatusCreated, pixKeyResponse)
}

func validateResponse(c *gin.Context, err error) bool {
	if err != nil {
		var eb businesserrors.BusinessErrors
		if !errors.As(err, &eb) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return true
		}
		out := make([]response.ErrorResponse, len(eb))
		for i, e := range eb {
			out[i] = response.ErrorResponse{Field: e.Field, Message: e.Message}
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return true
	}
	return false
}

func validateRequest(c *gin.Context, request modelrequest.RegisterPixKeyRequest) bool {
	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]response.ErrorResponse, len(ve))
			for i, e := range ve {
				out[i] = response.ErrorResponse{Field: e.Field(), Message: response.GetErrorMsg(e)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return true
	}
	return false
}
