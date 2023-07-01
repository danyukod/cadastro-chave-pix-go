package account

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/holder"
)

type accountDomain struct {
	accountType AccountType
	number      int
	agency      int
	holder      holder.HolderDomainInterface
}

func (a *accountDomain) Validate() error {
	if a.number <= 0 || a.number > 99999999 {
		return &ErrInvalidAccountNumber{}
	}
	if a.agency <= 0 || a.agency > 9999 {
		return &ErrInvalidAccountAgency{}
	}
	if a.accountType.EnumIndex() == 0 {
		return &ErrInvalidAccountType{}
	}
	if err := a.holder.Validate(); err != nil {
		return err
	}
	return nil
}

func (a *accountDomain) GetNumber() int {
	return a.number
}

func (a *accountDomain) GetAgency() int {
	return a.agency
}

func (a *accountDomain) GetAccountType() AccountType {
	return a.accountType
}

func (a *accountDomain) GetHolder() holder.HolderDomainInterface {
	return a.holder
}
