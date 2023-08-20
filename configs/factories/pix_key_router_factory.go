package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/docs"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/factory"
	routes2 "github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewPixKeyRouterFactory(database *gorm.DB) error {

	pixKeyController := factory.NewPixKeyHandlerFactory(database)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	routes2.InitPixKeyRoutes(v1, pixKeyController)

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router.Run(":8080")
}
