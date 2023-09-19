package factory

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler"
	"gorm.io/gorm"
)

func NewPixKeyHandlerFactory(gormDB *gorm.DB) handler.PixKeyHandlerInterface {
	pixKeyPersistence := persistence.NewPixKeyPersistence(gormDB)
	registerPixKeyUseCase := commands.NewRegisterPixKeyService(pixKeyPersistence)
	findPixKeyByKeyUseCase := commands.NewFindPixKeyByKeyService(pixKeyPersistence)
	findPixKeyUseCase := commands.NewFindPixKeyService(pixKeyPersistence)
	return handler.NewPixKeyHandlerInterface(registerPixKeyUseCase, findPixKeyByKeyUseCase, findPixKeyUseCase)
}
