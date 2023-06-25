package entity

import "errors"

type Account struct {
	AccountType AccountType
	Number      int
	Agency      int
	Holder      Holder
}

func NewAccount(number int, agency int, accountType AccountType, holder Holder) (*Account, error) {
	account := Account{
		AccountType: accountType,
		Number:      number,
		Agency:      agency,
		Holder:      holder,
	}
	if err := account.Validate(); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *Account) Validate() error {
	if a.Number < 0 || a.Number > 99999999 {
		return errors.New("the account number is invalid")
	}
	if a.Agency < 0 || a.Agency > 9999 {
		return errors.New("the account agency is invalid")
	}
	if a.AccountType.EnumIndex() == 0 {
		return errors.New("the account type is invalid")
	}
	if err := a.Holder.Validate(); err != nil {
		return err
	}
	return nil
}
