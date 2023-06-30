package factories

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller"
	persistence2 "github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/repository"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/services"
	"gorm.io/gorm"
)

func NewPixKeyControllerFactory(database *gorm.DB) controller.PixKeyControllerInterface {
	persistence := persistence2.NewPixKeyPersistence(database)
	repo := repository.NewRegisterPixKeyRepository(persistence)
	registerPixKeyUsecase := services.NewRegisterPixKeyService(repo)
	return controller.NewPixKeyControllerInterface(registerPixKeyUsecase)
}
