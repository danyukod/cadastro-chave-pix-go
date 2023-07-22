package services

import (
	"errors"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
)

type RegisterPixKeyService struct {
	pixKeyRepository output.RegisterPixKeyRepository
}

func NewRegisterPixKeyService(
	pixKeyRepository output.RegisterPixKeyRepository) *RegisterPixKeyService {
	return &RegisterPixKeyService{pixKeyRepository}
}

func (r *RegisterPixKeyService) Execute(request requestpackage.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error) {
	var businessErrors businesserros.BusinessErrors
	holderDomain, err := holder.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)
	if err != nil {
		var be businesserros.BusinessErrors
		if !errors.As(err, &be) {
			return nil, err
		}
		businessErrors = businesserros.AppendErrors(businessErrors, be)
	}

	accountType := account.AccountTypeFromText(request.AccountType)
	accoutDomain, err := account.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		var be businesserros.BusinessErrors
		if !errors.As(err, &be) {
			return nil, err
		}
		businessErrors = businesserros.AppendErrors(businessErrors, be)
	}

	pixKeyType := pix_key.PixKeyTypeFromText(request.PixKeyType)
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyType, request.PixKey, accoutDomain)
	if err != nil {
		var be businesserros.BusinessErrors
		if !errors.As(err, &be) {
			return nil, err
		}
		businessErrors = businesserros.AppendErrors(businessErrors, be)
	}

	if businessErrors.HasErrors() {
		return nil, businessErrors
	}

	exists, err := r.pixKeyRepository.VerifyIfPixKeyAlreadyExists(pixKeyDomain.GetPixKeyType().String(), pixKeyDomain.GetPixKey())
	if err != nil {
		return nil, err
	}

	if exists {
		businessErrors = businesserros.AddError(businessErrors, *businesserros.NewBusinessError("Pix Key", "Chave pix ja cadastrada."))
		return nil, businessErrors
	}

	pixKeyDomain, err = r.pixKeyRepository.RegisterPixKey(pixKeyDomain)
	if err != nil {
		return nil, err
	}

	return &response.RegisterPixKeyResponse{
		Id:                    pixKeyDomain.GetID(),
		PixKeyType:            pixKeyDomain.GetPixKeyType().String(),
		PixKey:                pixKeyDomain.GetPixKey(),
		AccountType:           pixKeyDomain.GetAccount().GetAccountType().String(),
		AccountNumber:         pixKeyDomain.GetAccount().GetNumber(),
		AgencyNumber:          pixKeyDomain.GetAccount().GetAgency(),
		AccountHolderName:     pixKeyDomain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: pixKeyDomain.GetAccount().GetHolder().GetLastName(),
	}, nil

}
