package commands

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
)

type FindPixKeyUsecase interface {
	Execute(dto.FindPixKeyDTO) (model.PixKeyDomainInterface, error)
}

type FindPixKeyService struct {
	persistence persistence.PixKeyPersistenceInterface
}

func NewFindPixKeyService(persistence persistence.PixKeyPersistenceInterface) FindPixKeyUsecase {
	return &FindPixKeyService{
		persistence: persistence,
	}
}

func (p *FindPixKeyService) Execute(dto dto.FindPixKeyDTO) (model.PixKeyDomainInterface, error) {
	var businessErrors businesserros.BusinessErrors

	pixKeyDomain, err := p.persistence.FindById(dto.Key)
	if err != nil {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	return pixKeyDomain, nil
}
