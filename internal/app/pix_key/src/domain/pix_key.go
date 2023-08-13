package domain

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/value_object"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/handler/model/request"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() value_object.PixKeyType
	GetPixKey() string
	GetAccount() aggregate.AccountDomainInterface
	Validate() error
}

func NewPixKeyDomain(pixKeyType value_object.PixKeyType, pixKey string, account aggregate.AccountDomainInterface) (PixKeyDomainInterface, error) {
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
	holderDomain, err := aggregate.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountType := aggregate.AccountTypeFromText(request.AccountType)
	accoutDomain, err := aggregate.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain := pixKeyDomain{
		pixKeyType: value_object.PixKeyTypeFromText(request.PixKeyType),
		pixKey:     request.PixKey,
		account:    accoutDomain,
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
	var businessErrors application.BusinessErrors
	if p.pixKeyType.EnumIndex() == 0 {
		businessErrors = application.AddError(businessErrors, *application.NewBusinessError("Pix Key Type", "O tipo de chave esta invalido.", p.pixKeyType.String()))
	}
	if p.pixKey == "" || p.pixKeyType.PixKeyValidate(p.pixKey) == false {
		businessErrors = application.AddError(businessErrors, *application.NewBusinessError("Pix Key", "O valor da chave esta invalido.", p.pixKeyType.String()))
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
