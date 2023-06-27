package domain

import "github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"

type AccountDomainInterface interface {
	Validate() error
	GetNumber() int
	GetAgency() int
	GetAccountType() enum.AccountType
	GetHolder() HolderDomainInterface
}

func NewAccountDomain(number int, agency int, accountType enum.AccountType, holder HolderDomainInterface) (AccountDomainInterface, error) {
	account := accountDomain{
		accountType: accountType,
		number:      number,
		agency:      agency,
		holder:      holder,
	}
	if err := account.Validate(); err != nil {
		return nil, err
	}
	return &account, nil
}
