package model

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() value_object.PixKeyType
	GetPixKey() string
	GetAccount() aggregate.AccountDomainInterface
	Validate() error
}

func NewPixKeyDomain(pixKeyType string, pixKey string, account aggregate.AccountDomainInterface) (PixKeyDomainInterface, error) {
	pixKeyDomain := pixKeyDomain{
		pixKeyType: value_object.PixKeyTypeFromText(pixKeyType),
		pixKey:     pixKey,
		account:    account,
	}
	if err := pixKeyDomain.Validate(); err != nil {
		return nil, err
	}
	return &pixKeyDomain, nil
}

type pixKeyDomain struct {
	id         string
	pixKeyType value_object.PixKeyType
	pixKey     string
	account    aggregate.AccountDomainInterface
}

func (p *pixKeyDomain) Validate() error {
	var businessErrors value_object.BusinessErrors
	if p.pixKeyType.EnumIndex() == 0 {
		businessErrors = value_object.AddError(businessErrors, *value_object.NewBusinessError("Pix Key Type", "O tipo de chave esta invalido.", p.pixKeyType.String()))
	}
	if p.pixKey == "" || p.pixKeyType.PixKeyValidate(p.pixKey) == false {
		businessErrors = value_object.AddError(businessErrors, *value_object.NewBusinessError("Pix Key", "O valor da chave esta invalido.", p.pixKeyType.String()))
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

func (p *pixKeyDomain) GetPixKeyType() value_object.PixKeyType {
	return p.pixKeyType
}

func (p *pixKeyDomain) GetPixKey() string {
	return p.pixKey
}

func (p *pixKeyDomain) GetAccount() aggregate.AccountDomainInterface {
	return p.account
}

type PixKeyValidation interface {
	PixKeyValidate(pixKey string) bool
}
