package domain

import (
	account2 "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/aggregate/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/aggregate/holder"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/object_value"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller/model/request"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() object_value.PixKeyType
	GetPixKey() string
	GetAccount() account2.AccountDomainInterface
	Validate() error
}

func NewPixKeyDomain(pixKeyType object_value.PixKeyType, pixKey string, account account2.AccountDomainInterface) (PixKeyDomainInterface, error) {
	pixKeyDomain := pixKeyDomain{
		pixKeyType: pixKeyType,
		pixKey:     pixKey,
		account:    account,
	}
	if err := pixKeyDomain.Validate(); err != nil {
		return nil, err
	}
	return &pixKeyDomain, nil
}

func PixKeyDomainFromRequest(request request.RegisterPixKeyRequest) (PixKeyDomainInterface, error) {
	holderDomain, err := holder.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountType := account2.AccountTypeFromText(request.AccountType)
	accoutDomain, err := account2.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain := pixKeyDomain{
		pixKeyType: object_value.PixKeyTypeFromText(request.PixKeyType),
		pixKey:     request.PixKey,
		account:    accoutDomain,
	}
	if err := pixKeyDomain.Validate(); err != nil {
		return nil, err
	}
	return &pixKeyDomain, nil
}
