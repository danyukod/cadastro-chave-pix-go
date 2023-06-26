package repository

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/mapper"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
)

func (p pixKeyRepository) CreatePixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	pixKeyEntity = mapper.ConvertDomainToEntity(pixKeyDomain)

	p.database.Create(&pixKeyEntity)

	pixKeyDomain.SetID(pixKeyEntity.ID)

	return pixKeyDomain, nil
}
