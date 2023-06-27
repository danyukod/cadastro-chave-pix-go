package services

import (
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/repository"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"
)

type RegisterPixKeyService struct {
	pixKeyRepository repository.PixKeyRepository
}

func NewRegisterPixKeyService(
	pixKeyRepository repository.PixKeyRepository) *RegisterPixKeyService {
	return &RegisterPixKeyService{pixKeyRepository}
}

func (r *RegisterPixKeyService) Execute(request requestpackage.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error) {
	holderDomain, err := domain.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountType := enum.AccountTypeFromText(request.AccountType)
	accoutDomain, err := domain.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyType := enum.PixKeyTypeFromText(request.PixKeyType)
	pixKeyDomain, err := domain.NewPixKeyDomain(pixKeyType, request.PixKey, accoutDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain, err = r.pixKeyRepository.CreatePixKey(pixKeyDomain)
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
