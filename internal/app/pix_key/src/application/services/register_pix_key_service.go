package services

import (
	"errors"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
)

type RegisterPixKeyService struct {
	pixKeyRepository output.RegisterPixKeyRepository
}

var businessErrors businesserros.BusinessErrors

func NewRegisterPixKeyService(
	pixKeyRepository output.RegisterPixKeyRepository) *RegisterPixKeyService {
	return &RegisterPixKeyService{pixKeyRepository}
}

func (r *RegisterPixKeyService) Execute(request requestpackage.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error) {

	var be *businesserros.BusinessErrors

	pixKeyDomain, err := pix_key.NewPixKeyDomain(request)
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

	return PixKeyDomainToWebResponse(pixKeyDomain), nil

}

func checkErrors(err error, be *businesserros.BusinessErrors) bool {
	if err == nil {
		return false
	}
	if errors.As(err, &be) {
		businessErrors = businesserros.AppendErrors(businessErrors, *be)
		return false
	}
	return true
}

func PixKeyDomainToWebResponse(domain pix_key.PixKeyDomainInterface) *response.RegisterPixKeyResponse {
	return &response.RegisterPixKeyResponse{
		Id:                    domain.GetID(),
		PixKeyType:            domain.GetPixKeyType().String(),
		PixKey:                domain.GetPixKey(),
		AccountType:           domain.GetAccount().GetAccountType().String(),
		AccountNumber:         domain.GetAccount().GetNumber(),
		AgencyNumber:          domain.GetAccount().GetAgency(),
		AccountHolderName:     domain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: domain.GetAccount().GetHolder().GetLastName(),
	}
}
