package controller

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/input"
	"github.com/gin-gonic/gin"
)

func NewPixKeyControllerInterface(
	usecase input.RegisterPixKeyUsecase,
) PixKeyControllerInterface {
	return &pixKeyController{
		usecase: usecase,
	}
}

type PixKeyControllerInterface interface {
	RegisterPixKey(c *gin.Context)
}

type pixKeyController struct {
	usecase input.RegisterPixKeyUsecase
}
