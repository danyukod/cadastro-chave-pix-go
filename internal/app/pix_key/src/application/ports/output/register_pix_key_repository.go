package output

import (
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/output/database"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/output/database/errors"
)

type RegisterPixKeyRepository interface {
	RegisterPixKey(
		pixKeyDomain domain.PixKeyDomainInterface,
	) (domain.PixKeyDomainInterface, error)
	VerifyIfPixKeyAlreadyExists(pixKeyType string, pixKey string) error
}

type registerPixKeyRepository struct {
	pixKeyPersistence database.PixKeyPersistenceInterface
}

func NewRegisterPixKeyRepository(
	pixKeyPersistence database.PixKeyPersistenceInterface,
) RegisterPixKeyRepository {
	return &registerPixKeyRepository{
		pixKeyPersistence,
	}
}

func (p registerPixKeyRepository) RegisterPixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	pixKeyDomain, err := p.pixKeyPersistence.CreatePixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (p registerPixKeyRepository) VerifyIfPixKeyAlreadyExists(pixKeyType string, pixKey string) error {
	pixKeyDomain, err := p.pixKeyPersistence.FindPixKeyByKeyAndType(pixKeyType, pixKey)
	if err != nil && !errors.IsNoRecordError(err) {
		return err
	}

	if pixKeyDomain != nil {
		return createBusinessError(pixKey)
	}

	return nil
}

func createBusinessError(pixKey string) businesserros.BusinessErrors {
	var businessErrors businesserros.BusinessErrors
	businessErrors = append(businessErrors, *createPixKeyAlreadyExistsError(pixKey))
	return businessErrors
}

func createPixKeyAlreadyExistsError(pixKey string) *businesserros.BusinessError {
	return businesserros.NewBusinessError("PixKey", "Chave pix ja cadastrada.", pixKey)
}
