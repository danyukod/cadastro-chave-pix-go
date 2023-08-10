package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPixKeyRouterFactory(database *gorm.DB) error {

	pixKeyController := routes.NewPixKeyControllerFactory(database)

	router := gin.Default()
	routes.InitPixKeyRoutes(&router.RouterGroup, pixKeyController)

	return router.Run(":8080")
}
