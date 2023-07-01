package account

type ErrInvalidAccountType struct {
}

func (e *ErrInvalidAccountType) Error() string {
	return "the account type is invalid"
}
