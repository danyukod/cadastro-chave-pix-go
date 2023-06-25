package entity

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/entity/enum"
)

type PixKey struct {
	PixKeyType enum.PixKeyType
	PixKey     string
	Account    Account
}

func NewPixKey(pixKeyType enum.PixKeyType, pixKey string, account Account) PixKey {
	return PixKey{
		PixKeyType: pixKeyType,
		PixKey:     pixKey,
		Account:    account,
	}
}

func (p *PixKey) Validate() error {
	if p.PixKey == "" {
		return errors.New("the pix key is invalid")
	}
	if p.PixKeyType.EnumIndex() == 0 || p.PixKeyType.PixKeyValidate(p.PixKey) == false {
		return errors.New("the pix key type is invalid")
	}
	if err := p.Account.Validate(); err != nil {
		return errors.New("the account is invalid")
	}
	return nil
}
