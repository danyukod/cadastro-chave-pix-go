package commands

import (
	"errors"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapter/orm"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/handler/model/request"
)

type RegisterPixKeyUsecase interface {
	Execute(request.RegisterPixKeyRequest) (domain.PixKeyDomainInterface, error)
}

type RegisterPixKeyService struct {
	persistence orm.PixKeyPersistenceInterface
}

func NewRegisterPixKeyService(
	persistence orm.PixKeyPersistenceInterface) *RegisterPixKeyService {
	return &RegisterPixKeyService{persistence: persistence}
}

func (r *RegisterPixKeyService) Execute(request request.RegisterPixKeyRequest) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserros.BusinessErrors

	var be *businesserros.BusinessErrors

	pixKeyDomain, err := domain.PixKeyDomainFromRequest(request)
	if checkErrors(err, be) {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	err = r.VerifyIfPixKeyAlreadyExists(pixKeyDomain.GetPixKeyType().String(), pixKeyDomain.GetPixKey())
	if checkErrors(err, be) {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	pixKeyDomain, err = r.RegisterPixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil

}

func checkErrors(err error, be *businesserros.BusinessErrors) bool {
	var businessErrors businesserros.BusinessErrors

	if err == nil {
		return false
	}
	if errors.As(err, &be) {
		businessErrors = businesserros.AppendErrors(businessErrors, *be)
		return false
	}
	return true
}

func (r *RegisterPixKeyService) RegisterPixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	pixKeyDomain, err := r.persistence.CreatePixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return pixKeyDomain, nil
}

func (r *RegisterPixKeyService) VerifyIfPixKeyAlreadyExists(pixKeyType string, pixKey string) error {
	var businessErrors businesserros.BusinessErrors
	pixKeyDomain, err := r.persistence.FindPixKeyByKeyAndType(pixKeyType, pixKey)
	if err != nil {
		return err
	}

	if pixKeyDomain != nil {
		return businesserros.AddError(businessErrors, *businesserros.CreatePixKeyAlreadyExistsError(pixKey))
	}

	return nil
}
