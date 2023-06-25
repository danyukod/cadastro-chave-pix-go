package main

import (
	main2 "github.com/danyukod/cadastro-chave-pix-go"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/routes"
	"github.com/danyukod/cadastro-chave-pix-go/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	logger.Info("About to start PixKey API...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	pixKeyController := main2.initDependencies()

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, pixKeyController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
