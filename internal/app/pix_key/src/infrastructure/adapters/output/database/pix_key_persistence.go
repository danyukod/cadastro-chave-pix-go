package database

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/output/database/errors"
	"gorm.io/gorm"
)

type PixKeyPersistenceInterface interface {
	CreatePixKey(domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error)
	FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (domain.PixKeyDomainInterface, error)
}

func NewPixKeyPersistence(db *gorm.DB) PixKeyPersistenceInterface {
	return pixKeyPersistence{
		db,
	}
}

type pixKeyPersistence struct {
	db *gorm.DB
}

func (p pixKeyPersistence) CreatePixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	pixKeyEntity = entity.ConvertDomainToEntity(pixKeyDomain)

	err := p.db.Create(&pixKeyEntity).Error

	if err != nil {
		return nil, err
	}

	pixKeyDomain, err = entity.PixKeyDomainFromEntity(pixKeyEntity)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (p pixKeyPersistence) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (domain.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	err := p.db.Where("pix_key_type = ?", pixKeyType).Where("pix_key = ?", pixKey).First(&pixKeyEntity).Error
	if err != nil {
		return nil, errors.NewPersistenceError(pixKeyEntity.TableName(), err.Error(), "pix_key_type = "+pixKeyType+" pix_key = "+pixKey)
	}

	pixKeyDomain, err := entity.PixKeyDomainFromEntity(pixKeyEntity)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}
