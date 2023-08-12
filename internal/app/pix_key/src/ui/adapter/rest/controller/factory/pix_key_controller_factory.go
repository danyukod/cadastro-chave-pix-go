package factory

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/services/ports/commands"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapter/orm"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller"
	"gorm.io/gorm"
)

func NewPixKeyControllerFactory(gormDB *gorm.DB) controller.PixKeyControllerInterface {
	persistence := orm.NewPixKeyPersistence(gormDB)
	repo := persistence.NewRegisterPixKeyRepository(persistence)
	registerPixKeyUseCase := commands.NewRegisterPixKeyService(repo)
	findPixKeyUseCase := commands.NewFindPixKeyService()
	return controller.NewPixKeyControllerInterface(registerPixKeyUseCase, findPixKeyUseCase)
}
