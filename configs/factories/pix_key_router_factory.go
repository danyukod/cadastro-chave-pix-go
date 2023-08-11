package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller/factory"
	routes2 "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPixKeyRouterFactory(database *gorm.DB) error {

	pixKeyController := factory.NewPixKeyControllerFactory(database)

	router := gin.Default()
	routes2.InitPixKeyRoutes(&router.RouterGroup, pixKeyController)

	return router.Run(":8080")
}
