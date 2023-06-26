package domain

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() enum.PixKeyType
	GetPixKey() string
	GetAccount() Account
	Validate() error
}

func NewPixKeyDomain(pixKeyType enum.PixKeyType, pixKey string, account Account) PixKeyDomainInterface {
	return &pixKeyDomain{
		pixKeyType: pixKeyType,
		pixKey:     pixKey,
		account:    account,
	}
}
