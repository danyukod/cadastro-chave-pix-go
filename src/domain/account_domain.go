package domain

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"
)

type accountDomain struct {
	accountType enum.AccountType
	number      int
	agency      int
	holder      HolderDomainInterface
}

func (a *accountDomain) Validate() error {
	if a.number < 0 || a.number > 99999999 {
		return errors.New("the account number is invalid")
	}
	if a.agency < 0 || a.agency > 9999 {
		return errors.New("the account agency is invalid")
	}
	if a.accountType.EnumIndex() == 0 {
		return errors.New("the account type is invalid")
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

func (a *accountDomain) GetAccountType() enum.AccountType {
	return a.accountType
}

func (a *accountDomain) GetHolder() HolderDomainInterface {
	return a.holder
}
