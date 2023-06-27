package domain

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() enum.PixKeyType
	GetPixKey() string
	GetAccount() AccountDomainInterface
	Validate() error
}

func NewPixKeyDomain(pixKeyType enum.PixKeyType, pixKey string, account AccountDomainInterface) (PixKeyDomainInterface, error) {
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
