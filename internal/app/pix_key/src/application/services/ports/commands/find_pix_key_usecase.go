package commands

import (
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
)

type FindPixKeyUsecase interface {
	Execute(request.FindPixKeyRequest) (domain.PixKeyDomainInterface, error)
}

type FindPixKeyService struct {
}

func NewFindPixKeyService() FindPixKeyUsecase {
	return &FindPixKeyService{}
}

func (p *FindPixKeyService) Execute(request.FindPixKeyRequest) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserros.BusinessErrors

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	return nil, nil
}
