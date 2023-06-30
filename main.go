package main

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/factory"
	"github.com/danyukod/cadastro-chave-pix-go/src/shared/logger"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
)

func main() {
	logger.Info("About to start PixKey API...")

	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := factory.NewPixKeyDatabaseFactory()
	if err != nil {
		log.Fatal(err)
	}

	pixKeyController := factory.NewPixKeyControllerFactory(database)

	err = factory.NewPixKeyRouterFactory(pixKeyController)
	if err != nil {
		log.Fatal(err)
	}
}
