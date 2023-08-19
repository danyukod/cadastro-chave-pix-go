package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/configs"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/factory"
	routes2 "github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPixKeyRouterFactory(database *gorm.DB, config configs.Config) error {

	pixKeyController := factory.NewPixKeyHandlerFactory(database)

	router := gin.Default()
	routes2.InitPixKeyRoutes(&router.RouterGroup, pixKeyController, config.GetJWTSecret())

	return router.Run(":8080")
}
