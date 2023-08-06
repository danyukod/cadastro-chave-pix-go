package services

import (
	"errors"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
	businesserros "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/errors"
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

	holderDomain, err := holder.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)

	accountType := account.AccountTypeFromText(request.AccountType)
	accoutDomain, err := account.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if checkErrors(err, be) {
		return nil, err
	}

	pixKeyType := pix_key.PixKeyTypeFromText(request.PixKeyType)
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyType, request.PixKey, accoutDomain)
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
