package account

type ErrInvalidAccountType struct{}

type ErrInvalidAccountNumber struct{}

type ErrInvalidAccountAgency struct{}

func (e *ErrInvalidAccountType) Error() string {
	return "the account type is invalid"
}

func (e *ErrInvalidAccountNumber) Error() string {
	return "the account number is invalid"
}

func (e *ErrInvalidAccountAgency) Error() string {
	return "the account agency is invalid"
}
