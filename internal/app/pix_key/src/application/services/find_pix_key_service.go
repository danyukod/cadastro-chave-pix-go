package services

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/input"
)

type FindPixKeyService struct {
}

func NewFindPixKeyService() input.FindPixKeyUsecase {
	return &FindPixKeyService{}
}

func (p *FindPixKeyService) Execute(request.FindPixKeyRequest) (*response.FindPixKeyResponse, error) {
	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	return nil, nil
}
