package factory

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/services/ports/commands"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapter/orm"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/handler"
	"gorm.io/gorm"
)

func NewPixKeyHandlerFactory(gormDB *gorm.DB) handler.PixKeyHandlerInterface {
	persistence := orm.NewPixKeyPersistence(gormDB)
	registerPixKeyUseCase := commands.NewRegisterPixKeyService(persistence)
	findPixKeyUseCase := commands.NewFindPixKeyService(persistence)
	return handler.NewPixKeyHandlerInterface(registerPixKeyUseCase, findPixKeyUseCase)
}
