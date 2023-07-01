package repository

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
)

type registerPixKeyRepository struct {
	pixKeyPersistence output.PixKeyPersistenceInterface
}

func NewRegisterPixKeyRepository(
	pixKeyPersistence output.PixKeyPersistenceInterface,
) output.RegisterPixKeyRepository {
	return &registerPixKeyRepository{
		pixKeyPersistence,
	}
}

func (p registerPixKeyRepository) RegisterPixKey(pixKeyDomain pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
	pixKeyDomain, err := p.pixKeyPersistence.CreatePixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (p registerPixKeyRepository) VerifyIfPixKeyAlreadyExists(pixKeyType string) (bool, error) {
	pixKeyDomain, err := p.pixKeyPersistence.FindPixKeyByType(pixKeyType)
	if err != nil {
		return false, err
	}

	if pixKeyDomain != nil {
		return true, nil
	}

	return false, nil
}
