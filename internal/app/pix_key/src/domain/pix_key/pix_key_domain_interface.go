package pix_key

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() PixKeyType
	GetPixKey() string
	GetAccount() account.AccountDomainInterface
	Validate() error
}

func NewPixKeyDomain(pixKeyType PixKeyType, pixKey string, account account.AccountDomainInterface) (PixKeyDomainInterface, error) {
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

	accountType := account.AccountTypeFromText(request.AccountType)
	accoutDomain, err := account.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain := pixKeyDomain{
		pixKeyType: PixKeyTypeFromText(request.PixKeyType),
		pixKey:     request.PixKey,
		account:    accoutDomain,
	}
	if err := pixKeyDomain.Validate(); err != nil {
		return nil, err
	}
	return &pixKeyDomain, nil
}
