package model

import (
	shared2 "github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/request"
)

type PixKeyDomainInterface interface {
	GetID() string
	SetID(string)
	GetPixKeyType() shared2.PixKeyType
	GetPixKey() string
	GetAccount() shared2.AccountDomainInterface
	Validate() error
}

func NewPixKeyDomain(pixKeyType shared2.PixKeyType, pixKey string, account shared2.AccountDomainInterface) (PixKeyDomainInterface, error) {
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
	holderDomain, err := shared2.NewHolderDomain(request.AccountHolderName, request.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountType := shared2.AccountTypeFromText(request.AccountType)
	accoutDomain, err := shared2.NewAccountDomain(request.AccountNumber, request.AgencyNumber, accountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain := pixKeyDomain{
		pixKeyType: shared2.PixKeyTypeFromText(request.PixKeyType),
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
	pixKeyType shared2.PixKeyType
	pixKey     string
	account    shared2.AccountDomainInterface
}

func (p *pixKeyDomain) Validate() error {
	var businessErrors shared2.BusinessErrors
	if p.pixKeyType.EnumIndex() == 0 {
		businessErrors = shared2.AddError(businessErrors, *shared2.NewBusinessError("Pix Key Type", "O tipo de chave esta invalido.", p.pixKeyType.String()))
	}
	if p.pixKey == "" || p.pixKeyType.PixKeyValidate(p.pixKey) == false {
		businessErrors = shared2.AddError(businessErrors, *shared2.NewBusinessError("Pix Key", "O valor da chave esta invalido.", p.pixKeyType.String()))
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

func (p *pixKeyDomain) GetPixKeyType() shared2.PixKeyType {
	return p.pixKeyType
}

func (p *pixKeyDomain) GetPixKey() string {
	return p.pixKey
}

func (p *pixKeyDomain) GetAccount() shared2.AccountDomainInterface {
	return p.account
}
