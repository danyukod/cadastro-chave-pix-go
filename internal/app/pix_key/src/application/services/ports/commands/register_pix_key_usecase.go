package commands

import (
	"errors"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/services/ports/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
)

type RegisterPixKeyUsecase interface {
	Execute(request.RegisterPixKeyRequest) (domain.PixKeyDomainInterface, error)
}

type RegisterPixKeyService struct {
	pixKeyRepository persistence.RegisterPixKeyRepository
}

func NewRegisterPixKeyService(
	pixKeyRepository persistence.RegisterPixKeyRepository) *RegisterPixKeyService {
	return &RegisterPixKeyService{pixKeyRepository}
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

	err = r.pixKeyRepository.VerifyIfPixKeyAlreadyExists(pixKeyDomain.GetPixKeyType().String(), pixKeyDomain.GetPixKey())
	if checkErrors(err, be) {
		return nil, err
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	pixKeyDomain, err = r.pixKeyRepository.RegisterPixKey(pixKeyDomain)
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
