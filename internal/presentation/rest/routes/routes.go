package routes

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler"
	"github.com/danyukod/cadastro-user-go/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitPixKeyRoutes(
	r *gin.RouterGroup,
	handler handler.PixKeyHandlerInterface,
) {
	apiGroup := r.Group("pix/keys")
	{
		apiGroup.Use(middleware.TokenAuthMiddleware())
		apiGroup.GET("/:key", handler.FindPixKeyByKey)
		apiGroup.POST("/", handler.RegisterPixKey)
	}

}
