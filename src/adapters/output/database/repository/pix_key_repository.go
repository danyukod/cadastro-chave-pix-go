package repository

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
	"gorm.io/gorm"
)

type pixKeyRepository struct {
	database *gorm.DB
}

func NewPixKeyRepository(
	database *gorm.DB,
) PixKeyRepository {
	return &pixKeyRepository{
		database,
	}
}

type PixKeyRepository interface {
	CreatePixKey(
		pixKeyDomain domain.PixKeyDomainInterface,
	) (domain.PixKeyDomainInterface, error)
}
