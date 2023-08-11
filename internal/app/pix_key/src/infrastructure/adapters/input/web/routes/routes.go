package routes

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller"
	"github.com/gin-gonic/gin"
)

func InitPixKeyRoutes(
	r *gin.RouterGroup,
	pixKeyController controller.PixKeyControllerInterface,
) {
	r.GET("/pix/keys/:key", pixKeyController.FindPixKeyByKindAndKey)
	r.POST("/pix/keys", pixKeyController.RegisterPixKey)
}
