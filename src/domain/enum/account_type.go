package enum

import "strings"

type AccountType int

const (
	UNDEFINED_ACCOUNT AccountType = iota
	CORRENTE
	POUPANCA
)

func (t AccountType) String() string {
	return [...]string{"Undefined", "Corrente", "Poupanca"}[t-1]
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
