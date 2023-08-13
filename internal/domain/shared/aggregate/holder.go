package aggregate

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
)

type HolderDomainInterface interface {
	Validate() error
	GetName() string
	GetLastName() string
}

func NewHolderDomain(name string, lastName string) (HolderDomainInterface, error) {
	holder := holderDomain{
		name:     name,
		lastName: lastName,
	}
	if err := holder.Validate(); err != nil {
		return nil, err
	}
	return &holder, nil
}

type holderDomain struct {
	name     string
	lastName string
}

func (h *holderDomain) Validate() error {
	var businessErrors value_object.BusinessErrors
	if len(h.name) < 3 || len(h.name) > 50 {
		businessErrors = value_object.AddError(businessErrors, *value_object.NewBusinessError("Account Holder Name", "O nome do titular esta invalido.", "holderName"))
	}
	if businessErrors.Len() > 0 {
		return businessErrors
	}
	return nil
}

func (h *holderDomain) GetName() string {
	return h.name
}

func (h *holderDomain) GetLastName() string {
	return h.lastName
}
