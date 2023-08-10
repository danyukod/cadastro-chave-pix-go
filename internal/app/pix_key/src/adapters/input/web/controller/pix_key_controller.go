package controller

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/input"
	"github.com/gin-gonic/gin"
)

func NewPixKeyControllerInterface(
	registerPixKeyUsecase input.RegisterPixKeyUsecase,
	findPixKeyUsecase input.FindPixKeyUsecase,
) PixKeyControllerInterface {
	return &pixKeyController{
		registerPixKeyUsecase: registerPixKeyUsecase,
		findPixKeyUsecase:     findPixKeyUsecase,
	}
}

type PixKeyControllerInterface interface {
	RegisterPixKey(c *gin.Context)
	FindPixKeyByKindAndKey(c *gin.Context)
}

type pixKeyController struct {
	registerPixKeyUsecase input.RegisterPixKeyUsecase
	findPixKeyUsecase     input.FindPixKeyUsecase
}
