package controller

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/pixKey"
	modelrequest "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *pixKeyController) FindPixKeyByKindAndKey(c *gin.Context) {
	var request modelrequest.FindPixKeyRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.findPixKeyUsecase.Execute(request)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, pixKey.PixKeyDomainToFindWebResponse(pixKeyDomain))
}
