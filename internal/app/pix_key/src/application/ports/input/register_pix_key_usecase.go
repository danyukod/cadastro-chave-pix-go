package input

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
)

type RegisterPixKeyUsecase interface {
	Execute(request.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error)
}
