package account

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
)

type AccountDomainInterface interface {
	Validate() error
	GetNumber() int
	GetAgency() int
	GetAccountType() AccountType
	GetHolder() holder.HolderDomainInterface
}

func NewAccountDomain(number int, agency int, accountType AccountType, holder holder.HolderDomainInterface) (AccountDomainInterface, error) {
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
