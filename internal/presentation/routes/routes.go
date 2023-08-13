package routes

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler"
	"github.com/gin-gonic/gin"
)

func InitPixKeyRoutes(
	r *gin.RouterGroup,
	handler handler.PixKeyHandlerInterface,
) {
	r.GET("/pix/keys/:key", handler.FindPixKeyByKey)
	r.POST("/pix/keys", handler.RegisterPixKey)
}
