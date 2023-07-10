package database

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/mapper"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
	"gorm.io/gorm"
)

func NewPixKeyPersistence(db *gorm.DB) output.PixKeyPersistenceInterface {
	return pixKeyPersistence{
		db,
	}
}

type pixKeyPersistence struct {
	db *gorm.DB
}

func (p pixKeyPersistence) CreatePixKey(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
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

func (p pixKeyPersistence) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	err := p.db.Where("pix_key_type = ?", pixKeyType).Where("pix_key = ?", pixKey).First(&pixKeyEntity).Error
	if err != nil {
		return nil, err
	}

	pixKeyDomain, err := mapper.ConvertEntityToDomain(pixKeyEntity)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}
