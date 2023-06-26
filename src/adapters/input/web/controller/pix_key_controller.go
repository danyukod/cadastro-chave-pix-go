package controller

import (
	modelrequest "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/ports/input"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PixKeyController struct {
	RegisterPixKeyUseCase input.RegisterPixKeyUsecase
}

func NewPixKeyController(registerPixKeyUseCase input.RegisterPixKeyUsecase) *PixKeyController {
	return &PixKeyController{
		RegisterPixKeyUseCase: registerPixKeyUseCase,
	}
}

func (p *PixKeyController) RegisterPixKey(c *gin.Context) {
	var request modelrequest.RegisterPixKeyRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := p.RegisterPixKeyUseCase.Execute(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}
