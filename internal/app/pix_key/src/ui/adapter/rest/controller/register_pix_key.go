package controller

import (
	modelrequest "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *pixKeyController) RegisterPixKey(c *gin.Context) {
	var request modelrequest.RegisterPixKeyRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.registerPixKeyUsecase.Execute(request)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.PixKeyDomainToRegisterWebResponse(pixKeyDomain))
}
