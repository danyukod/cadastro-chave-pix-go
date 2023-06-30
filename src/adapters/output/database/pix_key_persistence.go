package database

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/mapper"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
	"gorm.io/gorm"
)

type PixKeyPersistence struct {
	db *gorm.DB
}

func NewPixKeyPersistence(db *gorm.DB) *PixKeyPersistence {
	return &PixKeyPersistence{
		db,
	}
}

func (p PixKeyPersistence) CreatePixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	pixKeyEntity = mapper.ConvertDomainToEntity(pixKeyDomain)

	err := p.db.Create(&pixKeyEntity).Error

	if err != nil {
		return nil, err
	}

	pixKeyDomain, err = mapper.ConvertEntityToDomain(pixKeyEntity)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}
