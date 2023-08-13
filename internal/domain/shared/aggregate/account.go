package aggregate

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
	"strings"
)

type AccountDomainInterface interface {
	Validate() error
	GetNumber() int
	GetAgency() int
	GetAccountType() AccountType
	GetHolder() HolderDomainInterface
}

func NewAccountDomain(number int, agency int, accountType string, holder HolderDomainInterface) (AccountDomainInterface, error) {
	account := accountDomain{
		accountType: AccountTypeFromText(accountType),
		number:      number,
		agency:      agency,
		holder:      holder,
	}
	if err := account.Validate(); err != nil {
		return nil, err
	}
	return &account, nil
}

type accountDomain struct {
	accountType AccountType
	number      int
	agency      int
	holder      HolderDomainInterface
}

func (a *accountDomain) Validate() error {
	var businessErrors value_object.BusinessErrors
	if a.accountType.EnumIndex() == 0 {
		businessErrors = value_object.AddError(businessErrors, *value_object.NewBusinessError("Account Type", "O tipo de conta esta invalido.", "accountType"))
	}
	if a.number <= 0 || a.number > 99999999 {
		businessErrors = value_object.AddError(businessErrors, *value_object.NewBusinessError("Account Number", "O numero da conta esta invalido.", "accountNumber"))
	}
	if a.agency <= 0 || a.agency > 9999 {
		businessErrors = value_object.AddError(businessErrors, *value_object.NewBusinessError("Agency Number", "O numero da agencia esta invalido.", "agencyNumber"))
	}
	if businessErrors.Len() > 0 {
		return businessErrors
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

func (a *accountDomain) GetHolder() HolderDomainInterface {
	return a.holder
}

type AccountType int

const (
	UNDEFINED_ACCOUNT AccountType = iota
	CORRENTE
	POUPANCA
)

func (t AccountType) String() string {
	return [...]string{"UNDEFINED", "CORRENTE", "POUPANCA"}[t]
}

func (t AccountType) EnumIndex() int {
	return int(t)
}

func (t *AccountType) UnmarshalText(text []byte) error {
	*t = AccountTypeFromText(string(text))
	return nil
}

func AccountTypeFromText(text string) AccountType {
	switch strings.ToLower(text) {
	case "corrente":
		return CORRENTE
	case "poupanca":
		return POUPANCA
	default:
		return UNDEFINED_ACCOUNT
	}
}
