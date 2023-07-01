package pix_key

import (
	"errors"
	account2 "github.com/danyukod/cadastro-chave-pix-go/src/domain/account"
)

type pixKeyDomain struct {
	id         string
	pixKeyType PixKeyType
	pixKey     string
	account    account2.AccountDomainInterface
}

func (p *pixKeyDomain) Validate() error {
	if p.pixKey == "" {
		return &ErrInvalidPixKey{}
	}
	if p.pixKeyType.EnumIndex() == 0 || p.pixKeyType.PixKeyValidate(p.pixKey) == false {
		return &ErrInvalidPixKeyType{}
	}
	if err := p.account.Validate(); err != nil {
		return errors.New("the account is invalid")
	}
	return nil
}

func (p *pixKeyDomain) GetID() string {
	return p.id
}

func (p *pixKeyDomain) SetID(id string) {
	p.id = id
}

func (p *pixKeyDomain) GetPixKeyType() PixKeyType {
	return p.pixKeyType
}

func (p *pixKeyDomain) GetPixKey() string {
	return p.pixKey
}

func (p *pixKeyDomain) GetAccount() account2.AccountDomainInterface {
	return p.account
}
