package main

import (
	"github.com/danyukod/cadastro-chave-pix-go/configs"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/factories"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/logger"
	"log"
)

func main() {
	logger.Info("About to start PixKey API...")

	conf, err := configs.LoadConfig("cmd/server")

	//err = configs.MigrateDatabase(conf)
	//if err != nil {
	//	log.Fatal(err)
	//}

	database, err := configs.NewPixKeyDatabaseFactory(conf)
	if err != nil {
		log.Fatal(err)
	}

	pixKeyController := factories.NewPixKeyControllerFactory(database)

	err = factories.NewPixKeyRouterFactory(pixKeyController)
	if err != nil {
		log.Fatal(err)
	}
}
