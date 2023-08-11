package account

import (
	application "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/aggregate/holder"
)

type accountDomain struct {
	accountType AccountType
	number      int
	agency      int
	holder      holder.HolderDomainInterface
}

func (a *accountDomain) Validate() error {
	var businessErrors application.BusinessErrors
	if a.accountType.EnumIndex() == 0 {
		businessErrors = application.AddError(businessErrors, *application.NewBusinessError("Account Type", "O tipo de conta esta invalido.", "accountType"))
	}
	if a.number <= 0 || a.number > 99999999 {
		businessErrors = application.AddError(businessErrors, *application.NewBusinessError("Account Number", "O numero da conta esta invalido.", "accountNumber"))
	}
	if a.agency <= 0 || a.agency > 9999 {
		businessErrors = application.AddError(businessErrors, *application.NewBusinessError("Agency Number", "O numero da agencia esta invalido.", "agencyNumber"))
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

func (a *accountDomain) GetHolder() holder.HolderDomainInterface {
	return a.holder
}
