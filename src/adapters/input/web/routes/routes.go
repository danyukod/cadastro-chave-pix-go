package routes

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	pixKeyController controller.PixKeyControllerInterface,
) {
	r.POST("/pix/keys", pixKeyController.RegisterPixKey)
}
