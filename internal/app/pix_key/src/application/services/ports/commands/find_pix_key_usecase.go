package commands

import (
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapter/orm"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
)

type FindPixKeyUsecase interface {
	Execute(request.FindPixKeyRequest) (domain.PixKeyDomainInterface, error)
}

type FindPixKeyService struct {
	persistence orm.PixKeyPersistenceInterface
}

func NewFindPixKeyService(persistence orm.PixKeyPersistenceInterface) FindPixKeyUsecase {
	return &FindPixKeyService{
		persistence: persistence,
	}
}

func (p *FindPixKeyService) Execute(r request.FindPixKeyRequest) (domain.PixKeyDomainInterface, error) {
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
