package controller

import (
	"errors"
	modelrequest "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (p *pixKeyController) RegisterPixKey(c *gin.Context) {
	var request modelrequest.RegisterPixKeyRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]response.ErrorResponse, len(ve))
			for i, e := range ve {
				out[i] = response.ErrorResponse{Field: e.Field(), Message: response.GetErrorMsg(e)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	pixKeyResponse, err := p.usecase.Execute(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pixKeyResponse)
}