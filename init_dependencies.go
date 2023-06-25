package cadastro_chave_pix_go

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/services"
)

func initDependencies() *controller.PixKeyController {
	registerPixKeyUsecase := services.NewRegisterPixKeyService()
	return controller.NewPixKeyController(registerPixKeyUsecase)
}
