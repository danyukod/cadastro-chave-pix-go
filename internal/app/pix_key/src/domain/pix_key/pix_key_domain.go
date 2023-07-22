package pix_key

import (
	account2 "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/errors"
)

type pixKeyDomain struct {
	id         string
	pixKeyType PixKeyType
	pixKey     string
	account    account2.AccountDomainInterface
}

func (p *pixKeyDomain) Validate() error {
	var businessErrors errors.BusinessErrors
	if p.pixKeyType.EnumIndex() == 0 {
		businessErrors = errors.AddError(businessErrors, *errors.NewBusinessError("Pix Key Type", "O tipo de chave esta invalido."))
	}
	if p.pixKey == "" || p.pixKeyType.PixKeyValidate(p.pixKey) == false {
		businessErrors = errors.AddError(businessErrors, *errors.NewBusinessError("Pix Key", "O valor da chave esta invalido."))
	}
	if businessErrors.Len() > 0 {
		return businessErrors
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
