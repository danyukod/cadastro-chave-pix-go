package main

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/routes"
	"github.com/danyukod/cadastro-chave-pix-go/src/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func main() {
	logger.Info("About to start PixKey API...")

	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dsn := os.Getenv("DATABASE_CONNECTION_STRING")

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	pixKeyController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, pixKeyController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
