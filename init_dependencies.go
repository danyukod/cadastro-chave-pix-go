package main

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/repository"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/services"
	"gorm.io/gorm"
)

func initDependencies(
	database *gorm.DB,
) *controller.PixKeyController {
	repo := repository.NewPixKeyRepository(database)
	registerPixKeyUsecase := services.NewRegisterPixKeyService(repo)
	return controller.NewPixKeyController(registerPixKeyUsecase)
}
