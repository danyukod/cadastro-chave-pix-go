package services

import (
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/ports/output"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
)

type RegisterPixKeyService struct {
	pixKeyRepository output.RegisterPixKeyRepository
}

func NewRegisterPixKeyService(
	pixKeyRepository output.RegisterPixKeyRepository) *RegisterPixKeyService {
	return &RegisterPixKeyService{pixKeyRepository}
}

func (r *RegisterPixKeyService) Execute(request requestpackage.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error) {
	holderDomain, err := holder.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountType := account.AccountTypeFromText(request.AccountType)
	accoutDomain, err := account.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyType := pix_key.PixKeyTypeFromText(request.PixKeyType)
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyType, request.PixKey, accoutDomain)
	if err != nil {
		return nil, err
	}

	exists, err := r.pixKeyRepository.VerifyIfPixKeyAlreadyExists(pixKeyDomain.GetPixKeyType().String())
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, &pix_key.ErrPixKeyAlreadyExists{}
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
