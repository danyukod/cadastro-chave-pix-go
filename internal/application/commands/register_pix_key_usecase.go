package commands

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/request"
)

type RegisterPixKeyUsecase interface {
	Execute(request.RegisterPixKeyRequest) (model.PixKeyDomainInterface, error)
}

type RegisterPixKeyService struct {
	persistence persistence.PixKeyPersistenceInterface
}

func NewRegisterPixKeyService(
	persistence persistence.PixKeyPersistenceInterface) *RegisterPixKeyService {
	return &RegisterPixKeyService{persistence: persistence}
}

func (r *RegisterPixKeyService) Execute(request request.RegisterPixKeyRequest) (model.PixKeyDomainInterface, error) {
	var businessErrors businesserros.BusinessErrors

	var be *businesserros.BusinessErrors

	pixKeyDomain, err := model.PixKeyDomainFromRequest(request)
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

func (r *RegisterPixKeyService) RegisterPixKey(pixKeyDomain model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
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
