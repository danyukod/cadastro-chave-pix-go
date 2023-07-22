package pix_key

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
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
