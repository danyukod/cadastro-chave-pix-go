package factory

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/input"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/output/database"
	"gorm.io/gorm"
)

func NewPixKeyControllerFactory(gormDB *gorm.DB) controller.PixKeyControllerInterface {
	persistence := database.NewPixKeyPersistence(gormDB)
	repo := output.NewRegisterPixKeyRepository(persistence)
	registerPixKeyUseCase := input.NewRegisterPixKeyService(repo)
	findPixKeyUseCase := input.NewFindPixKeyService()
	return controller.NewPixKeyControllerInterface(registerPixKeyUseCase, findPixKeyUseCase)
}
