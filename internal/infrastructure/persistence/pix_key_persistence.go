package persistence

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence/entity"
	orm_errors "github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence/errors"
	"gorm.io/gorm"
)

type PixKeyPersistenceInterface interface {
	CreatePixKey(model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error)
	FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (model.PixKeyDomainInterface, error)
	FindById(id string) (model.PixKeyDomainInterface, error)
	FindPixKey() ([]model.PixKeyDomainInterface, error)
}

func NewPixKeyPersistence(db *gorm.DB) PixKeyPersistenceInterface {
	return pixKeyPersistence{
		db,
	}
}

type pixKeyPersistence struct {
	db *gorm.DB
}

func (p pixKeyPersistence) FindById(pixKey string) (model.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	err := p.db.Where("id = ?", pixKey).First(&pixKeyEntity).Error
	if err != nil {
		return nil, orm_errors.NewPersistenceError(pixKeyEntity.TableName(), err.Error(), "pix_key = "+pixKey)
	}

	pixKeyDomain, err := entity.PixKeyDomainFromEntity(pixKeyEntity)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (p pixKeyPersistence) CreatePixKey(pixKeyDomain model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
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

func (p pixKeyPersistence) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (model.PixKeyDomainInterface, error) {
	var pixKeyEntity entity.PixKeyEntity

	err := p.db.Where("pix_key_type = ?", pixKeyType).Where("pix_key = ?", pixKey).First(&pixKeyEntity).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, orm_errors.NewPersistenceError(pixKeyEntity.TableName(), err.Error(), "pix_key_type = "+pixKeyType+" pix_key = "+pixKey)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	pixKeyDomain, err := entity.PixKeyDomainFromEntity(pixKeyEntity)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (p pixKeyPersistence) FindPixKey() ([]model.PixKeyDomainInterface, error) {
	var pixKeyEntity []entity.PixKeyEntity

	err := p.db.Find(&pixKeyEntity).Error
	if err != nil {
		return nil, orm_errors.NewPersistenceError("pix_key", err.Error(), "")
	}

	var pixKeyDomainList []model.PixKeyDomainInterface

	for _, e := range pixKeyEntity {
		domain, err := entity.PixKeyDomainFromEntity(e)
		if err != nil {
			return nil, err
		}
		pixKeyDomainList = append(pixKeyDomainList, domain)
	}

	return pixKeyDomainList, nil
}
