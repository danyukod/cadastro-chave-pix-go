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
	{
		group := r.Group("/pix-keys")
		{
			group.Use(middleware.TimeoutMiddleware(), middleware.TokenAuthMiddleware())
			group.GET("/:key", handler.FindPixKeyByKey)
			group.POST("/", handler.RegisterPixKey)
		}
	}

}
