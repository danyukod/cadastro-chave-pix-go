package commands

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/request"
)

type FindPixKeyUsecase interface {
	Execute(request.FindPixKeyRequest) (model.PixKeyDomainInterface, error)
}

type FindPixKeyService struct {
	persistence persistence.PixKeyPersistenceInterface
}

func NewFindPixKeyService(persistence persistence.PixKeyPersistenceInterface) FindPixKeyUsecase {
	return &FindPixKeyService{
		persistence: persistence,
	}
}

func (p *FindPixKeyService) Execute(r request.FindPixKeyRequest) (model.PixKeyDomainInterface, error) {
	var businessErrors businesserros.BusinessErrors

	pixKeyDomain, err := p.persistence.FindById(r.PixKey)
	if err != nil {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	return pixKeyDomain, nil
}
