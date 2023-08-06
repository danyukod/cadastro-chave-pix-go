package repository

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/errors"
	"github.com/go-playground/validator/v10"
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

func (p registerPixKeyRepository) VerifyIfPixKeyAlreadyExists(pixKeyType string, pixKey string) error {
	pixKeyDomain, err := p.pixKeyPersistence.FindPixKeyByKeyAndType(pixKeyType, pixKey)
	if err != nil {
		return err
	}

	if pixKeyDomain != nil {
		return createBusinessError(pixKey)
	}

	return nil
}

func createBusinessError(pixKey string) businesserros.BusinessErrors {
	var businessErrors businesserros.BusinessErrors
	businessErrors = append(businessErrors, createPixKeyAlreadyExistsError(pixKey))
	return businessErrors
}

func createPixKeyAlreadyExistsError(pixKey string) validator.FieldError {
	return businesserros.NewBusinessError("PixKey", "Chave pix ja cadastrada.", pixKey)
}
