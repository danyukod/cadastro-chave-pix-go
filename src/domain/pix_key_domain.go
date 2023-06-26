package domain

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"
)

type pixKeyDomain struct {
	id         string
	pixKeyType enum.PixKeyType
	pixKey     string
	account    Account
}

func (p *pixKeyDomain) Validate() error {
	if p.pixKey == "" {
		return errors.New("the pix key is invalid")
	}
	if p.pixKeyType.EnumIndex() == 0 || p.pixKeyType.PixKeyValidate(p.pixKey) == false {
		return errors.New("the pix key type is invalid")
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

func (p *pixKeyDomain) GetPixKeyType() enum.PixKeyType {
	return p.pixKeyType
}

func (p *pixKeyDomain) GetPixKey() string {
	return p.pixKey
}

func (p *pixKeyDomain) GetAccount() Account {
	return p.account
}
