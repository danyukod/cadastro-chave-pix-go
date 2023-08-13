package factory

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler"
	"gorm.io/gorm"
)

func NewPixKeyHandlerFactory(gormDB *gorm.DB) handler.PixKeyHandlerInterface {
	persistence := persistence.NewPixKeyPersistence(gormDB)
	registerPixKeyUseCase := commands.NewRegisterPixKeyService(persistence)
	findPixKeyUseCase := commands.NewFindPixKeyService(persistence)
	return handler.NewPixKeyHandlerInterface(registerPixKeyUseCase, findPixKeyUseCase)
}
