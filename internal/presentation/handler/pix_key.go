package handler

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands"
	modelrequest "github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PixKeyHandlerInterface interface {
	FindPixKeyByKey(c *gin.Context)
	RegisterPixKey(c *gin.Context)
}

func NewPixKeyHandlerInterface(
	registerPixKeyUsecase commands.RegisterPixKeyUsecase,
	findPixKeyUsecase commands.FindPixKeyUsecase,
) PixKeyHandlerInterface {
	return &handler{
		registerPixKeyUsecase: registerPixKeyUsecase,
		findPixKeyUsecase:     findPixKeyUsecase,
	}
}

type handler struct {
	registerPixKeyUsecase commands.RegisterPixKeyUsecase
	findPixKeyUsecase     commands.FindPixKeyUsecase
}

func (p *handler) RegisterPixKey(c *gin.Context) {
	var request modelrequest.RegisterPixKeyRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.registerPixKeyUsecase.Execute(request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.PixKeyDomainToRegisterWebResponse(pixKeyDomain))
}

func (p *handler) FindPixKeyByKey(c *gin.Context) {
	var request modelrequest.FindPixKeyRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	pixKeyDomain, err := p.findPixKeyUsecase.Execute(request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, response.PixKeyDomainToFindWebResponse(pixKeyDomain))
}
