package factory

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/routes"
	"github.com/gin-gonic/gin"
)

func NewPixKeyRouterFactory(pixKeyController controller.PixKeyControllerInterface) error {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, pixKeyController)

	return router.Run(":8080")
}
