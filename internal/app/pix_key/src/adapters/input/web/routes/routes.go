package routes

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	pixKeyController controller.PixKeyControllerInterface,
) {
	r.POST("/pix/keys", pixKeyController.RegisterPixKey)
}
