package main

import (
	"github.com/danyukod/cadastro-chave-pix-go/configs"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/factories"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/logger"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := startPixKeyAPI()
	if err != nil {
		return
	}
}

func startPixKeyAPI() error {
	logger.Info("About to start PixKey API...")

	conf, err := loadConfig("cmd/server")
	if err != nil {
		logger.Error("Failed to load config: ", err)
		return err
	}

	if err := configs.MigrateDatabase(*conf); err != nil {
		logger.Error("Database migration failed: ", err)
		return err
	}

	database := newPixKeyDatabaseFactory(conf)

	pixKeyController := factories.NewPixKeyControllerFactory(database)

	if err := factories.NewPixKeyRouterFactory(pixKeyController); err != nil {
		logger.Error("Failed to start PixKey Router: ", err)
		return err
	}

	return nil
}

func loadConfig(filePath string) (*configs.Config, error) {
	conf, err := configs.LoadConfig(filePath)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func newPixKeyDatabaseFactory(conf *configs.Config) *gorm.DB {
	database, err := configs.NewPixKeyDatabaseFactory(*conf)
	checkForError(err)
	return database
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
