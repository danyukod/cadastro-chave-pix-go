package main

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/factories"
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

	database, err := factories.NewPixKeyDatabaseFactory()
	if err != nil {
		log.Fatal(err)
	}

	pixKeyController := factories.NewPixKeyControllerFactory(database)

	err = factories.NewPixKeyRouterFactory(pixKeyController)
	if err != nil {
		log.Fatal(err)
	}
}
