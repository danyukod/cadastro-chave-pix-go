package input

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/response"
)

type RegisterPixKeyUsecase interface {
	Execute(request.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error)
}
